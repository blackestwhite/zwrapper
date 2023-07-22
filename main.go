package main

import (
	"log"

	"github.com/blackestwhite/zwrapper/config"
	"github.com/blackestwhite/zwrapper/db"
	"github.com/blackestwhite/zwrapper/gateway"
	"github.com/blackestwhite/zwrapper/router"
	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config.Load()
	db.Connect()
	gateway.Initiate()

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.LoadHTMLGlob("./templates/*")
	router.Setup(engine)
	log.Panic(engine.Run(":8080"))
}
