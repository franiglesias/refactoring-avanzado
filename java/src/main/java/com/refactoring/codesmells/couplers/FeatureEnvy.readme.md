# Feature Envy

Envidia de características.

## Definición

Una clase usa la información de otra clase colaboradora para hacer cálculos o tomar decisiones, sugiriendo que la segunda clase debería exponer esos comportamientos. Al depender de la estructura de la colaboradora, la clase cliente queda acoplada.

## Ejemplo

```java
public static class Customer {
    private final String name;
    private final String email;
    private final String address;
    private final String phone;

    public Customer(String name, String email, String address, String phone) {
        this.name = name;
        this.email = email;
        this.address = address;
        this.phone = phone;
    }

    public String getName() {
        return name;
    }

    public String getEmail() {
        return email;
    }

    public String getAddress() {
        return address;
    }

    public String getPhone() {
        return phone;
    }
}

public static class InvoiceService {
    public void sendInvoice(Invoice invoice) {
        // This method knows too much about Customer's internal structure
        // and uses Customer's data more than its own

        // Formatting customer info
        String customerInfo = String.format(
                "Customer: %s%nEmail: %s%nAddress: %s%nPhone: %s",
                invoice.getCustomer().getName(),
                invoice.getCustomer().getEmail(),
                invoice.getCustomer().getAddress(),
                invoice.getCustomer().getPhone()
        );

        // Validating customer data
        if (invoice.getCustomer().getEmail() == null || invoice.getCustomer().getEmail().isEmpty()) {
            throw new IllegalArgumentException("customer email is required");
        }
        if (invoice.getCustomer().getAddress() == null || invoice.getCustomer().getAddress().isEmpty()) {
            throw new IllegalArgumentException("customer address is required");
        }

        System.out.printf("Sending invoice %s for %.2f to:%n%s%n",
                invoice.getInvoiceNumber(), invoice.getAmount(), customerInfo);
    }
}
```

## Ejercicio

Añade validación de número de teléfono y formateo especial para clientes VIP.

## Problemas que encontrarás

Probablemente, seguirás añadiendo lógica en `InvoiceService` que depende de detalles internos de `Customer`, esparciendo reglas en el lugar equivocado y volviendo frágiles los cambios.
