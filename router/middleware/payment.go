package middleware

import (
	"net/http"

	"github.com/blackestwhite/presenter"
	"github.com/blackestwhite/zwrapper/service"
	"github.com/gin-gonic/gin"
)

func Permitted() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessTokenService := service.AccessTokenService{}
		accessToken := c.GetHeader("x-zwrapper-access-token")
		if len(accessToken) < 36 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, presenter.Std{
				Ok:               false,
				ErrorCode:        http.StatusUnauthorized,
				ErrorDescription: "access token is not valid",
			})
			return
		}

		_, err := accessTokenService.GetByToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, presenter.Std{
				Ok:               false,
				ErrorCode:        http.StatusUnauthorized,
				ErrorDescription: "access token is not valid",
			})
			return
		}

		c.Next()
	}
}
