package handler

import "github.com/gin-gonic/gin"

type AdminHandler struct{}

func SetupAdmin(r *gin.RouterGroup) *AdminHandler {
	adminHandler := &AdminHandler{}
	return adminHandler
}
