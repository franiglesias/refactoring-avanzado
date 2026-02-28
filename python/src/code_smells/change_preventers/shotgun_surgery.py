"""
Shotgun Surgery Code Smell Example

A single change requires making small modifications across many different classes.
When you need to change the tax rate, you must modify multiple classes.

Refactoring suggestion: Move Method, Move Field, Inline Class
"""
from typing import TypedDict


class LineItem(TypedDict):
    name: str
    price: float
    qty: int


class PriceCalculator:
    """Tax calculation logic duplicated here"""

    def total_with_tax(self, items: list[LineItem]) -> float:
        subtotal = sum(item['price'] * item['qty'] for item in items)
        tax = subtotal * 0.21
        return subtotal + tax


class InvoiceService:
    """Tax calculation logic duplicated here"""

    def create_total(self, items: list[LineItem]) -> float:
        base = sum(item['price'] * item['qty'] for item in items)
        vat = base * 0.21
        return base + vat


class SalesReport:
    """Tax calculation logic duplicated here"""

    def summarize(self, items: list[LineItem]) -> str:
        total_sum = sum(item['price'] * item['qty'] for item in items)
        tax = total_sum * 0.21
        total = total_sum + tax
        return f"total={total:.2f}"


class LoyaltyPoints:
    """Tax calculation logic duplicated here"""

    def points(self, items: list[LineItem]) -> int:
        base = sum(item['price'] * item['qty'] for item in items)
        with_tax = base + base * 0.21
        return int(with_tax // 10)


def demo_shotgun(items: list[LineItem]) -> tuple[float, float]:
    calc = PriceCalculator().total_with_tax(items)
    inv = InvoiceService().create_total(items)
    return (calc, inv)
