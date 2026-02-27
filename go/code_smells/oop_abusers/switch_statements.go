package oop_abusers

import "fmt"

// PaymentMethod represents different payment methods
type PaymentMethod string

const (
	CreditCard PaymentMethod = "credit_card"
	PayPal     PaymentMethod = "paypal"
	BankTransfer PaymentMethod = "bank_transfer"
	Cash       PaymentMethod = "cash"
)

// PaymentProcessor demonstrates switch statement code smell
type PaymentProcessor struct{}

// ProcessPayment processes a payment using switch statements
func (pp *PaymentProcessor) ProcessPayment(method PaymentMethod, amount float64) error {
	switch method {
	case CreditCard:
		fmt.Printf("Processing credit card payment for $%.2f\n", amount)
		fmt.Println("Validating card...")
		fmt.Println("Contacting payment gateway...")
		fmt.Println("Payment successful")
		return nil

	case PayPal:
		fmt.Printf("Processing PayPal payment for $%.2f\n", amount)
		fmt.Println("Redirecting to PayPal...")
		fmt.Println("Payment successful")
		return nil

	case BankTransfer:
		fmt.Printf("Processing bank transfer for $%.2f\n", amount)
		fmt.Println("Generating transfer reference...")
		fmt.Println("Payment pending confirmation")
		return nil

	case Cash:
		fmt.Printf("Processing cash payment for $%.2f\n", amount)
		fmt.Println("Payment received")
		return nil

	default:
		return fmt.Errorf("unknown payment method: %s", method)
	}
}

// CalculateFee calculates processing fee - another switch on the same type
func (pp *PaymentProcessor) CalculateFee(method PaymentMethod, amount float64) float64 {
	switch method {
	case CreditCard:
		return amount * 0.029 // 2.9% for credit cards
	case PayPal:
		return amount * 0.034 // 3.4% for PayPal
	case BankTransfer:
		return 2.5 // flat fee for bank transfers
	case Cash:
		return 0 // no fee for cash
	default:
		return 0
	}
}

// GetPaymentMethodName gets the display name - yet another switch
func (pp *PaymentProcessor) GetPaymentMethodName(method PaymentMethod) string {
	switch method {
	case CreditCard:
		return "Tarjeta de Crédito"
	case PayPal:
		return "PayPal"
	case BankTransfer:
		return "Transferencia Bancaria"
	case Cash:
		return "Efectivo"
	default:
		return "Desconocido"
	}
}

// Exercise: Refactor this code to use polymorphism instead of switch statements
// 1. Create a PaymentStrategy interface with methods: Process, CalculateFee, GetName
// 2. Implement the interface for each payment method
// 3. Replace the switch statements with polymorphic calls
// 4. Notice how adding a new payment method becomes easier
