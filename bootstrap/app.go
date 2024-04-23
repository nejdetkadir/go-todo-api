package bootstrap

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
)

type Application struct {
	Env *Env
	DB  *gorm.DB
}

func NewApp() Application {
	app := &Application{}
	app.Env = GetEnvironmentVariables()
	app.DB = NewPSQLConnection(app.Env)

	return *app
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
	log.Println("The " + app.Env.AppName + " is running in " + app.Env.AppEnv + " mode")
}

func (app *Application) Init() *fiber.App {
	fiberConfig := getFiberConfig(app.Env)
	fiberApp := fiber.New(*fiberConfig)

	return fiberApp
}

func (app *Application) Run(fiberApp *fiber.App) {
	app.OnStartup()
	defer app.OnShutdown()

	log.Fatal(fiberApp.Listen(fmt.Sprintf(":%s", app.Env.Port)))
}

func getFiberConfig(env *Env) *fiber.Config {
	if env.AppEnv == "development" {
		return &fiber.Config{}
	}

	return &fiber.Config{}
}
