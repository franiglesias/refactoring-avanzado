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
