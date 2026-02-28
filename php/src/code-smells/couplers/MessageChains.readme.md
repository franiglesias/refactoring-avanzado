# Message Chains

Cadenas de mensajes.

## Definición

La navegación profunda por grafos de objetos acopla a los clientes a la estructura de los intermediarios y conduce a código frágil.

## Ejemplo

```php
<?php

declare(strict_types=1);

namespace CodeSmells\Couplers;

class Level2
{
    public function __construct(private int $value)
    {
    }

    public function getValue(): int
    {
        return $this->value;
    }
}

class Level1
{
    public function __construct(private Level2 $next)
    {
    }

    public function getNext(): Level2
    {
        return $this->next;
    }
}

class Root
{
    public function __construct(private Level1 $next)
    {
    }

    public function getNext(): Level1
    {
        return $this->next;
    }
}

function readDeep(Root $root): int
{
    return $root->getNext()->getNext()->getValue();
}
```

## Ejercicio

Inserta un nuevo `Level` entre `Root` y `Level1`, o reubica `getValue`.

## Problemas que encontrarás

Observa cómo cada cliente que usa `$root->getNext()->getNext()->getValue()` debe cambiar, revelando cómo las cadenas de mensajes vuelven costosas refactorizaciones simples.
