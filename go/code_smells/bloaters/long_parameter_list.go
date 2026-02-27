package bloaters

import "fmt"

// CreateUser demonstrates long parameter list code smell
func CreateUser(
	username string,
	email string,
	firstName string,
	lastName string,
	age int,
	country string,
	city string,
	address string,
	postalCode string,
	phoneNumber string,
	isActive bool,
	role string,
) error {
	// Validation
	if username == "" || email == "" {
		return fmt.Errorf("username and email are required")
	}

	// Business logic
	fmt.Printf("Creating user: %s (%s)\n", username, email)
	fmt.Printf("Name: %s %s\n", firstName, lastName)
	fmt.Printf("Location: %s, %s, %s, %s\n", address, city, postalCode, country)
	fmt.Printf("Contact: %s\n", phoneNumber)
	fmt.Printf("Age: %d, Role: %s, Active: %t\n", age, role, isActive)

	return nil
}

// Exercise: Refactor this code to reduce the parameter list
// 1. Group related parameters into structs (e.g., Address, PersonalInfo)
// 2. Consider using the Builder pattern
// 3. Think about which parameters are truly required vs optional
