package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadENV() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file")
		} else {
			log.Println(".env file loaded")
		}
	}
}
