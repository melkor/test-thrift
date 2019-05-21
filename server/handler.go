package server

import (
	"context"
	"math"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/melkor/test-thrift/payment"
	"github.com/sirupsen/logrus"
)

type PaymentHandler struct {
}

func NewPaymentHandler() *PaymentHandler {
	return &PaymentHandler{}
}

func (p *PaymentHandler) Authorize(c context.Context, card *payment.CreditCard) (bool, error) {
	spew.Dump(card)
	if card.Number < int64(math.Pow10(15)) {
		logrus.Infoln("number card is to small")
		return false, nil
	}

	if card.Number > int64(math.Pow10(16)) {
		logrus.Infoln("number card is to big")
		return false, nil
	}

	if card.Cryptogram < int32(math.Pow10(2)) {
		logrus.Infoln("crypto card is to small")
		return false, nil
	}

	if card.Cryptogram > int32(math.Pow10(3)) {
		logrus.Infoln("crypto card is to big")
		return false, nil
	}

	if card.Holder == "" {
		logrus.Infoln("name is empty")
		return false, nil
	}

	if card.Date < time.Now().Unix() {
		logrus.Infoln("card expired")
		return false, nil
	}

	return true, nil
}

func (p *PaymentHandler) Pay(c context.Context, card *payment.CreditCard, amount int32) (bool, error) {
	return true, nil
}
