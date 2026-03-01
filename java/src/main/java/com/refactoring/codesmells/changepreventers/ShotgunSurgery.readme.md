# Shotgun Surgery

Cirugía de escopeta.

## Definición

Un cambio requiere realizar muchas modificaciones pequeñas en múltiples clases. La lógica relacionada está esparcida por todo el código en lugar de estar centralizada.

## Ejemplo

```java
public static class PriceCalculator {
    public double totalWithTax(List<LineItem> items) {
        double subtotal = items.stream()
            .mapToDouble(i -> i.price * i.qty)
            .sum();
        double tax = subtotal * 0.21;
        return subtotal + tax;
    }
}

public static class InvoiceService {
    public double createTotal(List<LineItem> items) {
        double base = items.stream()
            .mapToDouble(i -> i.price * i.qty)
            .sum();
        double vat = base * 0.21;
        return base + vat;
    }
}

public static class SalesReport {
    public String summarize(List<LineItem> items) {
        double sum = items.stream()
            .mapToDouble(i -> i.price * i.qty)
            .sum();
        double tax = sum * 0.21;
        double total = sum + tax;
        return String.format("total=%.2f", total);
    }
}

public static class LoyaltyPoints {
    public int points(List<LineItem> items) {
        double base = items.stream()
            .mapToDouble(i -> i.price * i.qty)
            .sum();
        double withTax = base + base * 0.21;
        return (int) Math.floor(withTax / 10);
    }
}
```

## Ejercicio

Cambia la tasa de impuestos de 0.21 a 0.19 y añade una lógica de descuento del 5%.

## Problemas que encontrarás

Tendrás que modificar múltiples clases, corriendo el riesgo de olvidar actualizar alguna de ellas y generar inconsistencias en el sistema.
