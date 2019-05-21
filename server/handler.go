package server

import (
	"context"

	"github.com/melkor/test-thrift/payment"
)

type PaymentHandler struct {
}

func NewPaymentHandler() *PaymentHandler {
	return &PaymentHandler{}
}

func (p *PaymentHandler) Authorize(c context.Context, card *payment.CreditCard) (bool, error) {
	return true, nil
}

func (p *PaymentHandler) Pay(c context.Context, card *payment.CreditCard, amount int32) (bool, error) {
	return true, nil
}
