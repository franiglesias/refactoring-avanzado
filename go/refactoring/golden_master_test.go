package refactoring

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestReceiptPrinter_WithFixedRandomness tests the receipt printer with controlled randomness
func TestReceiptPrinter_WithFixedRandomness(t *testing.T) {
	// Arrange
	printer := NewReceiptPrinter()
	// Use a fixed seed for reproducible tests
	printer.rng = rand.New(rand.NewSource(42))

	order := GenerateOrder("ORD-001", "Ana", 3, 2)

	// Act
	receipt := printer.Print(order)

	// Assert - verify the receipt contains expected elements
	assert.Contains(t, receipt, "Recibo ORD-001")
	assert.Contains(t, receipt, "Libro: Clean Code")
	assert.Contains(t, receipt, "Café en grano 1kg")
	assert.Contains(t, receipt, "Cuaderno A5")
	assert.Contains(t, receipt, "Subtotal:")
	assert.Contains(t, receipt, "Impuestos:")
	assert.Contains(t, receipt, "TOTAL:")
}

// TestReceiptPrinter_CalculatesCorrectSubtotal tests subtotal calculation
func TestReceiptPrinter_CalculatesCorrectSubtotal(t *testing.T) {
	// Arrange
	printer := NewReceiptPrinter()
	printer.rng = rand.New(rand.NewSource(100)) // Seed that doesn't trigger discount

	order := Order{
		ID:           "TEST-001",
		CustomerName: "Test Customer",
		Items: []OrderItem{
			{SKU: "TEST-1", Description: "Test Item", UnitPrice: 10.0, Quantity: 2, Category: "general"},
		},
	}

	// Act
	receipt := printer.Print(order)

	// Assert
	assert.Contains(t, receipt, "20.00") // 10 * 2
}

// TestRoundMoney tests the money rounding function
func TestRoundMoney(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"rounds up", 10.556, 10.56},
		{"rounds down", 10.554, 10.55},
		{"no rounding needed", 10.55, 10.55},
		{"handles zero", 0.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := roundMoney(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// MockReceiptPrinter is a testable version with injectable date and random
type MockReceiptPrinter struct {
	*ReceiptPrinter
	fixedDate     time.Time
	fixedDiscount float64
}

func (m *MockReceiptPrinter) GetCurrentDate() time.Time {
	return m.fixedDate
}

func (m *MockReceiptPrinter) discount() float64 {
	return m.fixedDiscount
}

// TestReceiptPrinter_WithMockedDependencies shows how to test with controlled dependencies
func TestReceiptPrinter_WithMockedDependencies(t *testing.T) {
	// Arrange
	mockPrinter := &MockReceiptPrinter{
		ReceiptPrinter: NewReceiptPrinter(),
		fixedDate:      time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
		fixedDiscount:  0.0,
	}

	order := GenerateOrder("ORD-123", "Luis", 2, 1)

	// Act
	receipt := mockPrinter.Print(order)

	// Assert
	assert.Contains(t, receipt, "Recibo ORD-123")
	assert.Contains(t, receipt, "01/01/24")
}
