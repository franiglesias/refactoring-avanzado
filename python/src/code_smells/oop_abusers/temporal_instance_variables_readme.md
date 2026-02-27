# Temporal Instance Variables

Variables de instancia temporales.

## Definición

Este smell ocurre cuando un objeto tiene campos que solo están llenos (o tienen sentido) en ciertas etapas de su ciclo de vida. Esto suele indicar un acoplamiento temporal, donde los métodos deben llamarse en un orden específico para que el objeto sea válido, dejando al objeto en un estado inconsistente fuera de esa secuencia.

## Ejemplo

`PizzaOrder` utiliza variables de instancia que solo son válidas entre la llamada a `start()` y `place()`.

```typescript
export class PizzaOrder {
  private size?: 'S' | 'M' | 'L'
  private toppings: string[] = []
  private address?: string

  start(size: 'S' | 'M' | 'L') {
    this.size = size
    this.toppings = []
    this.address = this.address
  }

  addTopping(topping: string) {
    if (!this.size) {
      return
    }
    this.toppings.push(topping)
  }

  setDeliveryAddress(address: string) {
    this.address = address
  }

  place(): string {
    const summary = `Pizza ${this.size ?? '?'} to ${this.address ?? 'UNKNOWN'} with [${this.toppings.join(', ')}]`
    this.size = undefined
    this.address = undefined
    this.toppings = []
    return summary
  }
}

export function demoPizzaOrder(): string {
  const o = new PizzaOrder()
  o.start('L')
  o.addTopping('pepperoni')
  o.addTopping('mushroom')
  o.setDeliveryAddress('123 Main St')
  return o.place()
}
```

## Ejercicio

Añade una validación para que no se pueda llamar a `place()` si no se ha añadido al menos un ingrediente.

## Problemas que encontrarás

Te darás cuenta de que el objeto es una "máquina de estados" frágil. Si un cliente olvida llamar a `start()` o intenta llamar a `addTopping()` fuera de orden, el sistema puede fallar silenciosamente o requerir comprobaciones constantes de nulidad en cada método.
