"""
Message Chains Code Smell Example

A client asks an object for another object, which the client then asks for yet another object,
forming a chain: a.get_b().get_c().get_d()

This creates tight coupling and makes code fragile to changes in the chain.

Refactoring suggestion: Hide Delegate
"""


class Level2:
    """Deepest level in the chain"""

    def __init__(self, value: int) -> None:
        self._value = value

    def get_value(self) -> int:
        return self._value


class Level1:
    """Middle level in the chain"""

    def __init__(self, next_level: Level2) -> None:
        self._next = next_level

    def get_next(self) -> Level2:
        return self._next


class Root:
    """Root level in the chain"""

    def __init__(self, next_level: Level1) -> None:
        self._next = next_level

    def get_next(self) -> Level1:
        return self._next


def read_deep(root: Root) -> int:
    """
    Message chain: root.get_next().get_next().get_value()

    This client knows too much about the internal structure.
    Any change in the chain breaks this code.
    """
    return root.get_next().get_next().get_value()
