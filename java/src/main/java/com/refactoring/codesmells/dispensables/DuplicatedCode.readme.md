# Duplicated Code

Código duplicado.

## Definición

El mismo código (o muy similar) aparece en múltiples lugares. Los cambios deben replicarse en todos los lugares, aumentando el riesgo de inconsistencias y errores.

## Ejemplo

```java
public static class EmailService {

    public void sendWelcomeEmail(String email, String name) {
        // Duplicated validation
        if (email == null || email.isEmpty()) {
            throw new IllegalArgumentException("email is required");
        }
        if (!email.contains("@")) {
            throw new IllegalArgumentException("invalid email format");
        }

        // Duplicated formatting
        String subject = "Welcome!";
        String body = String.format("Hello %s, welcome to our service!", name);

        // Duplicated sending logic
        System.out.printf("Sending email to: %s%n", email);
        System.out.printf("Subject: %s%n", subject);
        System.out.printf("Body: %s%n", body);
        System.out.println("Email sent successfully");
    }

    public void sendPasswordResetEmail(String email, String resetToken) {
        // Duplicated validation
        if (email == null || email.isEmpty()) {
            throw new IllegalArgumentException("email is required");
        }
        if (!email.contains("@")) {
            throw new IllegalArgumentException("invalid email format");
        }

        // Duplicated formatting
        String subject = "Password Reset";
        String body = String.format("Use this token to reset your password: %s", resetToken);

        // Duplicated sending logic
        System.out.printf("Sending email to: %s%n", email);
        System.out.printf("Subject: %s%n", subject);
        System.out.printf("Body: %s%n", body);
        System.out.println("Email sent successfully");
    }

    public void sendOrderConfirmationEmail(String email, String orderId, double amount) {
        // ... more duplicated code
    }
}
```

## Ejercicio

Añade validación adicional de dominio de email y soporte para CC/BCC en todos los métodos de envío.

## Problemas que encontrarás

Tendrás que replicar los mismos cambios en múltiples lugares, con alto riesgo de olvidar actualizar alguno y generar comportamiento inconsistente.
