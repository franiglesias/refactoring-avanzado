# Primitive Obsession

Obsesión primitiva.

## Definición

Conceptos de dominio se modelan con primitivos, lo que obliga a esparcir reglas de validación, formato, y todo tipo de comportamiento, por todo el código.

## Ejemplo

```typescript
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
```

## Ejercicio

Introduce soporte para diferentes monedas, para enviar la facture por email, y para formatear la dirección en función del país.

## Problemas que encontrarás

Dado que los primitivos no nos permiten garantizar la integridad de sus valores, tendrás que introducir validaciones en muchos lugares, incluso de forma repetida. Algunos datos siempre viajan juntos (Data Clump), por lo que tienes que asegurarte de que permanecen juntos.

Para formatear de forma diferente basándote en algún dato arbitrario tendrás que introducir lógica de decisión.
