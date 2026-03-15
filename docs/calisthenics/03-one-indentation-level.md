# Un Solo Nivel de Indentación

Regla 3 de Object Calisthenics

## Definición

Cada método o función debería contener un único nivel de indentación. Esto significa que no debe haber estructuras de control anidadas (loops dentro de loops, condicionales dentro de loops, etc.).

## Descripción

**Múltiples niveles de indentación** son un indicador claro de que un método está haciendo demasiado. Cada nivel adicional de indentación representa:

1. **Una responsabilidad adicional**: El método mezcla diferentes niveles de abstracción
2. **Mayor complejidad ciclomática**: Más caminos de ejecución posibles
3. **Dificultad para comprender**: Hay que mantener múltiples contextos mentales simultáneamente
4. **Código difícil de testear**: Necesitas configurar estados complejos para alcanzar ramas profundas

La regla de "un solo nivel de indentación" fuerza la **composición de métodos pequeños**, cada uno operando en un único nivel de abstracción. En lugar de tener un método que hace todo, tienes un método que **orquesta** llamadas a otros métodos especializados.

Beneficios inmediatos:
- **Claridad**: Cada método cuenta una historia simple
- **Reusabilidad**: Métodos pequeños pueden reutilizarse
- **Testing**: Cada método puede testearse aisladamente
- **Mantenibilidad**: Cambios localizados en métodos específicos

Esta regla está estrechamente relacionada con el **principio de composición** y el patrón **Composed Method**: un método debe estar compuesto enteramente de llamadas a otros métodos del mismo nivel de abstracción.

## Síntomas

- Métodos con más de un nivel de llaves/indentación
- Loops anidados (for dentro de for, while dentro de for, etc.)
- Condicionales dentro de loops
- Loops dentro de condicionales dentro de loops
- Scroll vertical necesario para ver todo el método
- Dificultad para nombrar el método (hace demasiadas cosas)
- Variables temporales que solo se usan en un nivel de anidación
- Comentarios que separan "secciones" dentro del método
- Más de 3-4 niveles de indentación visibles

## Ejemplo

### Antes (Violación)

```pseudocode
function processOrders(orders) {
  for each order in orders {
    if (order.isValid()) {
      for each item in order.items {
        if (item.inStock()) {
          if (item.quantity > 0) {
            price = item.price * item.quantity
            if (order.hasDiscount()) {
              discount = price * order.discountRate
              price = price - discount
            }
            order.total = order.total + price
          }
        } else {
          log("Item out of stock: " + item.name)
        }
      }
      if (order.total > 0) {
        saveOrder(order)
        sendConfirmation(order.customerEmail)
      }
    } else {
      log("Invalid order: " + order.id)
    }
  }
}
```

**Problemas**:
- 5 niveles de indentación (8 puntos de complejidad ciclomática)
- Mezcla validación, cálculo, persistencia y notificación
- Imposible entender sin leer todo el método
- Difícil de testear (muchos casos edge)

### Después (Cumplimiento)

```pseudocode
function processOrders(orders) {
  for each order in orders {
    processOrder(order)
  }
}

function processOrder(order) {
  if (not order.isValid()) {
    logInvalidOrder(order)
    return
  }

  calculateOrderTotal(order)

  if (order.total > 0) {
    saveOrder(order)
    sendConfirmation(order.customerEmail)
  }
}

function calculateOrderTotal(order) {
  order.total = 0
  for each item in order.items {
    addItemToTotal(order, item)
  }
}

function addItemToTotal(order, item) {
  if (not item.inStock()) {
    logOutOfStock(item)
    return
  }

  if (item.quantity <= 0) {
    return
  }

  price = calculateItemPrice(order, item)
  order.total = order.total + price
}

function calculateItemPrice(order, item) {
  price = item.price * item.quantity

  if (order.hasDiscount()) {
    price = applyDiscount(price, order.discountRate)
  }

  return price
}

function applyDiscount(price, discountRate) {
  discount = price * discountRate
  return price - discount
}

function logInvalidOrder(order) {
  log("Invalid order: " + order.id)
}

function logOutOfStock(item) {
  log("Item out of stock: " + item.name)
}
```

