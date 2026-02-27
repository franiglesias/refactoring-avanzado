export type Product = { id: string; price: number }

export function addProduct(products: Product[], product: Product): void {
  const exists = products.some((p) => p.id === product.id)
  if (!exists) products.push(product)
}

export function totalPrice(products: Product[]): number {
  return products.map((p) => p.price).reduce((a, b) => a + b, 0)
}

export function removeProduct(products: Product[], productId: string): Product[] {
  return products.filter((p) => p.id !== productId)
}
