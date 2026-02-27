# Data clump

Grupo de datos.

## Definición

El mismo grupo de campos de datos viaja junto por muchos lugares, lo que sugiere un Value Object faltante y duplicación.

## Ejemplo

```typescript
export class Invoice {
  private readonly customerName: string
  private readonly customerCity: string
  private readonly customerStreet: string
  private readonly customerZip: string

  constructor(
    customerName: string,
    customerStreet: string,
    customerCity: string,
    customerZip: string,
  ) {
    this.customerZip = customerZip
    this.customerStreet = customerStreet
    this.customerCity = customerCity
    this.customerName = customerName
  }

  print(): string {
    return (
      `Factura para: ${this.customerName}\n` +
      `Dirección: ${this.customerStreet}, ${this.customerCity}, ${this.customerZip}`
    )
  }
}

export class ShippingLabel {
  private readonly customerName: string
  private readonly customerStreet: string
  private readonly customerCity: string
  private readonly customerZip: string

  constructor(
    customerName: string,
    customerStreet: string,
    customerCity: string,
    customerZip: string,
  ) {
    this.customerZip = customerZip
    this.customerCity = customerCity
    this.customerStreet = customerStreet
    this.customerName = customerName
  }

  print(): string {
    return (
      `Enviar a: ${this.customerName}\n` +
      `${this.customerStreet}, ${this.customerCity}, ${this.customerZip}`
    )
  }
}
```

## Ejercicio

Añade país y provincia y reglas de formateo internacional de la dirección.

## Problemas que encontrarás

Necesitarás modificar constructores, impresores y cualquier lugar que pase estos campos juntos,
multiplicando la superficie de cambio.
