# No Usar ELSE

Regla 2 de Object Calisthenics

## Definición

El código no debe contener cláusulas `else`. En su lugar, debe usar retornos tempranos, cláusulas de guarda, polimorfismo u otras técnicas que eviten la bifurcación explícita del flujo de control.

## Descripción

Las cláusulas **else** parecen inofensivas, pero introducen varios problemas sutiles:

1. **Ocultan el flujo principal**: El código más importante queda anidado dentro de bloques condicionales
2. **Aumentan la complejidad ciclomática**: Cada else añade un camino de ejecución alternativo
3. **Dificultan la comprensión**: Requieren mantener múltiples contextos mentales simultáneamente
4. **Esconden reglas de negocio**: Las condiciones importantes quedan enterradas en estructuras if-else
5. **Complican el testing**: Más caminos de ejecución = más casos de test necesarios

El problema fundamental es que **else invierte la lógica natural de validación**. En lugar de validar condiciones y fallar rápido, procesamos el camino feliz dentro de un bloque if, y el manejo de errores o casos especiales queda en el else.

Eliminar else fuerza a:
- Expresar las validaciones claramente al inicio (guard clauses)
- Hacer explícitas las reglas de negocio
- Reducir la anidación de código
- Separar caminos de ejecución incompatibles (polimorfismo)

## Síntomas

- Bloques `if-else` anidados (if-else dentro de otro if-else)
- Métodos con múltiples niveles de indentación por condicionales
- Dificultad para nombrar métodos porque manejan múltiples casos
- Código donde el "camino feliz" está en el `if` y los errores en el `else`
- Lógica duplicada entre bloques `if` y `else`
- Estructuras `if-else-if-else` en cadena
- Tests que necesitan configurar diferentes contextos para cada rama
- Comentarios como "// caso especial" o "// caso normal" separando bloques

## Ejemplo

### Antes (Violación)

```pseudocode
function processPayment(user, amount) {
  if (user is not null) {
    if (user.isActive) {
      if (amount > 0) {
        if (user.balance >= amount) {
          user.balance = user.balance - amount
          saveUser(user)
          return "Payment processed"
        } else {
          return "Insufficient funds"
        }
      } else {
        return "Invalid amount"
      }
    } else {
      return "User is not active"
    }
  } else {
    return "User not found"
  }
}

function calculateDiscount(customer) {
  if (customer.type == "VIP") {
    return 20
  } else {
    if (customer.type == "PREMIUM") {
      return 10
    } else {
      if (customer.type == "REGULAR") {
        return 5
      } else {
        return 0
      }
    }
  }
}
```

### Después (Cumplimiento)

```pseudocode
// Opción 1: Guard Clauses (cláusulas de guarda)
function processPayment(user, amount) {
  if (user is null) {
    return "User not found"
  }

  if (not user.isActive) {
    return "User is not active"
  }

  if (amount <= 0) {
    return "Invalid amount"
  }

  if (user.balance < amount) {
    return "Insufficient funds"
  }

  user.balance = user.balance - amount
  saveUser(user)
  return "Payment processed"
}

// Opción 2: Polimorfismo (para casos como calculateDiscount)
interface CustomerType {
  method getDiscount()
}

class VIPCustomer implements CustomerType {
  method getDiscount() {
    return 20
  }
}

class PremiumCustomer implements CustomerType {
  method getDiscount() {
    return 10
  }
}

class RegularCustomer implements CustomerType {
  method getDiscount() {
    return 5
  }
}

class GuestCustomer implements CustomerType {
  method getDiscount() {
    return 0
  }
}

// O con patrón Strategy/Command
class DiscountCalculator {
  discountMap = {
    "VIP": 20,
    "PREMIUM": 10,
    "REGULAR": 5,
    "GUEST": 0
  }

  method calculateDiscount(customerType) {
    return discountMap.get(customerType, 0)
  }
}
```

**Diferencias clave**:
- **Guard clauses al inicio**: Validaciones explícitas que retornan temprano
- **Camino feliz sin anidar**: El flujo principal está al nivel de indentación base
- **Sin else**: Cada validación retorna inmediatamente si falla
- **Legibilidad mejorada**: Se lee de arriba a abajo, sin saltos mentales
- **Polimorfismo para variaciones**: Casos mutuamente excluyentes se separan en tipos

## Ejercicio

**Tarea**: En el código proporcionado en tu lenguaje, elimina todas las cláusulas `else` usando guard clauses, early returns, polimorfismo o estructuras de datos según corresponda.

**Criterios de éxito**:
1. No quedan cláusulas `else` en el código
2. El flujo principal está al nivel de indentación más bajo
3. Las validaciones están explícitas al inicio de cada función
4. El código se lee de arriba a abajo sin saltos
5. La complejidad ciclomática se ha reducido

## Problemas que Encontrarás

### 1. Sensación de "código menos eficiente"

Múltiples `return` pueden parecer menos eficientes que un solo punto de salida, pero:
- La diferencia de rendimiento es imperceptible
- La ganancia en legibilidad es enorme
- Los compiladores modernos optimizan ambos casos igual

### 2. Violación de "single exit point"

La regla de "un solo punto de salida" es obsoleta. Proviene de lenguajes sin garbage collection donde había que liberar recursos manualmente. Los retornos tempranos son más claros.

### 3. Casos else legítimos con operador ternario

