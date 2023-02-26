package paulo

import (
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
	"github.com/candylaserknight/paulo/render"
	"github.com/candylaserknight/paulo/session"
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
	Session  *scs.SessionManager
	DB       Database
	JetViews *jet.Set
	config   config
}

type config struct {
	port        string
	renderer    string
	cookie      cookieConfig
	sessionType string
	database    databaseConfig
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

	// connect to database
	if os.Getenv("DATABASE_TYPE") != "" {
		db, err := p.OpenDB(os.Getenv("DATABASE_TYPE"), p.BuildDSN())
		if err != nil {
			errorLog.Println(err)
			os.Exit(1)
		}
		p.DB = Database{
			DataType: os.Getenv("DATABASE_TYPE"),
			Pool:     db,
		}
	}

	p.InfoLog = infoLog
	p.ErrorLog = errorLog
	p.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	p.Version = version
	p.RootPath = rootPath
	p.Routes = p.routes().(*chi.Mux)

	p.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
		cookie: cookieConfig{
			name:     os.Getenv("COOKIE_NAME"),
			lifetime: os.Getenv("COOKIE_LIFETIME"),
			persist:  os.Getenv("COOKIE_PERSIST"),
			secure:   os.Getenv("COOKIE_SECURE"),
			domain:   os.Getenv("COOKIE_DOMAIN"),
		},
		sessionType: os.Getenv("SESSION_TYPE"),
		database: databaseConfig{
			database: os.Getenv("DATABASE_TYPE"),
			dsn:      p.BuildDSN(),
		},
	}

	// create session

	sess := session.Session{
		CookieLifetime: p.config.cookie.lifetime,
		CookiePersist:  p.config.cookie.persist,
		CookieName:     p.config.cookie.name,
		CookieDomain:   p.config.cookie.domain,
		SessionType:    p.config.sessionType,
	}

	p.Session = sess.InitSession()

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

	defer p.DB.Pool.Close()

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
		Session:  p.Session,
	}

	p.Render = &myRenderer
}

func (p *Paulo) BuildDSN() string {
	var dsn string

	switch os.Getenv("DATABASE_TYPE") {
	case "postgres", "postgresql":
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s timezone=UTC connect_timeout=5",
			os.Getenv("DATABASE_HOST"),
			os.Getenv("DATABASE_PORT"),
			os.Getenv("DATABASE_USER"),
			os.Getenv("DATABASE_NAME"),
			os.Getenv("DATABASE_SSL_MODE"))

		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("%s password=%s", dsn, os.Getenv("DATABASE_PASS"))
		}

	default:
	}

	return dsn
}
