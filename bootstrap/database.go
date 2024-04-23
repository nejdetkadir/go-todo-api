package bootstrap

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func NewPSQLConnection(env *Env) *gorm.DB {
	db, err := gorm.Open(postgres.Open(env.DatabaseURL), getGormConfig(env))
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	
	return db
}

func getGormConfig(env *Env) *gorm.Config {
	if env.AppEnv == "development" {
		return &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}

	return &gorm.Config{}
}
