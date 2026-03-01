# Primitive Obsession

Obsesión primitiva.

## Definición

Conceptos de dominio se modelan con primitivos, lo que obliga a esparcir reglas de validación, formato y comportamiento por todo el código.

## Descripción

**Primitive Obsession** ocurre cuando usamos tipos primitivos (strings, numbers, booleans) para representar conceptos del dominio que merecen ser objetos por derecho propio. En lugar de crear una clase `Email` o `Money`, usamos strings y numbers, perdiendo la oportunidad de encapsular:
- Reglas de validación
- Formato y conversión
- Comportamiento relacionado
- Seguridad de tipos

Este smell aparece porque es más fácil usar un string que crear una nueva clase. Sin embargo, esta conveniencia inicial se convierte en un problema cuando:
- La misma validación se repite en múltiples lugares
- El formato inconsistente causa bugs
- No hay protección del compilador contra uso incorrecto
- Las reglas de negocio quedan dispersas

## Síntomas

- Validaciones repetidas del mismo tipo de dato en múltiples lugares
- Strings que representan conceptos con formato específico (email, URL, código postal)
- Numbers que representan cantidades con unidades (dinero, medidas)
- Grupos de primitivos que siempre viajan juntos (Data Clump)
- Lógica de formato dispersa en diferentes métodos
- Comentarios explicando restricciones o formato de variables primitivas
- Verificaciones de tipo o formato antes de usar el valor

## Ejemplo

```pseudocode
class Order {
  customerName: string
  customerEmail: string
  address: string
  totalAmount: number
  currency: string

  function sendInvoice() {
    // Validación dispersa
    if (not customerEmail.contains("@")) {
      throw error "Email inválido"
    }

    if (address is empty) {
      throw error "No se ha indicado dirección"
    }

    if (totalAmount <= 0) {
      throw error "El monto debe ser mayor que cero"
    }

    print "Factura enviada a " + customerName
    print "Dirección: " + address
    print "Total: " + totalAmount + " " + currency
  }
}

// La validación se repite en cada lugar que use email
function validateEmail(email: string) {
  if (not email.contains("@")) {
    throw error "Email inválido"
  }
}

// El formato se repite en cada lugar que muestre dinero
function formatMoney(amount: number, currency: string) {
  return amount.toString() + " " + currency
}
```

## Ejercicio

Introduce soporte para diferentes monedas, para enviar la factura por email, y para formatear la dirección en función del país.

## Problemas que encontrarás

Dado que los primitivos no nos permiten garantizar la integridad de sus valores, tendrás que introducir validaciones en muchos lugares, incluso de forma repetida. Algunos datos siempre viajan juntos (Data Clump), por lo que tienes que asegurarte de que permanecen juntos. Para formatear de forma diferente basándote en algún dato arbitrario tendrás que introducir lógica de decisión.

## Proceso de Refactoring

### 1. Identificar primitivos que representan conceptos
- Busca strings que representen valores con formato específico
- Busca numbers que representen cantidades con significado
- Identifica grupos de primitivos que siempre aparecen juntos

### 2. Crear Value Objects
- Crea una clase para cada concepto del dominio
- Ejemplo: `Email`, `Money`, `Address`, `PhoneNumber`
- Hazla inmutable si es posible

### 3. Encapsular validación en el constructor
- Mueve todas las validaciones al constructor del Value Object
- El objeto garantiza que siempre está en un estado válido
- Si la construcción falla, lanza una excepción clara

### 4. Añadir comportamiento relacionado
- Métodos de formato: `email.format()`, `money.formatWithCurrency()`
- Métodos de conversión: `money.convertTo(currency)`
- Métodos de comparación: `money.isGreaterThan(other)`
- Operaciones: `money.add(other)`, `money.multiply(factor)`

### 5. Reemplazar primitivos progresivamente
- Comienza por una clase o método
- Reemplaza el primitivo con el Value Object
- Actualiza las llamadas y elimina validaciones redundantes
- Ejecuta tests después de cada cambio

### 6. Eliminar validaciones y lógica duplicada
- Una vez que usas Value Objects, la validación está centralizada
- Elimina las verificaciones redundantes dispersas por el código
- La lógica de formato vive en un solo lugar

## Técnicas de Refactoring Aplicables

- **Replace Data Value with Object**: Convertir primitivos en objetos
- **Replace Type Code with Class**: Para códigos o enumeraciones
- **Extract Class**: Crear Value Objects para grupos de datos
- **Introduce Parameter Object**: Si los primitivos se pasan como parámetros
- **Replace Primitive with Object**: Técnica general para este smell

## Beneficios

- **Validación centralizada**: Reglas de negocio en un solo lugar
- **Tipo más seguro**: El compilador previene errores de tipo
- **Código más expresivo**: `Money` es más claro que `number`
- **Comportamiento encapsulado**: Lógica relacionada vive junta
- **Menos duplicación**: Formato y conversión en un solo lugar
- **Inmutabilidad**: Value Objects inmutables son más seguros
- **Mejor dominio**: El código refleja conceptos del negocio

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/bloaters/primitive-obsession.ts) - [README](../../typescript/src/code-smells/bloaters/primitive-obsession.readme.md)
- [Go](../../go/code_smells/bloaters/primitive_obsession.go) - [README](../../go/code_smells/bloaters/primitive_obsession.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/bloaters/PrimitiveObsession.java) - [README](../../java/src/main/java/com/refactoring/codesmells/bloaters/PrimitiveObsession.readme.md)
- [PHP](../../php/src/code-smells/bloaters/PrimitiveObsession.php) - [README](../../php/src/code-smells/bloaters/PrimitiveObsession.readme.md)
- [Python](../../python/src/code_smells/bloaters/primitive_obsession.py) - [README](../../python/src/code_smells/bloaters/primitive_obsession_readme.md)
- [C#](../../csharp/src/code-smells/bloaters/PrimitiveObsession.cs) - [README](../../csharp/src/code-smells/bloaters/primitive-obsession.readme.md)

## Referencias en Español

- [Primitive Obsession](https://franiglesias.github.io/primitive-obsession/) - Análisis exhaustivo del problema de usar primitivos para conceptos del dominio
- [Encapsular primitivos y colecciones](https://franiglesias.github.io/encapsulate/) - Técnicas para crear Value Objects y encapsular primitivos
- [Refactor cotidiano (4). Sustituye escalares por objetos](https://franiglesias.github.io/everyday-refactor-4/) - Práctica diaria de refactorización de primitivos
- [Value Objects con Doctrine](https://franiglesias.github.io/doctrine-vo-gal/) - Implementación de Value Objects en contexto de persistencia
- [Object Calisthenics para mejorar el diseño de las clases](https://franiglesias.github.io/calistenics-and-value-objects/) - Ejercicios para crear mejores Value Objects

## Referencias

- [Refactoring Guru - Primitive Obsession](https://refactoring.guru/smells/primitive-obsession)
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
- Eric Evans - "Domain-Driven Design" - Value Objects
- Vaughn Vernon - "Implementing Domain-Driven Design"
