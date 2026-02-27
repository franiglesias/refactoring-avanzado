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
