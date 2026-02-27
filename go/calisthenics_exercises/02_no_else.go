package calisthenics_exercises

// CalculateDiscount calculates discount using if-else
// Rule: Don't use the ELSE keyword
func CalculateDiscount(customerType string, amount float64) float64 {
	if customerType == "premium" {
		return amount * 0.20
	} else if customerType == "gold" {
		return amount * 0.15
	} else if customerType == "silver" {
		return amount * 0.10
	} else {
		return amount * 0.05
	}
}

// Exercise: Refactor this code to eliminate all else keywords
// Techniques to consider:
// 1. Early returns
// 2. Guard clauses
// 3. Polymorphism (strategy pattern)
// 4. Look-up tables or maps
