package service

import (
	"github.com/blackestwhite/zwrapper/entity"
	"github.com/blackestwhite/zwrapper/repository"
)

type PaymentService struct {
	paymentRepository repository.PaymentRepository
}

func NewPaymentService(paymentRepo repository.PaymentRepository) *PaymentService {
	return &PaymentService{
		paymentRepository: paymentRepo,
	}
}

func (p *PaymentService) Create(instance entity.Payment) (inserted entity.Payment, err error) {
	return p.paymentRepository.CreatePayment(instance)
}

func (p *PaymentService) Get(id string) (payment entity.Payment, err error) {
	return p.paymentRepository.GetPayment(id)
}
