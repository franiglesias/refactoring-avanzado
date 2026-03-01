# Long Parameter List

Lista de parámetros larga.

## Definición

Un método tiene una lista muy larga de parámetros, lo que dificulta su comprensión y uso. Suele indicar que falta una abstracción o agrupación de datos relacionados.

## Ejemplo

```java
public static void createUser(
        String username,
        String email,
        String firstName,
        String lastName,
        int age,
        String country,
        String city,
        String address,
        String postalCode,
        String phoneNumber,
        boolean isActive,
        String role
) {
    // Validation
    if (username == null || username.isEmpty() || email == null || email.isEmpty()) {
        throw new IllegalArgumentException("username and email are required");
    }

    // Business logic
    System.out.printf("Creating user: %s (%s)%n", username, email);
    System.out.printf("Name: %s %s%n", firstName, lastName);
    System.out.printf("Location: %s, %s, %s, %s%n", address, city, postalCode, country);
    System.out.printf("Contact: %s%n", phoneNumber);
    System.out.printf("Age: %d, Role: %s, Active: %b%n", age, role, isActive);
}
```

## Ejercicio

Añade más campos relacionados con el usuario (fecha de nacimiento, número de identificación fiscal, preferencias de idioma).

## Problemas que encontrarás

Cada nuevo parámetro alarga aún más la firma del método, haciendo el código cada vez más difícil de leer y mantener. Los cambios en el orden de los parámetros pueden generar bugs sutiles.
