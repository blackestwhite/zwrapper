package handler

import "github.com/gin-gonic/gin"

type PaymentHandler struct{}

func SetupPayment(r *gin.RouterGroup) *PaymentHandler {
	paymentHandler := &PaymentHandler{}
	return paymentHandler
}
