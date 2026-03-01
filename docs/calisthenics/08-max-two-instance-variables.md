# Máximo de Dos Variables de Instancia

Regla 8 de Object Calisthenics

## Definición

Cada clase debe tener un máximo de dos variables de instancia. Si una clase necesita más, probablemente está gestionando múltiples conceptos que deberían encapsularse en objetos separados.

## Descripción

Esta es quizás la **regla más estricta y controversial** de Object Calisthenics. A primera vista parece imposible, pero su propósito es forzar **cohesión extrema** y **descomposición en Value Objects**.

El problema de clases con muchas variables de instancia:

1. **Baja cohesión**: Variables no relacionadas entre sí en la misma clase
2. **Conceptos ocultos**: Grupos de variables que representan conceptos del dominio sin nombrar
3. **Primitive Obsession**: Usar muchos primitivos en lugar de objetos
4. **Dificultad para testear**: Muchas combinaciones posibles de estado
5. **Violación de SRP**: Clase gestiona múltiples responsabilidades

Cuando una clase tiene 5-10 variables de instancia, típicamente:
- Algunas variables están relacionadas entre sí (representan un concepto)
- La clase gestiona 2-3 conceptos diferentes simultáneamente
- Hay primitivos que deberían ser Value Objects

La solución es **agrupar variables relacionadas en Value Objects** o **extraer responsabilidades a clases separadas**.

Esta regla es extrema intencionalmente para entrenar el instinto de descomposición. En código real, 3-4 variables puede ser razonable, pero si superas 5-6, probablemente hay problemas de diseño.

## Síntomas

- Clases con más de 5-6 variables de instancia
- Variables primitivas que se pasan juntas a métodos
- Grupos de variables que siempre se validan/usan juntas
- Nombres de variables con prefijos comunes (`user_name`, `user_email`, `user_age`)
- Métodos que solo usan un subconjunto de las variables
- Constructor con muchos parámetros
- Dificultad para testear (muchas combinaciones de estado)
- Setters para múltiples variables relacionadas

## Ejemplo

### Antes (Violación)

```pseudocode
class User {
  // 10 variables de instancia
  string firstName
  string lastName
  string email
  string phone
  string street
  string city
  string postalCode
  string country
  int age
  string membershipType

  constructor(firstName, lastName, email, phone, street, city, postalCode, country, age, membershipType) {
    // Validaciones duplicadas para conceptos relacionados
    if (not contains(email, "@")) throw "Invalid email"
    if (age < 0 or age > 150) throw "Invalid age"
    if (postalCode.length != 5) throw "Invalid postal code"

    this.firstName = firstName
    this.lastName = lastName
    this.email = email
    this.phone = phone
    this.street = street
    this.city = city
    this.postalCode = postalCode
    this.country = country
    this.age = age
    this.membershipType = membershipType
  }

  method getFullName() {
    return this.firstName + " " + this.lastName
  }

  method getFullAddress() {
    return this.street + ", " + this.city + " " + this.postalCode + ", " + this.country
  }

  method isAdult() {
    return this.age >= 18
  }

  method isPremiumMember() {
    return this.membershipType == "PREMIUM"
  }

  method canReceivePromotions() {
    return this.email != null and contains(this.email, "@")
  }

  // Muchos métodos que operan sobre subconjuntos de variables
}
```

**Problemas**:
- 10 variables de instancia (demasiadas)
- Varios conceptos mezclados: nombre, contacto, dirección, edad, membresía
- Primitivos sin encapsular
- Validaciones duplicadas
- Métodos que agrupan variables relacionadas (síntoma de conceptos ocultos)

### Después (Cumplimiento)

