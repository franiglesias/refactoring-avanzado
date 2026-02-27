export class PizzaOrder {
  private size?: 'S' | 'M' | 'L' | undefined
  private toppings: string[] = []
  private address?: string | undefined

  start(size: 'S' | 'M' | 'L') {
    this.size = size
    this.toppings = []
    this.address = undefined
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
