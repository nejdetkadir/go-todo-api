package main

import (
	"go-todo-api/api/routes"
	"go-todo-api/bootstrap"
)

func main() {
	app := bootstrap.NewApp()
	fiberApp := app.Init()

	routes.Setup(fiberApp)

	app.Run(fiberApp)
}
