export type LineItem = { name: string; price: number; qty: number }

export class PriceCalculator {
  totalWithTax(items: LineItem[]): number {
    const subtotal = items.reduce((s, i) => s + i.price * i.qty, 0)
    const tax = subtotal * 0.21
    return subtotal + tax
  }
}

export class InvoiceService {
  createTotal(items: LineItem[]): number {
    const base = items.reduce((s, i) => s + i.price * i.qty, 0)
    const vat = base * 0.21
    return base + vat
  }
}

export class SalesReport {
  summarize(items: LineItem[]): string {
    const sum = items.reduce((s, i) => s + i.price * i.qty, 0)
    const tax = sum * 0.21
    const total = sum + tax
    return `total=${total.toFixed(2)}`
  }
}

export class LoyaltyPoints {
  points(items: LineItem[]): number {
    const base = items.reduce((s, i) => s + i.price * i.qty, 0)
    const withTax = base + base * 0.21
    return Math.floor(withTax / 10)
  }
}

export function demoShotgun(items: LineItem[]): [number, number] {
  const calc = new PriceCalculator().totalWithTax(items)
  const inv = new InvoiceService().createTotal(items)
  return [calc, inv]
}
