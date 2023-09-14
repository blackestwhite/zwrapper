package config

import (
	"os"

	"github.com/blackestwhite/zwrapper/utils"
)

var (
	MERCHANT_ID    string
	ADMIN_USERNAME string
	ADMIN_PASSWORD string
	BASE_URL       string
)

func Load() {
	MERCHANT_ID = os.Getenv("MERCHANT_ID")
	ADMIN_USERNAME = os.Getenv("ADMIN_USERNAME")
	ADMIN_PASSWORD = utils.Hash(os.Getenv("ADMIN_PASSWORD"))
	BASE_URL = os.Getenv("BASE_URL")
}
