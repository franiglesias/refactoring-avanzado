"""
Alternative Classes with Different Interfaces Code Smell Example

Two classes perform similar functions but have different method names.
This makes them hard to use interchangeably and violates consistency.

Refactoring suggestion: Rename Method, Extract Superclass, Extract Interface
"""
from typing import Literal


class TextLogger:
    """Logs messages using 'log' method"""

    def log(self, message: str) -> None:
        print(f"[text] {message}")


class MessageWriter:
    """
    Does the same thing as TextLogger but uses 'write' method.

    This inconsistency makes it harder to use them interchangeably
    and requires conditional logic to handle both.
    """

    def write(self, entry: str) -> None:
        print(f"[text] {entry}")


def use_alt_classes(choice: Literal['logger', 'writer'], msg: str) -> None:
    """
    Client code must know about both interfaces.

    If these classes had the same interface, polymorphism could be used instead.
    """
    if choice == 'logger':
        TextLogger().log(msg)
    else:
        MessageWriter().write(msg)
