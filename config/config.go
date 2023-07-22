package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	JWT_SECRET     string
	MERCHANT_ID    string
	ADMIN_USERNAME string
	ADMIN_PASSWORD string
	BASE_URL       string
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	JWT_SECRET = os.Getenv("JWT_SECRET")
	MERCHANT_ID = os.Getenv("MERCHANT_ID")
	ADMIN_USERNAME = os.Getenv("ADMIN_USERNAME")
	ADMIN_PASSWORD = os.Getenv("ADMIN_PASSWORD")
	BASE_URL = os.Getenv("BASE_URL")
}
