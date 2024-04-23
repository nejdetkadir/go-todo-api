package bootstrap

import (
	"go-todo-api/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewPSQLConnection(env EnvType) *gorm.DB {
	db, err := gorm.Open(postgres.Open(env.GetDatabaseURL()), getGormConfig(env))
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	return db
}

func getGormConfig(env EnvType) *gorm.Config {
	if env.GetAppEnv() == "development" {
		return &gorm.Config{
			//Logger: logger.Default.LogMode(logger.Info),
		}
	}

	return &gorm.Config{}
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&domain.Todo{})

	if err != nil {
		log.Fatal("Error while migrating the database: ", err)
	}
}
