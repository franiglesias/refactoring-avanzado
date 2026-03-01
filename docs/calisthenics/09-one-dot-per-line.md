# No Más de un Punto por Línea

Regla 9 de Object Calisthenics

## Definición

Cada línea de código debe contener como máximo un punto (`.`) de acceso a miembros. Esto previene cadenas de llamadas que revelan la estructura interna de objetos y viola la **Ley de Demeter**.

## Descripción

Las **cadenas de llamadas** (también llamadas "train wrecks" o "accidentes de tren") son líneas como:

```pseudocode
user.getAddress().getCity().getName()
order.getCustomer().getPaymentMethod().getLastDigits()
```

Estos patrones revelan varios problemas:

1. **Violación de encapsulación**: El código conoce la estructura interna profunda de objetos
2. **Alto acoplamiento**: Dependes no solo de `user`, sino también de `Address` y `City`
3. **Ley de Demeter violada**: "Solo habla con tus amigos inmediatos"
4. **Difícil de cambiar**: Cambios en la estructura interna rompen código cliente
5. **Responsabilidad mal ubicada**: El cliente toma decisiones que deberían estar en el objeto

La **Ley de Demeter** (o principio del "menor conocimiento") dice que un método solo debería llamar métodos de:
- El objeto mismo (`this`)
- Parámetros del método
- Objetos que crea
- Variables de instancia

**NO** debería llamar métodos sobre objetos retornados por otros métodos.

La solución es **Tell, Don't Ask**: En lugar de preguntar por objetos internos y operar sobre ellos, dile al objeto qué hacer.

**Excepciones aceptables**:
- Fluent interfaces deliberados (builders, DSLs)
- Cadenas sobre objetos inmutables de utilidad (strings, colecciones funcionales)
- Llamadas a métodos estáticos o constructores

## Síntomas

- Múltiples puntos en una línea (`a.b().c().d()`)
- Cadenas de getters (`getX().getY().getZ()`)
- Navegación profunda por estructura de objetos
- Código que "sabe demasiado" sobre objetos internos
- NullPointerException frecuentes en cadenas largas
- Tests que requieren mockear múltiples niveles de objetos
- Código frágil que rompe con cambios internos de objetos

## Ejemplo

### Antes (Violación)

```pseudocode
class OrderService {
  method processOrder(order) {
    // Violación: múltiples puntos revelando estructura
    customerEmail = order.getCustomer().getContactInfo().getEmail().getValue()
    sendEmail(customerEmail, "Order confirmation")

    // Violación: navegación profunda
    if (order.getCustomer().getAddress().getCountry().getCode() == "US") {
      applyUSTax(order)
    }

    // Violación: lógica de negocio basada en estructura interna
    if (order.getPayment().getMethod().getType() == "CREDIT_CARD" and
        order.getPayment().getMethod().getCard().getExpiryDate() < today()) {
      throw "Credit card expired"
    }

    // Violación: cálculos sobre estructura profunda
    total = 0
    for each item in order.getItems().getList() {
      total = total + item.getProduct().getPrice().getAmount()
    }

    // Violación: modificación profunda
    order.getShipping().getAddress().setStreet("New street")
  }
}

class ReportGenerator {
  method generateCustomerReport(customer) {
    // Violación: acceso a estructura profunda para formateo
    report = "Name: " + customer.getProfile().getName().getFirstName() + " " +
                         customer.getProfile().getName().getLastName() + "\n"
    report = report + "Email: " + customer.getProfile().getContact().getEmail().getValue() + "\n"
    report = report + "City: " + customer.getProfile().getContact().getAddress().getCity().getName() + "\n"
    return report
  }
}
```

**Problemas**:
- Alto acoplamiento a estructura interna
- Múltiples niveles de navegación
- Lógica de negocio en el lugar equivocado
- Frágil a cambios en estructura
- Difícil de testear (muchos niveles de mocking)

### Después (Cumplimiento)

