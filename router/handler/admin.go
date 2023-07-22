package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/blackestwhite/presenter"
	"github.com/blackestwhite/zwrapper/config"
	"github.com/blackestwhite/zwrapper/db"
	"github.com/blackestwhite/zwrapper/entity"
	"github.com/blackestwhite/zwrapper/utils"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type AdminHandler struct{}

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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	u, err := uuid.NewV4()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, presenter.Std{
			Ok:               false,
			ErrorCode:        http.StatusInternalServerError,
			ErrorDescription: err.Error(),
		})
		return
	}
	accessToken.Token = u.String()

	res, err := db.Client.Database("api").Collection("tokens").InsertOne(ctx, accessToken)
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
			"id":       res.InsertedID,
			"consumer": accessToken.Consumer,
			"token":    accessToken.Token,
		},
	})
}
