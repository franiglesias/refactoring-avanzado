# Feature Envy

Envidia de características.

## Definición

Una clase usa la información de otra clase colaboradora para hacer cálculos o tomar decisiones, sugiriendo que la segunda clase debería exponer esos comportamientos. Al depender de la estructura de la colaboradora, la clase cliente queda acoplada.

## Descripción

**Feature Envy** ocurre cuando un método parece más interesado en otra clase que en la clase donde reside. El método accede extensivamente a los datos de otro objeto, realiza operaciones con ellos y hace múltiples llamadas a getters.

Este smell sugiere que el comportamiento está en el lugar equivocado. El método "envidia" las características de otra clase porque necesita trabajar intensamente con sus datos. Esto viola el principio "Tell, Don't Ask" y crea acoplamiento fuerte entre clases.

Señales típicas:
- Un método hace múltiples llamadas a getters de otro objeto
- La lógica depende de la estructura interna de otra clase
- Cambios en la clase colaboradora requieren cambios en el cliente
- El método usa más datos de otra clase que de la propia

## Síntomas

- Métodos que hacen muchas llamadas a getters de otro objeto
- Lógica de decisión basada en el estado interno de otra clase
- Cálculos que usan principalmente datos de otro objeto
- El método conoce detalles íntimos de la estructura de otra clase
- Cambios en la clase colaboradora frecuentemente requieren cambios en el cliente
- El nombre del método sugiere que pertenece a otra clase

## Ejemplo

```pseudocode
class Customer {
  name: string
  street: string
  city: string
  zip: string
}

class ShippingCalculator {
  function cost(customer: Customer): number {
    // "Envidia" los datos de Customer para tomar decisiones
    base = customer.zip.startsWith("9") ? 10 : 20
    distant = customer.city.length > 6 ? 5 : 0
    return base + distant
  }
}

// Uso
customer = new Customer("John", "Main St", "New York", "90210")
calculator = new ShippingCalculator()
cost = calculator.cost(customer)  // calculator "envidia" a customer
```

## Ejercicio

Añade envío gratis para clientes en ciertas ciudades y un recargo de fin de semana.

## Problemas que encontrarás

Probablemente, seguirás añadiendo condiciones dentro de `ShippingCalculator` que dependen de detalles internos de `Customer`, esparciendo reglas en el lugar equivocado y volviendo frágiles los cambios.

## Proceso de Refactoring

### 1. Identificar el método envidioso
- Encuentra métodos que acceden extensivamente a otra clase
- Cuenta cuántas llamadas hace a getters de otro objeto
- Verifica si la lógica es realmente responsabilidad de la otra clase

### 2. Mover el método a donde pertenece
- Mueve el método a la clase cuyos datos usa principalmente
- Si usa datos de múltiples clases, muévelo a la que más usa
- El método debe vivir cerca de los datos que manipula

### 3. Invertir la dependencia
- En lugar de: `calculator.cost(customer)`
- Usa: `customer.calculateShippingCost()`
- La clase que tiene los datos expone el comportamiento

### 4. Aplicar "Tell, Don't Ask"
- No preguntes por datos para tomar decisiones
- Dile al objeto lo que necesitas y deja que él lo haga
- Ejemplo: en lugar de `if (customer.isVip()) discount = ...`
- Usa: `discount = customer.calculateDiscount()`

### 5. Extraer clase si es necesario
- Si el método usa datos de ambas clases equitativamente
- Considera crear una tercera clase que contenga el comportamiento
- Ejemplo: `ShippingPolicy` que conoce tanto `Customer` como reglas de envío

### 6. Usar polimorfismo en lugar de condicionales
- Si hay muchas condiciones basadas en tipos o estados
- Considera usar polimorfismo o el patrón Strategy
- Cada tipo de customer puede tener su propia lógica de cálculo

## Técnicas de Refactoring Aplicables

- **Move Method**: Mover el método a la clase que contiene los datos
- **Extract Method**: Extraer partes del método antes de moverlo
- **Extract Class**: Crear una clase nueva si el comportamiento no encaja en ninguna existente
- **Replace Conditional with Polymorphism**: Si hay condicionales basados en tipo
- **Introduce Parameter Object**: Si se pasan muchos datos de un objeto

## Beneficios

- **Mejor encapsulación**: Los datos y el comportamiento viven juntos
- **Menos acoplamiento**: Los clientes no conocen detalles internos
- **Código más intuitivo**: El comportamiento está donde se espera
- **Más fácil de cambiar**: Cambios en lógica de negocio afectan una sola clase
- **Tell, Don't Ask**: Objetos responsables de su propio comportamiento
- **Cohesión mejorada**: Cada clase tiene responsabilidades relacionadas

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/couplers/feature-envy.ts) - [README](../../typescript/src/code-smells/couplers/feature-envy.readme.md)
- [Go](../../go/code_smells/couplers/feature_envy.go) - [README](../../go/code_smells/couplers/feature_envy.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/couplers/FeatureEnvy.java) - [README](../../java/src/main/java/com/refactoring/codesmells/couplers/FeatureEnvy.readme.md)
- [PHP](../../php/src/code-smells/couplers/FeatureEnvy.php) - [README](../../php/src/code-smells/couplers/FeatureEnvy.readme.md)
- [Python](../../python/src/code_smells/couplers/feature_envy.py) - [README](../../python/src/code_smells/couplers/feature_envy_readme.md)
- [C#](../../csharp/src/code-smells/couplers/FeatureEnvy.cs) - [README](../../csharp/src/code-smells/couplers/feature-envy.readme.md)

## Referencias en Español

- [Refactor cotidiano (8). Dónde poner el conocimiento](https://franiglesias.github.io/everyday-refactor-8/) - Guía para ubicar correctamente el comportamiento cerca de los datos
- [Refactor cotidiano (6). Tell, Don't Ask y Ley de Demeter](https://franiglesias.github.io/everyday-refactor-6/) - Explicación del principio Tell, Don't Ask para evitar Feature Envy

## Referencias

- [Refactoring Guru - Feature Envy](https://refactoring.guru/smells/feature-envy)
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
- The Pragmatic Programmer - "Tell, Don't Ask" principle
