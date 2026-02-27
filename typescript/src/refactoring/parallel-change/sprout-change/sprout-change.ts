export type CartItem = {
  id: string
  price: number
  qty: number
  category?: 'general' | 'books' | 'food'
}
export type Region = 'US' | 'EU'

interface TaxPolicy {

}
// Regla existente: un único impuesto plano por región; los libros y la comida están exentos en la UE
function calculateUSTax(cart: CartItem[]) {
  const usSubtotal = cart.reduce((s, it) => s + it.price * it.qty, 0)

  return usSubtotal * 0.07 // 7% plano
}

function calculateEUTax(cart: CartItem[]) {
  const taxable = cart
    .filter((it) => it.category !== 'books' && it.category !== 'food')
    .reduce((s, it) => s + it.price * it.qty, 0)
  return taxable * 0.2 // 20% plano solo sobre los ítems gravables
}

function calculateDefault(cart: CartItem[]) {
  return 0
}

// (reglas embebidas en línea)
export function calculateTotal(cart: CartItem[], region: Region): number {
  const subtotal = cart.reduce((s, it) => s + it.price * it.qty, 0)

  let tax = 0
  if (region === 'US') {
    tax = calculateUSTax(cart);
  } else if (region === 'EU') {
    // exenciones ingenuas en línea
    tax = calculateEUTax(cart);
  } else {
    tax = calculateDefault(cart);
  }

  return roundCurrency(subtotal + tax)
}

export function roundCurrency(amount: number): number {
  return Math.round(amount * 100) / 100
}

// Uso de ejemplo, mantenido simple para estudiantes
export function demoSprout(): number {
  const cart: CartItem[] = [
    {id: 'p1', price: 10, qty: 2, category: 'general'},
    {id: 'b1', price: 20, qty: 1, category: 'books'},
  ]
  return calculateTotal(cart, 'EU')
}
