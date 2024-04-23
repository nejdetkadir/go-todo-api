package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-todo-api/bootstrap"
)

func DefineHealthCheckRoutes(container *bootstrap.Container) {
	container.FiberApp.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":      "ok",
			"environment": container.Env.GetAppEnv(),
		})
	})
}
