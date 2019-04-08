# Instalação

```go 
go get github.com/lucasmdomingues/gocielo
```

# Exemplos

### Cartão de crédito - transação simples.

```go
import (
	"fmt"
	"gocielo"
)

func main() {
	
	// ID Merchant & Merchant Key
  merchant := NewMerchant("ID", "KEY")

	// Number, Holder, Expiration Date, Brand, Security Code, Save Card
	card := NewCreditCard("1234123412341231", "Teste Holder", "12/2030", "Visa", "123", false)
	
	// Amount,Installments, Soft Descriptor
	payment := NewPayment(15700, 1, "123456789ABCD")

	// Sandbox, Merchant, Customer Name, Card, Payment, Order ID
	order, err := SendCreditCardPayment(false, merchant, "Comprador crédito simples", card, payment, "2014111703")
	if err != nil {
		fmt.Println(err)
    return
	}
  
  fmt.Println(order.Payment.ReturnMessage)
}
```
