# Switch Statements

Sentencias switch.

## Definición

El uso excesivo de `switch` o múltiples `if/else` basados en un código de tipo suele ser una señal de que falta polimorfismo. El problema principal es que cada vez que se añade una nueva variante (un nuevo tipo), hay que buscar y modificar todos los bloques `switch` dispersos por la aplicación.

## Ejemplo

`PaymentProcessor` utiliza múltiples `switch` para decidir cómo procesar pagos, calcular comisiones y obtener nombres según el tipo de método de pago.

```go
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
```

## Ejercicio

Añade un nuevo tipo de método de pago (`cryptocurrency`) con reglas de procesamiento, comisión y nombre especiales.

## Problemas que encontrarás

Tendrás que modificar todos los `switch` dispersos por el código y cualquier otro lugar que dependa del tipo de método de pago. A medida que el sistema crece, olvidar actualizar uno de estos puntos genera bugs difíciles de detectar.
