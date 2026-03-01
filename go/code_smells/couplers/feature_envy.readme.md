# Feature Envy

Envidia de características.

## Definición

Una clase usa la información de otra clase colaboradora para hacer cálculos o tomar decisiones, sugiriendo que la segunda clase debería exponer esos comportamientos. Al depender de la estructura de la colaboradora, la clase cliente queda acoplada.

## Ejemplo

`InvoiceService` se mete en los datos de `Customer` para tomar decisiones, lo que indica que el comportamiento quizá debería pertenecer a `Customer`.

```go
package couplers

import "fmt"

// Customer represents a customer
type Customer struct {
	Name    string
	Email   string
	Address string
	Phone   string
}

// Invoice represents an invoice
type Invoice struct {
	InvoiceNumber string
	Amount        float64
	Customer      Customer
}

// InvoiceService demonstrates feature envy code smell
// The method is more interested in the Customer class than its own class
type InvoiceService struct{}

// SendInvoice sends an invoice to a customer
// This method is "envious" of Customer's data
func (is *InvoiceService) SendInvoice(invoice Invoice) error {
	// This method knows too much about Customer's internal structure
	// and uses Customer's data more than its own

	// Formatting customer info
	customerInfo := fmt.Sprintf(
		"Customer: %s\nEmail: %s\nAddress: %s\nPhone: %s",
		invoice.Customer.Name,
		invoice.Customer.Email,
		invoice.Customer.Address,
		invoice.Customer.Phone,
	)

	// Validating customer data
	if invoice.Customer.Email == "" {
		return fmt.Errorf("customer email is required")
	}
	if invoice.Customer.Address == "" {
		return fmt.Errorf("customer address is required")
	}

	fmt.Printf("Sending invoice %s for %.2f to:\n%s\n",
		invoice.InvoiceNumber, invoice.Amount, customerInfo)

	return nil
}
```

## Ejercicio

Añade validación de formato de email y validación de dirección completa (calle, ciudad, código postal).

## Problemas que encontrarás

Probablemente, seguirás añadiendo condiciones dentro de `InvoiceService` que dependen de detalles internos de `Customer`, esparciendo reglas en el lugar equivocado y volviendo frágiles los cambios.
