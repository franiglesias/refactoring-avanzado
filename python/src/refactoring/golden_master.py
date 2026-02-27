from dataclasses import dataclass
from typing import Optional, Literal
from datetime import datetime
import random


@dataclass
class OrderItem:
    sku: str
    description: str
    unit_price: float
    quantity: int
    category: Optional[Literal['general', 'food', 'books']] = None


@dataclass
class Order:
    id: str
    customer_name: str
    items: list[OrderItem]


class ReceiptPrinter:
    """Do not change this function at the beginning of the exercise; first create the Golden Master."""

    def print(self, order: Order) -> str:
        now = self.get_current_date()

        header = f"Recibo {order.id} - {now.strftime('%x')} {now.strftime('%X')}"

        subtotal = 0
        lines = []
        for idx, it in enumerate(order.items):
            line_total = round_money(it.unit_price * it.quantity)
            subtotal = round_money(subtotal + line_total)
            lines.append(f"{idx + 1}. {it.description} ({it.sku}) x{it.quantity} = ${line_total:.2f}")

        lucky_discount_pct = self.discount()
        lucky_discount = round_money(subtotal * lucky_discount_pct)

        taxable_general = sum(
            i.unit_price * i.quantity
            for i in order.items
            if i.category != 'books' and i.category != 'food'
        )
        food_tax = sum(
            i.unit_price * i.quantity * 0.03
            for i in order.items
            if i.category == 'food'
        )
        general_tax = taxable_general * 0.07
        taxes = round_money(general_tax + food_tax)

        total = round_money(subtotal - lucky_discount + taxes)

        summary = [
            f"Subtotal: ${subtotal:.2f}",
            (f"Descuento de la suerte: -${lucky_discount:.2f} ({lucky_discount_pct * 100:.2f}%)"
             if lucky_discount > 0
             else "Descuento de la suerte: $0.00 (0.00%)"),
            f"Impuestos: ${taxes:.2f}",
            f"TOTAL: ${total:.2f}",
        ]

        return '\n'.join([header, *lines, '---', *summary])

    def discount(self) -> float:
        lucky_discount_pct = 0
        if random.random() < 0.1:
            lucky_discount_pct = random.random() * 0.05
        return lucky_discount_pct

    def get_current_date(self) -> datetime:
        return datetime.now()


def round_money(n: float) -> float:
    return round(n * 100) / 100


# Products catalog
PRODUCTS: list[dict] = [
    {'sku': 'BK-001', 'description': 'Libro: Clean Code', 'unit_price': 30, 'category': 'books'},
    {'sku': 'FD-010', 'description': 'Café en grano 1kg', 'unit_price': 12.5, 'category': 'food'},
    {'sku': 'GN-777', 'description': 'Cuaderno A5', 'unit_price': 5.2, 'category': 'general'},
    {'sku': 'GN-123', 'description': 'Bolígrafos (pack 10)', 'unit_price': 3.9, 'category': 'general'},
    {'sku': 'FD-222', 'description': 'Té verde 200g', 'unit_price': 6.75, 'category': 'food'},
]

CUSTOMERS = ['Ana', 'Luis', 'Mar', 'Iván', 'Sofía']


# Utility to generate Orders
def generate_order(
    id: str,
    customer_name: str,
    num_items: int,
    quantity: int,
) -> Order:
    items = []
    for i in range(num_items):
        p = PRODUCTS[i]
        items.append(OrderItem(
            sku=p['sku'],
            description=p['description'],
            unit_price=p['unit_price'],
            quantity=quantity,
            category=p.get('category')
        ))

    return Order(id=id, customer_name=customer_name, items=items)
