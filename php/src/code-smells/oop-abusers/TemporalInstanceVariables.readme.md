# Temporal Instance Variables

Variables de instancia temporales.

## Definición

Este smell ocurre cuando un objeto tiene campos que solo están llenos (o tienen sentido) en ciertas etapas de su ciclo de vida. Esto suele indicar un acoplamiento temporal, donde los métodos deben llamarse en un orden específico para que el objeto sea válido, dejando al objeto en un estado inconsistente fuera de esa secuencia.

## Ejemplo

`PizzaOrder` utiliza variables de instancia que solo son válidas entre la llamada a `start()` y `place()`.

```php
<?php

declare(strict_types=1);

namespace CodeSmells\OopAbusers;

class PizzaOrder
{
    /** @var 'S'|'M'|'L'|null */
    private ?string $size = null;
    /** @var array<string> */
    private array $toppings = [];
    private ?string $address = null;

    /**
     * @param 'S'|'M'|'L' $size
     */
    public function start(string $size): void
    {
        $this->size = $size;
        $this->toppings = [];
    }

    public function addTopping(string $topping): void
    {
        if (!$this->size) {
            return;
        }
        $this->toppings[] = $topping;
    }

    public function setDeliveryAddress(string $address): void
    {
        $this->address = $address;
    }

    public function place(): string
    {
        $summary = "Pizza {$this->size} to {$this->address} with [" . implode(', ', $this->toppings) . "]";
        $this->size = null;
        $this->address = null;
        $this->toppings = [];
        return $summary;
    }
}

function demoPizzaOrder(): string
{
    $o = new PizzaOrder();
    $o->start('L');
    $o->addTopping('pepperoni');
    $o->addTopping('mushroom');
    $o->setDeliveryAddress('123 Main St');
    return $o->place();
}
```

## Ejercicio

Añade una validación para que no se pueda llamar a `place()` si no se ha añadido al menos un ingrediente.

## Problemas que encontrarás

Te darás cuenta de que el objeto es una "máquina de estados" frágil. Si un cliente olvida llamar a `start()` o intenta llamar a `addTopping()` fuera de orden, el sistema puede fallar silenciosamente o requerir comprobaciones constantes de nulidad en cada método.
