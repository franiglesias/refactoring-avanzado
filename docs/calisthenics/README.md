# Object Calisthenics

Es un conjunto de nueve reglas de diseño que, aplicadas al código, mejoran su calidad, mantenibilidad y orientación a objetos.

## Origen

Las **Object Calisthenics** fueron propuestas por **Jeff Bay** en el libro **"The ThoughtWorks Anthology"** (2008). El nombre hace referencia a los ejercicios gimnásticos de calistenia, sugiriendo que son prácticas que "entrenan" y fortalecen las habilidades de diseño de código orientado a objetos.

[La calistenia busca la adquisición de resistencia, agilidad, equilibrio, coordinación y flexibilidad a través de un entrenamiento equilibrado con el propio peso corporal.](https://es.wikipedia.org/wiki/Calistenia). De un modo análogo, la _Objects Calisthenics_ busca incrementar la calidad del código mediante prácticas basadas en atender a las propiedades del código mismo.

Estas reglas son restricciones deliberadas que nos fuerzan a pensar de manera diferente sobre el diseño de nuestro código. Al igual que los ejercicios físicos desarrollan fuerza y flexibilidad, estas prácticas desarrollan código más robusto, flexible y expresivo.

## Filosofía

Object Calisthenics no son reglas absolutas que deban aplicarse ciegamente en todo momento. Son **herramientas de entrenamiento** y **guías de diseño** que:

- Revelan problemas de diseño ocultos, que son en sí mismos code-smells.
- Fuerzan a crear abstracciones más limpias, atacando síntomas propios de malas abstracciones.
- Mejoran la expresividad del código, ayudándonos a ser más claros y expresivos.
- Reducen el acoplamiento, favoreciendo la encapsulación.
- Aumentan la cohesión, contribuyendo a agrupar el código que trabaja junto.
- Facilitan el testing.

La clave está en **entender el propósito** de cada regla y aplicarlas con criterio según el contexto.

## Las 9 Reglas

| # | Regla                                | Nivel        | Documento |
|---|--------------------------------------|--------------|-----------|
| 1 | No usar abreviaturas                 | Principiante | [Ver documento](01-dont-use-abbreviations.md) |
| 2 | No usar ELSE                         | Principiante | [Ver documento](02-dont-use-else.md) |
| 3 | Un solo nivel de indentación         | Intermedio   | [Ver documento](03-one-indentation-level.md) |
| 4 | Empaquetar primitivos                | Intermedio   | [Ver documento](04-wrap-primitives.md) |
| 5 | Colecciones de primera clase         | Intermedio   | [Ver documento](05-first-class-collections.md) |
| 6 | No usar getters y setters            | Avanzado     | [Ver documento](06-no-getters-setters.md) |
| 7 | Mantener las unidades pequeñas       | Intermedio   | [Ver documento](07-small-entities.md) |
| 8 | Máximo de dos variables de instancia | Avanzado     | [Ver documento](08-max-two-instance-variables.md) |
| 9 | No más de un punto por línea         | Avanzado     | [Ver documento](09-one-dot-per-line.md) |

## Cómo Usar Estas Reglas

### 1. Como Herramienta de Aprendizaje

Cuando estés aprendiendo diseño orientado a objetos, aplica estas reglas estrictamente en proyectos de práctica o ejercicios de código. Te forzarán a pensar de manera diferente y descubrir patrones que de otro modo no verías.

### 2. Como Detector de Code Smells

Si tu código viola varias de estas reglas, probablemente tenga problemas de diseño subyacentes. Usa las reglas como checklist para identificar áreas que necesitan refactoring.

### 3. Como Guía de Revisión de Código

Durante code reviews, estas reglas proporcionan criterios objetivos para discutir la calidad del diseño más allá de preferencias personales.

### 4. Con Pragmatismo

En código de producción, algunas violaciones pueden estar justificadas por razones de performance, integraciones con frameworks, o simplicidad. Lo importante es que sean **decisiones conscientes**, no accidentes.

## Nivel de Aplicación Recomendado

### Principiante (Reglas 1-2)

Estas reglas mejoran inmediatamente la legibilidad sin requerir conocimientos profundos de diseño:

- **No usar abreviaturas**: Nombres claros y expresivos
- **No usar ELSE**: Retornos tempranos e introducción de _guard clauses_

**Objetivo**: Código que se lee como prosa, flujo más claro.

### Intermedio (Reglas 3-5, 7)

Requieren entender abstracciones y responsabilidades:

- **Un solo nivel de indentación**: Extracción de métodos, composición.
- **Empaquetar primitivos**: Value Objects, protección de las invariantes de dominio.
- **Colecciones de primera clase**: Encapsulación de comportamiento de colecciones para proteger sus propias invariantes y no acoplar el código a una estructura de datos.
- **Mantener las unidades pequeñas**: Favorecer _Single Responsibility Principle_.

**Objetivo**: Clases y métodos cohesivos con responsabilidades claras.

### Avanzado (Reglas 6, 8-9)

Requieren comprensión profunda de encapsulación y _Tell Don't Ask_:

- **No usar getters y setters**: Priorizar el comportamiento sobre los datos, Tell Don't Ask: pedir a los objetos que hagan cosas.
- **Máximo de dos variables de instancia**: Cohesión extrema, composición.
- **No más de un punto por línea**: Ley de Demeter, ocultar delegaciones y desacoplar objetos sin usar nuestro conocimiento de su estructura.

**Objetivo**: Diseño orientado a objetos puro, alta cohesión y bajo acoplamiento.

## Combinación con Otras Prácticas

Object Calisthenics se complementan perfectamente con:

- **SOLID Principles**: Las reglas ayudan a cumplir SRP, OCP y DIP naturalmente.
- **Clean Code**: Amplían los principios de código limpio con restricciones específicas.
- **TDD**: El código que sigue estas reglas es inherentemente más testeable.
- **Domain-Driven Design**: Fomentan la creación de modelos de dominio ricos.

## Ejercicios Prácticos

Este repositorio incluye ejercicios implementados en **estos lenguajes**:

- [TypeScript](../../typescript/src/calisthenics-exercises/)
- [Go](../../go/calisthenics_exercises/)
- [Java](../../java/src/main/java/com/refactoring/calisthenics/)
- [PHP](../../php/src/calisthenics-exercises/)
- [Python](../../python/src/calisthenics_exercises/)
- [C#](../../csharp/src/calisthenics-exercises/)

Cada ejercicio presenta código que viola una regla específica y te desafía a refactorizarlo para cumplirla.

## Beneficios Generales

Al aplicar Object Calisthenics de manera consistente:

1. **Código más legible**: Se lee como prosa, no necesitas descifrarlo o intentar ejecutarlo mentalmente. Es fácil escribir código que funcione, lo difícil es escribir código que sea legible.
2. **Mejor diseño**: Emergen abstracciones naturales y responsabilidades claras.
3. **Testing más fácil**: Clases pequeñas y cohesivas son simples de testear, incluso triviales.
4. **Menos bugs**: Menos complejidad = menos lugares para esconder errores.
5. **Mantenibilidad**: Cambios localizados, bajo acoplamiento, nos ayuda a controlar la complejidad.
6. **Onboarding más rápido**: Código autoexplicativo reduce curva de aprendizaje.

## Referencias en Español

### Artículos Generales sobre Object Calisthenics

- [Calistenias para objetos de valor](https://franiglesias.github.io/calistenics-and-value-objects/) - Aplicación de calisthenics a Value Objects
- [Object Calisthenics para adelgazar las clases](https://franiglesias.github.io/calistenics-and-small-classes/) - Técnicas para mantener clases pequeñas y cohesivas
- [Calistenia Object Calisthenics, 10](https://franiglesias.github.io/calisthenics-10/) - Análisis de las reglas y su aplicación práctica

### Artículos Relacionados por Concepto

- [Cómo poner nombres](https://franiglesias.github.io/naming-things/) - Sobre nombres expresivos y claros
- [Métodos largos](https://franiglesias.github.io/long-method/) - Problema de métodos largos y refactoring
- [Primitive obsession](https://franiglesias.github.io/primitive-obsession/) - Sobre encapsular primitivos
- [Refactor cotidiano (3): extraer para aclarar, parte 1](https://franiglesias.github.io/everyday-refactor-3/) - Sobre eliminación de ELSE
- [Refactor cotidiano (6): cuéntame, no me preguntes (tell, don't ask)](https://franiglesias.github.io/everyday-refactor-6/) - Principio Tell Don't Ask aplicado

## Referencias en Inglés

### Libros

- **"The ThoughtWorks Anthology"** - Jeff Bay - Capítulo original sobre Object Calisthenics
- **"Refactoring: Improving the Design of Existing Code"** - Martin Fowler - Técnicas de refactoring aplicables
- **"Clean Code: A Handbook of Agile Software Craftsmanship"** - Robert C. Martin - Principios complementarios

### Artículos y Recursos Online

- [Object Calisthenics - William Durand](https://williamdurand.fr/2013/06/03/object-calisthenics/) - Análisis detallado de las 9 reglas
- [Your Code as a Crime Scene - Adam Tornhill](https://www.adamtornhill.com/) - Análisis de calidad de código
- [Law of Demeter](https://en.wikipedia.org/wiki/Law_of_Demeter) - Principio relacionado con "un punto por línea"

### Videos y Conferencias

- [Object Calisthenics - Rafael Dohms](https://www.youtube.com/results?search_query=object+calisthenics+rafael+dohms) - Charlas sobre aplicación práctica
- [Clean Code Talks - Google](https://www.youtube.com/results?search_query=clean+code+talks+google) - Principios de diseño relacionados

## Conclusión

Object Calisthenics son **ejercicios de diseño** que, como todo ejercicio, requieren práctica deliberada. No se trata de seguir reglas dogmáticamente, sino de **entrenar tu instinto de diseño** para crear código más limpio y mantenible.

Comienza con las reglas más simples, practica con los ejercicios incluidos en este repositorio, y gradualmente incorpora las más avanzadas. Con el tiempo, muchos de estos principios se volverán naturales en tu forma de escribir código.

**Recuerda**: El objetivo no es perfección absoluta, sino mejora continua del diseño de tu código.
