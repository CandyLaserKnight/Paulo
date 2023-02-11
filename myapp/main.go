package main

import (
	"github.com/candylaserknight/paulo"
	"myapp/handlers"
)

type application struct {
	App      *paulo.Paulo
	Handlers *handlers.Handlers
}

func main() {
	p := initApplication()
	p.App.ListenAndServe()
}
