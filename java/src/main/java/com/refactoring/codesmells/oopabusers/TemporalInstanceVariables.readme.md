# Temporal Instance Variables

Variables de instancia temporales.

## Definición

Una clase usa variables de instancia que solo son válidas durante un período específico del flujo de trabajo, en lugar de mantener un estado consistente. Esto genera confusión sobre el estado válido del objeto y puede causar errores difíciles de detectar.

## Ejemplo

```java
public static class PizzaOrder {
    private String size;
    private List<String> toppings = new ArrayList<>();
    private String address;

    public void start(String size) {
        this.size = size;
        this.toppings = new ArrayList<>();
        this.address = null;
    }

    public void addTopping(String topping) {
        if (this.size == null) {
            return;
        }
        this.toppings.add(topping);
    }

    public void setDeliveryAddress(String address) {
        this.address = address;
    }

    public String place() {
        String sizeStr = this.size != null ? this.size : "?";
        String addressStr = this.address != null ? this.address : "UNKNOWN";
        String summary = String.format("Pizza %s to %s with [%s]",
            sizeStr, addressStr, String.join(", ", this.toppings));

        // Reset temporal state
        this.size = null;
        this.address = null;
        this.toppings = new ArrayList<>();

        return summary;
    }
}

public static String demoPizzaOrder() {
    PizzaOrder o = new PizzaOrder();
    o.start("L");
    o.addTopping("pepperoni");
    o.addTopping("mushroom");
    o.setDeliveryAddress("123 Main St");
    return o.place();
}
```

## Ejercicio

Añade validación de estado para prevenir el uso incorrecto (ej. llamar a `place()` sin haber llamado a `start()` primero).

## Problemas que encontrarás

El estado temporal hace que el objeto sea frágil y propenso a errores. Es difícil razonar sobre el estado válido del objeto en diferentes puntos del flujo.
