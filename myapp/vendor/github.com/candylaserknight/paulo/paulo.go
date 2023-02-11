package paulo

import (
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"github.com/candylaserknight/paulo/render"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const version = "1.0.0"

type Paulo struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Routes   *chi.Mux
	Render   *render.Render
	JetViews *jet.Set
	config   config
}

type config struct {
	port     string
	renderer string
}

func (p *Paulo) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	err := p.Init(pathConfig)
	if err != nil {
		return err
	}

	err = p.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	// read .env
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return nil
	}

	// create loggers
	infoLog, errorLog := p.startLoggers()
	p.InfoLog = infoLog
	p.ErrorLog = errorLog
	p.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	p.Version = version
	p.RootPath = rootPath
	p.Routes = p.routes().(*chi.Mux)

	p.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}

	var views = jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
		jet.InDevelopmentMode(),
	)

	p.JetViews = views
	p.createRenderer()
	return nil
}

func (p *Paulo) Init(c initPaths) error {
	root := c.rootPath
	for _, path := range c.folderNames {
		// create folder if it doesn't exist
		err := p.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

// ListenAndServe starts the web server
func (p *Paulo) ListenAndServe() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     p.ErrorLog,
		Handler:      p.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}
	p.InfoLog.Printf("Listening on port %s", os.Getenv("PORT"))
	err := srv.ListenAndServe()
	p.ErrorLog.Fatal(err)
}
func (p *Paulo) checkDotEnv(path string) error {
	err := p.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}
	return nil
}

func (p *Paulo) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}

func (p *Paulo) createRenderer() {
	myRenderer := render.Render{
		Renderer: p.config.renderer,
		RootPath: p.RootPath,
		Port:     p.config.port,
		JetViews: p.JetViews,
	}

	p.Render = &myRenderer
}
