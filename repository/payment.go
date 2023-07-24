package repository

import "github.com/blackestwhite/zwrapper/entity"

type PaymentRepository interface {
	CreatePayment(payment entity.Payment) (entity.Payment, error)
	GetPayment(id string) (entity.Payment, error)
}
