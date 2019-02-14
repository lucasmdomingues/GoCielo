package types

type Order struct {
	MerchantOrderID string    `json:"MerchantOrderId"`
	Customer        *Customer `json:"Customer"`
	Payment         *Payment  `json:"Payment"`
}
