package sprout_change

import "math"

type CartItem struct {
	ID       string
	Price    float64
	Qty      int
	Category string
}

type Region string

const (
	RegionUS Region = "US"
	RegionEU Region = "EU"
)

type TaxPolicy interface {
}

// CalculateTotal: regla existente: un único impuesto plano por región; los libros y la comida están exentos en la UE
func CalculateTotal(cart []CartItem, region Region) float64 {
	subtotal := 0.0
	for _, it := range cart {
		subtotal += it.Price * float64(it.Qty)
	}

	tax := 0.0
	if region == RegionUS {
		tax = subtotal * 0.07 // 7% plano
	} else if region == RegionEU {
		// exenciones ingenuas en línea
		taxable := 0.0
		for _, it := range cart {
			if it.Category != "books" && it.Category != "food" {
				taxable += it.Price * float64(it.Qty)
			}
		}
		tax = taxable * 0.2 // 20% plano solo sobre los ítems gravables
	}

	return RoundCurrency(subtotal + tax)
}

func RoundCurrency(amount float64) float64 {
	return math.Round(amount*100) / 100
}

// DemoSprout: uso de ejemplo, mantenido simple para estudiantes
func DemoSprout() float64 {
	cart := []CartItem{
		{ID: "p1", Price: 10, Qty: 2, Category: "general"},
		{ID: "b1", Price: 20, Qty: 1, Category: "books"},
	}
	return CalculateTotal(cart, RegionEU)
}
