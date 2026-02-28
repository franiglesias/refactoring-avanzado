# Parallel Inheritance Hierarchy

## Definición

Este smell ocurre cuando al cambiar una jerarquía de herencia, tienes que cambiar obligatoriamente otra jerarquía relacionada. Esto indicaría que ambas jerarquías están acopladas entre sí y no pueden evolucionar independientemente.

## Ejemplo

Agregar un nuevo componente de UI obliga a agregar métodos correspondientes en cada renderer, haciendo que ambas jerarquías tengan que crecer al unísono.

```php
<?php

declare(strict_types=1);

namespace CodeSmells\ChangePreventers;

abstract class Component
{
    abstract public function draw(Renderer $renderer): string;
}

class Button extends Component
{
    public function __construct(public string $label)
    {
    }

    public function draw(Renderer $renderer): string
    {
        return $renderer->renderButton($this);
    }
}

class TextBox extends Component
{
    public function __construct(public string $text)
    {
    }

    public function draw(Renderer $renderer): string
    {
        return $renderer->renderTextBox($this);
    }
}

abstract class Renderer
{
    abstract public function renderButton(Button $b): string;

    abstract public function renderTextBox(TextBox $t): string;
}

class HtmlRenderer extends Renderer
{
    public function renderButton(Button $b): string
    {
        return "<button>{$b->label}</button>";
    }

    public function renderTextBox(TextBox $t): string
    {
        return "<input value=\"{$t->text}\"/>";
    }
}

class MarkdownRenderer extends Renderer
{
    public function renderButton(Button $b): string
    {
        return "[{$b->label}]";
    }

    public function renderTextBox(TextBox $t): string
    {
        return "_{$t->text}_";
    }
}

function demoParallelHierarchy(): array
{
    $comps = [new Button('Save'), new TextBox('name')];
    $renderer = new HtmlRenderer();
    return array_map(fn($c) => $c->draw($renderer), $comps);
}
```

## Ejercicio

Añade un componente `Image` que muestre una imagen.

## Problemas que encontrarás

Necesitarás añadir Image y renderImage a Renderer, e implementarlo en todos los renderers, mostrando cambios en paralelo.
