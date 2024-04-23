package bootstrap

import "github.com/gofiber/fiber/v2"

type Container struct {
	Env      *Env
	FiberApp *fiber.App
}

func NewContainer(app *Application, fiberApp *fiber.App) *Container {
	return &Container{
		Env:      app.Env,
		FiberApp: fiberApp,
	}
}
