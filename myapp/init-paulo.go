package main

import (
	"log"
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

	pau.InfoLog.Println("Debug is set to", pau.Debug)

	app := &application{
		App: pau,
	}

	return app
}
