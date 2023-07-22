package gateway

import (
	"log"

	"github.com/blackestwhite/go-zarinpal"
	"github.com/blackestwhite/zwrapper/config"
)

var Instance *zarinpal.Zarinpal

func Initiate() {
	i, err := zarinpal.NewZarinpal(config.MERCHANT_ID, false)
	if err != nil {
		log.Panic(err)
	}
	Instance = i
}
