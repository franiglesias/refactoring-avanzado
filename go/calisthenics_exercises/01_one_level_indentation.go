package calisthenics_exercises

import "fmt"

// ProcessOrders demonstrates multiple levels of indentation
// Rule: Only one level of indentation per method
func ProcessOrders(orders []map[string]interface{}) {
	for _, order := range orders {
		if order["status"] == "pending" {
			if order["total"].(float64) > 0 {
				if order["customer"] != nil {
					customer := order["customer"].(map[string]interface{})
					if customer["email"] != "" {
						fmt.Printf("Processing order for %s\n", customer["email"])
					}
				}
			}
		}
	}
}

// Exercise: Refactor this code to have only one level of indentation
// Techniques to consider:
// 1. Extract methods
// 2. Early returns
// 3. Guard clauses
// 4. Replace nested conditionals with boolean expressions
