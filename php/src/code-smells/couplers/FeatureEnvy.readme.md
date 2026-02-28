# Feature Envy

Envidia de características.

## Definición

Una clase usa la información de otra clase colaboradora para hacer cálculos o tomar decisiones, sugiriendo que la segunda clase debería exponer esos comportamientos. Al depender de la estructura de la colaboradora, la clase cliente queda acoplada.

## Ejemplo

`ShippingCalculator` se mete en los datos de `Customer` para tomar decisiones, lo que indica que el comportamiento quizá debería pertenecer a `Customer`.

```php
<?php

declare(strict_types=1);

namespace CodeSmells\Couplers;

class Customer
{
    public function __construct(
        public string $name,
        public string $street,
        public string $city,
        public string $zip
    ) {
    }
}

class ShippingCalculator
{
    public function cost(Customer $customer): int
    {
        $base = str_starts_with($customer->zip, '9') ? 10 : 20;
        $distant = strlen($customer->city) > 6 ? 5 : 0;
        return $base + $distant;
    }
}

function demoFeatureEnvy(Customer $c): int
{
    return (new ShippingCalculator())->cost($c);
}
```

## Ejercicio

Añade envío gratis para clientes en ciertas ciudades y un recargo de fin de semana.

## Problemas que encontrarás

Probablemente, seguirás añadiendo condiciones dentro de `ShippingCalculator` que dependen de detalles internos de `Customer`, esparciendo reglas en el lugar equivocado y volviendo frágiles los cambios.
