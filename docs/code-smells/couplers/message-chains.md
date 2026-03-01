# Message Chains

Cadenas de mensajes.

## Definición

La navegación profunda por grafos de objetos acopla a los clientes a la estructura de los intermediarios y conduce a código frágil.

## Descripción

**Message Chains** ocurren cuando un cliente solicita un objeto, luego solicita otro objeto de ese objeto, y así sucesivamente, formando una larga cadena de llamadas. Este patrón se ve como múltiples llamadas a getters encadenadas: `a.getB().getC().getD().doSomething()`.

El problema principal es el acoplamiento extremo: el cliente debe conocer la estructura completa del grafo de objetos para navegar hasta el objeto que finalmente necesita. Si la estructura cambia en cualquier punto de la cadena, todos los clientes deben actualizarse.

Este smell viola el principio de Demeter (Law of Demeter): "Solo habla con tus amigos inmediatos". Las cadenas de mensajes crean dependencias transitivas que hacen el código frágil y difícil de mantener.

## Síntomas

- Múltiples llamadas a métodos getter encadenadas
- Navegación profunda a través de objetos: `obj.getA().getB().getC()`
- El cliente conoce la estructura interna del grafo de objetos
- Cambios en la estructura intermedia rompen múltiples clientes
- Código que parece "caminar" por la estructura de datos
- Tests que requieren configurar cadenas completas de objetos mock

## Ejemplo

```pseudocode
class Level2 {
  value: number

  function getValue(): number {
    return value
  }
}

class Level1 {
  next: Level2

  function getNext(): Level2 {
    return next
  }
}

class Root {
  next: Level1

  function getNext(): Level1 {
    return next
  }
}

// Uso con cadena de mensajes
function readDeep(root: Root): number {
  return root.getNext().getNext().getValue()
  // El cliente conoce 3 niveles de estructura
}
```

## Ejercicio

Inserta un nuevo `Level` entre `Root` y `Level1`, o reubica `getValue`.

## Problemas que encontrarás

Observa cómo cada cliente que usa `root.getNext().getNext().getValue()` debe cambiar, revelando cómo las cadenas de mensajes vuelven costosas refactorizaciones simples.

## Proceso de Refactoring

### 1. Identificar las cadenas
- Busca patrones de múltiples llamadas a getters encadenadas
- Identifica qué información realmente necesita el cliente
- Mapea todos los lugares donde aparece la misma cadena

### 2. Aplicar Hide Delegate
- Añade un método en el objeto raíz que encapsule la navegación
- El método delega internamente pero oculta la estructura
- Ejemplo: `root.getValue()` en lugar de `root.getNext().getNext().getValue()`

### 3. Extraer métodos de conveniencia
- Crea métodos en objetos intermedios para acortar cadenas
- Ejemplo: `level1.getValue()` que delega a `level2.getValue()`
- Cada objeto ofrece operaciones de alto nivel

### 4. Considerar si la estructura es correcta
- A veces las cadenas indican que falta una abstracción
- Pregúntate si realmente necesitas toda esa estructura
- Considera aplanar o simplificar el grafo de objetos

### 5. Usar el patrón Facade
- Si múltiples clientes hacen navegaciones similares
- Crea una Facade que proporcione una interfaz simple
- La Facade encapsula toda la navegación compleja

### 6. Aplicar Law of Demeter
- Un método solo debe llamar a:
  - Métodos del objeto mismo
  - Métodos de parámetros directos
  - Métodos de objetos que crea
  - Métodos de campos directos
- No navegues a través de objetos retornados

## Técnicas de Refactoring Aplicables

- **Hide Delegate**: Ocultar la navegación entre objetos
- **Extract Method**: Crear métodos que encapsulen la navegación
- **Move Method**: Mover comportamiento más cerca de los datos
- **Inline Method**: Si la cadena es corta y usada una vez
- **Replace Query with Parameter**: Si el objeto final puede pasarse directamente

## Beneficios

- **Menos acoplamiento**: Los clientes no conocen la estructura interna
- **Código más robusto**: Cambios en estructura no afectan clientes
- **Mejor encapsulación**: La navegación está oculta
- **Más fácil de refactorizar**: Puedes reorganizar objetos internamente
- **Código más legible**: Llamadas simples en lugar de cadenas complejas
- **Tests más simples**: No necesitas configurar estructuras complejas

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/couplers/message-chains.ts) - [README](../../typescript/src/code-smells/couplers/message-chains.readme.md)
- [Go](../../go/code_smells/couplers/message_chains.go) - [README](../../go/code_smells/couplers/message_chains.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/couplers/MessageChains.java) - [README](../../java/src/main/java/com/refactoring/codesmells/couplers/MessageChains.readme.md)
- [PHP](../../php/src/code-smells/couplers/MessageChains.php) - [README](../../php/src/code-smells/couplers/MessageChains.readme.md)
- [Python](../../python/src/code_smells/couplers/message_chains.py) - [README](../../python/src/code_smells/couplers/message_chains_readme.md)
- [C#](../../csharp/src/code-smells/couplers/MessageChains.cs) - [README](../../csharp/src/code-smells/couplers/message-chains.readme.md)

## Referencias en Español

- [Refactor cotidiano (6). Tell, Don't Ask y Ley de Demeter](https://franiglesias.github.io/everyday-refactor-6/) - Explicación de la Ley de Demeter para evitar cadenas de mensajes
- [Más allá de SOLID, los cimientos (Law of Demeter)](https://franiglesias.github.io/beyond-solid-2/) - Fundamentos del principio de mínimo conocimiento

## Referencias

- [Refactoring Guru - Message Chains](https://refactoring.guru/smells/message-chains)
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
- The Pragmatic Programmer - Law of Demeter (Principle of Least Knowledge)
