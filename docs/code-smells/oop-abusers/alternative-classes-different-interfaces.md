# Alternative Classes with Different Interfaces

Clases alternativas con interfaces diferentes.

## Definición

Ocurre cuando dos clases realizan tareas similares o son conceptualmente intercambiables, pero exponen nombres de métodos diferentes. Esto impide el uso de polimorfismo y obliga a los clientes a escribir código condicional para decidir qué método llamar dependiendo de la clase que estén usando.

## Descripción

**Alternative Classes with Different Interfaces** es un smell que aparece cuando tienes múltiples clases que hacen esencialmente lo mismo pero con interfaces incompatibles. Aunque las clases son funcionalmente similares o intercambiables, no pueden ser usadas polimórficamente porque tienen métodos con nombres o firmas diferentes.

Este problema surge cuando:
- Diferentes equipos o desarrolladores crean soluciones similares independientemente
- Se integra código de diferentes fuentes sin unificar interfaces
- Se evolucionan clases similares en direcciones divergentes
- No se aplica un diseño orientado a interfaces desde el principio

El resultado es código cliente lleno de condicionales que verifican tipos o flags para decidir qué método llamar, perdiendo los beneficios del polimorfismo.

## Síntomas

- Dos o más clases hacen cosas similares pero con métodos con nombres diferentes
- Código cliente con condicionales basados en tipo de clase
- No puedes usar polimorfismo porque las interfaces no coinciden
- Duplicación de lógica en el código cliente que usa estas clases
- Dificultad para intercambiar implementaciones
- Necesidad de conocer la clase concreta en tiempo de compilación
- Tests duplicados para funcionalidad similar

## Ejemplo

```pseudocode
class TextLogger {
  function log(message: string) {
    print "[text] " + message
  }
}

class MessageWriter {
  function write(entry: string) {
    print "[text] " + entry
  }
}

// Código cliente forzado a usar condicionales
function useAltClasses(choice: string, msg: string) {
  if (choice == "logger") {
    logger = new TextLogger()
    logger.log(msg)  // Diferente nombre de método
  } else {
    writer = new MessageWriter()
    writer.write(msg)  // Diferente nombre de método
  }
}

// No podemos usar polimorfismo porque las interfaces difieren
```

## Ejercicio

Añade logging con marca de tiempo a ambas implementaciones y permite que el cliente pueda intercambiarlas en tiempo de ejecución sin usar condicionales.

## Problemas que encontrarás

Al no compartir una interfaz común, te verás obligado a duplicar lógica en métodos con nombres distintos y a esparcir sentencias `if/else` o `switch` en los clientes, haciendo que cambios simples se vuelvan tediosos y propensos a errores.

## Proceso de Refactoring

### 1. Identificar clases alternativas
- Busca clases que realizan funciones similares
- Identifica clases que podrían ser intercambiables
- Mapea qué métodos corresponden entre clases

### 2. Analizar diferencias de interfaz
- Compara los métodos de ambas clases
- Identifica métodos equivalentes con nombres diferentes
- Nota diferencias en firmas de métodos
- Identifica funcionalidad que existe en una pero no en la otra

### 3. Definir interfaz común
- Crea una interfaz o clase abstracta
- Usa nombres de métodos que tengan sentido para ambas implementaciones
- Define firmas de métodos consistentes
- Incluye todos los métodos necesarios para la funcionalidad común

### 4. Unificar nombres de métodos
- Usa **Rename Method** para hacer los nombres consistentes
- Elige el nombre más descriptivo y claro
- Actualiza todas las llamadas existentes

### 5. Hacer que las clases implementen la interfaz común
- Ambas clases deben implementar la interfaz definida
- Adapta las firmas de métodos si es necesario
- Asegúrate de que el comportamiento es equivalente

### 6. Refactorizar código cliente
- Reemplaza condicionales basados en tipo con polimorfismo
- Usa la interfaz común en lugar de clases concretas
- Elimina código duplicado en el cliente
- Aprovecha la intercambiabilidad para testing y configuración

## Técnicas de Refactoring Aplicables

- **Extract Interface**: Definir una interfaz común
- **Extract Superclass**: Si hay implementación común, extraer a superclase
- **Rename Method**: Unificar nombres de métodos
- **Add Parameter/Remove Parameter**: Hacer las firmas compatibles
- **Replace Conditional with Polymorphism**: Eliminar condicionales en cliente

## Beneficios

- **Polimorfismo**: Usar interfaces en lugar de condicionales
- **Intercambiabilidad**: Fácil cambiar entre implementaciones
- **Código cliente más simple**: Sin necesidad de conocer tipos concretos
- **Mejor testabilidad**: Fácil crear mocks e implementaciones de test
- **Extensibilidad**: Añadir nuevas implementaciones sin cambiar clientes
- **Consistencia**: Interfaz uniforme en todo el sistema
- **Open/Closed Principle**: Abierto a extensión, cerrado a modificación

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/oop-abusers/alternative-classes-different-interfaces.ts) - [README](../../typescript/src/code-smells/oop-abusers/alternative-classes-different-interfaces.readme.md)
- [Go](../../go/code_smells/oop_abusers/alternative_classes_different_interfaces.go) - [README](../../go/code_smells/oop_abusers/alternative_classes_different_interfaces.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/oopabusers/AlternativeClassesDifferentInterfaces.java) - [README](../../java/src/main/java/com/refactoring/codesmells/oopabusers/AlternativeClassesDifferentInterfaces.readme.md)
- [PHP](../../php/src/code-smells/oop-abusers/AlternativeClassesDifferentInterfaces.php) - [README](../../php/src/code-smells/oop-abusers/AlternativeClassesDifferentInterfaces.readme.md)
- [Python](../../python/src/code_smells/oop_abusers/alternative_classes_different_interfaces.py) - [README](../../python/src/code_smells/oop_abusers/alternative_classes_different_interfaces_readme.md)
- [C#](../../csharp/src/code-smells/oop-abusers/AlternativeClassesDifferentInterfaces.cs) - [README](../../csharp/src/code-smells/oop-abusers/alternative-classes-different-interfaces.readme.md)

## Referencias en Español

- [Los principios SOLID (Interface Segregation)](https://franiglesias.github.io/principios-solid/) - Explicación del principio de segregación de interfaces para crear contratos cohesivos

## Referencias

- [Refactoring Guru - Alternative Classes with Different Interfaces](https://refactoring.guru/smells/alternative-classes-with-different-interfaces)
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
- Gang of Four - "Design Patterns" - Strategy pattern
