# Parallel Inheritance Hierarchy

Jerarquía de herencia paralela.

## Definición

Este smell ocurre cuando al cambiar una jerarquía de herencia, tienes que cambiar obligatoriamente otra jerarquía relacionada. Esto indica que ambas jerarquías están acopladas entre sí y no pueden evolucionar independientemente.

## Descripción

**Parallel Inheritance Hierarchy** es un caso especial de **Shotgun Surgery** aplicado a jerarquías de clases. Ocurre cuando dos jerarquías de herencia crecen en paralelo: cada vez que añades una subclase en una jerarquía, debes añadir una subclase correspondiente en la otra.

Este problema aparece frecuentemente cuando se intenta separar responsabilidades usando herencia, pero las responsabilidades están demasiado acopladas. Por ejemplo:
- Componentes de UI y sus renderizadores
- Modelos de dominio y sus serializadores
- Comandos y sus handlers
- Entidades y sus validadores

El resultado es duplicación de estructura y un esfuerzo de mantenimiento multiplicado.

## Síntomas

- Dos jerarquías de clases con estructuras similares
- Los nombres de las clases en ambas jerarquías son paralelos (Button/ButtonRenderer)
- Añadir una subclase en una jerarquía requiere añadir otra en la paralela
- Las subclases en ambas jerarquías tienen una relación uno-a-uno
- Métodos en una jerarquía referencian específicamente clases de la otra
- Cambios en una jerarquía siempre requieren cambios en la otra

## Ejemplo

```pseudocode
// Primera jerarquía: Componentes
abstract class Component {
  abstract function draw(renderer)
}

class Button extends Component {
  label: string

  function draw(renderer) {
    return renderer.renderButton(this)
  }
}

class TextBox extends Component {
  text: string

  function draw(renderer) {
    return renderer.renderTextBox(this)
  }
}

// Segunda jerarquía paralela: Renderers
abstract class Renderer {
  abstract function renderButton(button)
  abstract function renderTextBox(textbox)
}

class HtmlRenderer extends Renderer {
  function renderButton(button) {
    return "<button>" + button.label + "</button>"
  }

  function renderTextBox(textbox) {
    return "<input value='" + textbox.text + "'/>"
  }
}

class MarkdownRenderer extends Renderer {
  function renderButton(button) {
    return "[" + button.label + "]"
  }

  function renderTextBox(textbox) {
    return "_" + textbox.text + "_"
  }
}
```

## Ejercicio

Añade un componente `Image` que muestre una imagen.

## Problemas que encontrarás

Necesitarás añadir `Image` a la jerarquía de componentes y `renderImage` a la clase base `Renderer`, e implementarlo en todos los renderers existentes (HtmlRenderer, MarkdownRenderer), mostrando cambios en paralelo.

## Proceso de Refactoring

### 1. Identificar el patrón de duplicación
- Mapea las relaciones entre ambas jerarquías
- Identifica qué clases dependen de cuáles
- Verifica si realmente necesitas ambas jerarquías

### 2. Considerar el patrón Visitor
- Si las operaciones (renderización, validación) varían frecuentemente
- Implementa el patrón Visitor para externalizar las operaciones
- Los componentes aceptan visitantes en lugar de conocer los renderers

### 3. Usar composición en lugar de herencia
- Reemplaza una de las jerarquías con estrategias o plugins
- Ejemplo: en lugar de subclases de Renderer, usa objetos de estrategia
- Los componentes delegan a estrategias intercambiables

### 4. Aplicar el patrón Bridge
- Separa la abstracción de la implementación
- Permite que ambas dimensiones varíen independientemente
- Ejemplo: Component (abstracción) y RenderingStrategy (implementación)

### 5. Colapsar una jerarquía
- Si una jerarquía solo existe para soportar la otra
- Considera eliminarla y usar configuración o datos
- Ejemplo: usar una tabla de mapeo en lugar de subclases

### 6. Usar reflexión o metadatos
- En lenguajes dinámicos, considera usar nombres convencionales
- Ejemplo: `render("button", data)` en lugar de `renderButton(button)`
- Reduce el acoplamiento estructural entre jerarquías

## Técnicas de Refactoring Aplicables

- **Move Method**: Mover comportamiento a una sola jerarquía
- **Move Field**: Consolidar datos en un solo lugar
- **Replace Inheritance with Delegation**: Usar composición en lugar de herencia
- **Bridge Pattern**: Desacoplar abstracción de implementación
- **Visitor Pattern**: Externalizar operaciones sobre estructuras
- **Strategy Pattern**: Reemplazar subclases con estrategias configurables

## Beneficios

- **Evolución independiente**: Las jerarquías pueden cambiar sin afectarse
- **Menos duplicación**: No necesitas replicar estructura
- **Añadir características es más fácil**: Un solo punto de cambio
- **Mejor testabilidad**: Puedes testear cada jerarquía independientemente
- **Flexibilidad**: Mezclar y combinar implementaciones sin restricciones
- **Cumplimiento de OCP**: Abierto a extensión, cerrado a modificación

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/change-preventers/parallel-inheritance-hierarchy.ts) - [README](../../typescript/src/code-smells/change-preventers/parallel-inheritance-hierarchy.readme.md)
- [Go](../../go/code_smells/change_preventers/parallel_inheritance_hierarchy.go) - [README](../../go/code_smells/change_preventers/parallel_inheritance_hierarchy.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/changepreventers/ParallelInheritanceHierarchy.java) - [README](../../java/src/main/java/com/refactoring/codesmells/changepreventers/ParallelInheritanceHierarchy.readme.md)
- [PHP](../../php/src/code-smells/change-preventers/ParallelInheritanceHierarchy.php) - [README](../../php/src/code-smells/change-preventers/ParallelInheritanceHierarchy.readme.md)
- [Python](../../python/src/code_smells/change_preventers/parallel_inheritance_hierarchy.py) - [README](../../python/src/code_smells/change_preventers/parallel_inheritance_hierarchy_readme.md)
- [C#](../../csharp/src/code-smells/change-preventers/ParallelInheritanceHierarchy.cs) - [README](../../csharp/src/code-smells/change-preventers/parallel-inheritance-hierarchy.readme.md)

## Referencias en Español

- [Polimorfismo y extensibilidad de objetos](https://franiglesias.github.io/polimorfismo-y-extensibilidad-de-objetos/) - Uso de polimorfismo para evitar jerarquías paralelas y mejorar la extensibilidad

## Referencias

- [Refactoring Guru - Parallel Inheritance Hierarchies](https://refactoring.guru/smells/parallel-inheritance-hierarchies)
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
- Gang of Four - "Design Patterns" - Bridge, Visitor, Strategy patterns
