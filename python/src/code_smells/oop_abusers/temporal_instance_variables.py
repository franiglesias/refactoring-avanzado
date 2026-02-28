"""
Temporal Instance Variables Code Smell Example

Instance variables that are only valid during certain phases of an object's lifecycle.
This creates temporal coupling and makes the object's state inconsistent.

Refactoring suggestion: Replace Method with Method Object, Extract Class
"""
from typing import Literal


class PizzaOrder:
    """
    Instance variables (size, toppings, address) are only meaningful
    between start() and place() calls.

    The object goes through phases: uninitialized -> started -> placed -> uninitialized
    This temporal coupling makes the class fragile and error-prone.
    """

    def __init__(self) -> None:
        self._size: Literal['S', 'M', 'L'] | None = None
        self._toppings: list[str] = []
        self._address: str | None = None

    def start(self, size: Literal['S', 'M', 'L']) -> None:
        """Phase 1: Initialize order"""
        self._size = size
        self._toppings = []
        self._address = None

    def add_topping(self, topping: str) -> None:
        """Phase 2: Build order (only valid after start())"""
        if not self._size:
            return
        self._toppings.append(topping)

    def set_delivery_address(self, address: str) -> None:
        """Phase 2: Build order (only valid after start())"""
        self._address = address

    def place(self) -> str:
        """Phase 3: Finalize and reset"""
        summary = f"Pizza {self._size or '?'} to {self._address or 'UNKNOWN'} with [{', '.join(self._toppings)}]"
        # Reset temporal state - object returns to uninitialized phase
        self._size = None
        self._address = None
        self._toppings = []
        return summary


def demo_pizza_order() -> str:
    order = PizzaOrder()
    order.start('L')
    order.add_topping('pepperoni')
    order.add_topping('mushroom')
    order.set_delivery_address('123 Main St')
    return order.place()
