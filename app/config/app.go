package config

import (
	"os"
	"strconv"
)

type AppConfig struct {
	APP  string
	HOST string
	PORT int

	JwtKey string
}

func NewAppConfig() *AppConfig {
	appConfig := &AppConfig{
		APP:    os.Getenv("APP"),
		HOST:   os.Getenv("HOST"),
		JwtKey: os.Getenv("JWT_KEY"),
	}
	appConfig.PORT, _ = strconv.Atoi(os.Getenv("PORT"))

	return appConfig
}
