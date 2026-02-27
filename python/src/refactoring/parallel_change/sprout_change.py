from dataclasses import dataclass
from typing import Optional, Literal


@dataclass
class CartItem:
    id: str
    price: float
    qty: int
    category: Optional[Literal['general', 'books', 'food']] = None


Region = Literal['US', 'EU']


# Regla existente: un único impuesto plano por región; los libros y la comida están exentos en la UE
# (reglas embebidas en línea)
def calculate_total(cart: list[CartItem], region: Region) -> float:
    subtotal = sum(it.price * it.qty for it in cart)

    tax = 0.0
    if region == 'US':
        tax = subtotal * 0.07  # 7% plano
    elif region == 'EU':
        # exenciones ingenuas en línea
        taxable = sum(
            it.price * it.qty
            for it in cart
            if it.category not in ('books', 'food')
        )
        tax = taxable * 0.2  # 20% plano solo sobre los ítems gravables

    return round_currency(subtotal + tax)


def round_currency(amount: float) -> float:
    return round(amount * 100) / 100


# Uso de ejemplo, mantenido simple para estudiantes
def demo_sprout() -> float:
    cart = [
        CartItem(id='p1', price=10, qty=2, category='general'),
        CartItem(id='b1', price=20, qty=1, category='books'),
    ]
    return calculate_total(cart, 'EU')
