package databasego

type Config struct {
	DRIVER   string `yaml:"DRIVER"`
	DATABASE string `yaml:"DATABASE"`
	HOST     string `yaml:"HOST"`
	PORT     int    `yaml:"PORT"`
	USER     string `yaml:"USER"`
	PASSWORD string `yaml:"PASSWORD"`
}