```pseudocode
// Opción 1: Tell, Don't Ask - Delegar comportamiento
class OrderService {
  EmailService emailService

  method processOrder(order) {
    // Un solo punto: le decimos al order qué hacer
    order.sendConfirmationEmail(emailService)

    // Un solo punto: el order sabe aplicar sus propios impuestos
    order.applyTaxesIfNeeded()

    // Un solo punto: el order valida su propio pago
    order.validatePayment()

    // Un solo punto: el order calcula su propio total
    total = order.calculateTotal()
  }
}

class Order {
  Customer customer
  Payment payment
  OrderItems items
  Shipping shipping

  method sendConfirmationEmail(emailService) {
    // El order delega al customer
    this.customer.sendOrderConfirmation(emailService, this)
  }

  method applyTaxesIfNeeded() {
    // El order pregunta al customer por el país (un solo punto)
    if (this.customer.isFromUS()) {
      this.applyUSTax()
    }
  }

  method validatePayment() {
    // El payment sabe validarse a sí mismo
    this.payment.ensureValid()
  }

  method calculateTotal() {
    // Los items saben calcular su total
    return this.items.calculateTotal()
  }

  private method applyUSTax() {
    // Lógica de impuestos
  }
}

class Customer {
  Profile profile
  Address address

  method sendOrderConfirmation(emailService, order) {
    // Profile sabe enviar email
    this.profile.sendEmail(emailService, "Order confirmation", order)
  }

  method isFromUS() {
    // Address sabe su país
    return this.address.isCountry("US")
  }

  method formatForReport() {
    // Customer sabe cómo formatearse para reporte
    return this.profile.formatForReport() + "\n" +
           this.address.formatForReport()
  }
}

class Profile {
  Name name
  ContactInfo contact

  method sendEmail(emailService, subject, body) {
    // ContactInfo sabe su email
    email = this.contact.emailAddress()
    emailService.send(email, subject, body)
  }

  method formatForReport() {
    return "Name: " + this.name.full() + "\n" +
           "Email: " + this.contact.emailAddress()
  }
}

class ContactInfo {
  Email email
  Address address

  method emailAddress() {
    return this.email.value()
  }

  method formatForReport() {
    return this.address.cityName()
  }
}

class Address {
  City city
  Country country

  method isCountry(countryCode) {
    return this.country.hasCode(countryCode)
  }

  method cityName() {
    return this.city.name()
  }

  method formatForReport() {
    return "City: " + this.city.name()
  }
}

class Payment {
  PaymentMethod method

  method ensureValid() {
    // PaymentMethod sabe validarse
    this.method.ensureNotExpired()
  }
}

class PaymentMethod {
  CreditCard card

  method ensureNotExpired() {
    // CreditCard sabe si está expirada
    if (this.card.isExpired()) {
      throw "Credit card expired"
    }
  }
}

class OrderItems {
  array items

  method calculateTotal() {
    total = 0
    for each item in this.items {
      // Un solo punto: cada item sabe su precio
      total = total + item.totalPrice()
    }
    return total
  }
}

class OrderItem {
  Product product
  int quantity

  method totalPrice() {
    // Product sabe su precio
    return this.product.price() * this.quantity
  }
}

class ReportGenerator {
  method generateCustomerReport(customer) {
    // Un solo punto: customer sabe formatearse
    return customer.formatForReport()
  }
}
```

**Mejoras**:
- Cada objeto expone comportamiento, no estructura
- Responsabilidad distribuida apropiadamente
- Un solo nivel de indirección por línea
- Bajo acoplamiento
- Fácil de testear (mockear solo dependencias directas)
- Cambios internos no afectan código cliente

## Ejercicio

**Tarea**: En el código proporcionado en tu lenguaje, identifica cadenas de llamadas con múltiples puntos y refactorízalas delegando comportamiento apropiadamente.

