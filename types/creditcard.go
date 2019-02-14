package types

type CreditCard struct {
	CardNumber     string `json:"CardNumber"`
	Holder         string `json:"Holder"`
	ExpirationDate string `json:"ExpirationDate"`
	SaveCard       bool   `json:"SaveCard"`
	Brand          string `json:"Brand"`
	SecurityCode   string `json:"SecurityCode"`
}
