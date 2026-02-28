package refactoring

import (
	"testing"
	"time"

	"github.com/bradleyjkemp/cupaloy/v2"
)

var counter = 0

// ReceiptPrinterWithoutDiscountForTest is a test version with controlled date and no discount
type ReceiptPrinterWithoutDiscountForTest struct {
	*ReceiptPrinter
}

func (r *ReceiptPrinterWithoutDiscountForTest) GetCurrentDate() time.Time {
	return time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC)
}

func (r *ReceiptPrinterWithoutDiscountForTest) discount() float64 {
	return 0.0
}

// ReceiptPrinterWithDiscountForTest is a test version with controlled date and fixed discount
type ReceiptPrinterWithDiscountForTest struct {
	*ReceiptPrinter
}

func (r *ReceiptPrinterWithDiscountForTest) GetCurrentDate() time.Time {
	return time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC)
}

func (r *ReceiptPrinterWithDiscountForTest) discount() float64 {
	return 0.05
}

func TestShouldPrintAReceipt(t *testing.T) {
	// Given a customer
	customer := "Ana"
	// Given a number of items
	item := 1
	// Given quantity
	quantity := 1

	counter++
	order := GenerateOrder("ORD-1", customer, item, quantity)
	printer := &ReceiptPrinterWithoutDiscountForTest{
		ReceiptPrinter: NewReceiptPrinter(),
	}
	receipt := printer.Print(order)

	cupaloy.SnapshotT(t, receipt)
}
