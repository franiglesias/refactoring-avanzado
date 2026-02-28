# Data clump

Grupo de datos.

## Definición

El mismo grupo de campos de datos viaja junto por muchos lugares, lo que sugiere un Value Object faltante y duplicación.

## Ejemplo

```php
<?php

declare(strict_types=1);

namespace CodeSmells\Bloaters;

class Invoice
{
    private readonly string $customerName;
    private readonly string $customerCity;
    private readonly string $customerStreet;
    private readonly string $customerZip;

    public function __construct(
        string $customerName,
        string $customerStreet,
        string $customerCity,
        string $customerZip
    ) {
        $this->customerName = $customerName;
        $this->customerStreet = $customerStreet;
        $this->customerCity = $customerCity;
        $this->customerZip = $customerZip;
    }

    public function print(): string
    {
        return "Factura para: {$this->customerName}\n" .
               "Dirección: {$this->customerStreet}, {$this->customerCity}, {$this->customerZip}";
    }
}

class ShippingLabel
{
    private readonly string $customerName;
    private readonly string $customerStreet;
    private readonly string $customerCity;
    private readonly string $customerZip;

    public function __construct(
        string $customerName,
        string $customerStreet,
        string $customerCity,
        string $customerZip
    ) {
        $this->customerName = $customerName;
        $this->customerStreet = $customerStreet;
        $this->customerCity = $customerCity;
        $this->customerZip = $customerZip;
    }

    public function print(): string
    {
        return "Enviar a: {$this->customerName}\n" .
               "{$this->customerStreet}, {$this->customerCity}, {$this->customerZip}";
    }
}
```

## Ejercicio

Añade país y provincia y reglas de formateo internacional de la dirección.

## Problemas que encontrarás

Necesitarás modificar constructores, impresores y cualquier lugar que pase estos campos juntos, multiplicando la superficie de cambio.
