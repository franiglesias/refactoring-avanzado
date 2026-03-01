# Long Method

Método largo.

## Definición

Un método en una clase es muy largo.

## Ejemplo

```java
public void process(Order order) {
    // Validar el pedido
    if (order.items == null || order.items.isEmpty()) {
        System.out.println("El pedido no tiene productos");
        return;
    }

    // Validar precios y cantidades
    for (OrderItem item : order.items) {
        if (item.price < 0 || item.quantity <= 0) {
            System.out.println("Producto inválido en el pedido");
            return;
        }
    }

    // Constantes de negocio
    final double TAX_RATE = 0.21; // 21% IVA
    final double FREE_SHIPPING_THRESHOLD = 50;
    final double SHIPPING_FLAT = 5;

    // Calcular subtotal
    double subtotal = 0;
    for (OrderItem item : order.items) {
        subtotal += item.price * item.quantity;
    }

    // Descuento por cliente VIP (10% del subtotal)
    double discount = 0;
    if (order.customerType == CustomerType.VIP) {
        discount = roundMoney(subtotal * 0.1);
        System.out.println("Descuento VIP aplicado");
    }

    // Base imponible
    double taxable = Math.max(0, subtotal - discount);

    // Impuestos
    double tax = roundMoney(taxable * TAX_RATE);

    // Envío
    double shipping = taxable >= FREE_SHIPPING_THRESHOLD ? 0 : SHIPPING_FLAT;

    // Total
    double total = roundMoney(taxable + tax + shipping);

    // Actualizar el pedido
    order.subtotal = roundMoney(subtotal);
    order.discount = discount;
    order.tax = tax;
    order.shipping = shipping;
    order.total = total;

    // Simular guardado en base de datos...
    // Enviar correo de confirmación...
    // Imprimir resumen...
}
```

## Ejercicio

Añade soporte de cupones con expiración y multi-moneda (USD/EUR) con reglas de redondeo distintas.

## Problemas que encontrarás

Tienes que tocar diferentes secciones dentro del método, lo que genera riesgo de cambios indeseados y aumenta el esfuerzo de mantenimiento.