**Criterios de éxito**:
1. Ninguna línea tiene más de un punto (excepto fluent APIs legítimos)
2. El comportamiento está encapsulado en los objetos apropiados
3. No hay navegación profunda por estructuras de objetos
4. Se aplica Tell, Don't Ask consistentemente
5. La Ley de Demeter se respeta

## Problemas que Encontrarás

### 1. "Necesito ese dato profundo"

Si realmente necesitas información profunda, hay varias opciones:

**Opción A**: Delega la responsabilidad
```pseudocode
// Mal
if (order.getCustomer().getAddress().getCity().getPopulation() > 1000000) {
  applyBigCityDiscount()
}

// Bien
if (order.isFromBigCity()) {
  applyBigCityDiscount()
}
```

**Opción B**: Usa Double Dispatch
```pseudocode
order.applyDiscountBasedOnLocation(discountCalculator)
```

**Opción C**: Si es para presentación, usa DTOs en boundaries
```pseudocode
// En capa de presentación, no en dominio
dto = OrderDTO.fromDomain(order)
display(dto.customerCityName)
```

### 2. "Ahora tengo métodos delegadores vacíos"

Sí, tendrás métodos que solo delegan. Esto es correcto porque:
- Ocultan la estructura interna
- Son puntos de extensión futuros
- Hacen el código cliente más expresivo
- Facilitan cambios futuros en la estructura

### 3. "¿Qué pasa con fluent interfaces?"

Fluent interfaces y builders son excepciones legítimas:
```pseudocode
query = QueryBuilder.new()
  .select("*")
  .from("users")
  .where("age > 18")
  .build()
```

Esto está bien porque:
- Es una API deliberadamente diseñada así
- Cada método retorna el mismo objeto (builder pattern)
- No está revelando estructura interna

### 4. "¿Y las operaciones funcionales en colecciones?"

También son excepciones aceptables:
```pseudocode
result = list
  .filter(x => x > 10)
  .map(x => x * 2)
  .reduce((a, b) => a + b)
```

Esto está bien porque:
- Son transformaciones sobre inmutables
- Es parte de la API estándar del lenguaje
- No está acoplando a estructura interna

### 5. "Creo muchos métodos triviales"

Sí, pero son importantes porque:
- Documentan la intención
- Desacoplan de la estructura
- Son puntos de extensión
- Facilitan testing

## Proceso de Aplicación

### 1. Identificar cadenas de llamadas

Busca líneas con múltiples puntos:
```pseudocode
a.b().c().d()
```

### 2. Analizar la intención

Para cada cadena:
- ¿Qué está intentando hacer el código?
- ¿Es obtener información?
- ¿Es ejecutar una acción?
- ¿Es validar algo?

### 3. Decidir dónde debe estar la responsabilidad

Pregúntate: "¿Qué objeto debería conocer/hacer esto?"
- Típicamente, el objeto más profundo en la cadena tiene el conocimiento
- El objeto de nivel superior debe delegar

### 4. Crear métodos delegadores

En cada nivel de la cadena, crea un método que delegue al siguiente:

```pseudocode
// Antes
city = customer.getProfile().getContact().getAddress().getCity()

// Paso 1: Customer delega a Profile
class Customer {
  method city() {
    return this.profile.city()
  }
}

// Paso 2: Profile delega a ContactInfo
class Profile {
  method city() {
    return this.contact.city()
  }
}

// Paso 3: ContactInfo delega a Address
class ContactInfo {
  method city() {
    return this.address.city()
  }
}

// Uso
city = customer.city()  // Un solo punto
```

### 5. Mover lógica al lugar apropiado

Si hay lógica de negocio, muévela al objeto que tiene el conocimiento:

