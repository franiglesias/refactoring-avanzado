# Empaquetar Primitivos

Regla 4 de Object Calisthenics

## Definición

Los tipos primitivos del lenguaje (strings, números, booleanos, fechas, etc.) que representan conceptos del dominio deben encapsularse en objetos que protejan sus invariantes y añadan comportamiento relacionado.

## Descripción

**Primitive Obsession** es el code smell de usar tipos primitivos para representar conceptos del dominio que tienen reglas de validación, formato o comportamiento específico.

Problemas de usar primitivos directamente:

1. **Sin validación**: Un string puede ser vacío cuando representa un email, un número puede ser negativo cuando representa una edad
2. **Validación duplicada**: Cada lugar que usa el primitivo debe validar de nuevo
3. **Comportamiento esparcido**: Lógica relacionada con el concepto está duplicada en múltiples lugares
4. **Sin expresividad del dominio**: `string` no comunica que es un email, ISBN, DNI, o dirección postal
5. **Acoplamiento a representación**: Cambiar de string a objeto complejo requiere modificar todo el código cliente
6. **Testing complicado**: Hay que recordar crear datos válidos manualmente en cada test

La solución es crear **Value Objects**: objetos inmutables que:
- Encapsulan un valor primitivo
- Validan invariantes en construcción
- Proporcionan comportamiento relacionado con el concepto
- Expresan el dominio explícitamente
- Son comparables por valor, no por referencia

Los Value Objects son uno de los patrones fundamentales de **Domain-Driven Design**.

## Síntomas

- Strings representando emails, URLs, códigos postales, DNI, ISBN, etc.
- Números representando dinero, porcentajes, edades, cantidades con unidades
- Booleanos con nombres poco claros (`active`, `enabled`, `flag`)
- Fechas sin timezone o validación de rangos
- Validaciones del tipo `if (email.contains("@"))` repetidas
- Comentarios explicando qué representa un primitivo
- Parámetros con nombres genéricos (`value`, `data`, `info`)
- Métodos utilitarios estáticos para validar (`EmailValidator.isValid(string)`)
- Métodos que reciben muchos parámetros primitivos relacionados
- Lógica de formato duplicada (`formatCurrency`, `parseCurrency` en varios lugares)

## Ejemplo

### Antes (Violación)

```pseudocode
class User {
  string email
  int age
  string phone
  decimal balance
  string country

  constructor(email, age, phone, balance, country) {
    // Sin validación o validación superficial
    this.email = email
    this.age = age
    this.phone = phone
    this.balance = balance
    this.country = country
  }
}

class OrderService {
  method processOrder(userId, amount, currency) {
    // Validaciones repetidas
    if (amount <= 0) {
      throw "Invalid amount"
    }
    if (currency != "USD" and currency != "EUR") {
      throw "Invalid currency"
    }

    // Lógica de conversión duplicada
    if (currency == "EUR") {
      amount = amount * 1.1  // Conversión hardcoded
    }

    user = findUser(userId)
    if (user.balance < amount) {
      throw "Insufficient funds"
    }

    user.balance = user.balance - amount
    saveUser(user)
  }

  method displayPrice(amount, currency) {
    // Lógica de formato duplicada
    if (currency == "USD") {
      return "$" + formatDecimal(amount, 2)
    }
    if (currency == "EUR") {
      return formatDecimal(amount, 2) + "€"
    }
  }
}
```

**Problemas**:
- Email sin validar (puede ser vacío, sin @, formato incorrecto)
- Age sin validar (puede ser negativo, 500 años, etc.)
- Phone sin formato estandarizado
- Balance sin validación (puede ser negativo)
- Currency como string (sin validación, lógica duplicada)
- Conversión de moneda hardcoded y duplicada
- Formato de precio duplicado

### Después (Cumplimiento)

