package change_preventers

import "fmt"

// Code smell: Shotgun Surgery [Cirugía de escopeta].
// La misma regla de impuestos está duplicada en muchas clases;
// cambiarla requiere ediciones en múltiples lugares.

// Ejercicio: Cambia el impuesto del 21% al 18.5% con redondeo a 2 decimales.

// Tendrás que buscar cada copia y asegurar un redondeo consistente en todas partes,
// destacando cómo la duplicación convierte un cambio pequeño en muchas ediciones arriesgadas.

type LineItem struct {
	Name  string
	Price float64
	Qty   int
}

type PriceCalculator struct{}

func (p *PriceCalculator) TotalWithTax(items []LineItem) float64 {
	subtotal := 0.0
	for _, item := range items {
		subtotal += item.Price * float64(item.Qty)
	}
	tax := subtotal * 0.21
	return subtotal + tax
}

type InvoiceService struct{}

func (i *InvoiceService) CreateTotal(items []LineItem) float64 {
	base := 0.0
	for _, item := range items {
		base += item.Price * float64(item.Qty)
	}
	vat := base * 0.21
	return base + vat
}

type SalesReport struct{}

func (s *SalesReport) Summarize(items []LineItem) string {
	sum := 0.0
	for _, item := range items {
		sum += item.Price * float64(item.Qty)
	}
	tax := sum * 0.21
	total := sum + tax
	return fmt.Sprintf("total=%.2f", total)
}

type LoyaltyPoints struct{}

func (l *LoyaltyPoints) Points(items []LineItem) int {
	base := 0.0
	for _, item := range items {
		base += item.Price * float64(item.Qty)
	}
	withTax := base + base*0.21
	return int(withTax / 10)
}

func DemoShotgun(items []LineItem) (float64, float64) {
	calc := PriceCalculator{}.TotalWithTax(items)
	inv := InvoiceService{}.CreateTotal(items)
	return calc, inv
}
