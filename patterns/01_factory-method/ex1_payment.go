package main

import "fmt"

// Step 1: Create a common function using interface
type PaymentMethod interface {
	Pay(amount float32) string
	Status() string
}

// Step 2: List all method using in source code
type PaymentType int

// Define constant value
const (
	Cash PaymentType = iota
	DebitCard
)

// Step 3: Define struct payment methods
type CashPM struct{}
type DebitCardPM struct{}

// Step 4: Override Pay methods with receiver type is CashPM
func (c *CashPM) Pay(amount float32) string {
	fmt.Printf("Pay by Cash method: %f\n", amount)
	return "Success"
}

// Step 6: If we add a new methods, we will receive notification
// that we missing implementing a new methods
func (c *CashPM) Status() string {
	return "OK"
}

// Step 4: Override Pay methods with receiver type is DebitCardPM
func (c *DebitCardPM) Pay(amount float32) string {
	fmt.Printf("Pay by DebitCardPM method: %f\n", amount)
	return "Success"
}

func (c *DebitCardPM) Status() string {
	return "OK"
}

// Step 5: Implement Payment Method
func GetPaymentMethod(t PaymentType) PaymentMethod {
	switch t {
	case Cash:
		return new(CashPM)
	case DebitCard:
		return new(DebitCardPM)
	default:
		return nil
	}
}

func main() {
	// usage
	payment := GetPaymentMethod(DebitCard)
	payment.Pay(20)
	fmt.Println(payment.Status())

	payment2 := GetPaymentMethod(Cash)
	payment2.Pay(50)
	fmt.Print(payment2)
}
