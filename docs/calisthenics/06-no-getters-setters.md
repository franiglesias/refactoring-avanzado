# No Usar Getters y Setters

Regla 6 de Object Calisthenics

## Definición

Los objetos no deben exponer su estado interno mediante getters y setters. En su lugar, deben proporcionar comportamiento que opere sobre ese estado, siguiendo el principio **Tell, Don't Ask**.

## Descripción

**Getters y setters** parecen inocuos, incluso se consideran "buenas prácticas" en algunos contextos. Sin embargo, son una forma sutil de **romper la encapsulación**. Al exponer el estado interno:

1. **Acoplamiento estructural**: El código cliente depende de la estructura interna del objeto, dificultando la posibilidad de cambiarla en el futuro
2. **Lógica esparcida**: Comportamiento que debería estar en el objeto acaba en el código cliente
3. **Violación de Tell, Don't Ask**: Preguntas por datos para tomar decisiones, en lugar de decirle al objeto qué hacer
4. **Objetos anémicos**: Clases que son solo contenedores de datos sin comportamiento
5. **Dificulta evolución**: Cambiar la estructura interna requiere cambiar todo el código cliente, posiblemente en muchos lugares

La diferencia fundamental:

**Enfoque procedural (getters/setters)**:
```pseudocode
if (user.getBalance() > price.getAmount()) {
  user.setBalance(user.getBalance() - price.getAmount())
}
```

