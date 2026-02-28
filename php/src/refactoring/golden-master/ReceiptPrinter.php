<?php

namespace Refactoring\GoldenMaster;

class OrderItem
{
    public function __construct(
        public string $sku,
        public string $description,
        public float $unitPrice,
        public int $quantity,
        public ?string $category = null
    ) {}
}

class Order
{
    public function __construct(
        public string $id,
        public string $customerName,
        public array $items
    ) {}
}

class ReceiptPrinter
{
    // Do not change this function at the beginning of the exercise; first create the Golden Master.
    public function print(Order $order): string
    {
        $now = $this->getCurrentDate();

        $header = sprintf("Recibo %s - %s", $order->id, $now->format('m/d/y H:i:s'));

        $subtotal = 0;
        $lines = [];
        $index = 1;

        foreach ($order->items as $item) {
            $lineTotal = $this->roundMoney($item->unitPrice * $item->quantity);
            $subtotal = $this->roundMoney($subtotal + $lineTotal);
            $lines[] = sprintf(
                "%d. %s (%s) x%d = $%.2f",
                $index++,
                $item->description,
                $item->sku,
                $item->quantity,
                $lineTotal
            );
        }

        $luckyDiscountPct = $this->discount();
        $luckyDiscount = $this->roundMoney($subtotal * $luckyDiscountPct);

        $taxableGeneral = 0.0;
        $foodTax = 0.0;

        foreach ($order->items as $item) {
            if ($item->category !== 'books' && $item->category !== 'food') {
                $taxableGeneral += $item->unitPrice * $item->quantity;
            }
            if ($item->category === 'food') {
                $foodTax += $item->unitPrice * $item->quantity * 0.03;
            }
        }

        $generalTax = $taxableGeneral * 0.07;
        $taxes = $this->roundMoney($generalTax + $foodTax);

        $total = $this->roundMoney($subtotal - $luckyDiscount + $taxes);

        $discountLine = $luckyDiscount > 0
            ? sprintf("Descuento de la suerte: -$%.2f (%.2f%%)", $luckyDiscount, $luckyDiscountPct * 100)
            : "Descuento de la suerte: $0.00 (0.00%)";

        $summary = [
            sprintf("Subtotal: $%.2f", $subtotal),
            $discountLine,
            sprintf("Impuestos: $%.2f", $taxes),
            sprintf("TOTAL: $%.2f", $total)
        ];

        $parts = array_merge([$header], $lines, ['---'], $summary);

        return implode("\n", $parts);
    }

    protected function discount(): float
    {
        $luckyDiscountPct = 0.0;
        if (mt_rand() / mt_getrandmax() < 0.1) {
            $luckyDiscountPct = (mt_rand() / mt_getrandmax()) * 0.05;
        }
        return $luckyDiscountPct;
    }

    protected function getCurrentDate(): \DateTime
    {
        return new \DateTime();
    }

    private function roundMoney(float $n): float
    {
        return round($n * 100) / 100;
    }
}

function generateOrder(string $id, string $customerName, int $numItems, int $quantity): Order
{
    $products = [
        ['sku' => 'BK-001', 'description' => 'Libro: Clean Code', 'unitPrice' => 30, 'category' => 'books'],
        ['sku' => 'FD-010', 'description' => 'Café en grano 1kg', 'unitPrice' => 12.5, 'category' => 'food'],
        ['sku' => 'GN-777', 'description' => 'Cuaderno A5', 'unitPrice' => 5.2, 'category' => 'general'],
        ['sku' => 'GN-123', 'description' => 'Bolígrafos (pack 10)', 'unitPrice' => 3.9, 'category' => 'general'],
        ['sku' => 'FD-222', 'description' => 'Té verde 200g', 'unitPrice' => 6.75, 'category' => 'food'],
    ];

    $items = [];
    for ($i = 0; $i < $numItems && $i < count($products); $i++) {
        $p = $products[$i];
        $items[] = new OrderItem(
            $p['sku'],
            $p['description'],
            $p['unitPrice'],
            $quantity,
            $p['category']
        );
    }

    return new Order($id, $customerName, $items);
}
