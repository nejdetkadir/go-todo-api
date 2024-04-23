package bootstrap

import (
	"github.com/spf13/viper"
	"log"
)

type EnvType interface {
	GetAppName() string
	GetAppEnv() string
	GetDatabaseURL() string
	GetPort() string
}

type Env struct {
	AppName     string `mapstructure:"APP_NAME"`
	AppEnv      string `mapstructure:"APP_ENV"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	Port        string `mapstructure:"PORT"`
}

func GetEnvironmentVariables() EnvType {
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

func (e *Env) GetAppName() string {
	return e.AppName
}

func (e *Env) GetAppEnv() string {
	return e.AppEnv
}

func (e *Env) GetDatabaseURL() string {
	return e.DatabaseURL
}

func (e *Env) GetPort() string {
	return e.Port
}