```pseudocode
// Value Object: Name (2 variables → 1 objeto)
class Name {
  private string firstName
  private string lastName

  constructor(firstName, lastName) {
    this.ensureNotEmpty(firstName, "First name")
    this.ensureNotEmpty(lastName, "Last name")
    this.firstName = firstName
    this.lastName = lastName
  }

  method full() {
    return this.firstName + " " + this.lastName
  }

  method first() {
    return this.firstName
  }

  method last() {
    return this.lastName
  }

  private method ensureNotEmpty(value, field) {
    if (value == null or value.trim() == "") {
      throw field + " cannot be empty"
    }
  }
}

// Value Object: Email
class Email {
  private string value

  constructor(email) {
    if (not contains(email, "@")) {
      throw "Invalid email format"
    }
    this.value = toLowerCase(email)
  }

  method getValue() {
    return this.value
  }

  method domain() {
    return split(this.value, "@")[1]
  }
}

// Value Object: Phone
class PhoneNumber {
  private string value

  constructor(phone) {
    this.validate(phone)
    this.value = phone
  }

  private method validate(phone) {
    // Validación específica
    if (phone.length < 10) {
      throw "Invalid phone number"
    }
  }

  method getValue() {
    return this.value
  }
}

// Value Object: Address (4 variables → 1 objeto)
class Address {
  private string street
  private string city
  private PostalCode postalCode
  private Country country

  constructor(street, city, postalCode, country) {
    this.street = street
    this.city = city
    this.postalCode = postalCode
    this.country = country
  }

  method full() {
    return this.street + ", " + this.city.name() + " " +
           this.postalCode.value() + ", " + this.country.name()
  }

  method getCity() {
    return this.city
  }

  method getCountry() {
    return this.country
  }
}

// Value Object: PostalCode
class PostalCode {
  private string value

  constructor(code) {
    if (code.length != 5) {
      throw "Invalid postal code"
    }
    this.value = code
  }

  method value() {
    return this.value
  }
}

// Value Object: Country
class Country {
  private string code

  constructor(code) {
    this.ensureValidCountry(code)
    this.code = code
  }

  private method ensureValidCountry(code) {
    // Validación de código de país
  }

  method name() {
    // Retornar nombre completo del país
    return countryNames[this.code]
  }

  method code() {
    return this.code
  }
}

// Value Object: Age
class Age {
  private int years

  constructor(years) {
    if (years < 0 or years > 150) {
      throw "Invalid age"
    }
    this.years = years
  }

  method value() {
    return this.years
  }

  method isAdult() {
    return this.years >= 18
  }

  method isSenior() {
    return this.years >= 65
  }
}

// Value Object: MembershipType
class MembershipType {
  private string type

  constructor(type) {
    allowedTypes = ["BASIC", "PREMIUM", "VIP"]
    if (not contains(allowedTypes, type)) {
      throw "Invalid membership type"
    }
    this.type = type
  }

  method isPremium() {
    return this.type == "PREMIUM"
  }

  method isVIP() {
    return this.type == "VIP"
  }

  method getDiscountPercentage() {
    discounts = {"BASIC": 0, "PREMIUM": 10, "VIP": 20}
    return discounts[this.type]
  }
}

// Clase User refactorizada: 2 variables de instancia
class User {
  UserProfile profile        // Variable 1: información personal
  UserMembership membership  // Variable 2: información de membresía

  constructor(profile, membership) {
    this.profile = profile
    this.membership = membership
  }

  method canReceivePromotions() {
    return this.profile.hasValidEmail()
  }

  method getDiscount() {
    return this.membership.getDiscountPercentage()
  }

  method profile() {
    return this.profile
  }

  method membership() {
    return this.membership
  }
}

// Clase auxiliar: UserProfile agrupa información personal
class UserProfile {
  Name name           // Variable 1
  ContactInfo contact // Variable 2

  constructor(name, contact) {
    this.name = name
    this.contact = contact
  }

  method fullName() {
    return this.name.full()
  }

  method hasValidEmail() {
    return this.contact.hasValidEmail()
  }

  method address() {
    return this.contact.address()
  }
}

// Clase auxiliar: ContactInfo agrupa información de contacto
class ContactInfo {
  Email email        // Variable 1
  PhoneNumber phone  // Variable 2
  Address address    // Podría moverse a otra clase si es necesario

  constructor(email, phone, address) {
    this.email = email
    this.phone = phone
    this.address = address
  }

  method email() {
    return this.email
  }

  method phone() {
    return this.phone
  }

  method address() {
    return this.address
  }

  method hasValidEmail() {
    return this.email != null
  }
}

// Clase auxiliar: UserMembership agrupa información de membresía
class UserMembership {
  MembershipType type  // Variable 1
  Age age              // Variable 2

  constructor(type, age) {
    this.type = type
    this.age = age
  }

  method isPremium() {
    return this.type.isPremium()
  }

  method getDiscountPercentage() {
    return this.type.getDiscountPercentage()
  }

  method isAdult() {
    return this.age.isAdult()
  }
}
```

