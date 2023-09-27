package config

import (
	"api-tabungan/infrastructure/database"
	"api-tabungan/infrastructure/logger"
	shared_constant "api-tabungan/infrastructure/shared/constant"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvironmentConfig struct {
	Env      string
	App      App
	Database database.DatabaseConfig
	Log      logger.LogConfig
}

type App struct {
	Name    string
	Version string
	Port    int
}

type Log struct {
	Path      string
	Prefix    string
	Extension string
}

func LoadENVConfig() (config EnvironmentConfig, err error) {
	err = godotenv.Load()
	if err != nil {
		err = fmt.Errorf(shared_constant.ErrConvertStringToInt, err)
		return
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		err = fmt.Errorf(shared_constant.ErrConvertStringToInt, err)
		return
	}

	config = EnvironmentConfig{
		Env: os.Getenv("ENV"),
		App: App{
			Name:    os.Getenv("APP_NAME"),
			Version: os.Getenv("APP_VERSION"),
			Port:    port,
		},
		Database: database.DatabaseConfig{
			Dialect:        os.Getenv("DB_DIALECT"),
			Host:           os.Getenv("DB_HOST"),
			Name:           os.Getenv("DB_NAME"),
			Username:       os.Getenv("DB_USERNAME"),
			Password:       os.Getenv("DB_PASSWORD"),
			Port:           os.Getenv("DB_PORT"),
			SetMaxIdleConn: os.Getenv("DB_SET_MAX_IDLE_CONN"),
			SetMaxOpenConn: os.Getenv("DB_SET_MAX_OPEN_CONN"),
			SetMaxIdleTime: os.Getenv("DB_SET_MAX_IDLE_TIME"),
			SetMaxLifeTime: os.Getenv("DB_SET_MAX_LIFE_TIME"),
		},
		Log: logger.LogConfig{
			Path:      os.Getenv("LOG_PATH"),
			Prefix:    os.Getenv("LOG_PREFIX"),
			Extension: os.Getenv("LOG_EXT"),
		},
	}

	return
}
