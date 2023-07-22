package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	JWT_SECRET  string
	MERCHANT_ID string
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	JWT_SECRET = os.Getenv("JWT_SECRET")
	MERCHANT_ID = os.Getenv("MERCHANT_ID")
}
