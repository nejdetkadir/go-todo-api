package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	apiGroup := app.Group("/api")
	v1 := apiGroup.Group("/v1")

	DefineHelloRoutes(v1)
}
