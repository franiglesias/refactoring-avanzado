# Temporal Instance Variables

Variables de instancia temporales.

## Definición

Este smell ocurre cuando un objeto tiene campos que solo están llenos (o tienen sentido) en ciertas etapas de su ciclo de vida. Esto suele indicar un acoplamiento temporal, donde los métodos deben llamarse en un orden específico para que el objeto sea válido, dejando al objeto en un estado inconsistente fuera de esa secuencia.

## Descripción

**Temporal Instance Variables** (variables de instancia temporales) aparecen cuando un objeto tiene campos que no siempre tienen un valor válido. Estos campos están "vacíos" (null, undefined, valores por defecto) en algunas etapas del ciclo de vida del objeto y solo se llenan cuando se llaman ciertos métodos en el orden correcto.

Este smell indica varios problemas:
- El objeto no siempre está en un estado consistente
- Hay acoplamiento temporal: métodos deben llamarse en orden específico
- Los invariantes del objeto no están protegidos
- El objeto actúa como una máquina de estados frágil
- Es fácil usar el objeto incorrectamente

El resultado es código propenso a errores donde olvidar llamar un método o llamarlos en el orden incorrecto causa fallos sutiles o comportamiento indefinido.

## Síntomas

- Campos que son opcionales o nullable pero representan conceptos requeridos
- Métodos que verifican si los campos están inicializados antes de usarlos
- Comentarios que explican en qué orden llamar los métodos
- Necesidad de llamar `initialize()` o `setup()` después de construcción
- Campos que se llenan y vacían durante el ciclo de vida del objeto
- Errores de NullPointerException o similar por campos no inicializados
- El objeto tiene estados "válido" e "inválido" no expresados explícitamente

## Ejemplo

```pseudocode
class PizzaOrder {
  size: string?              // Solo válido después de start()
  toppings: string[] = []    // Solo válido entre start() y place()
  address: string?           // Solo válido después de setDeliveryAddress()

  function start(size: string) {
    this.size = size
    this.toppings = []
    // address puede estar indefinido
  }

  function addTopping(topping: string) {
    if (not this.size) {
      // Verificación necesaria por estado inválido
      return
    }
    this.toppings.add(topping)
  }

  function setDeliveryAddress(address: string) {
    this.address = address
  }

  function place(): string {
    summary = "Pizza " + (this.size ?? "?") +
              " to " + (this.address ?? "UNKNOWN") +
              " with [" + this.toppings.join(", ") + "]"

    // Limpiar estado temporal
    this.size = null
    this.address = null
    this.toppings = []

    return summary
  }
}

// Uso: orden incorrecto causa problemas
order = new PizzaOrder()
order.addTopping("pepperoni")  // Falla silenciosamente: no hay size
order.start("L")
order.addTopping("mushroom")    // Ahora sí funciona
order.place()                   // address es UNKNOWN
```

## Ejercicio

Añade una validación para que no se pueda llamar a `place()` si no se ha añadido al menos un ingrediente.

## Problemas que encontrarás

Te darás cuenta de que el objeto es una "máquina de estados" frágil. Si un cliente olvida llamar a `start()` o intenta llamar a `addTopping()` fuera de orden, el sistema puede fallar silenciosamente o requerir comprobaciones constantes de nulidad en cada método.

## Proceso de Refactoring

### 1. Identificar variables temporales
- Busca campos que son null/undefined en algunos momentos
- Identifica campos que se llenan y vacían durante el ciclo de vida
- Mapea las etapas del ciclo de vida del objeto

### 2. Extraer objetos de estado
- Crea clases separadas para cada estado del ciclo de vida
- Ejemplo: `PizzaOrderDraft`, `PizzaOrderReady`, `PlacedOrder`
- Cada clase tiene solo los campos que son válidos en ese estado

### 3. Usar Builder pattern
- Si el objeto necesita construcción paso a paso
- Crea un builder que acumula el estado temporal
- El builder retorna un objeto inmutable completamente inicializado
- Ejemplo: `new PizzaOrderBuilder().size("L").addTopping("pepperoni").build()`

### 4. Reemplazar con objeto de parámetros
- Si las variables temporales son solo para pasar entre métodos
- Agrúpalas en un objeto de parámetros
- Pásalo explícitamente en lugar de almacenarlo en estado

### 5. Aplicar State pattern
- Si el objeto tiene comportamiento que varía según el estado
- Usa el patrón State para hacer explícitos los estados
- Cada estado es una clase con su propio comportamiento

### 6. Hacer objetos inmutables
- Diseña objetos que son válidos desde la construcción
- No permitas modificación después de crear
- Usar builders o factories para construcción compleja

## Técnicas de Refactoring Aplicables

- **Replace Method with Method Object**: Extraer máquina de estados a objeto separado
- **Extract Class**: Crear clases para cada estado del ciclo de vida
- **Builder Pattern**: Para construcción paso a paso segura
- **State Pattern**: Para comportamiento que varía según estado
- **Introduce Parameter Object**: Para datos temporales que se pasan entre métodos
- **Self Encapsulate Field**: Para controlar acceso a campos temporales

## Beneficios

- **Objetos siempre válidos**: No hay estados inconsistentes
- **Sin acoplamiento temporal**: No importa el orden de operaciones
- **Type safety**: El compilador previene uso incorrecto
- **Código más claro**: Estados explícitos en lugar de implícitos
- **Menos bugs**: Imposible olvidar inicializar
- **Mejor encapsulación**: Los invariantes están protegidos
- **Inmutabilidad**: Objetos más seguros y predecibles

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/oop-abusers/temporal-instance-variables.ts) - [README](../../typescript/src/code-smells/oop-abusers/temporal-instance-variables.readme.md)
- [Go](../../go/code_smells/oop_abusers/temporal_instance_variables.go) - [README](../../go/code_smells/oop_abusers/temporal_instance_variables.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/oopabusers/TemporalInstanceVariables.java) - [README](../../java/src/main/java/com/refactoring/codesmells/oopabusers/TemporalInstanceVariables.readme.md)
- [PHP](../../php/src/code-smells/oop-abusers/TemporalInstanceVariables.php) - [README](../../php/src/code-smells/oop-abusers/TemporalInstanceVariables.readme.md)
- [Python](../../python/src/code_smells/oop_abusers/temporal_instance_variables.py) - [README](../../python/src/code_smells/oop_abusers/temporal_instance_variables_readme.md)
- [C#](../../csharp/src/code-smells/oop-abusers/TemporalInstanceVariables.cs) - [README](../../csharp/src/code-smells/oop-abusers/temporal-instance-variables.readme.md)

## Referencias en Español

- [Object Calisthenics para mejorar el diseño de las clases](https://franiglesias.github.io/calistenics-and-value-objects/) - Incluye técnicas para mantener objetos en estado válido consistente

## Referencias

- [Refactoring Guru - Temporary Field](https://refactoring.guru/smells/temporary-field)
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
- Joshua Bloch - "Effective Java" - Builder pattern, Immutability
- Gang of Four - "Design Patterns" - State, Builder patterns
