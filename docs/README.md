# Documentación de Refactoring Avanzado

Esta documentación contiene explicaciones genéricas y detalladas de todos los ejercicios del curso de Refactoring Avanzado, independientes del lenguaje de programación.

## Estructura de la Documentación

Cada ejercicio incluye:
- **Definición y Descripción**: Qué es el code smell o regla y por qué es importante
- **Síntomas**: Cómo identificarlo en tu código
- **Ejemplo en pseudocódigo**: Demostración genérica del problema
- **Ejercicio práctico**: Tarea específica para experimentar el problema
- **Proceso de Refactoring/Aplicación**: Pasos detallados para aplicar la técnica
- **Técnicas aplicables**: Técnicas de refactoring que puedes usar
- **Beneficios**: Ventajas de aplicar la práctica
- **Versiones por lenguaje**: Enlaces a implementaciones en TypeScript, Go, Java, PHP, Python y C#
- **Referencias en español**: Artículos del blog de Fran Iglesias
- **Referencias**: Recursos en inglés (libros, artículos)

## Contenido del Curso

### Code Smells

Los code smells son indicadores de problemas más profundos en el diseño del código. No son bugs, pero sugieren debilidades que pueden dificultar el mantenimiento y la evolución del software.

**[Ver catálogo completo de Code Smells](code-smells/README.md)** - 21 code smells organizados en 5 categorías

#### Bloaters

Code smells en los que se complica el cambio en el código por producir unidades demasiado grandes o por obligarnos a introducir mucho código auxiliar.

- [Data Clump](code-smells/bloaters/data-clump.md) - Grupo de datos que siempre viajan juntos
- [Large Class](code-smells/bloaters/large-class.md) - Clase con demasiadas responsabilidades
- [Long Method](code-smells/bloaters/long-method.md) - Método excesivamente largo
- [Long Parameter List](code-smells/bloaters/long-parameter-list.md) - Método con muchos parámetros
- [Primitive Obsession](code-smells/bloaters/primitive-obsession.md) - Uso excesivo de tipos primitivos

#### Change Preventers

Code smells que hacen que cualquier cambio sea costoso e incluso arriesgado al obligarnos a intervenir en muchos lugares del código a la vez.

- [Divergent Change](code-smells/change-preventers/divergent-change.md) - Clase que cambia por múltiples razones
- [Parallel Inheritance Hierarchy](code-smells/change-preventers/parallel-inheritance-hierarchy.md) - Jerarquías que crecen en paralelo
- [Shotgun Surgery](code-smells/change-preventers/shotgun-surgery.md) - Cambio que requiere modificar múltiples clases

#### Couplers

Code smells en los que cambios en una unidad fuerzan cambios en otra que tiene un acoplamiento muy fuerte.

- [Feature Envy](code-smells/couplers/feature-envy.md) - Método que usa más datos de otra clase
- [Inappropriate Intimacy](code-smells/couplers/inappropriate-intimacy.md) - Clases demasiado acopladas
- [Message Chains](code-smells/couplers/message-chains.md) - Cadenas largas de llamadas
- [Middleman](code-smells/couplers/middleman.md) - Clase que solo delega

#### Dispensables

Code smells debidos a código innecesario, que introduce ruido dificultando la inteligibilidad del código.

- [Comments](code-smells/dispensables/comments.md) - Comentarios que compensan código poco claro
- [Data Class](code-smells/dispensables/data-class.md) - Clase sin comportamiento
- [Dead Code](code-smells/dispensables/dead-code.md) - Código que nunca se ejecuta
- [Duplicated Code](code-smells/dispensables/duplicated-code.md) - Código repetido
- [Lazy Class](code-smells/dispensables/lazy-class.md) - Clase que hace muy poco

#### OOP Abusers

Code smells debido a la aplicación inadecuada de la orientación a objetos.

