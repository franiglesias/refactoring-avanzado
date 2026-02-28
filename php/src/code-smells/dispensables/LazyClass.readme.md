# Lazy Class

Clase perezosa.

## Definición

Una clase perezosa es aquella que no aporta suficiente valor para justificar su existencia. Suelen ser clases que solo envuelven una operación trivial o que tienen muy poca responsabilidad, añadiendo una complejidad innecesaria al sistema.

## Ejemplo

La clase `ShippingLabelBuilder` solo tiene un método que realiza una concatenación de strings simple, algo que podría resolverse con una función pura.

```php
<?php

declare(strict_types=1);

namespace CodeSmells\Dispensables;

class Address
{
    public function __construct(
        public string $name,
        public string $line1,
        public ?string $city = null
    ) {
    }
}

class ShippingLabelBuilder
{
    public function build(Address $a): string
    {
        return "{$a->name} — {$a->line1}" . ($a->city ? ", {$a->city}" : '');
    }
}

function printShippingLabel(): void
{
    $address = new Address('John Doe', '123 Main St', 'New York');

    $labelBuilder = new ShippingLabelBuilder();
    $label = $labelBuilder->build($address);
    echo $label . "\n";
}
```

## Ejercicio

Reescribe el código para eliminar la necesidad de la clase `ShippingLabelBuilder`.

## Problemas que encontrarás

Mantener una estructura de clase para una lógica tan simple te obliga a instanciar objetos innecesariamente y añade capas de abstracción que dificultan la legibilidad del código sin ofrecer beneficios a cambio.
