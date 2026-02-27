from dataclasses import dataclass


@dataclass
class OrderItem:
    name: str
    price: float
    quantity: int


@dataclass
class PrintJob:
    title: str
    items: list[dict]
    subtotal: float
    discount: float
    tax: float
    shipping: float
    total: float
    currency: str
    formatted: dict
    metadata: dict


class OrderService:
    def process(self, order: dict):
        # Validar el pedido
        if not order.get('items') or len(order['items']) == 0:
            print('El pedido no tiene productos')
            return

        # Validar precios y cantidades
        for item in order['items']:
            if item['price'] < 0 or item['quantity'] <= 0:
                print('Producto inválido en el pedido')
                return

        # Constantes de negocio (simples por ahora)
        TAX_RATE = 0.21  # 21% IVA
        FREE_SHIPPING_THRESHOLD = 50
        SHIPPING_FLAT = 5

        # Calcular subtotal
        subtotal = 0
        for item in order['items']:
            subtotal += item['price'] * item['quantity']

        # Descuento por cliente VIP (10% del subtotal)
        discount = 0
        if order.get('customer_type') == 'VIP':
            discount = round_money(subtotal * 0.1)
            print('Descuento VIP aplicado')

        # Base imponible
        taxable = max(0, subtotal - discount)

        # Impuestos
        tax = round_money(taxable * TAX_RATE)

        # Envío
        shipping = 0 if taxable >= FREE_SHIPPING_THRESHOLD else SHIPPING_FLAT

        # Total
        total = round_money(taxable + tax + shipping)

        # Actualizar el pedido (side-effects requeridos)
        order['subtotal'] = round_money(subtotal)
        order['discount'] = discount
        order['tax'] = tax
        order['shipping'] = shipping
        order['total'] = total

        # Registrar en la base de datos (simulado)
        print(f"[DB] Guardando pedido con total {format_money(total)}")

        # Enviar correo de confirmación
        print(f"[MAIL] Enviando correo a {order.get('customer_email')}")

        # Imprimir recibo
        print(f"[PRN] Imprimiendo recibo para pedido con total {format_money(total)}")


def round_money(n: float) -> float:
    return round(n * 100) / 100


def format_money(n: float) -> str:
    return f"${n:.2f}"