**Mejoras**:
- User tiene solo 2 variables de instancia
- Cada concepto está encapsulado en un Value Object
- Validaciones centralizadas
- Comportamiento relacionado agrupado
- Fácil de testear (cada Value Object independientemente)
- Expresividad del dominio maximizada
- Sin primitivos expuestos

## Ejercicio

**Tarea**: En el código proporcionado en tu lenguaje, identifica clases con más de dos variables de instancia y refactorízalas agrupando variables relacionadas en Value Objects.

**Criterios de éxito**:
1. Ninguna clase tiene más de 2 variables de instancia (o máximo 3-4 si es absolutamente necesario)
2. Variables relacionadas están agrupadas en Value Objects
3. Cada Value Object tiene validación y comportamiento apropiado
4. Los conceptos del dominio están explícitamente nombrados
5. No hay primitivos sin encapsular

## Problemas que Encontrarás

### 1. "Es imposible con solo 2 variables"

No es imposible, pero requiere pensar diferente:
- Agrupa variables relacionadas en objetos
- Separa responsabilidades en clases distintas
- Usa composición agresivamente

En práctica, 3-4 variables puede ser razonable, pero más de 5-6 es sospechoso.

### 2. "Demasiadas clases pequeñas"

Sí, tendrás muchos Value Objects. Esto es **bueno**:
- Cada uno es testeable independientemente
- Cada uno encapsula validación y comportamiento
- El dominio está explícitamente modelado
- Es más fácil entender `Address` que 4 strings sueltos

### 3. "¿Dónde pongo todos estos Value Objects?"

Organización sugerida:
```
domain/
  model/
    User.java
    UserProfile.java
    UserMembership.java
  valueobjects/
    Name.java
    Email.java
    Address.java
    Age.java
    MembershipType.java
```

### 4. "Performance de tantos objetos pequeños"

El overhead es mínimo:
- Los lenguajes modernos están optimizados para muchos objetos pequeños
- Los Value Objects son típicamente inmutables (fácil de optimizar)
- La ganancia en mantenibilidad supera infinitamente cualquier costo
- Si tienes un cuello de botella real, optimiza ese punto específico

### 5. "¿Qué pasa con clases de infraestructura?"

Esta regla aplica principalmente a **clases de dominio**. Clases de infraestructura (controladores, repositorios, configuración) pueden tener más variables si:
- Cada variable es una dependencia inyectada
- La clase tiene una responsabilidad clara
- Las variables no representan conceptos del dominio sin encapsular

## Proceso de Aplicación

### 1. Identificar clases con muchas variables

Busca clases con más de 5-6 variables de instancia.

### 2. Agrupar variables relacionadas

Para cada clase:
- Lista todas las variables
- Identifica cuáles se usan juntas
- Busca nombres con prefijos comunes
- Identifica conceptos del dominio ocultos

### 3. Crear Value Objects

Por cada grupo de variables relacionadas:
- Crea un Value Object que las encapsule
- Añade validación en el constructor
- Añade comportamiento relacionado
- Hazlo inmutable

### 4. Extraer responsabilidades

Si después de agrupar aún tienes muchas variables:
- Probablemente la clase tiene múltiples responsabilidades
- Separa en clases distintas (ej: UserProfile, UserMembership)
- Usa composición

### 5. Reemplazar primitivos

- Cada primitivo que representa un concepto del dominio → Value Object
- Strings (emails, URLs, códigos) → Value Objects
- Números (dinero, cantidades, porcentajes) → Value Objects
- Booleanos (estados) → considerar State pattern o Value Objects

