# Lazy Class

Clase perezosa.

## Definición

Una clase perezosa es aquella que no aporta suficiente valor para justificar su existencia. Suelen ser clases que solo envuelven una operación trivial o que tienen muy poca responsabilidad, añadiendo una complejidad innecesaria al sistema.

## Ejemplo

```java
public static class Address {
    String name;
    String line1;
    String city;

    public Address(String name, String line1, String city) {
        this.name = name;
        this.line1 = line1;
        this.city = city;
    }
}

public static class ShippingLabelBuilder {
    public String build(Address a) {
        String cityPart = a.city != null ? ", " + a.city : "";
        return String.format("%s — %s%s", a.name, a.line1, cityPart);
    }
}

public static void printShippingLabel() {
    Address address = new Address("John Doe", "123 Main St", "New York");
    ShippingLabelBuilder labelBuilder = new ShippingLabelBuilder();
    String label = labelBuilder.build(address);
    System.out.println(label);
}
```

## Ejercicio

Reescribe el código para eliminar la necesidad de la clase `ShippingLabelBuilder`.

## Problemas que encontrarás

Mantener una estructura de clase para una lógica tan simple te obliga a instanciar objetos innecesariamente y añade capas de abstracción que dificultan la legibilidad del código sin ofrecer beneficios a cambio.