```pseudocode
// Value Object: Email
class Email {
  private string value

  constructor(email) {
    if (not contains(email, "@")) {
      throw "Invalid email format"
    }
    if (length(email) < 5) {
      throw "Email too short"
    }
    this.value = toLowerCase(email)  // Normalización
  }

  method getValue() {
    return this.value
  }

  method getDomain() {
    return split(this.value, "@")[1]
  }

  method equals(other) {
    return this.value == other.value
  }
}

// Value Object: Age
class Age {
  private int value

  constructor(years) {
    if (years < 0) {
      throw "Age cannot be negative"
    }
    if (years > 150) {
      throw "Age unrealistic"
    }
    this.value = years
  }

  method getValue() {
    return this.value
  }

  method isAdult() {
    return this.value >= 18
  }

  method isSenior() {
    return this.value >= 65
  }
}

// Value Object: Money
class Money {
  private decimal amount
  private Currency currency

  constructor(amount, currency) {
    if (amount < 0) {
      throw "Amount cannot be negative"
    }
    this.amount = amount
    this.currency = currency
  }

  method add(other) {
    if (this.currency != other.currency) {
      throw "Cannot add different currencies"
    }
    return new Money(this.amount + other.amount, this.currency)
  }

  method subtract(other) {
    if (this.currency != other.currency) {
      throw "Cannot subtract different currencies"
    }
    result = this.amount - other.amount
    if (result < 0) {
      throw "Insufficient funds"
    }
    return new Money(result, this.currency)
  }

  method convertTo(targetCurrency, exchangeRate) {
    convertedAmount = this.amount * exchangeRate
    return new Money(convertedAmount, targetCurrency)
  }

  method format() {
    return this.currency.format(this.amount)
  }

  method isGreaterThan(other) {
    if (this.currency != other.currency) {
      throw "Cannot compare different currencies"
    }
    return this.amount > other.amount
  }
}

// Value Object: Currency
class Currency {
  private string code
  private string symbol

  constructor(code) {
    allowedCurrencies = ["USD", "EUR", "GBP"]
    if (not contains(allowedCurrencies, code)) {
      throw "Unsupported currency: " + code
    }
    this.code = code
    this.symbol = this.getSymbolForCode(code)
  }

  private method getSymbolForCode(code) {
    symbols = {"USD": "$", "EUR": "€", "GBP": "£"}
    return symbols[code]
  }

  method format(amount) {
    formatted = formatDecimal(amount, 2)
    if (this.code == "USD") {
      return this.symbol + formatted
    }
    return formatted + this.symbol
  }

  method equals(other) {
    return this.code == other.code
  }
}

// Uso mejorado
class User {
  Email email
  Age age
  PhoneNumber phone
  Money balance
  Country country

  constructor(email, age, phone, balance, country) {
    // Validación automática en construcción de Value Objects
    this.email = email
    this.age = age
    this.phone = phone
    this.balance = balance
    this.country = country
  }

  method canPurchase(price) {
    return this.balance.isGreaterThan(price)
  }

  method deduct(amount) {
    this.balance = this.balance.subtract(amount)
  }
}

class OrderService {
  method processOrder(userId, price) {
    user = findUser(userId)

    if (not user.canPurchase(price)) {
      throw "Insufficient funds"
    }

    user.deduct(price)
    saveUser(user)
  }

  method displayPrice(price) {
    return price.format()  // Money sabe formatearse
  }
}
```

**Mejoras**:
- Validación centralizada en cada Value Object
- Comportamiento relacionado encapsulado (formato, comparación, operaciones)
- Tipo de dominio explícito (`Email` vs `string`)
- Sin duplicación de lógica de validación o formato
- Inmutabilidad garantizada
- Testing simplificado (crear valores válidos es fácil)

## Ejercicio

**Tarea**: En el código proporcionado en tu lenguaje, identifica todos los primitivos que representan conceptos del dominio y encapsúlalos en Value Objects con validación e invariantes apropiadas.

**Criterios de éxito**:
1. Todos los conceptos del dominio están representados por clases, no primitivos
2. Cada Value Object valida sus invariantes en construcción
3. Comportamiento relacionado está encapsulado en el Value Object
4. Los Value Objects son inmutables
5. No hay validaciones duplicadas en el código

## Problemas que Encontrarás

### 1. "Es demasiado verbose para algo tan simple"

Un email es solo un string, ¿por qué crear una clase? Porque:
- La validación se hace una vez, se usa mil veces
- El tipo explícito documenta el dominio
- Cambios futuros están localizados (ej: añadir validación más estricta)
- El código cliente es más limpio (no hay validaciones everywhere)

### 2. Dónde poner el límite

No todo primitivo necesita envoltorio:
- **Sí**: Conceptos de dominio con reglas (Email, Money, ISBN, DNI)
- **No**: Primitivos genéricos sin reglas (índice de array, contadores temporales)
- **Criterio**: ¿Tiene reglas de validación o comportamiento específico del dominio?

### 3. Performance

Los Value Objects tienen un pequeño overhead, pero:
- Es despreciable en la mayoría de aplicaciones
- Los beneficios de mantenibilidad superan cualquier costo
- Si tienes un cuello de botella real, optimiza ese punto específico
- Los lenguajes modernos optimizan objetos pequeños e inmutables

### 4. Compatibilidad con frameworks

Algunos frameworks esperan tipos primitivos (serialización JSON, ORMs, etc.):
- Usa adaptadores en la capa de infraestructura
- Convierte entre Value Objects y primitivos en los boundaries
- El dominio debe ser independiente de frameworks

### 5. Value Objects con múltiples valores

A veces un concepto requiere varios valores (ej: Coordenada geográfica = latitud + longitud):
- Agrúpalos en un solo Value Object
- Valida que la combinación sea válida
- Proporciona operaciones que tengan sentido para el concepto completo

## Proceso de Aplicación

### 1. Identificar candidatos

