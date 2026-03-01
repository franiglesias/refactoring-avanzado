# Divergent Change

Cambio divergente.

## Definición

Una clase tiene múltiples razones para cambiar, lo que normalmente indica que se ocupa de muchas responsabilidades que se deberían separar en clases especialistas más pequeñas.

## Descripción

**Divergent Change** ocurre cuando una sola clase necesita ser modificada de diferentes maneras por diferentes razones. Es una violación directa del Principio de Responsabilidad Única (Single Responsibility Principle). La clase se convierte en un punto focal de cambios que no están relacionados entre sí.

Por ejemplo, una clase puede cambiar:
- Cuando cambia el formato de exportación de datos
- Cuando cambia la lógica de validación
- Cuando cambia el método de persistencia
- Cuando cambia el sistema de mensajería

Cada tipo de cambio debería afectar a una sola clase, no a múltiples responsabilidades dentro de la misma clase.

## Síntomas

- Al hacer un cambio pequeño, necesitas modificar múltiples métodos en la clase
- Diferentes programadores modifican la misma clase por razones no relacionadas
- La clase tiene métodos que claramente pertenecen a diferentes categorías
- Cambios en requisitos externos (formato, base de datos, APIs) siempre tocan esta clase
- La clase tiene dependencias en múltiples sistemas externos no relacionados
- Tests de diferentes aspectos requieren configuraciones completamente diferentes

## Ejemplo

```pseudocode
class ProfileManager {
  users: Map

  // Responsabilidad: Validación
  function register(user) {
    if (user.name is empty) {
      throw error "invalid name"
    }
    if (not user.email.contains("@")) {
      throw error "invalid email"
    }
    users.add(user.id, user)
  }

  function updateEmail(id, newEmail) {
    if (not newEmail.contains("@")) {
      throw error "invalid email"
    }
    user = users.get(id)
    if (user is null) {
      throw error "not found"
    }
    users.update(id, user with email = newEmail)
  }

  // Responsabilidad: Exportación
  function exportAsJson() {
    return toJson(users.values())
  }

  function exportAsCsv() {
    rows = ["id,name,email"]
    for each user in users.values() {
      rows.add(user.id + "," + user.name + "," + user.email)
    }
    return rows.join("\n")
  }

  // Responsabilidad: Mensajería
  function sendWelcomeEmail(user) {
    return "Welcome " + user.name + "! Sent to " + user.email
  }
}
```

## Ejercicio

Añade un número de teléfono con validación, inclúyelo en las exportaciones y envía un SMS.

## Problemas que encontrarás

Tocarás validación, almacenamiento, exportAsJson/Csv y mensajería en un solo lugar, demostrando cómo un cambio fuerza ediciones en responsabilidades no relacionadas.

## Proceso de Refactoring

### 1. Identificar las diferentes razones de cambio
- Analiza los métodos de la clase
- Agrupa métodos por el tipo de cambio que los afectaría
- Identifica las responsabilidades distintas (validación, persistencia, exportación, etc.)

### 2. Crear clases especializadas
- Para cada grupo de responsabilidad, crea una nueva clase
- Dale un nombre que refleje su única responsabilidad
- Ejemplo: `UserValidator`, `UserRepository`, `UserExporter`, `UserNotifier`

### 3. Extraer métodos a las nuevas clases
- Mueve los métodos relacionados a la clase apropiada
- Asegúrate de que cada clase tiene acceso a los datos que necesita
- Usa inyección de dependencias si las clases necesitan colaborar

### 4. Actualizar la clase original
- Si es necesario, mantén la clase original como coordinadora
- Delega a las clases especializadas en lugar de hacer el trabajo directamente
- Considera si la clase original sigue siendo necesaria

### 5. Reorganizar dependencias
- Cada clase especializada debe depender solo de lo que necesita
- Ejemplo: `UserValidator` no necesita saber nada de exportación
- Reduce el acoplamiento entre responsabilidades

### 6. Validar con tests
- Los tests deberían poder enfocarse en una sola responsabilidad
- Es más fácil testear validación sin tener que configurar exportación
- Los cambios en una responsabilidad no deberían romper tests de otra

## Técnicas de Refactoring Aplicables

- **Extract Class**: Crear nuevas clases para cada responsabilidad
- **Move Method**: Mover métodos a la clase donde pertenecen
- **Move Field**: Mover campos a la clase que los usa
- **Extract Interface**: Definir contratos claros para cada responsabilidad
- **Introduce Facade**: Si necesitas una interfaz simple sobre clases especializadas

## Beneficios

- **Cambios localizados**: Cada cambio afecta solo a una clase
- **Mejor organización**: Código relacionado vive junto
- **Tests más enfocados**: Tests simples que prueban una sola cosa
- **Paralelización**: Diferentes desarrolladores pueden trabajar en diferentes responsabilidades
- **Menos conflictos**: Cambios en áreas diferentes no colisionan
- **Comprensión más fácil**: Cada clase tiene un propósito claro
- **Cumplimiento de SRP**: Una razón para cambiar por clase

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/change-preventers/divergent-change.ts) - [README](../../typescript/src/code-smells/change-preventers/divergent-change.readme.md)
- [Go](../../go/code_smells/change_preventers/divergent_change.go) - [README](../../go/code_smells/change_preventers/divergent_change.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/changepreventers/DivergentChange.java) - [README](../../java/src/main/java/com/refactoring/codesmells/changepreventers/DivergentChange.readme.md)
- [PHP](../../php/src/code-smells/change-preventers/DivergentChange.php) - [README](../../php/src/code-smells/change-preventers/DivergentChange.readme.md)
- [Python](../../python/src/code_smells/change_preventers/divergent_change.py) - [README](../../python/src/code_smells/change_preventers/divergent_change_readme.md)
- [C#](../../csharp/src/code-smells/change-preventers/DivergentChange.cs) - [README](../../csharp/src/code-smells/change-preventers/divergent-change.readme.md)

## Referencias en Español

- [Los principios SOLID (SRP)](https://franiglesias.github.io/principios-solid/) - Explicación detallada del Principio de Responsabilidad Única
- [Más allá de SOLID, los cimientos](https://franiglesias.github.io/beyond-solid-2/) - Fundamentos de diseño que soportan SOLID y la cohesión

## Referencias

- [Refactoring Guru - Divergent Change](https://refactoring.guru/smells/divergent-change)
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
- Robert C. Martin - "Agile Software Development" - Single Responsibility Principle
