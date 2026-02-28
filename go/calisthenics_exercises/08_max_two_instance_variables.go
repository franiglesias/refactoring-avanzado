package calisthenics_exercises

// Regla 8: Máximo de dos variables de instancia por clase
// Cuando una clase tiene muchas variables de instancia es posible que estas puedan
// agruparse representando conceptos. A menudo, la clase gestiona uno o dos conceptos
// del dominio, pero necesitamos varios valores primitivos para representar cada uno de ellos.

// Ejercicio: Agrupa las variables de instancia en value objects que representen
// conceptos del dominio.

type CheckoutSession struct {
	cartItems       []CartItem
	customerId      *string
	shippingAddress *string
	billingAddress  *string
	couponCode      *string
	paymentMethod   *string // "CARD" o "PAYPAL"
	currency        string
	taxRate         float64
}

type CartItem struct {
	ID    string
	Price float64
	Qty   int
}

func NewCheckoutSession() *CheckoutSession {
	return &CheckoutSession{
		cartItems: []CartItem{},
		currency:  "USD",
		taxRate:   0.21,
	}
}

func (c *CheckoutSession) AddItem(id string, price float64, qty int) {
	c.cartItems = append(c.cartItems, CartItem{ID: id, Price: price, Qty: qty})
}

func (c *CheckoutSession) Total() float64 {
	subtotal := 0.0
	for _, item := range c.cartItems {
		subtotal += item.Price * float64(item.Qty)
	}

	discount := 0.0
	if c.couponCode != nil {
		discount = 10 // lógica de descuento primitiva
	}

	taxed := (subtotal - discount) * (1 + c.taxRate)

	if c.currency == "USD" {
		return taxed
	}
	return taxed * 0.9 // conversión simulada
}
