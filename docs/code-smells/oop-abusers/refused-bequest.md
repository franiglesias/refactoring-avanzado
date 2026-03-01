# Refused Bequest

Herencia rechazada.

## Definición

Este smell aparece cuando una subclase hereda de una clase base, pero ignora o sobrescribe con excepciones/operaciones vacías gran parte de lo que hereda. Esto sugiere que la relación de herencia no es adecuada ("is-a") o que la jerarquía de clases necesita ser replanteada, posiblemente usando composición o extrayendo una interfaz más pequeña.

## Descripción

**Refused Bequest** (herencia rechazada) ocurre cuando una subclase hereda comportamiento que no quiere o no puede usar apropiadamente. La subclase "rechaza" la herencia proporcionando implementaciones vacías, lanzando excepciones, o simplemente ignorando métodos heredados.

Este smell indica que la relación "is-a" no es apropiada. La subclase no es realmente un tipo especial de la superclase si rechaza parte significativa de su comportamiento o interfaz.

Situaciones comunes:
- Una subclase implementa métodos como no-ops o lanzando excepciones
- La subclase solo usa una pequeña parte de la funcionalidad heredada
- La herencia se usó para reutilización de código, no por relación conceptual
- La jerarquía fue diseñada incorrectamente desde el principio

## Síntomas

- Métodos heredados que lanzan `NotImplementedException` o similares
- Métodos sobrescritos con implementaciones vacías
- La subclase ignora la mayoría de los métodos de la superclase
- Comentarios como "este método no se usa en esta subclase"
- Violaciones del Principio de Sustitución de Liskov
- La subclase solo usa una fracción de la interfaz heredada
- La herencia se usó solo para compartir código, no por relación conceptual

## Ejemplo

```pseudocode
interface Resettable {
  function reset()
}

interface Controller {
  function start()
  function stop()
}

class BaseController implements Controller, Resettable {
  function start() {
    print "starting"
  }

  function stop() {
    print "stopping"
  }

  function reset() {
    print "resetting"
  }
}

class ReadOnlyController implements Controller {
  // Rechaza la herencia - no puede implementar apropiadamente
  function start() {
    // Implementación vacía - rechaza el comportamiento
  }

  function stop() {
    // Implementación vacía - rechaza el comportamiento
  }
}

function demoRefusedBequest(readonly: boolean) {
  controller = readonly ? new ReadOnlyController() : new BaseController()
  controller.start()  // ReadOnly no hace nada
  controller.stop()   // ReadOnly no hace nada
}
```

## Ejercicio

Añade un método de ciclo de vida `pause` a la interfaz `Controller` y haz que `start` y `stop` sean obligatorios con lógica real.

## Problemas que encontrarás

`ReadOnlyController` se verá forzado a implementar métodos que no tienen sentido para su propósito, lo que te obligará a lanzar excepciones o dejar implementaciones vacías que violan el Principio de Sustitución de Liskov.

## Proceso de Refactoring

### 1. Identificar herencia rechazada
- Busca métodos sobrescritos con implementaciones vacías
- Encuentra métodos que lanzan excepciones de "no implementado"
- Identifica subclases que usan poco de la superclase

### 2. Evaluar la relación "is-a"
- ¿La subclase realmente ES un tipo de la superclase?
- ¿Puede sustituir a la superclase en cualquier contexto?
- ¿O solo se heredó para reutilizar código?

### 3. Reemplazar herencia con composición
- Si la relación no es "is-a" genuina
- Usa composición: la subclase contiene una instancia de la clase
- Delega solo a los métodos que realmente necesita
- Ejemplo: `ReadOnlyController` contiene un `ViewerCapability`

### 4. Extraer interfaz más pequeña
- Si múltiples subclases rechazan los mismos métodos
- Crea una interfaz más pequeña con solo lo esencial
- Las subclases implementan solo lo que necesitan
- Usa **Extract Interface** o **Extract Subclass**

### 5. Mover comportamiento común
- Si el comportamiento compartido es mínimo
- Extráelo a una clase utilidad o helper
- Elimina la herencia innecesaria

### 6. Aplicar Interface Segregation Principle
- Las interfaces deben ser pequeñas y cohesivas
- Los clientes no deben depender de métodos que no usan
- Divide interfaces grandes en interfaces más específicas

## Técnicas de Refactoring Aplicables

- **Replace Inheritance with Delegation**: Usar composición en lugar de herencia
- **Extract Subclass**: Si solo algunas subclases necesitan ciertos métodos
- **Extract Superclass**: Si la superclase tiene demasiada funcionalidad
- **Extract Interface**: Definir contratos más pequeños y específicos
- **Push Down Method**: Mover métodos solo a las subclases que los necesitan
- **Push Down Field**: Mover campos solo a donde se usan

## Beneficios

- **Mejor diseño**: Relaciones de herencia apropiadas
- **LSP cumplido**: Las subclases son sustituibles por la superclase
- **Interfaces claras**: Contratos que tienen sentido
- **Menos código muerto**: No hay métodos vacíos o excepciones
- **Composición flexible**: Más flexible que herencia rígida
- **Mejor testabilidad**: No necesitas lidiar con métodos no implementados
- **Comprensión más fácil**: La jerarquía refleja la realidad

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/oop-abusers/refused-bequest.ts) - [README](../../typescript/src/code-smells/oop-abusers/refused-bequest.readme.md)
- [Go](../../go/code_smells/oop_abusers/refused_bequest.go) - [README](../../go/code_smells/oop_abusers/refused_bequest.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/oopabusers/RefusedBequest.java) - [README](../../java/src/main/java/com/refactoring/codesmells/oopabusers/RefusedBequest.readme.md)
- [PHP](../../php/src/code-smells/oop-abusers/RefusedBequest.php) - [README](../../php/src/code-smells/oop-abusers/RefusedBequest.readme.md)
- [Python](../../python/src/code_smells/oop_abusers/refused_bequest.py) - [README](../../python/src/code_smells/oop_abusers/refused_bequest_readme.md)
- [C#](../../csharp/src/code-smells/oop-abusers/RefusedBequest.cs) - [README](../../csharp/src/code-smells/oop-abusers/refused-bequest.readme.md)

## Referencias en Español

- [Los principios SOLID (Liskov Substitution)](https://franiglesias.github.io/principios-solid/) - Explicación del principio de sustitución de Liskov y herencia correcta
- [Más allá de SOLID, los cimientos](https://franiglesias.github.io/beyond-solid-2/) - Fundamentos de diseño que incluyen composición sobre herencia

## Referencias

- [Refactoring Guru - Refused Bequest](https://refactoring.guru/smells/refused-bequest)
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
- Robert C. Martin - "Agile Software Development" - Liskov Substitution Principle
- Gang of Four - "Design Patterns" - Favor composition over inheritance