```pseudocode
max = (a > b) ? a : b
```

El operador ternario es aceptable para asignaciones simples. La regla aplica principalmente a bloques de código con lógica, no a expresiones simples.

### 4. if-else simétricos

```pseudocode
if (condition) {
  doSomething()
} else {
  doSomethingElse()
}
```

Si ambos bloques son igualmente importantes y no hay validación, considera:
- Invertir la condición para que el caso más común vaya primero
- Usar polimorfismo si representa variaciones de comportamiento
- Dividir en dos funciones separadas si son responsabilidades distintas

### 5. Cadenas if-else-if complejas

Son candidatos perfectos para **polimorfismo**, **patrón Strategy**, o **estructuras de datos** (diccionarios/mapas).

## Proceso de Aplicación

### 1. Identificar todas las cláusulas else

- Busca todas las apariciones de `else` en el código
- Incluye `else if` y operadores ternarios complejos
- Marca cada una con su contexto

### 2. Clasificar por tipo

**Tipo A - Validaciones**: if valida, else maneja error
→ Usar **guard clauses**

**Tipo B - Bifurcaciones simétricas**: if y else hacen cosas diferentes pero válidas
→ Considerar **polimorfismo** o **dividir el método**

**Tipo C - Selección entre múltiples opciones**: if-else-if-else en cadena
→ Usar **polimorfismo**, **patrón Strategy** o **estructura de datos**

### 3. Aplicar Guard Clauses (Tipo A)

```pseudocode
// Antes
if (valid) {
  process()
} else {
  handleError()
}

// Después
if (not valid) {
  handleError()
  return
}
process()
```

### 4. Aplicar Polimorfismo (Tipos B y C)

- Identifica el criterio de selección (tipo, estado, rol, etc.)
- Crea una interfaz o clase base
- Extrae cada rama a una implementación concreta
- Reemplaza el if-else con dispatch polimórfico

### 5. Usar Estructuras de Datos (Tipo C simple)

Si las ramas solo retornan valores sin lógica compleja:

```pseudocode
// Antes
if (type == "A") return 10
else if (type == "B") return 20
else return 0

// Después
values = {"A": 10, "B": 20}
return values.get(type, 0)
```

### 6. Verificar legibilidad

- Lee el código de arriba a abajo
- Debe fluir naturalmente: validaciones primero, procesamiento después
- Las reglas de negocio deben ser evidentes

## Técnicas de Refactoring Aplicables

- **Replace Nested Conditional with Guard Clauses**: Extraer validaciones al inicio
- **Replace Conditional with Polymorphism**: Usar herencia/interfaces para variaciones
- **Introduce Null Object**: Evitar validaciones null con objeto que representa ausencia
- **Replace Conditional with Strategy**: Encapsular variaciones de algoritmo
- **Extract Method**: Separar lógica de cada rama en métodos propios
- **Decompose Conditional**: Extraer condiciones complejas a métodos con nombres descriptivos

## Beneficios

### 1. Legibilidad Mejorada

El código se lee linealmente de arriba a abajo, sin saltos mentales entre ramas.

### 2. Camino Feliz Visible

El flujo principal de ejecución está al nivel más bajo de indentación, siendo inmediatamente visible.

### 3. Validaciones Explícitas

Las guard clauses hacen obvias todas las precondiciones que deben cumplirse.

### 4. Menos Complejidad Ciclomática

Menos caminos de ejecución = código más simple de entender y testear.

### 5. Mejor Testing

Cada guard clause puede testearse independientemente, y el camino feliz es un caso separado.

### 6. Facilita Refactoring Futuro

Código sin else anidado es más fácil de extraer, mover y reorganizar.

### 7. Fuerza Mejor Diseño

Eliminar else te obliga a pensar en polimorfismo y separación de responsabilidades.

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/calisthenics-exercises/dont-use-else.ts)
- [Go](../../go/calisthenics_exercises/02_no_else.go)
- [Java](../../java/src/main/java/com/refactoring/calisthenics/NoElse.java)
- [PHP](../../php/src/calisthenics-exercises/DontUseElse.php)
- [Python](../../python/src/calisthenics_exercises/dont_use_else.py)
- [C#](../../csharp/src/calisthenics-exercises/DontUseElse.cs)

## Referencias en Español

- [Refactor cotidiano (3): extraer para aclarar, parte 1](https://franiglesias.github.io/everyday-refactor-3/) - Sobre eliminación de else y extracción
- [Refactor cotidiano (4): condicionales](https://franiglesias.github.io/everyday-refactor-4/) - Técnicas para simplificar condicionales
- [Object Calisthenics para adelgazar las clases](https://franiglesias.github.io/calistenics-and-small-classes/) - Incluye discusión sobre eliminación de else

## Referencias

- **"Refactoring: Improving the Design of Existing Code"** - Martin Fowler - Replace Nested Conditional with Guard Clauses
- **"Clean Code"** - Robert C. Martin - Capítulo sobre funciones y estructuras de control
- [Object Calisthenics - William Durand](https://williamdurand.fr/2013/06/03/object-calisthenics/) - Regla #2
- [Flattening Arrow Code](https://blog.codinghorror.com/flattening-arrow-code/) - Jeff Atwood sobre reducir anidación
- **"Refactoring to Patterns"** - Joshua Kerievsky - Replace Conditional Logic with Strategy
