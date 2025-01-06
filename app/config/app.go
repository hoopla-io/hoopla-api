package config

import "github.com/itv-go/yamlreader"

type AppConfig struct {
	APP  string `yaml:"APP"`
	HOST string `yaml:"HOST"`
	PORT int    `yaml:"PORT"`

	JwtKey string `yaml:"JWT_KEY"`
}

func NewAppConfig() *AppConfig {
	read, err := yamlreader.Read("config/app.yaml", &AppConfig{})
	if err != nil {
		return nil
	}

	return read
}
