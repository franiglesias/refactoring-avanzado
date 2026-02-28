# Switch Statements

Sentencias switch.

## Definición

El uso excesivo de `switch` o múltiples `if/else` basados en un código de tipo suele ser una señal de que falta polimorfismo. El problema principal es que cada vez que se añade una nueva variante (un nuevo tipo), hay que buscar y modificar todos los bloques `switch` dispersos por la aplicación.

## Ejemplo

La función `calculatePay` utiliza un `switch` para decidir cómo calcular el salario según el tipo de empleado.

```php
<?php

declare(strict_types=1);

namespace CodeSmells\OopAbusers;

class EmployeeRecord
{
    /**
     * @param 'engineer'|'manager'|'sales' $kind
     */
    public function __construct(
        public string $kind,
        public float $base,
        public ?float $bonus = null,
        public ?float $commission = null
    ) {
    }
}

function calculatePay(EmployeeRecord $rec): float
{
    return match ($rec->kind) {
        'engineer' => $rec->base,
        'manager' => $rec->base + ($rec->bonus ?? 0),
        'sales' => $rec->base + ($rec->commission ?? 0),
        default => throw new \InvalidArgumentException("Unknown employee kind: {$rec->kind}")
    };
}

/**
 * @return array<float>
 */
function demoSwitchStatements(): array
{
    return [
        calculatePay(new EmployeeRecord('engineer', 1000)),
        calculatePay(new EmployeeRecord('manager', 1000, 200)),
        calculatePay(new EmployeeRecord('sales', 800, null, 500)),
    ];
}
```

## Ejercicio

Añade un nuevo tipo de empleado (`contractor`) con una regla de pago especial (ej. tarifa por horas).

## Problemas que encontrarás

Tendrás que modificar el `switch` (o `match`) y cualquier otro código que dependa del tipo de empleado. A medida que el sistema crece, olvidar actualizar uno de estos puntos genera bugs difíciles de detectar.
