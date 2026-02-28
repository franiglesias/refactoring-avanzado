# Primitive Obsession

Obsesión primitiva.

## Definición

Conceptos de dominio se modelan con primitivos, lo que obliga a esparcir reglas de validación, formato, y todo tipo de comportamiento, por todo el código.

## Ejemplo

```php
<?php

declare(strict_types=1);

namespace CodeSmells\Bloaters;

class Order
{
    public function __construct(
        private string $customerName,
        private string $customerEmail,
        private string $address,
        private float $totalAmount,
        private string $currency
    ) {
    }

    public function sendInvoice(): void
    {
        if (!str_contains($this->customerEmail, '@')) {
            throw new \InvalidArgumentException('Email inválido');
        }
        if (empty($this->address)) {
            throw new \InvalidArgumentException('No se ha indicado dirección');
        }
        if ($this->totalAmount <= 0) {
            throw new \InvalidArgumentException('El monto debe ser mayor que cero');
        }
        echo "Factura enviada a {$this->customerName} en {$this->address} por {$this->totalAmount} {$this->currency}\n";
    }
}
```

## Ejercicio

Introduce soporte para diferentes monedas, para enviar la factura por email, y para formatear la dirección en función del país.

## Problemas que encontrarás

Dado que los primitivos no nos permiten garantizar la integridad de sus valores, tendrás que introducir validaciones en muchos lugares, incluso de forma repetida. Algunos datos siempre viajan juntos (Data Clump), por lo que tienes que asegurarte de que permanecen juntos.

Para formatear de forma diferente basándote en algún dato arbitrario tendrás que introducir lógica de decisión.
