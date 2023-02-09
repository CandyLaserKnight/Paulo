package main

import (
	"log"
	"myapp/handlers"
	"os"

	"github.com/candylaserknight/paulo"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init paulo
	pau := &paulo.Paulo{}
	err = pau.New(path)
	if err != nil {
		log.Fatal(err)
	}

	pau.AppName = "myapp"

	myHandlers := &handlers.Handlers{
		App: pau,
	}

	app := &application{
		App:      pau,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()

	return app
}
