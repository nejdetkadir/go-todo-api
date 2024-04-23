package bootstrap

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	AppName     string `mapstructure:"APP_NAME"`
	AppEnv      string `mapstructure:"APP_ENV"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	Port        string `mapstructure:"PORT"`
}

func GetEnvironmentVariables() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	if env.AppEnv == "" {
		env.AppEnv = "development"
	}

	if env.Port == "" {
		env.Port = "3000"
	}

	return &env
}
