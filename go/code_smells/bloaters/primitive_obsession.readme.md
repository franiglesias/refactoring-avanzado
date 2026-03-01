# Primitive Obsession

Obsesión primitiva.

## Definición

Conceptos de dominio se modelan con primitivos, lo que obliga a esparcir reglas de validación, formato, y todo tipo de comportamiento, por todo el código.

## Ejemplo

```go
package bloaters

import "fmt"

// Order demonstrates primitive obsession code smell
// Using primitive types instead of domain objects
type Order struct {
	CustomerName  string
	CustomerEmail string
	Address       string
	TotalAmount   float64
	Currency      string
}

// NewOrder creates a new order
func NewOrder(customerName, customerEmail, address string, totalAmount float64, currency string) *Order {
	return &Order{
		CustomerName:  customerName,
		CustomerEmail: customerEmail,
		Address:       address,
		TotalAmount:   totalAmount,
		Currency:      currency,
	}
}

// SendInvoice sends an invoice to the customer
func (o *Order) SendInvoice() error {
	// Validation logic scattered everywhere instead of being in domain objects
	if !containsAt(o.CustomerEmail) {
		return fmt.Errorf("email inválido")
	}
	if o.Address == "" {
		return fmt.Errorf("no se ha indicado dirección")
	}
	if o.TotalAmount <= 0 {
		return fmt.Errorf("el monto debe ser mayor que cero")
	}
	fmt.Printf("Factura enviada a %s in %s por %.2f %s\n",
		o.CustomerName, o.Address, o.TotalAmount, o.Currency)
	return nil
}

func containsAt(email string) bool {
	for _, c := range email {
		if c == '@' {
			return true
		}
	}
	return false
}
```

## Ejercicio

Introduce soporte para diferentes monedas, para enviar la factura por email, y para formatear la dirección en función del país.

## Problemas que encontrarás

Dado que los primitivos no nos permiten garantizar la integridad de sus valores, tendrás que introducir validaciones en muchos lugares, incluso de forma repetida. Algunos datos siempre viajan juntos (Data Clump), por lo que tienes que asegurarte de que permanecen juntos.

Para formatear de forma diferente basándote en algún dato arbitrario tendrás que introducir lógica de decisión.
