package calisthenics_exercises

import "fmt"

// Order demonstrates using primitives directly
// Rule: Wrap all primitives and strings in their own types
type Order struct {
	ID           string
	CustomerName string
	Email        string
	Amount       float64
	Currency     string
}

// ValidateAndProcess validates and processes an order
func ValidateAndProcess(order Order) error {
	// Email validation
	if order.Email == "" || !containsAt(order.Email) {
		return fmt.Errorf("invalid email")
	}

	// Amount validation
	if order.Amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}

	fmt.Printf("Processing order %s for %s (%.2f %s)\n",
		order.ID, order.CustomerName, order.Amount, order.Currency)

	return nil
}

func containsAt(s string) bool {
	for _, c := range s {
		if c == '@' {
			return true
		}
	}
	return false
}

// Exercise: Refactor this code to wrap primitives in domain types
// Create types like:
// - Email (with validation)
// - Money (with amount and currency)
// - OrderID
// - CustomerName
// This makes the domain model explicit and validation automatic
