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

// Exercise: Refactor this code to use domain objects instead of primitives
// 1. Create an Email value object with validation
// 2. Create an Address value object
// 3. Create a Money value object to handle amount and currency together
// 4. Modify the Order to use these value objects
