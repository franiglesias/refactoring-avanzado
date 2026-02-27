class Order {
  constructor(
    private customerName: string,
    private customerEmail: string,
    private address: string,
    private totalAmount: number,
    private currency: string,
  ) {
  }

  sendInvoice() {
    if (!this.customerEmail.includes('@')) {
      throw new Error('Email inválido')
    }
    if (!this.address) {
      throw new Error('No se ha indicado dirección')
    }
    if (this.totalAmount <= 0) {
      throw new Error('El monto debe ser mayor que cero')
    }
    console.log(`Factura enviada a ${this.customerName} in ${this.address} por ${this.totalAmount} ${this.currency}`)
  }
}
