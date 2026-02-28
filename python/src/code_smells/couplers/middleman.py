"""
Middleman Code Smell Example

A class exists solely to delegate most of its work to another class.
It adds no value and just forwards calls.

Refactoring suggestion: Remove Middle Man
"""


class Catalog:
    """The actual class doing the work"""

    def __init__(self) -> None:
        self._items: dict[str, str] = {}

    def add(self, item_id: str, name: str) -> None:
        self._items[item_id] = name

    def find(self, item_id: str) -> str | None:
        return self._items.get(item_id)

    def list(self) -> list[str]:
        return list(self._items.values())


class Shop:
    """
    Middleman: just delegates everything to Catalog.

    This class adds no value - it's just a pass-through.
    Clients should probably use Catalog directly.
    """

    def __init__(self, catalog: Catalog) -> None:
        self._catalog = catalog

    def add(self, item_id: str, name: str) -> None:
        self._catalog.add(item_id, name)

    def find(self, item_id: str) -> str | None:
        return self._catalog.find(item_id)

    def list(self) -> list[str]:
        return self._catalog.list()


def demo_middleman() -> list[str]:
    catalog = Catalog()
    shop = Shop(catalog)
    shop.add('1', 'Book')
    shop.add('2', 'Pen')
    return shop.list()
