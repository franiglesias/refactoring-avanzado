# Shotgun Surgery

## Definición

Cuando necesito hacer un cambio tengo que hacerlo en muchos lugares del código que pueden estar alejados entre sí, incluso en distintos módulos.

## Ejemplo

La misma regla de impuestos está duplicada en muchas clases; cambiarla requiere ediciones en múltiples lugares.

```php
<?php

declare(strict_types=1);

namespace CodeSmells\ChangePreventers;

class LineItem
{
    public function __construct(
        public string $name,
        public float $price,
        public int $qty
    ) {}
}

class PriceCalculator
{
    /**
     * @param array<LineItem> $items
     */
    public function totalWithTax(array $items): float
    {
        $subtotal = array_reduce(
            $items,
            fn($s, $i) => $s + $i->price * $i->qty,
            0
        );
        $tax = $subtotal * 0.21;
        return $subtotal + $tax;
    }
}

class InvoiceService
{
    /**
     * @param array<LineItem> $items
     */
    public function createTotal(array $items): float
    {
        $base = array_reduce(
            $items,
            fn($s, $i) => $s + $i->price * $i->qty,
            0
        );
        $vat = $base * 0.21;
        return $base + $vat;
    }
}

class SalesReport
{
    /**
     * @param array<LineItem> $items
     */
    public function summarize(array $items): string
    {
        $sum = array_reduce(
            $items,
            fn($s, $i) => $s + $i->price * $i->qty,
            0
        );
        $tax = $sum * 0.21;
        $total = $sum + $tax;
        return "total=" . number_format($total, 2);
    }
}

class LoyaltyPoints
{
    /**
     * @param array<LineItem> $items
     */
    public function points(array $items): int
    {
        $base = array_reduce(
            $items,
            fn($s, $i) => $s + $i->price * $i->qty,
            0
        );
        $withTax = $base + $base * 0.21;
        return (int)floor($withTax / 10);
    }
}
```

## Ejercicio

Cambia el impuesto del 21% al 18.5% con redondeo a 2 decimales.

## Problemas que encontrarás

Tendrás que buscar cada copia y asegurar un redondeo consistente en todas partes, destacando cómo la duplicación convierte un cambio pequeño en muchas ediciones arriesgadas.
