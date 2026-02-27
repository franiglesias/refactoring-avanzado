export function processOrdersWithDiscounts(
  orders: Array<{ id: string; items: Array<{ price: number }>; customer: { isVip: boolean } }>,
): number {
  let total = 0
  for (const order of orders) {
    if (order.items && order.items.length > 0) {
      for (const item of order.items) {
        if (order.customer && order.customer.isVip) {
          if (item.price > 100) {
            total += item.price * 0.8 // gran descuento VIP
          } else {
            total += item.price * 0.9 // pequeño descuento VIP
          }
        } else {
          if (item.price > 100) {
            total += item.price * 0.95 // gran descuento regular
          } else {
            total += item.price // sin descuento
          }
        }
      }
    }
  }
  return total
}
