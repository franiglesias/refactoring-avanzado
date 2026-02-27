package refactoring

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// OrderItem represents an item in an order
type OrderItem struct {
	SKU         string
	Description string
	UnitPrice   float64
	Quantity    int
	Category    string // "general", "food", "books", or empty
}

// Order represents a customer order
type Order struct {
	ID           string
	CustomerName string
	Items        []OrderItem
}

// ReceiptPrinter prints receipts for orders
// Do not change this struct at the beginning of the exercise; first create the Golden Master.
type ReceiptPrinter struct {
	rng *rand.Rand
}

// NewReceiptPrinter creates a new ReceiptPrinter
func NewReceiptPrinter() *ReceiptPrinter {
	return &ReceiptPrinter{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Print generates a receipt string for the given order
func (rp *ReceiptPrinter) Print(order Order) string {
	now := rp.GetCurrentDate()

	header := fmt.Sprintf("Recibo %s - %s", order.ID, now.Format("01/02/06 15:04:05"))

	subtotal := 0.0
	var lines []string
	for idx, it := range order.Items {
		lineTotal := roundMoney(it.UnitPrice * float64(it.Quantity))
		subtotal = roundMoney(subtotal + lineTotal)
		lines = append(lines, fmt.Sprintf("%d. %s (%s) x%d = $%.2f",
			idx+1, it.Description, it.SKU, it.Quantity, lineTotal))
	}

	luckyDiscountPct := rp.discount()
	luckyDiscount := roundMoney(subtotal * luckyDiscountPct)

	taxableGeneral := 0.0
	foodTax := 0.0
	for _, i := range order.Items {
		if i.Category != "books" && i.Category != "food" {
			taxableGeneral += i.UnitPrice * float64(i.Quantity)
		}
		if i.Category == "food" {
			foodTax += i.UnitPrice * float64(i.Quantity) * 0.03
		}
	}
	generalTax := taxableGeneral * 0.07
	taxes := roundMoney(generalTax + foodTax)

	total := roundMoney(subtotal - luckyDiscount + taxes)

	discountLine := ""
	if luckyDiscount > 0 {
		discountLine = fmt.Sprintf("Descuento de la suerte: -$%.2f (%.2f%%)",
			luckyDiscount, luckyDiscountPct*100)
	} else {
		discountLine = "Descuento de la suerte: $0.00 (0.00%)"
	}

	summary := []string{
		fmt.Sprintf("Subtotal: $%.2f", subtotal),
		discountLine,
		fmt.Sprintf("Impuestos: $%.2f", taxes),
		fmt.Sprintf("TOTAL: $%.2f", total),
	}

	var parts []string
	parts = append(parts, header)
	parts = append(parts, lines...)
	parts = append(parts, "---")
	parts = append(parts, summary...)

	return strings.Join(parts, "\n")
}

func (rp *ReceiptPrinter) discount() float64 {
	luckyDiscountPct := 0.0
	if rp.rng.Float64() < 0.1 {
		luckyDiscountPct = rp.rng.Float64() * 0.05
	}
	return luckyDiscountPct
}

// GetCurrentDate returns the current date/time
func (rp *ReceiptPrinter) GetCurrentDate() time.Time {
	return time.Now()
}

func roundMoney(n float64) float64 {
	return float64(int(n*100+0.5)) / 100
}

// Products catalog
var Products = []map[string]interface{}{
	{"sku": "BK-001", "description": "Libro: Clean Code", "unit_price": 30.0, "category": "books"},
	{"sku": "FD-010", "description": "Café en grano 1kg", "unit_price": 12.5, "category": "food"},
	{"sku": "GN-777", "description": "Cuaderno A5", "unit_price": 5.2, "category": "general"},
	{"sku": "GN-123", "description": "Bolígrafos (pack 10)", "unit_price": 3.9, "category": "general"},
	{"sku": "FD-222", "description": "Té verde 200g", "unit_price": 6.75, "category": "food"},
}

// Customers list
var Customers = []string{"Ana", "Luis", "Mar", "Iván", "Sofía"}

// GenerateOrder creates a test order
func GenerateOrder(id, customerName string, numItems, quantity int) Order {
	var items []OrderItem
	for i := 0; i < numItems && i < len(Products); i++ {
		p := Products[i]
		item := OrderItem{
			SKU:         p["sku"].(string),
			Description: p["description"].(string),
			UnitPrice:   p["unit_price"].(float64),
			Quantity:    quantity,
			Category:    p["category"].(string),
		}
		items = append(items, item)
	}

	return Order{
		ID:           id,
		CustomerName: customerName,
		Items:        items,
	}
}
