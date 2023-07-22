package config

import (
	"log"
	"os"

	"github.com/blackestwhite/zwrapper/utils"
	"github.com/joho/godotenv"
)

var (
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

	MERCHANT_ID = os.Getenv("MERCHANT_ID")
	ADMIN_USERNAME = os.Getenv("ADMIN_USERNAME")
	ADMIN_PASSWORD = utils.Hash(os.Getenv("ADMIN_PASSWORD"))
	BASE_URL = os.Getenv("BASE_URL")
}
