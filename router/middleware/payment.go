package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/blackestwhite/presenter"
	"github.com/blackestwhite/zwrapper/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Permitted() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("x-zwrapper-access-token")
		if len(accessToken) < 36 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, presenter.Std{
				Ok:               false,
				ErrorCode:        http.StatusUnauthorized,
				ErrorDescription: "access token is not valid",
			})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		result := db.Client.Database("zwrapper").Collection("tokens").FindOne(ctx, bson.M{
			"token": accessToken,
		})
		if result.Err() != nil {
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
