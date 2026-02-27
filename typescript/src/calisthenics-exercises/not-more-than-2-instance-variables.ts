export class CheckoutSession {
  private cartItems: Array<{ id: string; price: number; qty: number }> = []
  private customerId: string | null = null
  private shippingAddress: string | null = null
  private billingAddress: string | null = null
  private couponCode: string | null = null
  private paymentMethod: 'CARD' | 'PAYPAL' | null = null
  private currency: string = 'USD'
  private taxRate: number = 0.21

  addItem(id: string, price: number, qty: number) {
    this.cartItems.push({id, price, qty})
  }

  total(): number {
    const subtotal = this.cartItems.reduce((sum, i) => sum + i.price * i.qty, 0)
    const discount = this.couponCode ? 10 : 0 // lógica de descuento primitiva
    const taxed = (subtotal - discount) * (1 + this.taxRate)
    return this.currency === 'USD' ? taxed : taxed * 0.9 // conversión simulada
  }
}
