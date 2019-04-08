package gocielo

import (
	"testing"
)

func TestSendCreditCardPayment(t *testing.T) {

	merchant := NewMerchant("ID", "KEY")

	card := NewCreditCard("1234123412341231", "Teste Holder", "12/2030", "Visa", "123", false)
	payment := NewPayment(15700, 1, "123456789ABCD")

	order, err := SendCreditCardPayment(merchant, "Comprador crédito simples", card, payment, "2014111703")
	if err != nil {
		t.Error(order)
	}

	t.Log(order.Payment.ReturnMessage)
}
