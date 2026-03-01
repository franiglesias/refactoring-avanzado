# Data Class

Clase de datos.

## Definición

Una clase de datos es aquella que solo contiene campos y métodos para acceder a ellos (getters/setters), sin poseer lógica de negocio propia. Esto suele derivar en modelos de dominio anémicos, donde el comportamiento está disperso en otros servicios o clases que manipulan estos datos.

## Descripción

**Data Class** es una clase que actúa como un simple contenedor de datos sin comportamiento significativo. Solo tiene campos, getters, setters, y quizás métodos básicos como `equals()` o `toString()`. Toda la lógica que opera sobre estos datos vive en otras clases.

Este smell es problemático porque:
- Viola el principio de encapsulación
- Separa datos y comportamiento que deberían vivir juntos
- Crea modelos de dominio anémicos
- Resulta en lógica dispersa en múltiples "servicios"
- Dificulta encontrar dónde está el comportamiento relacionado

Las Data Classes son comunes en arquitecturas que separan estrictamente "modelos" de "lógica de negocio", llevando a un estilo procedimental disfrazado de orientación a objetos.

## Síntomas

- La clase solo tiene campos y métodos de acceso (getters/setters)
- No hay lógica de negocio o validación en la clase
- Otras clases (servicios) contienen toda la lógica que opera sobre estos datos
- Los campos son públicos o tienen setters sin restricciones
- La clase no protege sus invariantes
- Múltiples clases manipulan directamente los datos
- Cambios en reglas de negocio requieren tocar múltiples servicios

## Ejemplo

```pseudocode
class UserRecord {
  id: string
  name: string
  email: string
  createdAt: Date

  // Solo constructor y getters/setters implícitos
}

class UserService {
  function createUser(name: string, email: string): UserRecord {
    // Validación en el servicio, no en el modelo
    if (not email.contains("@")) {
      throw error "Invalid email"
    }

    return new UserRecord(generateId(), name, email, now())
  }

  function updateUserEmail(user: UserRecord, newEmail: string) {
    // Validación repetida en otro lugar
    if (not newEmail.contains("@")) {
      throw error "Invalid email"
    }
    user.email = newEmail
  }
}

class UserReportGenerator {
  function generateUserSummary(user: UserRecord): string {
    // Lógica de presentación en otro lugar
    return "User " + user.name + " (" + user.email + ") created on " +
           user.createdAt.toLocaleDateString()
  }
}

// La lógica está dispersa en servicios separados
```

## Ejercicio

Implementa reglas de dominio adicionales, como requerir verificación de email o restringir el registro a ciertos dominios (por ejemplo, `company.com`).

## Problemas que encontrarás

Tendrás que modificar múltiples servicios y lugares que manipulan `UserRecord`. Esto demuestra cómo separar el comportamiento de los datos provoca que cambios simples se dispersen ampliamente por el código (Shotgun Surgery).

## Proceso de Refactoring

### 1. Identificar comportamiento relacionado
- Busca métodos en otras clases que operan sobre la Data Class
- Identifica validaciones, cálculos, y transformaciones
- Agrupa comportamiento por responsabilidad

### 2. Encapsular campos
- Hacer los campos privados
- Reemplaza acceso directo con métodos
- Elimina setters innecesarios o hazlos privados

### 3. Mover validación al modelo
- Mueve validaciones al constructor o métodos setter
- El objeto debe garantizar que siempre está en estado válido
- No permitas crear objetos inválidos

### 4. Mover comportamiento relacionado
- Usa **Move Method** para traer lógica de negocio a la clase
- La clase debe ser responsable de sus propias operaciones
- Ejemplo: `user.updateEmail()` en lugar de `service.updateEmail(user)`

### 5. Encapsular colecciones
- Si la clase tiene colecciones, no las expongas directamente
- Proporciona métodos para añadir, remover, consultar
- Protege la integridad de las colecciones

### 6. Aplicar Tell, Don't Ask
- Los clientes no deberían pedir datos para tomar decisiones
- El objeto debe ofrecer comportamiento, no solo datos
- Ejemplo: `user.canAccessResource()` en lugar de exponer roles y permisos

## Técnicas de Refactoring Aplicables

- **Move Method**: Mover comportamiento de servicios al modelo
- **Encapsulate Field**: Hacer campos privados
- **Remove Setting Method**: Eliminar setters innecesarios
- **Hide Method**: Ocultar métodos que solo son útiles internamente
- **Encapsulate Collection**: Proteger colecciones internas
- **Replace Data Value with Object**: Si hay primitivos que merecen ser objetos

## Beneficios

- **Encapsulación real**: Datos y comportamiento viven juntos
- **Validación centralizada**: Imposible tener objetos inválidos
- **Cohesión mejorada**: Todo lo relacionado está en un lugar
- **Menos duplicación**: La lógica vive en un solo lugar
- **Modelos ricos**: Objetos que representan conceptos del dominio
- **Más fácil de entender**: El comportamiento está donde lo esperas
- **Tell, Don't Ask**: Objetos responsables de sí mismos

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/dispensables/data-class.ts) - [README](../../typescript/src/code-smells/dispensables/data-class.readme.md)
- [Go](../../go/code_smells/dispensables/data_class.go) - [README](../../go/code_smells/dispensables/data_class.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/dispensables/DataClass.java) - [README](../../java/src/main/java/com/refactoring/codesmells/dispensables/DataClass.readme.md)
- [PHP](../../php/src/code-smells/dispensables/DataClass.php) - [README](../../php/src/code-smells/dispensables/DataClass.readme.md)
- [Python](../../python/src/code_smells/dispensables/data_class.py) - [README](../../python/src/code_smells/dispensables/data_class_readme.md)
- [C#](../../csharp/src/code-smells/dispensables/DataClass.cs) - [README](../../csharp/src/code-smells/dispensables/data-class.readme.md)

## Referencias en Español

- [Object Calisthenics para mejorar el diseño de las clases](https://franiglesias.github.io/calistenics-and-value-objects/) - Técnicas para crear objetos con comportamiento rico en lugar de data classes
- [Una introducción a Domain Driven Design (Entities, Value Objects)](https://franiglesias.github.io/ddd-intro/) - Conceptos de DDD para crear modelos de dominio ricos

## Referencias

- [Refactoring Guru - Data Class](https://refactoring.guru/smells/data-class)
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
- Martin Fowler - "Anemic Domain Model" (artículo)
- Eric Evans - "Domain-Driven Design" - Rich Domain Models