### 6. Actualizar constructor

```pseudocode
// Antes: 10 parámetros primitivos
constructor(firstName, lastName, email, phone, street, city, code, country, age, type)

// Después: 2 objetos compuestos
constructor(profile, membership)

// Con factory method para facilitar construcción
static method create(firstName, lastName, email, phone, street, city, code, country, age, type) {
  name = new Name(firstName, lastName)
  emailObj = new Email(email)
  phoneObj = new PhoneNumber(phone)
  address = new Address(street, city, new PostalCode(code), new Country(country))
  contact = new ContactInfo(emailObj, phoneObj, address)
  profile = new UserProfile(name, contact)

  ageObj = new Age(age)
  typeObj = new MembershipType(type)
  membership = new UserMembership(typeObj, ageObj)

  return new User(profile, membership)
}
```

### 7. Mover comportamiento apropiadamente

- Métodos que usan solo un subconjunto de variables → Muévelos al Value Object correspondiente
- Métodos que coordinan → Quedan en la clase principal

## Técnicas de Refactoring Aplicables

- **Extract Class**: Separar responsabilidades en clases distintas
- **Replace Data Value with Object**: Convertir primitivos en Value Objects
- **Introduce Parameter Object**: Agrupar parámetros relacionados
- **Preserve Whole Object**: Pasar objetos completos en lugar de partes
- **Move Method**: Mover comportamiento a Value Objects apropiados
- **Replace Type Code with Class**: Convertir códigos de tipo en objetos

## Beneficios

### 1. Alta Cohesión

Cada clase tiene variables fuertemente relacionadas entre sí.

### 2. Conceptos del Dominio Explícitos

Value Objects nombran conceptos que antes estaban implícitos.

### 3. Validación Centralizada

Cada Value Object valida sus propias invariantes.

### 4. Reutilización

Value Objects pueden usarse en múltiples contextos.

### 5. Testing Simplificado

Cada Value Object es testeable independientemente con menos combinaciones de estado.

### 6. Inmutabilidad

Value Objects típicamente son inmutables, eliminando bugs de mutación.

### 7. Evolución del Código

Cambiar representación de un concepto está localizado en su Value Object.

### 8. Sin Primitive Obsession

Todos los conceptos del dominio están modelados como objetos, no primitivos.

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/calisthenics-exercises/not-more-than-2-instance-variables.ts)
- [Go](../../go/calisthenics_exercises/08_max_two_instance_variables.go)
- [Java](../../java/src/main/java/com/refactoring/calisthenics/MaxTwoInstanceVariables.java)
- [PHP](../../php/src/calisthenics-exercises/NotMoreThanTwoInstanceVariables.php)
- [Python](../../python/src/calisthenics_exercises/not_more_than_2_instance_variables.py)
- [C#](../../csharp/src/calisthenics-exercises/NotMoreThan2InstanceVariables.cs)

## Referencias en Español

- [Calistenias para objetos de valor](https://franiglesias.github.io/calistenics-and-value-objects/) - Aplicación directa de esta regla
- [Primitive obsession](https://franiglesias.github.io/primitive-obsession/) - Problema que esta regla ayuda a resolver
- [Value Objects](https://franiglesias.github.io/value-objects/) - Guía sobre creación de Value Objects

## Referencias

- **"Domain-Driven Design"** - Eric Evans - Value Objects como building blocks
- **"Implementing Domain-Driven Design"** - Vaughn Vernon - Capítulo sobre Value Objects y composición
- [Object Calisthenics - William Durand](https://williamdurand.fr/2013/06/03/object-calisthenics/) - Regla #8
- **"Refactoring: Improving the Design of Existing Code"** - Martin Fowler - Replace Data Value with Object
- [Primitive Obsession - Refactoring Guru](https://refactoring.guru/smells/primitive-obsession)
- [Single Responsibility Principle](https://en.wikipedia.org/wiki/Single-responsibility_principle) - Principio relacionado
