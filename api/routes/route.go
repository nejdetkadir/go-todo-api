package routes

import (
	"go-todo-api/bootstrap"
)

func Setup(container *bootstrap.Container) {
	apiGroup := container.FiberApp.Group("/api")
	v1 := apiGroup.Group("/v1")

	DefineHelloRoutes(v1, container)
}