**Mejoras**:
- Cada método tiene máximo 1 nivel de indentación
- Cada método tiene una única responsabilidad clara
- Nombres descriptivos cuentan una historia
- Fácil de testear cada comportamiento aisladamente
- Fácil de modificar sin romper otras partes

La complejidad ciclomática de este código se calcula así:

| Función | Complejidad Ciclomática | Puntos de Decisión |
|---------|------------------------|-------------------|
| `processOrders(orders)` | 2 | 1 bucle for |
| `processOrder(order)` | 3 | 2 condicionales if |
| `calculateOrderTotal(order)` | 2 | 1 bucle for |
| `addItemToTotal(order, item)` | 3 | 2 condicionales if |
| `calculateItemPrice(order, item)` | 2 | 1 condicional if |
| `applyDiscount(price, discountRate)` | 1 | 0 (función lineal) |
| `logInvalidOrder(order)` | 1 | 0 (función lineal) |
| `logOutOfStock(item)` | 1 | 0 (función lineal) |
| **TOTAL** | **15** | **7 decisiones + 8 funciones** |

Si bien el total es de 15 puntos, puedes ver que la máxima complejidad es 3.

## Ejercicio

**Tarea**: En el código proporcionado en tu lenguaje, refactoriza los métodos para que cada uno tenga un único nivel de indentación. Extrae comportamiento a métodos auxiliares bien nombrados.

**Criterios de éxito**:
1. Ningún método tiene más de un nivel de indentación
2. Cada método tiene un nombre descriptivo que explica su responsabilidad
3. No hay loops anidados
4. No hay condicionales dentro de loops (o viceversa)
5. El código se lee como una narrativa de alto nivel

## Problemas que Encontrarás

### 1. "Método de un solo loop no puede tener un nivel"

Si el método solo tiene un loop simple sin lógica adicional, está bien. La regla se aplica cuando hay **anidación** (loop dentro de loop, condicional dentro de loop).

### 2. Explosión de métodos pequeños

Tendrás muchos métodos pequeños. Esto es **correcto**. Cada método debe hacer una cosa. Si tienes miedo de "demasiados métodos", probablemente tu clase tiene demasiadas responsabilidades (violación de SRP).

### 3. Métodos privados "triviales"

Métodos como `logOutOfStock(item)` parecen triviales, pero:
- Dan nombre a conceptos
- Son puntos de extensión futuros
- Facilitan el testing (puedes mockear la acción de logging)
- Hacen el código autoexplicativo

### 4. Performance por llamadas a métodos

Las llamadas a métodos tienen un costo, pero:
- Es despreciable en la mayoría de aplicaciones
- Los compiladores modernos hacen inline de métodos pequeños automáticamente
- La ganancia en mantenibilidad supera infinitamente cualquier pérdida de performance
- Si tienes un cuello de botella real, optimiza ese punto específico (no optimización prematura)

### 5. No saber dónde poner métodos auxiliares

Si los métodos extraídos:
- Usan datos principalmente de un objeto → Van en ese objeto
- Son operaciones de dominio → Considera crear una clase de servicio
- Son utilidades genéricas → Considera un módulo de utilidades
- Usan datos del método padre → Quizás necesitas un **Method Object** (extraer a una clase)

## Proceso de Aplicación

### 1. Identificar métodos con múltiples niveles

- Busca métodos con anidación visible
- Marca cualquier método que requiera scroll para verse completo
- Prioriza los que tienen más niveles de indentación

### 2. Analizar responsabilidades

Para cada método violador:
- Identifica las diferentes "cosas" que hace
- Busca bloques de código que puedan nombrarse
- Marca secciones separadas por comentarios (candidatos a extracción)

### 3. Extraer nivel más interno primero

Estrategia bottom-up:
- Identifica el bloque de código más profundamente anidado
- Extráelo a un método con nombre descriptivo
- Reemplaza el bloque con la llamada al nuevo método
- Repite hasta que solo quede un nivel

### 4. Aplicar guard clauses