```pseudocode
// Antes
if (customer.getProfile().getContact().getAddress().getCountry().getCode() == "US") {
  applyUSTax()
}

// Después
if (customer.isFromUS()) {
  applyUSTax()
}

// Implementación
class Customer {
  method isFromUS() {
    return this.profile.isFromUS()
  }
}

class Profile {
  method isFromUS() {
    return this.contact.isFromUS()
  }
}

class ContactInfo {
  method isFromUS() {
    return this.address.isFromUS()
  }
}

class Address {
  method isFromUS() {
    return this.country.isUS()
  }
}

class Country {
  method isUS() {
    return this.code == "US"
  }
}
```

### 6. Aplicar Tell, Don't Ask

En lugar de pedir datos y operar sobre ellos, pídele al objeto que haga la operación:

```pseudocode
// Antes (Ask)
email = customer.getProfile().getContact().getEmail().getValue()
emailService.send(email, "Subject", "Body")

// Después (Tell)
customer.sendEmail(emailService, "Subject", "Body")
```

### 7. Usar Hide Delegate refactoring

Oculta la estructura interna sistemáticamente:
- Cada objeto solo expone su API pública
- La estructura interna está completamente oculta
- Cambios internos no afectan código cliente

## Técnicas de Refactoring Aplicables

- **Hide Delegate**: Ocultar objetos internos, proporcionar métodos delegadores
- **Move Method**: Mover lógica al objeto que tiene el conocimiento
- **Extract Method**: Crear métodos para operaciones que cruzan múltiples objetos
- **Introduce Foreign Method**: Si no puedes modificar una clase, crea método en la tuya
- **Replace Temp with Query**: Eliminar variables temporales que almacenan resultados de cadenas

## Beneficios

### 1. Bajo Acoplamiento

El código cliente no depende de la estructura interna de objetos.

### 2. Encapsulación Real

La estructura interna está verdaderamente oculta.

### 3. Facilita Cambios

Modificar estructura interna no rompe código cliente.

### 4. Código Expresivo

Los métodos delegadores documentan intenciones claramente.

### 5. Testing Simplificado

Solo necesitas mockear dependencias directas, no cadenas profundas.

### 6. Menos NullPointerException

Sin cadenas largas, menos puntos donde null puede aparecer.

### 7. Ley de Demeter

Aplicación estricta del principio de "menor conocimiento".

### 8. Better Error Messages

Errores ocurren en el objeto apropiado, no en el código cliente.

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/calisthenics-exercises/only-one-dot-per-line.ts)
- [Go](../../go/calisthenics_exercises/05_one_dot_per_line.go)
- [Java](../../java/src/main/java/com/refactoring/calisthenics/OneDotPerLine.java)
- [PHP](../../php/src/calisthenics-exercises/OnlyOneDotPerLine.php)
- [Python](../../python/src/calisthenics_exercises/only_one_dot_per_line.py)
- [C#](../../csharp/src/calisthenics-exercises/OnlyOneDotPerLine.cs)

## Referencias en Español

- [Refactor cotidiano (6): cuéntame, no me preguntes (tell, don't ask)](https://franiglesias.github.io/everyday-refactor-6/) - Principio Tell Don't Ask aplicado
- [Object Calisthenics para adelgazar las clases](https://franiglesias.github.io/calistenics-and-small-classes/) - Incluye discusión sobre encapsulación

## Referencias

- **"The Pragmatic Programmer"** - Hunt & Thomas - Discusión sobre la Ley de Demeter
- **"Refactoring: Improving the Design of Existing Code"** - Martin Fowler - Hide Delegate refactoring
- [Object Calisthenics - William Durand](https://williamdurand.fr/2013/06/03/object-calisthenics/) - Regla #9
- [Law of Demeter](https://en.wikipedia.org/wiki/Law_of_Demeter) - Principio fundamental
- [Tell, Don't Ask - Martin Fowler](https://martinfowler.com/bliki/TellDontAsk.html) - Principio relacionado
- [Train Wreck Code](https://wiki.c2.com/?TrainWreck) - Anti-pattern que esta regla previene
- **"Growing Object-Oriented Software, Guided by Tests"** - Freeman & Pryce - Diseño basado en Tell Don't Ask