- [Alternative Classes with Different Interfaces](code-smells/oop-abusers/alternative-classes-different-interfaces.md) - Clases similares con interfaces diferentes
- [Refused Bequest](code-smells/oop-abusers/refused-bequest.md) - Subclase que no usa la herencia recibida
- [Switch Statements](code-smells/oop-abusers/switch-statements.md) - Switch sobre tipos en lugar de polimorfismo
- [Temporal Instance Variables](code-smells/oop-abusers/temporal-instance-variables.md) - Variables de instancia válidas solo en ciertas fases

### Object Calisthenics

Object Calisthenics es un conjunto de 9 reglas propuestas por Jeff Bay que, aplicadas a nuestro código, pueden mejorar su calidad y aproximarlo a un mejor diseño mediante prácticas disciplinadas.

**[Ver documentación completa de Object Calisthenics](calisthenics/README.md)**

Las 9 reglas son:

1. [No usar abreviaturas](calisthenics/01-dont-use-abbreviations.md) - Nombres expresivos y completos
2. [No usar ELSE](calisthenics/02-dont-use-else.md) - Simplificar flujo de control
3. [Un solo nivel de indentación](calisthenics/03-one-indentation-level.md) - Métodos enfocados
4. [Empaquetar primitivos](calisthenics/04-wrap-primitives.md) - Value Objects para conceptos del dominio
5. [Colecciones de primera clase](calisthenics/05-first-class-collections.md) - Encapsular comportamiento de colecciones
6. [No usar getters y setters](calisthenics/06-no-getters-setters.md) - Tell, Don't Ask
7. [Mantener las unidades pequeñas](calisthenics/07-small-entities.md) - Clases y métodos pequeños
8. [Máximo dos variables de instancia](calisthenics/08-max-two-instance-variables.md) - Cohesión alta
9. [No más de un punto por línea](calisthenics/09-one-dot-per-line.md) - Ley de Demeter

### Técnicas de Refactoring Avanzado

Técnicas especializadas para trabajar con código legacy sin tests y realizar cambios seguros en producción.

**[Ver documentación completa de Técnicas de Refactoring](refactoring/README.md)**

#### Golden Master
- [Golden Master](refactoring/golden-master/README.md) - Tests de caracterización para código legacy

#### Parallel Change
- [Introducción](refactoring/parallel-change/README.md) - Visión general de cambios en paralelo
- [Sprout Change](refactoring/parallel-change/sprout-change.md) - Hacer brotar nuevo código
- [Wrap Change](refactoring/parallel-change/wrap-change.md) - Envolver dependencias
- [Expand-Migrate-Contract](refactoring/parallel-change/expand-migrate-contract.md) - Cambios estructurales en 3 fases

## Implementaciones por Lenguaje

Cada ejercicio está implementado en 6 lenguajes de programación:

