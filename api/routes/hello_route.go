package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-todo-api/bootstrap"
)

func DefineHelloRoutes(router fiber.Router, container *bootstrap.Container) {
	router.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Hello, %s!", container.Env.AppName))
	})
}
