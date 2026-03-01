# Large Class

Clase grande.

## Definición

Una clase contiene muchas propiedades, muchos métodos o muchas líneas de código, acumulando muchas responsabilidades no relacionadas o que pueden responder a necesidades diferentes.

## Descripción

Un **Large Class** ocurre cuando una sola clase intenta hacer demasiado, violando el Principio de Responsabilidad Única (SRP). Esta clase suele tener:
- Demasiados campos (más de 10-15)
- Demasiados métodos (más de 15-20)
- Demasiadas líneas de código (más de 200-300)
- Responsabilidades que pertenecen a diferentes áreas del dominio

Este smell aparece gradualmente cuando se añaden características a una clase existente sin cuestionar si realmente pertenecen ahí. La clase se convierte en un "dios objeto" que sabe y hace demasiado.

## Síntomas

- La clase tiene muchos campos de instancia no relacionados entre sí
- Grupos de métodos que solo operan sobre un subconjunto de los campos
- Diferentes razones para modificar la clase (autenticación, notificaciones, permisos, etc.)
- Dificultad para encontrar métodos específicos dentro de la clase
- Tests que requieren mucha configuración para probar una sola funcionalidad
- Comentarios que marcan secciones diferentes dentro de la clase

## Ejemplo

```pseudocode
class UserAccount {
  // Campos de identidad
  name
  email
  password

  // Campos de autenticación
  lastLogin
  loginAttempts

  // Campos de notificaciones
  notifications[]

  // Campos de permisos
  isAdmin

  // Métodos de autenticación
  function login(password)
  function resetPassword(newPassword)

  // Métodos de perfil
  function updateEmail(newEmail)
  function updateName(newName)

  // Métodos de notificaciones
  function addNotification(message)
  function getNotifications()
  function clearNotifications()

  // Métodos de administración
  function promoteToAdmin()
  function revokeAdmin()
}
```

## Ejercicio

Añade soporte para autenticación de dos factores (2FA) y preferencias de notificación.

## Problemas que encontrarás

Tocarás autenticación, estado y notificaciones en una clase inflada, aumentando la probabilidad de romper comportamiento no relacionado. Cada nueva característica hace la clase más difícil de entender y mantener.

## Proceso de Refactoring

### 1. Identificar grupos de responsabilidades
- Analiza los campos y métodos de la clase
- Agrupa campos y métodos que trabajan juntos
- Identifica responsabilidades cohesivas (autenticación, perfil, notificaciones, etc.)

### 2. Extraer clases por responsabilidad
- Crea nuevas clases para cada grupo identificado
- Dale nombres que reflejen su responsabilidad única
- Ejemplo: `AuthenticationManager`, `NotificationCenter`, `UserProfile`

### 3. Mover campos y métodos
- Mueve los campos relacionados a la nueva clase
- Mueve los métodos que operan sobre esos campos
- Mantén la clase original como coordinadora si es necesario

### 4. Actualizar referencias
- Modifica la clase original para delegar a las nuevas clases
- Usa composición en lugar de mantener todo en un solo objeto
- Actualiza los clientes para usar las nuevas clases cuando sea apropiado

### 5. Ejecutar tests y validar
- Verifica que todos los tests pasen después de cada extracción
- Asegúrate de que el comportamiento observable no ha cambiado
- Refina los límites entre clases si es necesario

### 6. Repetir hasta lograr clases cohesivas
- Continúa extrayendo hasta que cada clase tenga una sola razón para cambiar
- Busca el equilibrio entre demasiadas clases pequeñas y pocas clases grandes

## Técnicas de Refactoring Aplicables

- **Extract Class**: Crear nuevas clases para grupos de responsabilidades relacionadas
- **Extract Subclass**: Si hay variaciones de comportamiento, considerar subclases
- **Extract Interface**: Definir contratos claros para cada responsabilidad
- **Move Method**: Mover métodos a la clase donde pertenecen
- **Move Field**: Mover campos a la clase que los usa
- **Replace Data Value with Object**: Convertir grupos de datos primitivos en objetos

## Beneficios

- **Responsabilidad única**: Cada clase tiene una razón clara para existir y cambiar
- **Más fácil de entender**: Clases pequeñas son más fáciles de comprender completamente
- **Más fácil de testear**: Tests más enfocados con menos configuración
- **Mejor reusabilidad**: Componentes pequeños son más fáciles de reutilizar
- **Menor acoplamiento**: Cambios en una responsabilidad no afectan otras
- **Navegación más fácil**: Encontrar código relevante es más rápido

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/bloaters/large-class.ts) - [README](../../typescript/src/code-smells/bloaters/large-class.readme.md)
- [Go](../../go/code_smells/bloaters/large_class.go) - [README](../../go/code_smells/bloaters/large_class.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/bloaters/LargeClass.java) - [README](../../java/src/main/java/com/refactoring/codesmells/bloaters/LargeClass.readme.md)
- [PHP](../../php/src/code-smells/bloaters/LargeClass.php) - [README](../../php/src/code-smells/bloaters/LargeClass.readme.md)
- [Python](../../python/src/code_smells/bloaters/large_class.py) - [README](../../python/src/code_smells/bloaters/large_class_readme.md)
- [C#](../../csharp/src/code-smells/bloaters/LargeClass.cs) - [README](../../csharp/src/code-smells/bloaters/large-class.readme.md)

## Referencias en Español

- [Object Calisthenics para adelgazar las clases](https://franiglesias.github.io/calistenics-and-small-classes/) - Técnicas prácticas para mantener clases pequeñas y cohesivas
- [Los principios SOLID (SRP)](https://franiglesias.github.io/principios-solid/) - Explicación del Principio de Responsabilidad Única y su aplicación

## Referencias

- [Refactoring Guru - Large Class](https://refactoring.guru/smells/large-class)
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
- Robert C. Martin - "Clean Code" - Single Responsibility Principle
