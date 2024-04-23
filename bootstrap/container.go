package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"go-todo-api/domain"
	"go-todo-api/repository"
	"go-todo-api/service"
)

type Container struct {
	Env            EnvType
	FiberApp       *fiber.App
	TodoRepository domain.TodoRepository
	TodoService    domain.TodoService
}

func NewContainer(app *Application, fiberApp *fiber.App) *Container {
	todoRepository := repository.NewTodoRepository(app)
	todoService := service.NewTodoService(todoRepository)

	return &Container{
		Env:            app.Env,
		FiberApp:       fiberApp,
		TodoRepository: todoRepository,
		TodoService:    todoService,
	}
}
