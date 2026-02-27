<?php

namespace CodeSmells\Dispensables;

class Product
{
    public function __construct(
        public string $name,
        public float $price,
        public int $stock
    ) {}
}

class ProductService
{
    public function applyDiscount(Product $product, float $discount): float
    {
        return $product->price * (1 - $discount);
    }

    public function isInStock(Product $product): bool
    {
        return $product->stock > 0;
    }

    public function sell(Product $product, int $quantity): void
    {
        if ($product->stock >= $quantity) {
            $product->stock -= $quantity;
        }
    }
}
