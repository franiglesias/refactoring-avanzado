# Parallel Inheritance Hierarchy

Jerarquía de herencia paralela.

## Definición

Cada vez que añades una subclase a una jerarquía, también debes añadir una subclase correspondiente a otra jerarquía. Las dos jerarquías crecen en paralelo, duplicando esfuerzo.

## Ejemplo

```java
public abstract static class Component {
    public abstract String draw(Renderer renderer);
}

public static class Button extends Component {
    private String label;

    public Button(String label) {
        this.label = label;
    }

    @Override
    public String draw(Renderer renderer) {
        return renderer.renderButton(this);
    }

    public String getLabel() {
        return label;
    }
}

public static class TextBox extends Component {
    private String text;

    public TextBox(String text) {
        this.text = text;
    }

    @Override
    public String draw(Renderer renderer) {
        return renderer.renderTextBox(this);
    }

    public String getText() {
        return text;
    }
}

public abstract static class Renderer {
    public abstract String renderButton(Button b);
    public abstract String renderTextBox(TextBox t);
}

public static class HtmlRenderer extends Renderer {
    @Override
    public String renderButton(Button b) {
        return String.format("<button>%s</button>", b.getLabel());
    }

    @Override
    public String renderTextBox(TextBox t) {
        return String.format("<input value=\"%s\"/>", t.getText());
    }
}

public static class MarkdownRenderer extends Renderer {
    @Override
    public String renderButton(Button b) {
        return String.format("[%s]", b.getLabel());
    }

    @Override
    public String renderTextBox(TextBox t) {
        return String.format("_%s_", t.getText());
    }
}
```

## Ejercicio

Añade un nuevo componente (Checkbox) y un nuevo renderizador (PlainText).

## Problemas que encontrarás

Cada nuevo componente requiere añadir métodos a todos los renderizadores, y cada nuevo renderizador debe implementar métodos para todos los componentes existentes.