Para reducir anidación de condicionales:
- Invierte validaciones para retornar temprano
- Mueve el "camino feliz" al nivel base de indentación
- Combina con regla #2 (No Else)

### 5. Extraer loops internos

```pseudocode
// Antes
for each order in orders {
  for each item in order.items {
    processItem(item)
  }
}

// Después
for each order in orders {
  processOrderItems(order)
}

function processOrderItems(order) {
  for each item in order.items {
    processItem(item)
  }
}
```

### 6. Renombrar para claridad

Los métodos extraídos deben tener nombres que:
- Describan qué hacen, no cómo
- Estén al mismo nivel de abstracción que el método padre
- Lean naturalmente cuando se usan

### 7. Reorganizar si es necesario

Si después de la extracción tienes muchos métodos auxiliares:
- Considera si pertenecen a otra clase
- Evalúa si necesitas un **Method Object**
- Verifica que tu clase no tenga demasiadas responsabilidades (SRP)

## Técnicas de Refactoring Aplicables

- **Extract Method**: Técnica principal para reducir indentación
- **Replace Nested Conditional with Guard Clauses**: Aplanar condicionales anidados
- **Decompose Conditional**: Extraer condiciones complejas
- **Replace Loop with Pipeline**: En lenguajes con funciones de orden superior (map, filter, reduce)
- **Replace Method with Method Object**: Para métodos muy complejos con muchas variables locales
- **Move Method**: Mover métodos auxiliares a la clase apropiada
- **Inline Temp**: Eliminar variables temporales innecesarias

## Beneficios

### 1. Legibilidad Dramáticamente Mejorada

Cada método se lee como una secuencia de pasos de alto nivel, sin detalles de implementación que distraigan.

### 2. Reusabilidad

Métodos pequeños y específicos pueden reutilizarse en otros contextos.

### 3. Testing Simplificado

Cada método puede testearse independientemente con casos simples, en lugar de configurar estados complejos.

### 4. Debugging Más Fácil

Stack traces más descriptivos, y es más fácil poner breakpoints en comportamiento específico.

### 5. Reduce Complejidad Ciclomática

Menos caminos de ejecución por método = menor complejidad cognitiva.

### 6. Facilita Code Review

Revisores pueden entender métodos individuales sin contexto completo del sistema.

### 7. Mejor Organización del Código

Fuerza a pensar en responsabilidades y abstracciones apropiadas.

### 8. Documentación Implícita

Los nombres de métodos documentan el propósito sin necesidad de comentarios.

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/calisthenics-exercises/one-indentation-level.ts)
- [Go](../../go/calisthenics_exercises/01_one_level_indentation.go)
- [Java](../../java/src/main/java/com/refactoring/calisthenics/OneLevelIndentation.java)
- [PHP](../../php/src/calisthenics-exercises/OneIndentationLevel.php)
- [Python](../../python/src/calisthenics_exercises/one_indentation_level.py)
- [C#](../../csharp/src/calisthenics-exercises/OneIndentationLevel.cs)

## Referencias en Español

- [Métodos largos](https://franiglesias.github.io/long-method/) - Incluye técnicas para reducir indentación
- [Ejercicio de refactor (2) Extraer hasta la última gota](https://franiglesias.github.io/ejercicio-de-refactor-2/) - Práctica de extracción de métodos
- [Object Calisthenics para adelgazar las clases](https://franiglesias.github.io/calistenics-and-small-classes/) - Incluye técnicas para mantener métodos simples

## Referencias

- **"Refactoring: Improving the Design of Existing Code"** - Martin Fowler - Extract Method, Compose Method
- **"Clean Code"** - Robert C. Martin - Capítulo sobre funciones: "Do One Thing"
- [Object Calisthenics - William Durand](https://williamdurand.fr/2013/06/03/object-calisthenics/) - Regla #3
- **"Smalltalk Best Practice Patterns"** - Kent Beck - Composed Method pattern
- [Flattening Arrow Code](https://blog.codinghorror.com/flattening-arrow-code/) - Jeff Atwood sobre reducir anidación
- [Cyclomatic Complexity](https://en.wikipedia.org/wiki/Cyclomatic_complexity) - Métrica de complejidad relacionada
