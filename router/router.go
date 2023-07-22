package router

import (
	"github.com/blackestwhite/zwrapper/router/handler"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	api := router.Group("/api")
	v1 := api.Group("/v1")

	handler.SetupAdmin(v1)
	handler.SetupPayment(v1)
}