**Enfoque orientado a objetos (Tell, Don't Ask)**:
```pseudocode
user.pay(price)
```

El primer enfoque:
- Pregunta por datos (`getBalance`, `getAmount`)
- Toma decisiones fuera del objeto
- Modifica estado desde afuera (`setBalance`)
- Acopla a la representación interna (balance como número)

El segundo:
- Le dice al objeto qué hacer (`pay`)
- El objeto toma sus propias decisiones
- El objeto modifica su propio estado
- Oculta la representación interna

Esta regla fuerza diseño **orientado a comportamiento** en lugar de **orientado a datos**.

## Síntomas

- Métodos que empiezan con `get` o `set`
- Propiedades públicas en clases
- Código cliente con cadenas de llamadas a getters
- Lógica de negocio fuera de los objetos que contienen los datos relevantes
- Clases "anémicas" con solo getters/setters y sin comportamiento
- Tests que verifican estado interno en lugar de comportamiento
- Condicionales que preguntan por estado antes de llamar métodos
- Métodos que retornan objetos mutables del estado interno
- Documentación que describe "qué campos tiene" en lugar de "qué hace"

## Ejemplo

### Antes (Violación)

```pseudocode
class BankAccount {
  private decimal balance
  private string accountNumber
  private boolean active

  method getBalance() {
    return this.balance
  }

  method setBalance(newBalance) {
    this.balance = newBalance
  }

  method isActive() {
    return this.active
  }

  method setActive(active) {
    this.active = active
  }

  method getAccountNumber() {
    return this.accountNumber
  }
}

class TransferService {
  method transfer(from, to, amount) {
    // Preguntamos por estado
    if (not from.isActive()) {
      throw "Source account is not active"
    }

    if (not to.isActive()) {
      throw "Destination account is not active"
    }

    if (from.getBalance() < amount) {
      throw "Insufficient funds"
    }

    // Modificamos estado desde fuera
    from.setBalance(from.getBalance() - amount)
    to.setBalance(to.getBalance() + amount)

    // Lógica de negocio fuera del objeto
    if (from.getBalance() < 100) {
      sendLowBalanceWarning(from.getAccountNumber())
    }
  }
}

class ReportService {
  method generateReport(accounts) {
    report = ""
    for each account in accounts {
      // Acceso a estructura interna
      report = report + account.getAccountNumber() + ": "
      report = report + "$" + account.getBalance()

      if (account.isActive()) {
        report = report + " (Active)"
      } else {
        report = report + " (Inactive)"
      }
    }
    return report
  }
}
```

**Problemas**:
- Balance expuesto y modificado desde fuera
- Validaciones fuera del objeto (isActive, balance >= amount)
- Lógica de negocio en servicios externos
- Acoplamiento a representación interna
- BankAccount es una clase anémica (solo datos, sin comportamiento)

### Después (Cumplimiento)

```pseudocode
class BankAccount {
  private decimal balance
  private string accountNumber
  private boolean active

  constructor(accountNumber, initialBalance) {
    this.accountNumber = accountNumber
    this.balance = initialBalance
    this.active = true
  }

  // Tell, don't ask: comportamiento en lugar de getters/setters
  method withdraw(amount) {
    this.ensureIsActive()
    this.ensureSufficientFunds(amount)

    this.balance = this.balance - amount

    if (this.hasLowBalance()) {
      this.notifyLowBalance()
    }
  }

  method deposit(amount) {
    this.ensureIsActive()
    this.ensureValidAmount(amount)

    this.balance = this.balance + amount
  }

  method deactivate() {
    this.active = false
  }

  method activate() {
    this.active = true
  }

  // Validaciones privadas - comportamiento interno
  private method ensureIsActive() {
    if (not this.active) {
      throw "Account is not active"
    }
  }

  private method ensureSufficientFunds(amount) {
    if (this.balance < amount) {
      throw "Insufficient funds"
    }
  }

  private method ensureValidAmount(amount) {
    if (amount <= 0) {
      throw "Amount must be positive"
    }
  }

  private method hasLowBalance() {
    return this.balance < 100
  }

  private method notifyLowBalance() {
    // Encapsula la lógica de notificación
    NotificationService.send(this.accountNumber, "Low balance warning")
  }

  // Comportamiento específico para reporting
  method formatForReport() {
    status = this.active ? "Active" : "Inactive"
    return this.accountNumber + ": $" + this.balance + " (" + status + ")"
  }

  // Si realmente necesitas exponer información, hazlo de forma controlada
  method canWithdraw(amount) {
    return this.active and this.balance >= amount
  }
}

class TransferService {
  method transfer(from, to, amount) {
    // Tell, don't ask: le decimos a los objetos qué hacer
    from.withdraw(amount)
    to.deposit(amount)

    // Las validaciones y lógica de negocio están dentro de los objetos
  }
}

class ReportService {
  method generateReport(accounts) {
    report = ""
    for each account in accounts {
      // El objeto sabe cómo representarse
      report = report + account.formatForReport() + "\n"
    }
    return report
  }
}

// Si necesitas transferir información entre sistemas (ej: API, DB),
// usa DTOs en la capa de infraestructura
class BankAccountDTO {
  string accountNumber
  decimal balance
  boolean active

  static method fromDomain(account) {
    // Convierte del objeto de dominio al DTO solo en el boundary
    // Este es el único lugar donde "preguntamos por datos"
    return new BankAccountDTO(
      account.accountNumber,
      account.balance,
      account.active
    )
  }
}
```

**Mejoras**:
- Comportamiento encapsulado en `BankAccount`
- Validaciones dentro del objeto
- Estado modificado solo internamente
- No hay _getters/setters_ públicos
- _Tell, Don't Ask_ aplicado consistentemente
- Lógica de negocio en el lugar apropiado
- DTOs solo en boundaries para serialización

## Ejercicio

**Tarea**: En el código proporcionado en tu lenguaje, elimina getters y setters reemplazándolos con métodos de comportamiento que sigan el principio Tell, Don't Ask.

**Criterios de éxito**:
1. No hay métodos que empiecen con `get` o `set`
2. No hay propiedades públicas
3. El comportamiento está encapsulado en los objetos que poseen los datos
4. El código cliente "le dice" a los objetos qué hacer, no "pregunta" por datos
5. Las validaciones están dentro de los objetos, no en código cliente

## Problemas que Encontrarás

### 1. "¿Cómo serializo/deserializo sin getters?"

En los **boundaries** (API, base de datos, UI), usa:
- **DTOs (Data Transfer Objects)**: Objetos simples solo para transferencia de datos
- **Mappers**: Convierten entre objetos de dominio y DTOs
- **Reflection** o anotaciones del framework (si es necesario)

El dominio debe ser puro; los DTOs son infraestructura.

### 2. "¿Cómo testeo sin verificar estado interno?"

Testea **comportamiento**, no estado:

```pseudocode
// Mal: testear estado
test "should set balance to 100" {
  account = new BankAccount()
  account.setBalance(100)
  assert account.getBalance() == 100
}

// Bien: testear comportamiento
test "should allow withdrawal given sufficient funds" {
  account = new BankAccount("123", 100)
  account.withdraw(50)
  // Verifica comportamiento: ¿puede hacer otra operación?
  assert account.canWithdraw(30) == true
}

test "should prevent withdrawal given insufficient funds" {
  account = new BankAccount("123", 30)
  assertThrows(() => account.withdraw(50))
}
```

### 3. "Frameworks requieren getters/setters"

Opciones:
- Usa DTOs en la capa de infraestructura, objetos de dominio puros internamente
- Configura el framework para usar reflection/field access directo
- Si es inevitable, haz getters/setters package-private o protected
- Considera si realmente necesitas ese framework

### 4. "¿Qué pasa con Value Objects?"

Value Objects **podrían** tener getters porque son inmutables. El riesgo de getters es cuando permiten mutación externa. En Value Objects, no hay mutación.

Aun así, deberían proporcionar métodos de comportamiento en lugar de getters cuando tenga sentido. Los métodos "mutadores" devuelven instancias nuevas con los valores cambiados.

Típicamente, un getter en un Value Object es para obtener una representación de su valor interno para pasarlo a otras capas.

### 5. "¿Cómo obtengo información para mostrar en UI?"

Opciones:
- **Query methods**: Métodos que retornan información preparada para presentación (`formatForDisplay()`)
- **DTOs**: En la capa de presentación, convierte objetos de dominio a DTOs
- **Double dispatch**: El objeto sabe cómo "visitarse" a sí mismo para diferentes propósitos
- **View Models**: Objetos específicos para la capa de presentación

## Proceso de Aplicación

### 1. Identificar getters/setters

Marca todos los métodos que:
- Empiezan con `get` o `set`
- Son propiedades públicas
- Retornan estado interno mutable

### 2. Analizar uso de getters

Para cada getter, encuentra dónde se usa:
- ¿Se usa para tomar decisiones? → Mueve la decisión al objeto
- ¿Se usa para cálculos? → Mueve el cálculo al objeto
- ¿Se usa para validación? → Mueve la validación al objeto
- ¿Se usa para formatear? → Mueve el formato al objeto
- ¿Se usa para serialización? → Usa DTOs en boundaries

### 3. Introducir métodos de comportamiento

Reemplaza:
```pseudocode
// Antes: Ask
if (user.getAge() >= 18) {
  allowAccess()
}

// Después: Tell
if (user.isAdult()) {
  allowAccess()
}

// Aún mejor: Tell completamente
user.accessIfAdult(() => allowAccess())
```

### 4. Eliminar setters

Los setters permiten modificación incontrolada. Reemplaza con:
- **Métodos de comportamiento**: `activate()`, `deactivate()`, no `setActive(bool)`
- **Inmutabilidad**: Retorna nueva instancia en lugar de modificar
- **Constructor/factory**: Establece estado válido desde el inicio

### 5. Aplicar Tell, Don't Ask sistemáticamente

Pregúntate: "¿Estoy preguntando por datos para tomar una decisión?"
- **Sí** → Mueve la decisión al objeto
- **No** → Probablemente está bien

### 6. Usar Double Dispatch si es necesario

Para operaciones complejas donde múltiples objetos interactúan:

```pseudocode
// En lugar de
if (order.getStatus() == "PENDING" and customer.getType() == "VIP") {
  applyVIPDiscount(order)
}

// Usa
order.processForCustomer(customer)  // Double dispatch

// Dentro de Order
method processForCustomer(customer) {
  customer.applyDiscountTo(this)
}

// Dentro de Customer
method applyDiscountTo(order) {
  if (this.isVIP()) {
    order.applyVIPDiscount()
  }
}
```

### 7. Crear DTOs solo en boundaries

```pseudocode
// Controller/API layer
method getAccount(accountId) {
  account = accountRepository.find(accountId)
  return AccountDTO.fromDomain(account)  // Conversión solo aquí
}
```

## Técnicas de Refactoring Aplicables

- **Move Method**: Mover lógica que usa getters al objeto que tiene los datos
- **Extract Method**: Crear métodos de comportamiento a partir de código que usa getters
- **Encapsulate Field**: Hacer campos privados, reemplazar acceso directo con comportamiento
- **Replace Data Value with Object**: Convertir clases anémicas en objetos con comportamiento
- **Introduce Parameter Object**: Si muchos getters se usan juntos, agrúpalos
- **Hide Delegate**: Ocultar objetos internos, proporcionar comportamiento en su lugar
- **Remove Setting Method**: Eliminar setters, hacer objetos inmutables o usar métodos específicos

## Beneficios

### 1. Encapsulación Real

El estado interno está verdaderamente protegido, no solo nominalmente privado.

### 2. Objetos Ricos en Comportamiento

Las clases tienen comportamiento significativo, no solo datos.

### 3. Facilita Evolución

Cambiar representación interna no rompe código cliente (porque no accede a ella).

### 4. Reduce Acoplamiento

El código cliente no depende de la estructura interna de los objetos.

### 5. Lógica de Negocio Centralizada

Validaciones y reglas están en los objetos apropiados, no esparcidas.

### 6. Código más Expresivo

`user.activate()` es más claro que `user.setActive(true)`.

### 7. Testing de Comportamiento

Tests verifican qué hace el objeto, no cómo está estructurado internamente.

### 8. Menos Bugs

Sin mutación externa descontrolada = menos estados inconsistentes.

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/calisthenics-exercises/no-getters-or-setters.ts)
- [Go](../../go/calisthenics_exercises/06_no_getters_setters.go)
- [Java](../../java/src/main/java/com/refactoring/calisthenics/NoGettersSetters.java)
- [PHP](../../php/src/calisthenics-exercises/NoGettersOrSetters.php)
- [Python](../../python/src/calisthenics_exercises/no_getters_or_setters.py)
- [C#](../../csharp/src/calisthenics-exercises/NoGettersOrSetters.cs)

## Referencias en Español

- [Refactor cotidiano (6): cuéntame, no me preguntes (tell, don't ask)](https://franiglesias.github.io/everyday-refactor-6/) - Sobre el principio Tell Don't Ask
- [Calistenias para objetos de valor](https://franiglesias.github.io/calistenics-and-value-objects/) - Incluye discusión sobre getters en Value Objects
- [Diseño orientado a objetos](https://franiglesias.github.io/tag/oop/) - Series de artículos sobre OOP

## Referencias

- **"Practical Object-Oriented Design"** - Sandi Metz - Tell, Don't Ask principle
- **"Growing Object-Oriented Software, Guided by Tests"** - Freeman & Pryce - Testing behavior, not state
- [Object Calisthenics - William Durand](https://williamdurand.fr/2013/06/03/object-calisthenics/) - Regla #6
- [Tell, Don't Ask - Martin Fowler](https://martinfowler.com/bliki/TellDontAsk.html) - Principio fundamental
- **"Domain-Driven Design"** - Eric Evans - Rich domain models vs anemic models
- [Anemic Domain Model - Martin Fowler](https://martinfowler.com/bliki/AnemicDomainModel.html) - Anti-pattern
- [Law of Demeter](https://en.wikipedia.org/wiki/Law_of_Demeter) - Principio relacionado
