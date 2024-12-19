package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port     string
	Postgres Postgres
}

type Postgres struct {
	DB       string
	Host     string
	Port     string
	User     string
	Password string
}

func Load(path string) Config {
	godotenv.Load(path + "/.env")
	conf := viper.New()

	conf.AutomaticEnv()

	cfg := Config{
		Port: conf.GetString("PORT"),
		Postgres: Postgres{
			DB:       conf.GetString("POSTGRES_DB"),
			Host:     conf.GetString("POSTGRES_HOST"),
			Port:     conf.GetString("POSTGRES_PORT"),
			User:     conf.GetString("POSTGRES_USER"),
			Password: conf.GetString("POSTGRES_PASSWORD"),
		},
	}

	return cfg
}
