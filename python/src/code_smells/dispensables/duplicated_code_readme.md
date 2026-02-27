# Duplicated Code

Código duplicado.

## Definición

El código duplicado ocurre cuando la misma estructura de código o lógica aparece en más de un lugar. Es uno de los code smells más comunes y peligrosos, ya que cualquier cambio en la lógica debe replicarse en todas las copias, aumentando el riesgo de inconsistencias.

## Ejemplo

Dos funciones realizan exactamente la misma lógica de cálculo de subtotal e impuestos, variando mínimamente en los nombres de las propiedades.

```typescript
export function calculateOrderTotalWithTax(
  items: { price: number; qty: number }[],
  taxRate: number,
): number {
  let subtotal = 0
  for (const item of items) {
    subtotal += item.price * item.qty
  }
  const tax = subtotal * taxRate
  return subtotal + tax
}

export function computeCartTotalIncludingTax(
  items: { price: number; quantity: number }[],
  taxRate: number,
): number {
  let subtotal = 0
  for (const item of items) {
    subtotal += item.price * item.quantity
  }
  const tax = subtotal * taxRate
  return subtotal + tax
}

export function demoDuplicatedCode(): [number, number] {
  const itemsA = [
    {price: 10, qty: 2},
    {price: 5, qty: 3},
  ]
  const itemsB = [
    {price: 10, quantity: 2},
    {price: 5, quantity: 3},
  ]
  return [calculateOrderTotalWithTax(itemsA, 0.21), computeCartTotalIncludingTax(itemsB, 0.21)]
}
```

## Ejercicio

Cambia la regla de impuestos para que sea escalonada (ej. 10% hasta $100 y 21% por encima).

## Problemas que encontrarás

Tendrás que actualizar múltiples implementaciones y recordar mantenerlas consistentes, lo que demuestra cómo la duplicación multiplica el esfuerzo y el riesgo de error humano.
