"""
Lazy Class Code Smell Example

A class that doesn't do enough to justify its existence.
It has minimal functionality and could be inlined or removed.

Refactoring suggestion: Inline Class, Collapse Hierarchy
"""
from typing import TypedDict


class Address(TypedDict):
    name: str
    line1: str
    city: str | None


class ShippingLabelBuilder:
    """
    Lazy class: only has one trivial method.

    This class adds no value - its functionality could be
    a simple function or moved elsewhere.
    """

    def build(self, address: Address) -> str:
        city_part = f", {address['city']}" if address.get('city') else ''
        return f"{address['name']} — {address['line1']}{city_part}"


def print_shipping_label() -> None:
    address: Address = {
        'name': 'John Doe',
        'line1': '123 Main St',
        'city': 'New York',
    }

    label_builder = ShippingLabelBuilder()
    label = label_builder.build(address)
    print(label)
