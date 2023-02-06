package main

import "github.com/candylaserknight/paulo"

type application struct {
	App *paulo.Paulo
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
