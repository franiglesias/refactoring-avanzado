export type OrderItem = {
  sku: string
  description: string
  unitPrice: number
  quantity: number
  category?: 'general' | 'food' | 'books'
}

export type Order = {
  id: string
  customerName: string
  items: OrderItem[]
}

export class ReceiptPrinter {
  // Do not change this function at the beginning of the exercise; first create the Golden Master.
  print(order: Order): string {
    const now = this.getCurrentDate()

    const header = `Recibo ${order.id} - ${now.toLocaleDateString()} ${now.toLocaleTimeString()}`

    let subtotal = 0
    let lines = order.items.map((it, idx) => {
      const lineTotal = round(it.unitPrice * it.quantity)
      subtotal = round(subtotal + lineTotal)
      return `${idx + 1}. ${it.description} (${it.sku}) x${it.quantity} = $${lineTotal.toFixed(2)}`
    })
    let luckyDiscountPct = this.discount();
    const luckyDiscount = round(subtotal * luckyDiscountPct)

    const taxableGeneral = order.items
      .filter((i) => i.category !== 'books')
      .reduce((s, i) => s + (i.category === 'food' ? 0 : i.unitPrice * i.quantity), 0)
    const foodTax = order.items
      .filter((i) => i.category === 'food')
      .reduce((s, i) => s + i.unitPrice * i.quantity * 0.03, 0)
    const generalTax = taxableGeneral * 0.07
    const taxes = round(generalTax + foodTax)

    const total = round(subtotal - luckyDiscount + taxes)

    const summary = [
      `Subtotal: $${subtotal.toFixed(2)}`,
      luckyDiscount > 0
        ? `Descuento de la suerte: -$${luckyDiscount.toFixed(2)} (${(luckyDiscountPct * 100).toFixed(2)}%)`
        : `Descuento de la suerte: $0.00 (0.00%)`,
      `Impuestos: $${taxes.toFixed(2)}`,
      `TOTAL: $${total.toFixed(2)}`,
    ]

    return [header, ...lines, '---', ...summary].join('\n')
  }

  protected discount() {
    let luckyDiscountPct = 0
    if (Math.random() < 0.1) {
      luckyDiscountPct = Math.random() * 0.05
    }
    return luckyDiscountPct;
  }

  protected getCurrentDate() {
    return new Date(Date.now());
  }
}

function round(n: number): number {
  return Math.round(n * 100) / 100
}

const products: Omit<OrderItem, 'quantity'>[] = [
  {sku: 'BK-001', description: 'Libro: Clean Code', unitPrice: 30, category: 'books'},
  {sku: 'FD-010', description: 'Café en grano 1kg', unitPrice: 12.5, category: 'food'},
  {sku: 'GN-777', description: 'Cuaderno A5', unitPrice: 5.2, category: 'general'},
  {sku: 'GN-123', description: 'Bolígrafos (pack 10)', unitPrice: 3.9, category: 'general'},
  {sku: 'FD-222', description: 'Té verde 200g', unitPrice: 6.75, category: 'food'},
]

const customers = ['Ana', 'Luis', 'Mar', 'Iván', 'Sofía']

// Utility to generate Orders
export function generateOrder(
  id: string,
  customerName: string,
  numItems: number,
  quantity: number,
): Order {
  const items: OrderItem[] = []
  for (let i = 0; i < numItems; i++) {
    const p = products[i]
    items.push({...p, quantity: quantity} as OrderItem) // 1..4 unidades
  }

  return {id, customerName, items}
}
