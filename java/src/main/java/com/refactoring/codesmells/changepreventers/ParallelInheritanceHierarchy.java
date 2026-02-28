package com.refactoring.codesmells.changepreventers;

import java.util.List;
import java.util.stream.Collectors;

/**
 * Code Smell: Parallel Inheritance Hierarchy [Jerarquía de herencia paralela]
 * Cada vez que añades una subclase de Component, también debes añadir una subclase de Renderer.
 * Las jerarquías Component y Renderer crecen en paralelo.
 *
 * Añadir un nuevo componente (ej: Checkbox) requiere también añadir un nuevo renderizador.
 */
public class ParallelInheritanceHierarchy {

    public static void main(String[] args) {
        List<String> result = demoParallelHierarchy();
        result.forEach(System.out::println);
    }

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

    public static List<String> demoParallelHierarchy() {
        List<Component> comps = List.of(new Button("Save"), new TextBox("name"));
        Renderer renderer = new HtmlRenderer();
        return comps.stream()
            .map(c -> c.draw(renderer))
            .collect(Collectors.toList());
    }
}
