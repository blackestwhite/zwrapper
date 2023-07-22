package handler

import (
	"net/http"

	"github.com/blackestwhite/presenter"
	"github.com/blackestwhite/zwrapper/config"
	"github.com/blackestwhite/zwrapper/entity"
	"github.com/blackestwhite/zwrapper/service"
	"github.com/blackestwhite/zwrapper/utils"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	accessTokenService service.AccessTokenService
}

func SetupAdmin(r *gin.RouterGroup) *AdminHandler {
	adminHandler := &AdminHandler{}
	adminHandler.initRoutes(r)
	return adminHandler
}

func (a *AdminHandler) initRoutes(r *gin.RouterGroup) {
	admin := r.Group("/admin")
	admin.POST("/newConsumer", a.newConsumer)
}

func (a *AdminHandler) newConsumer(c *gin.Context) {
	var accessToken entity.AccessToken
	err := c.Bind(&accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, presenter.Std{
			Ok:               false,
			ErrorCode:        http.StatusBadRequest,
			ErrorDescription: err.Error(),
		})
		return
	}

	username := c.Query("username")
	password := c.Query("password")
	if username != config.ADMIN_USERNAME {
		c.AbortWithStatusJSON(http.StatusUnauthorized, presenter.Std{
			Ok:               false,
			ErrorCode:        http.StatusUnauthorized,
			ErrorDescription: "wrong username",
		})
		return
	}
	if utils.Hash(password) != config.ADMIN_PASSWORD {
		c.AbortWithStatusJSON(http.StatusUnauthorized, presenter.Std{
			Ok:               false,
			ErrorCode:        http.StatusUnauthorized,
			ErrorDescription: "wrong password",
		})
		return
	}

	accessToken, err = a.accessTokenService.Create(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, presenter.Std{
			Ok:               false,
			ErrorCode:        http.StatusInternalServerError,
			ErrorDescription: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, presenter.Std{
		Ok: true,
		Result: gin.H{
			"message":  "token generated successfully",
			"consumer": accessToken.Consumer,
			"token":    accessToken.Token,
		},
	})
}
