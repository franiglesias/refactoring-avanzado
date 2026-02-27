# Feature Envy

Envidia de características.

## Definición

Una clase cliente la información de otra clase colaboradora para hacer cálculos o tomar decisiones, sugiriendo que la segunda clase debería exponer esos comportamientos. Al depender de la estructura de la colaboradora, la clase cliente queda acoplada.

## Ejemplo

`ShippingCalculator` se mete en los datos de `Customer` para tomar decisiones, lo que indica que el comportamiento quizá debería pertenecer a `Customer`.

```typescript
export class Customer {
  constructor(
    public name: string,
    public street: string,
    public city: string,
    public zip: string,
  ) {
  }
}

export class ShippingCalculator {
  cost(customer: Customer): number {
    const base = customer.zip.startsWith('9') ? 10 : 20
    const distant = customer.city.length > 6 ? 5 : 0
    return base + distant
  }
}

export function demoFeatureEnvy(c: Customer): number {
  return new ShippingCalculator().cost(c)
}
```

## Ejercicio

Añade envío gratis para clientes en ciertas ciudades y un recargo de fin de semana.

## Problemas que encontrarás

Probablemente, seguirás añadiendo condiciones dentro de `ShippingCalculator` que dependen de detalles internos de `Customer`, esparciendo reglas en el lugar equivocado y volviendo frágiles los cambios.
