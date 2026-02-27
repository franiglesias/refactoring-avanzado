# Middleman

Intermediario.

## Definición

Ocurre cuando una clase realiza una única acción: delegar el trabajo a otra clase. Si una clase existe solo como un "pasamanos" hacia otro objeto, es posible que estemos ante una capa de abstracción innecesaria que oscurece al colaborador real.

## Ejemplo

`Shop` hace poco más que delegar a `Catalog`, añadiendo una capa innecesaria que oculta dónde ocurre realmente la lógica.

```typescript
export class Catalog {
  private items = new Map<string, string>()

  add(id: string, name: string): void {
    this.items.set(id, name)
  }

  find(id: string): string | undefined {
    return this.items.get(id)
  }

  list(): string[] {
    return Array.from(this.items.values())
  }
}

export class Shop {
  constructor(private catalog: Catalog) {
  }

  add(id: string, name: string): void {
    this.catalog.add(id, name)
  }

  find(id: string): string | undefined {
    return this.catalog.find(id)
  }

  list(): string[] {
    return this.catalog.list()
  }
}
```

## Ejercicio

Añade una funcionalidad `searchByPrefix` en `Catalog` y propágala a través de `Shop`.

## Problemas que encontrarás

Añadirás métodos a `Shop` que solo pasan a través hacia `Catalog`, fomentando la duplicación accidental y ocultando dónde vive el comportamiento real cuando necesites cambiarlo después.
