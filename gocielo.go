package gocielo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const CIELO_PREFIX = "https://apisandbox.cieloecommerce.cielo.com.br"

func SendCreditCardPayment(merchant *Merchant, customerName string, card *CreditCard, payment *Payment, orderID string) (*Order, error) {

	data, err := json.Marshal(&Order{
		MerchantOrderID: orderID,
		Customer: &Customer{
			Name: customerName,
		},
		Payment: &Payment{
			Type:           "CreditCard",
			Amount:         payment.Amount,
			Installments:   payment.Installments,
			SoftDescriptor: payment.SoftDescriptor,
			CreditCard: &CreditCard{
				CardNumber:     card.CardNumber,
				Holder:         card.Holder,
				ExpirationDate: card.ExpirationDate,
				SecurityCode:   card.SecurityCode,
				Brand:          card.Brand,
			},
			Country:  "BRA",
			Currency: "BRL",
			Provider: "Simulado",
		}})
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/1/sales/", CIELO_PREFIX)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Add("MerchantId", merchant.Id)
	req.Header.Add("MerchantKey", merchant.Key)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var order *Order

	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func NewPayment(amount, installments int64, softDescriptor string) *Payment {
	return &Payment{
		Amount:         amount,
		Installments:   installments,
		SoftDescriptor: softDescriptor,
	}
}

func NewCreditCard(number, holder, expirationDate, brand, securityCode string, saveCard bool) *CreditCard {
	return &CreditCard{
		CardNumber:     number,
		Holder:         holder,
		ExpirationDate: expirationDate,
		SaveCard:       saveCard,
		Brand:          brand,
		SecurityCode:   securityCode,
	}
}

func NewMerchant(merchantID, merchantKey string) *Merchant {
	return &Merchant{
		Id:  merchantID,
		Key: merchantKey,
	}
}
