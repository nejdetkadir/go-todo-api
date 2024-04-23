package main

import (
	"go-todo-api/api/routes"
	"go-todo-api/bootstrap"
)

func main() {
	app := bootstrap.NewApp()
	fiberApp, container := app.Init()

	routes.Setup(container)

	app.Run(fiberApp)
}
