# Data Clump

Grupo de datos.

## Definición

El mismo grupo de campos de datos viaja junto por muchos lugares, lo que sugiere un Value Object faltante y duplicación.

## Ejemplo

```java
public static class ProductService {

    public void shipProduct(
            String productId,
            String street,
            String city,
            String postalCode,
            String country
    ) {
        System.out.printf("Shipping product %s to:%n", productId);
        System.out.printf("%s%n%s, %s%n%s%n", street, city, postalCode, country);
    }

    public boolean validateDeliveryAddress(
            String street,
            String city,
            String postalCode,
            String country
    ) {
        return street != null && !street.isEmpty() &&
                city != null && !city.isEmpty() &&
                postalCode != null && !postalCode.isEmpty() &&
                country != null && !country.isEmpty();
    }

    public double calculateShippingCost(
            String street,
            String city,
            String postalCode,
            String country,
            double weight
    ) {
        double baseCost = 10.0;
        if (!"Spain".equals(country)) {
            baseCost += 15.0;
        }
        if ("Madrid".equals(city) || "Barcelona".equals(city)) {
            baseCost += 2.0;
        }
        return baseCost + (weight * 0.5);
    }
}
```

## Ejercicio

Añade país y provincia y reglas de formateo internacional de la dirección.

## Problemas que encontrarás

Necesitarás modificar constructores, validadores y cualquier lugar que pase estos campos juntos, multiplicando la superficie de cambio.
