package calisthenics_exercises

// Regla 5: Colecciones de primera clase
// Los tipos colectivos (arrays, maps, slices, etc.) son una forma de representar
// colecciones de objetos, pero al igual que ocurre con los tipos primitivos,
// no nos protegen de problemas de validez o integridad de las invariantes del dominio.

// Ejercicio: Envuelve el slice de productos en un tipo que represente
// un concepto del dominio y proteja sus invariantes.

type Product struct {
	ID    string
	Price float64
}

func AddProduct(products []Product, product Product) []Product {
	for _, p := range products {
		if p.ID == product.ID {
			return products
		}
	}
	return append(products, product)
}

func TotalPrice(products []Product) float64 {
	total := 0.0
	for _, p := range products {
		total += p.Price
	}
	return total
}

func RemoveProduct(products []Product, productID string) []Product {
	result := []Product{}
	for _, p := range products {
		if p.ID != productID {
			result = append(result, p)
		}
	}
	return result
}
