# Parallel Inheritance Hierarchy

## Definición

Este smell ocurre cuando al cambiar una jerarquía de herencia, tienes que cambiar obligatoriamente otra jerarquía relacionada. Esto indicaría que ambas jerarquías están acopladas entre sí y no pueden evolucionar independientemente.

## Ejemplo

Agregar un nuevo componente de UI obliga a agregar métodos correspondientes en cada renderer, haciendo que ambas jerarquías tengan que crecer al unísono.

```go
type Renderer interface {
	RenderButton(b *Button) string
	RenderTextBox(t *TextBox) string
}

type Component interface {
	Draw(renderer Renderer) string
}

type Button struct {
	Label string
}

func (b *Button) Draw(renderer Renderer) string {
	return renderer.RenderButton(b)
}

type TextBox struct {
	Text string
}

func (t *TextBox) Draw(renderer Renderer) string {
	return renderer.RenderTextBox(t)
}

type HtmlRenderer struct{}

func (h *HtmlRenderer) RenderButton(b *Button) string {
	return fmt.Sprintf("<button>%s</button>", b.Label)
}

func (h *HtmlRenderer) RenderTextBox(t *TextBox) string {
	return fmt.Sprintf("<input value=\"%s\"/>", t.Text)
}

type MarkdownRenderer struct{}

func (m *MarkdownRenderer) RenderButton(b *Button) string {
	return fmt.Sprintf("[%s]", b.Label)
}

func (m *MarkdownRenderer) RenderTextBox(t *TextBox) string {
	return fmt.Sprintf("_%s_", t.Text)
}
```

## Ejercicio

Añade un componente `Image` que muestre una imagen.

## Problemas que encontrarás

Necesitarás añadir Image y renderImage a Renderer, e implementarlo en todos los renderers, mostrando cambios en paralelo.
