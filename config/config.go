package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB_HOST string
var DB_PORT string
var DB_USER string
var DB_PASSWORD string
var DB_NAME string
var SECRET string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	SECRET = os.Getenv("SECRET")

}
