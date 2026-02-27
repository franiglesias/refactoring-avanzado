package dispensables

import "fmt"

// EmailService demonstrates duplicated code smell
type EmailService struct{}

// SendWelcomeEmail sends a welcome email
func (es *EmailService) SendWelcomeEmail(email, name string) error {
	// Duplicated validation
	if email == "" {
		return fmt.Errorf("email is required")
	}
	if !containsAt(email) {
		return fmt.Errorf("invalid email format")
	}

	// Duplicated formatting
	subject := "Welcome!"
	body := fmt.Sprintf("Hello %s, welcome to our service!", name)

	// Duplicated sending logic
	fmt.Printf("Sending email to: %s\n", email)
	fmt.Printf("Subject: %s\n", subject)
	fmt.Printf("Body: %s\n", body)
	fmt.Println("Email sent successfully")

	return nil
}

// SendPasswordResetEmail sends a password reset email
func (es *EmailService) SendPasswordResetEmail(email, resetToken string) error {
	// Duplicated validation
	if email == "" {
		return fmt.Errorf("email is required")
	}
	if !containsAt(email) {
		return fmt.Errorf("invalid email format")
	}

	// Duplicated formatting
	subject := "Password Reset"
	body := fmt.Sprintf("Use this token to reset your password: %s", resetToken)

	// Duplicated sending logic
	fmt.Printf("Sending email to: %s\n", email)
	fmt.Printf("Subject: %s\n", subject)
	fmt.Printf("Body: %s\n", body)
	fmt.Println("Email sent successfully")

	return nil
}

// SendOrderConfirmationEmail sends an order confirmation email
func (es *EmailService) SendOrderConfirmationEmail(email, orderID string, amount float64) error {
	// Duplicated validation
	if email == "" {
		return fmt.Errorf("email is required")
	}
	if !containsAt(email) {
		return fmt.Errorf("invalid email format")
	}

	// Duplicated formatting
	subject := "Order Confirmation"
	body := fmt.Sprintf("Your order %s for $%.2f has been confirmed", orderID, amount)

	// Duplicated sending logic
	fmt.Printf("Sending email to: %s\n", email)
	fmt.Printf("Subject: %s\n", subject)
	fmt.Printf("Body: %s\n", body)
	fmt.Println("Email sent successfully")

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

// Exercise: Refactor this code to eliminate duplication
// 1. Extract the common validation logic
// 2. Extract the common email sending logic
// 3. Create a generic sendEmail method that the others can use