Busca primitivos que:
- Tienen validaciones específicas
- Se validan de la misma forma en múltiples lugares
- Representan conceptos del dominio (no son genéricos)
- Tienen comportamiento asociado (formato, cálculo, comparación)

### 2. Crear el Value Object

```pseudocode
class ConceptName {
  private primitiveType value

  constructor(value) {
    validate(value)  // Lanzar excepción si inválido
    this.value = value
  }

  method getValue() {
    return this.value
  }

  // Métodos de comportamiento específico del concepto

  method equals(other) {
    return this.value == other.value
  }
}
```

### 3. Hacer inmutable

- Todos los campos deben ser `private` y `final`/`readonly`/`const`
- No proporcionar setters
- Operaciones que "modifican" retornan nuevas instancias

### 4. Añadir validación

- Valida en el constructor
- Lanza excepciones descriptivas si los datos son inválidos
- Asegura que es imposible crear instancias inválidas

### 5. Añadir comportamiento

- Métodos de formato (`format`, `toString`)
- Métodos de comparación (`equals`, `isGreaterThan`)
- Operaciones del dominio (`add`, `subtract`, `convert`)
- Queries sobre el valor (`isValid`, `isExpired`, `hasExpired`)

### 6. Implementar igualdad por valor

```pseudocode
method equals(other) {
  if (other is null) return false
  if (other is not ConceptName) return false
  return this.value == other.value
}

method hashCode() {
  return hash(this.value)
}
```

### 7. Reemplazar uso de primitivos

- Cambia parámetros de métodos a usar Value Objects
- Cambia campos de clases a usar Value Objects
- Actualiza constructores y factories
- La compilación fallará donde falta conversión, lo cual es bueno (type safety)

### 8. Eliminar validaciones duplicadas

- Busca validaciones del primitivo en el código
- Elimínalas (ya no son necesarias, el Value Object garantiza validez)
- Simplifica el código cliente

## Técnicas de Refactoring Aplicables

- **Replace Primitive with Object**: Técnica principal
- **Replace Type Code with Class**: Si el primitivo es un código de tipo
- **Introduce Parameter Object**: Si varios primitivos relacionados se pasan juntos
- **Preserve Whole Object**: Pasar el Value Object completo en lugar de sus partes
- **Extract Class**: Para crear el Value Object inicial
- **Encapsulate Field**: Hacer privados los campos internos

## Beneficios

### 1. Validación Centralizada

Las reglas de validación están en un solo lugar, no duplicadas en todo el código.

### 2. Type Safety

El sistema de tipos del lenguaje previene errores (no puedes pasar un `Email` donde se espera un `PhoneNumber`).

### 3. Expresividad del Dominio

El código habla el lenguaje del dominio, no el lenguaje de tipos primitivos.

### 4. Encapsulación de Comportamiento

Lógica relacionada con el concepto está junto al dato, no esparcida.

### 5. Inmutabilidad

Value Objects inmutables eliminan bugs relacionados con mutación inesperada.

### 6. Testing Simplificado

Es fácil crear valores válidos en tests, y el Value Object mismo es testeable aisladamente.

### 7. Evolución del Código

Cambios en representación o reglas están localizados en el Value Object.

### 8. Menos Bugs

Imposible crear valores inválidos = menos bugs relacionados con datos incorrectos.

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/calisthenics-exercises/no-primitives.ts)
- [Go](../../go/calisthenics_exercises/03_wrap_primitives.go)
- [Java](../../java/src/main/java/com/refactoring/calisthenics/WrapPrimitives.java)
- [PHP](../../php/src/calisthenics-exercises/NoPrimitives.php)
- [Python](../../python/src/calisthenics_exercises/no_primitives.py)
- [C#](../../csharp/src/calisthenics-exercises/NoPrimitives.cs)

## Referencias en Español

- [Primitive obsession](https://franiglesias.github.io/primitive-obsession/) - Análisis detallado del smell y cómo resolverlo
- [Calistenias para objetos de valor](https://franiglesias.github.io/calistenics-and-value-objects/) - Aplicación de calisthenics a Value Objects
- [Value Objects](https://franiglesias.github.io/value-objects/) - Guía sobre creación de Value Objects

## Referencias

- **"Domain-Driven Design"** - Eric Evans - Value Objects como building block fundamental
- **"Implementing Domain-Driven Design"** - Vaughn Vernon - Capítulo sobre Value Objects
- **"Refactoring: Improving the Design of Existing Code"** - Martin Fowler - Replace Data Value with Object
- [Object Calisthenics - William Durand](https://williamdurand.fr/2013/06/03/object-calisthenics/) - Regla #4
- [Value Object - Martin Fowler](https://martinfowler.com/bliki/ValueObject.html) - Definición y características
- [Primitive Obsession - Refactoring Guru](https://refactoring.guru/smells/primitive-obsession)
