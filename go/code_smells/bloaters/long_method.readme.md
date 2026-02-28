# Long method

Método largo.

## Definición

Un método en una clase es muy largo.

## Ejemplo

```go
type OrderService struct{}

func (s *OrderService) Process(order *Order) {
	// Validar el pedido
	if order.Items == nil || len(order.Items) == 0 {
		fmt.Println("El pedido no tiene productos")
		return
	}

	// Validar precios y cantidades
	for _, item := range order.Items {
		if item.Price < 0 || item.Quantity <= 0 {
			fmt.Println("Producto inválido en el pedido")
			return
		}
	}

	// Constantes de negocio (simples por ahora)
	const TAX_RATE = 0.21           // 21% IVA
	const FREE_SHIPPING_THRESHOLD = 50.0
	const SHIPPING_FLAT = 5.0

	// Calcular subtotal
	subtotal := 0.0
	for _, item := range order.Items {
		subtotal += item.Price * float64(item.Quantity)
	}

	// Descuento por cliente VIP (10% del subtotal)
	discount := 0.0
	if order.CustomerType == "VIP" {
		discount = roundMoney(subtotal * 0.1)
		fmt.Println("Descuento VIP aplicado")
	}

	// Base imponible
	taxable := math.Max(0, subtotal-discount)

	// Impuestos
	tax := roundMoney(taxable * TAX_RATE)

	// Envío
	shipping := 0.0
	if taxable >= FREE_SHIPPING_THRESHOLD {
		shipping = 0
	} else {
		shipping = SHIPPING_FLAT
	}

	// Total
	total := roundMoney(taxable + tax + shipping)

	// ... Cientos de líneas más para:
	// - Registrar en la base de datos (simulado)
	// - Enviar correo de confirmación
	// - Imprimir resumen en impresora térmica
}
```

## Ejercicio

Añade soporte de cupones con expiración y multi-moneda (USD/EUR) con reglas de redondeo distintas.

## Problemas que encontrarás

Tienes que tocar diferentes secciones dentro del método, lo que genera riesgo de cambios indeseados
y aumenta el esfuerzo de mantenimiento.
