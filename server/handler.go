package server

import "context"

type PaymentHandler struct {
}

func NewPaymentHandler() *PaymentHandler {
	return &PaymentHandler{}
}

func (p *PaymentHandler) Authorize(c context.Context) (bool, error) {
	return true, nil
}

func (p *PaymentHandler) Pay(c context.Context, amount int32) (bool, error) {
	return true, nil
}
