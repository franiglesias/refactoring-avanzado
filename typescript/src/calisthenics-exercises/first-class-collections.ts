export type Product = { id: string; price: number }

export class ProductOperations {
  private products: Product[] = []

  addProduct(product: Product): void {
    const exists = this.products.some((p) => p.id === product.id)
    if (!exists) this.products.push(product)
  }

  totalPrice(): number {
    return this.products.map((p) => p.price).reduce((a, b) => a + b, 0)
  }

  removeProduct(productId: string): void {
    this.products = this.products.filter((p) => p.id !== productId)
  }
}
