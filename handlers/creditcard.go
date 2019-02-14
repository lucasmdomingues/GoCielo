package handlers

import (
	"encoding/json"
	"fmt"
	"gocielo/types"
	"net/http"
	"strconv"
)

func CreditCardHandlerSandBox(w http.ResponseWriter, r *http.Request) {

	var log types.Log

	customerName := r.FormValue("customerName")

	cardNumber := r.FormValue("cardNumber")
	cardHolder := r.FormValue("cardHolder")
	cardExpirationDate := r.FormValue("cardExpirationDate")
	cardSecurityCode := r.FormValue("cardSecurityCode")
	cardBrand := r.FormValue("cardBrand")

	amount, err := strconv.ParseInt(r.FormValue("amount"), 10, 64)
	if err != nil {
		log.Error(w, err)
		return
	}

	installments, err := strconv.ParseInt(r.FormValue("installments"), 10, 64)
	if err != nil {
		log.Error(w, err)
		return
	}

	softDescriptor := r.FormValue("softDescriptor")

	customer := &types.Customer{
		Name: customerName,
	}

	creditCard := &types.CreditCard{
		CardNumber:     cardNumber,
		Holder:         cardHolder,
		ExpirationDate: cardExpirationDate,
		SecurityCode:   cardSecurityCode,
		Brand:          cardBrand,
	}

	payment := &types.Payment{
		Type:           "CreditCard",
		Amount:         amount,
		Installments:   installments,
		SoftDescriptor: softDescriptor,
		CreditCard:     creditCard,
		Country:        "BRA",
		Currency:       "BRL",
		Provider:       "Simulado",
	}

	order := &types.Order{
		MerchantOrderID: "2014111703",
		Customer:        customer,
		Payment:         payment,
	}

	jsonValues, err := json.Marshal(order)
	if err != nil {
		log.Error(w, err)
		return
	}

	request := &types.Request{
		ResponseWriter: w,
		Route:          "https://apisandbox.cieloecommerce.cielo.com.br/1/sales",
		Method:         "POST",
		Values:         jsonValues,
	}

	merchant := &types.Merchant{
		Id:  "ea1c300b-22ae-46b2-b7ff-41348f72ef2b",
		Key: "JCFJRVUNHTPKSWFAXXYFYNGJSZEZFYHVRNAPLWBA",
	}

	response, err := request.SendRequest(merchant.Id, merchant.Key)
	if err != nil {
		log.Error(w, err)
		return
	}

	err = json.Unmarshal(response, order)
	if err != nil {
		log.Error(w, err)
		return
	}

	orderJSON, err := json.Marshal(order)
	if err != nil {
		log.Error(w, err)
		return
	}

	fmt.Fprint(w, string(orderJSON))
}
