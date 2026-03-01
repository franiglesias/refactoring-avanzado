# Duplicated Code

Código duplicado.

## Definición

El código duplicado ocurre cuando la misma estructura de código o lógica aparece en más de un lugar. Es uno de los code smells más comunes y peligrosos, ya que cualquier cambio en la lógica debe replicarse en todas las copias, aumentando el riesgo de inconsistencias.

## Ejemplo

Tres métodos realizan exactamente la misma lógica de validación y envío de emails, variando solo en el contenido del mensaje.

```go
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
```

## Ejercicio

Añade un nuevo tipo de email (p. ej., notificación de promoción) con las mismas validaciones y lógica de envío.

## Problemas que encontrarás

Tendrás que duplicar nuevamente toda la lógica de validación y envío, lo que demuestra cómo la duplicación multiplica el esfuerzo y el riesgo de error humano.
