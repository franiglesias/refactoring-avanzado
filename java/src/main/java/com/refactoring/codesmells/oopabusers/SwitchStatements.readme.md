# Switch Statements

Sentencias switch.

## Definición

El uso excesivo de `switch` o múltiples `if/else` basados en un código de tipo suele ser una señal de que falta polimorfismo. El problema principal es que cada vez que se añade una nueva variante (un nuevo tipo), hay que buscar y modificar todos los bloques `switch` dispersos por la aplicación.

## Ejemplo

```java
public enum PaymentMethod {
    CREDIT_CARD,
    PAYPAL,
    BANK_TRANSFER,
    CASH
}

public static class PaymentProcessor {

    public void processPayment(PaymentMethod method, double amount) {
        switch (method) {
            case CREDIT_CARD:
                System.out.printf("Processing credit card payment for $%.2f%n", amount);
                System.out.println("Validating card...");
                System.out.println("Contacting payment gateway...");
                System.out.println("Payment successful");
                break;

            case PAYPAL:
                System.out.printf("Processing PayPal payment for $%.2f%n", amount);
                System.out.println("Redirecting to PayPal...");
                System.out.println("Payment successful");
                break;

            case BANK_TRANSFER:
                System.out.printf("Processing bank transfer for $%.2f%n", amount);
                System.out.println("Generating transfer reference...");
                System.out.println("Payment pending confirmation");
                break;

            case CASH:
                System.out.printf("Processing cash payment for $%.2f%n", amount);
                System.out.println("Payment received");
                break;
        }
    }

    public double calculateFee(PaymentMethod method, double amount) {
        switch (method) {
            case CREDIT_CARD:
                return amount * 0.029;
            case PAYPAL:
                return amount * 0.034;
            case BANK_TRANSFER:
                return 2.5;
            case CASH:
                return 0;
            default:
                return 0;
        }
    }
}
```

## Ejercicio

Añade un nuevo método de pago (CRYPTOCURRENCY) con reglas especiales de procesamiento y comisiones.

## Problemas que encontrarás

Tendrás que modificar múltiples `switch` statements en diferentes métodos. A medida que el sistema crece, olvidar actualizar uno de estos puntos genera bugs difíciles de detectar.
