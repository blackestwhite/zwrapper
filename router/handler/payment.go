package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/blackestwhite/presenter"
	"github.com/blackestwhite/zwrapper/config"
	"github.com/blackestwhite/zwrapper/db"
	"github.com/blackestwhite/zwrapper/entity"
	"github.com/blackestwhite/zwrapper/gateway"
	"github.com/blackestwhite/zwrapper/repository"
	"github.com/blackestwhite/zwrapper/router/middleware"
	"github.com/blackestwhite/zwrapper/service"
	"github.com/blackestwhite/zwrapper/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type PaymentHandler struct {
	paymentService service.PaymentService
}

func SetupPayment(r *gin.RouterGroup) *PaymentHandler {
	paymentRepo := repository.NewMongoPaymentRepository(db.Client, "zwrapper", "payemnts")
	paymentService := service.NewPaymentService(paymentRepo)
	paymentHandler := &PaymentHandler{
		paymentService: *paymentService,
	}
	paymentHandler.initRoutes(r)
	return paymentHandler
}

func (p *PaymentHandler) initRoutes(r *gin.RouterGroup) {
	payment := r.Group("/payment")

	payment.GET("/landing", landing)
	payment.GET("/pay/:id", p.payPayment)

	payment.POST("/new", middleware.Permitted(), p.newPayment)
	payment.POST("/verify/:id", middleware.Permitted(), p.verifyPayment)
}

func landing(c *gin.Context) {
	authority := c.Query("Authority")
	status := c.Query("Status")
	paid := false

	filter := bson.M{
		"authority": authority,
	}
	var p entity.Payment
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	db.Client.Database("zwrapper").Collection("payments").FindOne(ctx, filter).Decode(&p)

	if status == "OK" {
		paid = true
		_, _, _, err := gateway.Instance.PaymentVerification(p.Amount, authority)
		if err != nil {
			// log.Println(status, v, ref, st, err)
			paid = false
		}
	}

	if paid {
		go func() {
			wp := entity.WebhookPayload{
				ID:     p.ID.Hex(),
				Amount: p.Amount,
				Key:    p.Key,
				URL:    p.Webhook,
			}
			err := utils.Post(wp)
			if err != nil {
				// try again on failure
				utils.Post(wp)
			}
		}()
	}

	c.HTML(200, "index.html", gin.H{
		"succeeded": paid,
		"failed":    !paid,
		"next":      p.Next,
	})
}

func (p *PaymentHandler) newPayment(c *gin.Context) {
	var payment entity.Payment
	err := json.NewDecoder(c.Request.Body).Decode(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, presenter.Std{
			Ok:               false,
			ErrorCode:        http.StatusInternalServerError,
			ErrorDescription: err.Error(),
		})
		return
	}

	_, auth, _, err := gateway.Instance.NewPaymentRequest(payment.Amount, config.BASE_URL+"/api/v1/payment/landing", payment.Description, "", "")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, presenter.Std{
			Ok:               false,
			ErrorCode:        http.StatusInternalServerError,
			ErrorDescription: err.Error(),
		})
		return
	}

	payment.Authority = auth
	payment.Key = c.GetHeader("x-zwrapper-access-token")

	payment, err = p.paymentService.Create(payment)
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
			"id": payment.ID.Hex(),
		},
	})
}

func (p *PaymentHandler) payPayment(c *gin.Context) {
	payment, err := p.paymentService.Get(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, presenter.Std{
			Ok:               false,
			ErrorCode:        http.StatusInternalServerError,
			ErrorDescription: err.Error(),
		})
		return
	}

	gateway.Instance.RefreshAuthority(payment.Authority, 1800)

	c.Redirect(http.StatusPermanentRedirect, gateway.Instance.PaymentEndpoint+payment.Authority)
}

func (p *PaymentHandler) verifyPayment(c *gin.Context) {
	payment, err := p.paymentService.Get(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, presenter.Std{
			Ok:               false,
			ErrorCode:        http.StatusInternalServerError,
			ErrorDescription: err.Error(),
		})
		return
	}

	verified, ref, status, err := gateway.Instance.PaymentVerification(payment.Amount, payment.Authority)
	if err != nil {
		c.JSON(200, presenter.Std{
			Ok: true,
			Result: gin.H{
				"paid":     false,
				"verified": verified,
				"ref":      ref,
				"status":   status,
				"error":    err.Error(),
			},
		})
		return
	}
	log.Println(verified, ref, status, err)
	if err != nil {
		return
	}
	c.JSON(200, presenter.Std{
		Ok: true,
		Result: gin.H{
			"paid": true,
			"ref":  ref,
		},
	})
}
