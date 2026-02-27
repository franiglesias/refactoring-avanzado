# Parallel Inheritance Hierarchy

## Definición

Este smell ocurre cuando al cambiar una jerarquía de herencia, tienes que cambiar obligatoriamente otra jerarquía relacionada. Esto indicaría que ambas jerarquías están acopladas entre sí y no pueden evolucionar independientemente.

## Ejemplo

Agregar un nuevo componente de UI obliga a agregar métodos correspondientes en cada renderer, haciendo que ambas jerarquías tengan que crecer al unísono.

```typescript
export abstract class Component {
  abstract draw(renderer: Renderer): string
}

export class Button extends Component {
  constructor(public label: string) {
    super()
  }

  draw(renderer: Renderer): string {
    return renderer.renderButton(this)
  }
}

export class TextBox extends Component {
  constructor(public text: string) {
    super()
  }

  draw(renderer: Renderer): string {
    return renderer.renderTextBox(this)
  }
}

export abstract class Renderer {
  abstract renderButton(b: Button): string

  abstract renderTextBox(t: TextBox): string
}

export class HtmlRenderer extends Renderer {
  renderButton(b: Button): string {
    return `<button>${b.label}</button>`
  }

  renderTextBox(t: TextBox): string {
    return `<input value="${t.text}"/>`
  }
}

export class MarkdownRenderer extends Renderer {
  renderButton(b: Button): string {
    return `[${b.label}]`
  }

  renderTextBox(t: TextBox): string {
    return `_${t.text}_`
  }
}

export function demoParallelHierarchy(): string[] {
  const comps: Component[] = [new Button('Save'), new TextBox('name')]
  const renderer = new HtmlRenderer()
  return comps.map((c) => c.draw(renderer))
}
```

## Ejercicio

Añade un componente `Image` que muestre una imagen.

## Problemas que encontrarás

Necesitarás añadir Image y renderImage a Renderer, e implementarlo en todos los renderers, mostrando cambios en paralelo.
