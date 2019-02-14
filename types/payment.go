package types

type Payment struct {
	ServiceTaxAmount    int64         `json:"ServiceTaxAmount"`
	Installments        int64         `json:"Installments"`
	Interest            int64         `json:"Interest"`
	Capture             bool          `json:"Capture"`
	Authenticate        bool          `json:"Authenticate"`
	Recurrent           bool          `json:"Recurrent"`
	CreditCard          *CreditCard   `json:"CreditCard"`
	ProofOfSale         string        `json:"ProofOfSale"`
	Tid                 string        `json:"Tid"`
	AuthorizationCode   string        `json:"AuthorizationCode"`
	PaymentID           string        `json:"PaymentId"`
	Type                string        `json:"Type"`
	Amount              int64         `json:"Amount"`
	SoftDescriptor      string        `json:"SoftDescriptor"`
	Provider            string        `json:"Provider"`
	Currency            string        `json:"Currency"`
	Country             string        `json:"Country"`
	ExtraDataCollection []interface{} `json:"ExtraDataCollection"`
	Status              int64         `json:"Status"`
	ReturnCode          string        `json:"ReturnCode"`
	ReturnMessage       string        `json:"ReturnMessage"`
	Links               *[]Link       `json:"Links"`
	ReceivedDate        string        `json:"ReceivedDate"`
	IsSplitted          bool          `json:"IsSplitted"`
}

type Link struct {
	Method string `json:"Method"`
	Rel    string `json:"Rel"`
	Href   string `json:"Href"`
}
