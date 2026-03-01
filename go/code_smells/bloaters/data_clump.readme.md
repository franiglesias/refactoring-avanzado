# Data clump

Grupo de datos.

## Definición

El mismo grupo de campos de datos viaja junto por muchos lugares, lo que sugiere un Value Object faltante y duplicación.

## Ejemplo

```go
package bloaters

import "fmt"

// ProductService demonstrates data clump code smell
// The same group of data (street, city, postalCode, country) appears together repeatedly
type ProductService struct{}

// ShipProduct ships a product to an address
func (ps *ProductService) ShipProduct(
	productID string,
	street string,
	city string,
	postalCode string,
	country string,
) error {
	fmt.Printf("Shipping product %s to:\n", productID)
	fmt.Printf("%s\n%s, %s\n%s\n", street, city, postalCode, country)
	return nil
}

// ValidateDeliveryAddress validates a delivery address
func (ps *ProductService) ValidateDeliveryAddress(
	street string,
	city string,
	postalCode string,
	country string,
) bool {
	return street != "" && city != "" && postalCode != "" && country != ""
}

// CalculateShippingCost calculates shipping cost based on address
func (ps *ProductService) CalculateShippingCost(
	street string,
	city string,
	postalCode string,
	country string,
	weight float64,
) float64 {
	baseCost := 10.0
	// International shipping costs more
	if country != "Spain" {
		baseCost += 15.0
	}
	// City surcharge
	if city == "Madrid" || city == "Barcelona" {
		baseCost += 2.0
	}
	return baseCost + (weight * 0.5)
}
```

## Ejercicio

Añade país y provincia y reglas de formateo internacional de la dirección.

## Problemas que encontrarás

Necesitarás modificar constructores, impresores y cualquier lugar que pase estos campos juntos, multiplicando la superficie de cambio.
