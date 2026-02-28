# Middleman

Intermediario.

## Definición

Ocurre cuando una clase realiza una única acción: delegar el trabajo a otra clase. Si una clase existe solo como un "pasamanos" hacia otro objeto, es posible que estemos ante una capa de abstracción innecesaria que oscurece al colaborador real.

## Ejemplo

`Shop` hace poco más que delegar a `Catalog`, añadiendo una capa innecesaria que oculta dónde ocurre realmente la lógica.

```php
<?php

declare(strict_types=1);

namespace CodeSmells\Couplers;

class Catalog
{
    /** @var array<string, string> */
    private array $items = [];

    public function add(string $id, string $name): void
    {
        $this->items[$id] = $name;
    }

    public function find(string $id): ?string
    {
        return $this->items[$id] ?? null;
    }

    /**
     * @return array<string>
     */
    public function list(): array
    {
        return array_values($this->items);
    }
}

class Shop
{
    public function __construct(private Catalog $catalog)
    {
    }

    public function add(string $id, string $name): void
    {
        $this->catalog->add($id, $name);
    }

    public function find(string $id): ?string
    {
        return $this->catalog->find($id);
    }

    /**
     * @return array<string>
     */
    public function list(): array
    {
        return $this->catalog->list();
    }
}
```

## Ejercicio

Añade una funcionalidad `searchByPrefix` en `Catalog` y propágala a través de `Shop`.

## Problemas que encontrarás

Añadirás métodos a `Shop` que solo pasan a través hacia `Catalog`, fomentando la duplicación accidental y ocultando dónde vive el comportamiento real cuando necesites cambiarlo después.
