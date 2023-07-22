package main

import (
	"log"

	"github.com/blackestwhite/zwrapper/config"
	"github.com/blackestwhite/zwrapper/db"
	"github.com/blackestwhite/zwrapper/gateway"
	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config.Load()
	db.Connect()
	gateway.Initiate()

	router := gin.New()
	router.LoadHTMLGlob("./templates/*")
	log.Panic(router.Run(":8080"))
}
