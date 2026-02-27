<?php

namespace CodeSmells\ChangePreventers;

class ProductPrice
{
    public function __construct(
        public float $basePrice
    ) {}

    public function getPrice(): float
    {
        return $this->basePrice * 1.21; // IVA hardcoded
    }
}

class Invoice
{
    public function calculateTotal(float $subtotal): float
    {
        return $subtotal * 1.21; // IVA hardcoded
    }
}

class Receipt
{
    public function calculateTax(float $amount): float
    {
        return $amount * 0.21; // IVA hardcoded
    }
}
