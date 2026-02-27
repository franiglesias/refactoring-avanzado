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
