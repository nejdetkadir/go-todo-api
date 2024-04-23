package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-todo-api/api/handlers"
	"go-todo-api/bootstrap"
)

func DefineTodoRoutes(router fiber.Router, container *bootstrap.Container, validator CustomValidator) {
	handler := handlers.NewTodoHandler(container, &validator)

	router.Get("/todos", handler.GetTodos)
	router.Get("/todos/deleted", handler.GetDeletedTodos)
	router.Get("/todos/:id", handler.GetTodoById)
	router.Post("/todos", handler.CreateTodo)
	router.Put("/todos/:id", handler.UpdateTodoById)
	router.Delete("/todos/:id", handler.DeleteTodoById)
	router.Patch("/todos/:id/complete", handler.MarkAsCompleted)
	router.Patch("/todos/:id/uncomplete", handler.MarkAsUncompleted)
	router.Patch("/todos/:id/recover", handler.RecoverTodo)
}
