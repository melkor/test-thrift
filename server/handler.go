package server

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/melkor/test-thrift/payment"
)

type PaymentHandler struct {
}

func NewPaymentHandler() *PaymentHandler {
	return &PaymentHandler{}
}

func (p *PaymentHandler) Authorize(c context.Context, card *payment.CreditCard) (bool, error) {
	spew.Dump(card)
	if card.Number < int64(math.Pow10(15)) {
		return logError("number card is to small")
	}

	if card.Number > int64(math.Pow10(16)) {
		return logError("number card is to big")
	}

	if card.Cryptogram < int32(math.Pow10(2)) {
		return logError("crypto card is to small")
	}

	if card.Cryptogram > int32(math.Pow10(3)) {
		return logError("crypto card is to big")
	}

	if card.Holder == "" {
		return logError("name is empty")
	}

	if card.Date < time.Now().Unix() {
		return logError("card expired")
	}

	return true, nil
}

func logError(message string) (bool, error) {
	log(message)
	return false, errors.New(message)
}

func (p *PaymentHandler) Pay(c context.Context, card *payment.CreditCard, amount int32) (bool, error) {
	return true, nil
}