- **[TypeScript](../typescript/)** - JavaScript con tipos estáticos
- **[Go](../go/)** - Lenguaje compilado con concurrencia nativa
- **[Java](../java/)** - Lenguaje orientado a objetos clásico
- **[PHP](../php/)** - Lenguaje interpretado para web
- **[Python](../python/)** - Lenguaje interpretado multiparadigma
- **[C#](../csharp/)** - Lenguaje orientado a objetos de Microsoft

## Cómo Usar Este Material

### Para Aprender

1. Lee la documentación genérica del code smell o regla
2. Estudia el ejemplo en pseudocódigo
3. Elige tu lenguaje preferido y estudia la implementación
4. Intenta hacer el ejercicio propuesto
5. Experimenta el problema al modificar el código
6. Aplica el proceso de refactoring sugerido

### Para Enseñar

1. Presenta el concepto con la documentación genérica
2. Muestra ejemplos en el lenguaje que usen tus estudiantes
3. Pide que completen el ejercicio sin refactorizar primero
4. Discute los problemas encontrados
5. Guía el proceso de refactoring paso a paso
6. Compara el código antes y después

### Para Referencia

- Usa la documentación para identificar smells en tu código
- Consulta el proceso de refactoring cuando encuentres un smell
- Revisa las implementaciones en diferentes lenguajes para entender patrones comunes

## Técnicas de Refactoring

Las técnicas mencionadas en la documentación provienen del catálogo de Martin Fowler:

- **Extract Method/Function**: Extraer código a un método nuevo
- **Extract Class**: Crear una clase nueva con responsabilidad específica
- **Move Method/Field**: Mover comportamiento a la clase correcta
- **Introduce Parameter Object**: Agrupar parámetros en un objeto
- **Replace Conditional with Polymorphism**: Usar herencia en lugar de condicionales
- **Replace Type Code with Strategy/State**: Usar patrones de diseño
- **Hide Delegate**: Ocultar dependencias internas
- **Remove Middle Man**: Eliminar delegación innecesaria
- **Inline Method/Class**: Eliminar abstracciones innecesarias

## Principios de Diseño

Los refactorings buscan mejorar el código según estos principios:

- **Single Responsibility Principle (SRP)**: Una clase debe tener una única razón para cambiar
- **Open/Closed Principle (OCP)**: Abierto para extensión, cerrado para modificación
- **Liskov Substitution Principle (LSP)**: Los subtipos deben ser sustituibles por sus tipos base
- **Interface Segregation Principle (ISP)**: Interfaces específicas mejor que una general
- **Dependency Inversion Principle (DIP)**: Depender de abstracciones, no de concreciones
- **Don't Repeat Yourself (DRY)**: Evitar duplicación
- **Tell, Don't Ask**: Los objetos deben hacer cosas, no exponer su estado
- **Law of Demeter**: Hablar solo con amigos cercanos

## Recursos Adicionales

### Recursos en Español

#### Blog de Fran Iglesias
El blog [franiglesias.github.io](https://franiglesias.github.io) contiene más de 40 artículos en español sobre refactoring, testing y principios de diseño:

**Series completas:**
- [Refactor Cotidiano (8 artículos)](https://franiglesias.github.io/everyday-refactor-1/) - Técnicas de refactoring diario
- [Ejercicio de Refactoring (4 artículos)](https://franiglesias.github.io/ejercicio-de-refactor-1/) - Ejercicio práctico paso a paso
- [Outside-in y BDD (6+ artículos)](https://franiglesias.github.io/outside-in-with-behat-phpspec/) - Behavior Driven Development
- [Más allá de SOLID (4 artículos)](https://franiglesias.github.io/beyond-solid/) - Principios de diseño avanzados

**Artículos destacados:**
- [Primitive Obsession](https://franiglesias.github.io/primitive-obsession/) - Code smell y refactoring
- [Métodos largos](https://franiglesias.github.io/long-method/) - Análisis y soluciones
- [Cobertura de test rápida con Golden Master](https://franiglesias.github.io/approval_testing/) - Técnica para código legacy
- [Object Calisthenics](https://franiglesias.github.io/calistenics-and-value-objects/) - Reglas de diseño
- [Los principios SOLID](https://franiglesias.github.io/principios-solid/) - Explicación completa
- [Introducción a DDD](https://franiglesias.github.io/ddd-intro/) - Domain Driven Design

**Libros:**
- [Aprende Test Driven Development](https://franiglesias.github.io/new-book/) - Libro completo en español sobre TDD

**Nota:** Cada ejercicio en esta documentación incluye referencias específicas a artículos relevantes del blog.

### Recursos en Inglés

- [Refactoring Guru](https://refactoring.guru/) - Catálogo visual de refactorings y patrones
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
- Robert C. Martin - "Clean Code: A Handbook of Agile Software Craftsmanship"
- Michael Feathers - "Working Effectively with Legacy Code"
- [Refactoring.com](https://refactoring.com/catalog/) - Catálogo de técnicas de refactoring
- Jeff Bay - "Object Calisthenics" (The ThoughtWorks Anthology)

## Contribuir

Si encuentras errores o mejoras en la documentación:

1. Revisa la documentación genérica en `docs/`
2. Verifica las implementaciones específicas de cada lenguaje
3. Abre un issue describiendo el problema
4. O propón cambios mediante pull request

---

**Licencia**: Material educativo del Curso de Refactoring Avanzado

**Última actualización**: 2025
