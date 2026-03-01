# Middleman

Intermediario.

## Definición

Ocurre cuando una clase realiza una única acción: delegar el trabajo a otra clase. Si una clase existe solo como un "pasamanos" hacia otro objeto, es posible que estemos ante una capa de abstracción innecesaria que oscurece al colaborador real.

## Descripción

**Middleman** es el opuesto de **Message Chains**. Mientras que Message Chains significa que un cliente conoce demasiado sobre la estructura interna, Middleman significa que una clase existe solo para ocultar otro objeto sin añadir valor real.

Este smell aparece cuando se aplica excesivamente el principio "Hide Delegate". La encapsulación es buena, pero cuando una clase se convierte en un simple envoltorio que solo pasa llamadas a otra clase sin añadir lógica, comportamiento o valor, entonces es un intermediario innecesario.

Señales de que es un problema:
- Más de la mitad de los métodos de la clase solo delegan
- La clase no tiene estado significativo propio
- La clase no añade validación, transformación o lógica
- Los clientes podrían hablar directamente con la clase delegada

## Síntomas

- La mayoría de los métodos de la clase solo delegan a otra clase
- La clase no tiene lógica de negocio propia
- La clase no transforma o valida datos antes de delegar
- Añadir funcionalidad requiere añadir métodos de paso en el middleman
- La clase existe solo para cumplir una interfaz o patrón
- Los nombres de los métodos son idénticos en el middleman y el delegado
- La clase no tiene estado significativo más allá de la referencia al delegado

## Ejemplo

```pseudocode
class Catalog {
  items: Map

  function add(id: string, name: string) {
    items.set(id, name)
  }

  function find(id: string): string {
    return items.get(id)
  }

  function list(): string[] {
    return Array.from(items.values())
  }
}

class Shop {
  catalog: Catalog  // Solo referencia al delegado

  // Todos los métodos solo delegan
  function add(id: string, name: string) {
    catalog.add(id, name)  // Sin lógica adicional
  }

  function find(id: string): string {
    return catalog.find(id)  // Sin transformación
  }

  function list(): string[] {
    return catalog.list()  // Sin valor agregado
  }
}

// Los clientes podrían usar Catalog directamente
```

## Ejercicio

Añade una funcionalidad `searchByPrefix` en `Catalog` y propágala a través de `Shop`.

## Problemas que encontrarás

Añadirás métodos a `Shop` que solo pasan a través hacia `Catalog`, fomentando la duplicación accidental y ocultando dónde vive el comportamiento real cuando necesites cambiarlo después.

## Proceso de Refactoring

### 1. Identificar el middleman
- Cuenta cuántos métodos solo delegan sin añadir valor
- Si más del 50% de los métodos son delegación pura, es un middleman
- Verifica si la clase tiene lógica de negocio propia

### 2. Evaluar si el middleman aporta valor
- ¿Añade validación o transformación?
- ¿Proporciona una interfaz simplificada sobre un subsistema complejo?
- ¿Coordina entre múltiples colaboradores?
- Si no aporta nada de esto, probablemente es innecesario

### 3. Aplicar Remove Middle Man
- Expón la clase delegada directamente a los clientes
- Los clientes llaman a `catalog.add()` en lugar de `shop.add()`
- Elimina los métodos de paso del middleman

### 4. Mover lógica útil antes de eliminar
- Si algunos métodos sí añaden valor, muévelos
- Extrae esa lógica a la clase delegada o a otra clase
- Solo elimina después de preservar lo valioso

### 5. Considerar Replace Delegation with Inheritance
- Si el middleman representa una especialización
- Considera hacer que herede o implemente la interfaz del delegado
- Pero cuidado: solo si la relación "is-a" tiene sentido

### 6. Mantener facade si simplifica complejidad
- Si el middleman oculta un subsistema complejo
- Y proporciona una interfaz simple y cohesiva
- Entonces no es un middleman problemático, es un Facade válido

## Técnicas de Refactoring Aplicables

- **Remove Middle Man**: Exponer la clase delegada directamente
- **Inline Method**: Eliminar métodos que solo delegan
- **Replace Delegation with Inheritance**: Si la relación "is-a" tiene sentido
- **Extract Class**: Si hay algo de lógica valiosa, sepárala primero

## Beneficios

- **Menos indirección**: Código más directo y fácil de seguir
- **Menos clases**: Reduce la complejidad innecesaria
- **Performance**: Elimina llamadas de método innecesarias
- **Claridad**: Obvio dónde está la funcionalidad real
- **Menos mantenimiento**: Menos código que mantener sincronizado
- **Navegación más fácil**: Ir directamente a la implementación real

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/couplers/middleman.ts) - [README](../../typescript/src/code-smells/couplers/middleman.readme.md)
- [Go](../../go/code_smells/couplers/middleman.go) - [README](../../go/code_smells/couplers/middleman.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/couplers/Middleman.java) - [README](../../java/src/main/java/com/refactoring/codesmells/couplers/Middleman.readme.md)
- [PHP](../../php/src/code-smells/couplers/Middleman.php) - [README](../../php/src/code-smells/couplers/Middleman.readme.md)
- [Python](../../python/src/code_smells/couplers/middleman.py) - [README](../../python/src/code_smells/couplers/middleman_readme.md)
- [C#](../../csharp/src/code-smells/couplers/Middleman.cs) - [README](../../csharp/src/code-smells/couplers/middleman.readme.md)

## Referencias en Español

- [Desacoplarse del sistema](https://franiglesias.github.io/decoupling_from_system/) - Cuándo usar intermediarios apropiadamente y cuándo son innecesarios

## Referencias

- [Refactoring Guru - Middle Man](https://refactoring.guru/smells/middle-man)
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
- Gang of Four - "Design Patterns" - Facade pattern (el uso correcto de intermediarios)
