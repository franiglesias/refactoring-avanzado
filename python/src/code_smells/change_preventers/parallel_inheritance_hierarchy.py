"""
Parallel Inheritance Hierarchy Code Smell Example

When you create a subclass in one hierarchy, you must create a corresponding
subclass in another hierarchy. The hierarchies mirror each other.

Refactoring suggestion: Move Method and Move Field to eliminate one hierarchy
"""
from abc import ABC, abstractmethod


class Component(ABC):
    """Abstract component hierarchy"""

    @abstractmethod
    def draw(self, renderer: 'Renderer') -> str:
        pass


class Button(Component):
    """Each component type needs a corresponding renderer method"""

    def __init__(self, label: str) -> None:
        self.label = label

    def draw(self, renderer: 'Renderer') -> str:
        return renderer.render_button(self)


class TextBox(Component):
    """Each component type needs a corresponding renderer method"""

    def __init__(self, text: str) -> None:
        self.text = text

    def draw(self, renderer: 'Renderer') -> str:
        return renderer.render_text_box(self)


class Renderer(ABC):
    """Abstract renderer hierarchy mirrors Component hierarchy"""

    @abstractmethod
    def render_button(self, button: Button) -> str:
        pass

    @abstractmethod
    def render_text_box(self, textbox: TextBox) -> str:
        pass


class HtmlRenderer(Renderer):
    """For each Component subclass, we need a render method here"""

    def render_button(self, button: Button) -> str:
        return f"<button>{button.label}</button>"

    def render_text_box(self, textbox: TextBox) -> str:
        return f'<input value="{textbox.text}"/>'


class MarkdownRenderer(Renderer):
    """For each Component subclass, we need a render method here"""

    def render_button(self, button: Button) -> str:
        return f"[{button.label}]"

    def render_text_box(self, textbox: TextBox) -> str:
        return f"_{textbox.text}_"


def demo_parallel_hierarchy() -> list[str]:
    components: list[Component] = [Button('Save'), TextBox('name')]
    renderer = HtmlRenderer()
    return [component.draw(renderer) for component in components]
