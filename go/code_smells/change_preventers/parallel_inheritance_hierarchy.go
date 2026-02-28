package change_preventers

import "fmt"

// Code smell: Parallel Inheritance Hierarchy [Jerarquía de herencia paralela].
// Agregar un nuevo componente de UI obliga a agregar métodos correspondientes en cada renderer,
// haciendo que ambas jerarquías tengan que crecer al unísono.

// Ejercicio: Añade un componente Image que muestre una imagen.

// Necesitarás añadir Image y renderImage a Renderer, e implementarlo en todos los renderers,
// mostrando cambios en paralelo.

type Renderer interface {
	RenderButton(b *Button) string
	RenderTextBox(t *TextBox) string
}

type Component interface {
	Draw(renderer Renderer) string
}

// Button component
type Button struct {
	Label string
}

func (b *Button) Draw(renderer Renderer) string {
	return renderer.RenderButton(b)
}

// TextBox component
type TextBox struct {
	Text string
}

func (t *TextBox) Draw(renderer Renderer) string {
	return renderer.RenderTextBox(t)
}

// HtmlRenderer implementation
type HtmlRenderer struct{}

func (h *HtmlRenderer) RenderButton(b *Button) string {
	return fmt.Sprintf("<button>%s</button>", b.Label)
}

func (h *HtmlRenderer) RenderTextBox(t *TextBox) string {
	return fmt.Sprintf("<input value=\"%s\"/>", t.Text)
}

// MarkdownRenderer implementation
type MarkdownRenderer struct{}

func (m *MarkdownRenderer) RenderButton(b *Button) string {
	return fmt.Sprintf("[%s]", b.Label)
}

func (m *MarkdownRenderer) RenderTextBox(t *TextBox) string {
	return fmt.Sprintf("_%s_", t.Text)
}

func DemoParallelHierarchy() []string {
	comps := []Component{
		&Button{Label: "Save"},
		&TextBox{Text: "name"},
	}
	renderer := &HtmlRenderer{}
	results := make([]string, len(comps))
	for i, c := range comps {
		results[i] = c.Draw(renderer)
	}
	return results
}
