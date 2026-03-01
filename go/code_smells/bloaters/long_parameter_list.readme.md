# Long Parameter List

Lista larga de parámetros.

## Definición

Una función recibe más de tres o cuatro parámetros.

## Ejemplo

```go
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
```

## Ejercicio

Añade dos opciones más (p. ej., locale y pageSize) al usuario.

## Problemas que encontrarás

Con más de tres parámetros es difícil recordar con exactitud cuáles son, el orden o el tipo de cada uno. Añadir parámetros no hace más que aumentar la dificultad de uso y mantenimiento.
