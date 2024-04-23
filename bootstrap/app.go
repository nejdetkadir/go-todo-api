package bootstrap

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-todo-api/domain"
	"gorm.io/gorm"
	"log"
)

type ApplicationType interface {
	GetEnv() EnvType
	GetDB() *gorm.DB
	GetContainer(fiberApp *fiber.App) *Container
	Run(fiberApp *fiber.App)
	Init() (*fiber.App, *Container)
}

type Application struct {
	Env EnvType
	DB  *gorm.DB
}

func NewApp() ApplicationType {
	app := &Application{}
	app.Env = GetEnvironmentVariables()
	app.DB = NewPSQLConnection(app.Env)

	return app
}

func (app *Application) OnShutdown() {
	sqlDB, err := app.DB.DB()
	if err != nil {
		log.Fatal("Error getting database connection: ", err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatal("Error closing database connection: ", err)
	}
}

func (app *Application) OnStartup() {
	log.Println("The " + app.Env.GetAppName() + " is running in " + app.Env.GetAppEnv() + " mode")

	AutoMigrate(app.DB)
}

func (app *Application) Init() (*fiber.App, *Container) {
	fiberConfig := getFiberConfig(app.Env)
	fiberApp := fiber.New(*fiberConfig)

	container := app.GetContainer(fiberApp)

	return fiberApp, container
}

func (app *Application) Run(fiberApp *fiber.App) {
	app.OnStartup()
	defer app.OnShutdown()

	log.Fatal(fiberApp.Listen(fmt.Sprintf(":%s", app.Env.GetPort())))
}

func getFiberConfig(env EnvType) *fiber.Config {
	/* custom config for development if needed
	if env.GetAppEnv() == "development" {
		return &fiber.Config{}
	}
	*/

	return &fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			message := "An unexpected error occurred"

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
				message = e.Message
			}

			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

			return c.Status(code).JSON(domain.GlobalErrorResponse{
				Status:  code,
				Message: message,
			})
		},
	}
}

func (app *Application) GetContainer(fiberApp *fiber.App) *Container {
	return NewContainer(app, fiberApp)
}

func (app *Application) GetEnv() EnvType {
	return app.Env
}

func (app *Application) GetDB() *gorm.DB {
	return app.DB
}
