# Lazy Class

Clase perezosa.

## Definición

Una clase perezosa es aquella que no aporta suficiente valor para justificar su existencia. Suelen ser clases que solo envuelven una operación trivial o que tienen muy poca responsabilidad, añadiendo una complejidad innecesaria al sistema.

## Ejemplo

La clase `ShippingLabelBuilder` solo tiene un método que realiza una concatenación de strings simple, algo que podría resolverse con una función pura.

```typescript
export type Address = { name: string; line1: string; city?: string }

export class ShippingLabelBuilder {
  build(a: Address): string {
    return `${a.name} — ${a.line1}${a.city ? ', ' + a.city : ''}`
  }
}

export function printShippingLabel() {
  const address: Address = {
    name: 'John Doe',
    line1: '123 Main St',
    city: 'New York',
  }

  const labelBuilder = new ShippingLabelBuilder()
  const label = labelBuilder.build(address)
  console.log(label)
}
```

## Ejercicio

Reescribe el código para eliminar la necesidad de la clase `ShippingLabelBuilder`.

## Problemas que encontrarás

Mantener una estructura de clase para una lógica tan simple te obliga a instanciar objetos innecesariamente y añade capas de abstracción que dificultan la legibilidad del código sin ofrecer beneficios a cambio.
