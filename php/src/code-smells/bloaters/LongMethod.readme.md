# Long method

Método largo.

## Definición

Un método en una clase es muy largo.

## Ejemplo

```php
<?php

declare(strict_types=1);

namespace RefactoringAvanzado\CodeSmells\Bloaters;

class OrderService
{
    public function process(Order $order): void
    {
        // Validar el pedido
        if (!$order->items || count($order->items) === 0) {
            echo "El pedido no tiene productos\n";
            return;
        }

        // Validar precios y cantidades
        foreach ($order->items as $item) {
            if ($item['price'] < 0 || $item['quantity'] <= 0) {
                echo "Producto inválido en el pedido\n";
                return;
            }
        }

        // Constantes de negocio (simples por ahora)
        $TAX_RATE = 0.21; // 21% IVA
        $FREE_SHIPPING_THRESHOLD = 50;
        $SHIPPING_FLAT = 5;

        // Calcular subtotal
        $subtotal = 0;
        foreach ($order->items as $item) {
            $subtotal += $item['price'] * $item['quantity'];
        }

        // Descuento por cliente VIP (10% del subtotal)
        $discount = 0;
        if ($order->customerType === 'VIP') {
            $discount = $this->roundMoney($subtotal * 0.1);
            echo "Descuento VIP aplicado\n";
        }

        // Base imponible
        $taxable = max(0, $subtotal - $discount);

        // Impuestos
        $tax = $this->roundMoney($taxable * $TAX_RATE);

        // Envío
        $shipping = $taxable >= $FREE_SHIPPING_THRESHOLD ? 0 : $SHIPPING_FLAT;

        // Total
        $total = $this->roundMoney($taxable + $tax + $shipping);

        // ... (continúa con muchas más líneas de código para DB, email, impresión)
    }
}
```

El ejemplo completo tiene más de 400 líneas dentro de un solo método `process()`.

## Ejercicio

Añade soporte de cupones con expiración y multi-moneda (USD/EUR) con reglas de redondeo distintas.

## Problemas que encontrarás

Tienes que tocar diferentes secciones dentro del método, lo que genera riesgo de cambios indeseados y aumenta el esfuerzo de mantenimiento.
