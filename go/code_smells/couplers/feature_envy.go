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

// Exercise: Refactor this code to eliminate feature envy
// 1. Move the customer-related logic to the Customer struct
// 2. Create methods like Customer.FormatInfo() and Customer.Validate()
// 3. Keep the InvoiceService focused on invoice-related operations
