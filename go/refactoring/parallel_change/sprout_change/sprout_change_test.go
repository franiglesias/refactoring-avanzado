package sprout_change

import "testing"

var cart = []CartItem{
	{ID: "p1", Price: 10, Qty: 2, Category: "general"},
	{ID: "b1", Price: 20, Qty: 1, Category: "books"},
	{ID: "f1", Price: 15, Qty: 4, Category: "food"},
}

func executeSubject(cart []CartItem, region Region) float64 {
	return CalculateTotal(cart, region)
}

func TestShouldCalculateTotalForEU(t *testing.T) {
	result := executeSubject(cart, RegionEU)
	if result != 104 {
		t.Errorf("expected 104, got %v", result)
	}
}

func TestShouldCalculateTotalForUS(t *testing.T) {
	result := executeSubject(cart, RegionUS)
	if result != 107 {
		t.Errorf("expected 107, got %v", result)
	}
}
