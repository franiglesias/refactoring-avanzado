# Curso Refactoring Avanzado

Ejemplos y ejercicios del Curso de Refactoring Avanzado.

## 📚 Documentación

- **[Documentación Genérica](./docs/README.md)**: Explicaciones detalladas independientes del lenguaje con procesos de refactoring paso a paso y ejemplos en pseudocódigo.
- **[Object Calisthenics](./docs/calisthenics/README.md)**: Guía completa de las 9 reglas de Object Calisthenics con ejemplos y ejercicios.
- **[Índice de Code Smells](./docs/code-smells/README.md)**: Catálogo de 21 code smells organizados por categoría.
- **[Técnicas de Refactoring Avanzado](./docs/refactoring/README.md)**: Técnicas para manejar el refactoring de código existente.

## Versiones disponibles

Este curso está disponible en los siguientes lenguajes:

- **[TypeScript](./typescript/README.md)**: Versión con Node.js, TypeScript y Vitest
- **[Python](./python/README.md)**: Versión con Python 3.11+ y pytest
- **[Java](./java/README.md)**: Versión con Java 11+, Maven, JUnit 5 y AssertJ
- **[Go](./go/README.md)**: Versión con Go 1.21+ y testing estándar
- **[C#](./csharp/README.md)**: Versión con .NET 8, C# 12 y xUnit
- **[PHP](./php/README.md)**: Versión con PHP 8.2+, Composer, PHPUnit y Docker

## Preparación

Para trabajar con el material del curso, te recomendamos no clonarlo, sino descargar la versión `.zip`. De este modo, evitarás tener gestión de versiones global de todo el proyecto y podrás limitarlo a cada versión del lenguaje que quieras practicar.

Para preparar el entorno de trabajo, sigue las instrucciones específicas de cada versión según el lenguaje que prefieras usar. En general, te recomendamos ponerlo bajo un gestor de versiones tanto para hacer pruebas, como para trabajar de forma sistemática en los ejercicios avanzados.

## Contenido

Este curso incluye ejemplos y ejercicios organizados en las siguientes categorías:

### Mantenimiento diario de código: Calistenia

Un conjunto de 9 reglas para escribir código nuevo o evaluar código existente y modificarlo para acercarlo a un mejor diseño:

1. Un nivel de indentación por método
2. No uses ELSE
3. Envuelve primitivos
4. Colecciones de primera clase
5. Un punto por línea
6. No uses abreviaciones
7. Mantén las entidades pequeñas
8. No más de 2 variables de instancia
9. Sin getters ni setters

Aplicar las reglas de calistenia nos ayuda a mantener bajo control la entropía, es decir, el desorden que se genera de forma orgánica al escribir código.

### Code Smells

En estos ejercicios se presenta cada _smell_ con un ejemplo de código y se propone un problema de código para solucionar.

Cada propuesta presenta una dificultad debida al _code smell_, que deberías abordar antes de nada con un refactor para reducir el coste de cambio. El propósito de estos ejercicios es ilustrar cómo el _code smell_, incrementa el coste de intervención en el código.

**Sugerencias para realizar los ejercicios:**

1. Introduce tests para caracterizar el comportamiento actual del código
2. Intenta resolver el ejercicio sin refactorizar primero
3. Realiza un refactor para reducir el coste del cambio
4. Completa el ejercicio tras el refactor

Usa el gestor de versiones para consolidar cambios a medida que progresas en la refactorización.

#### Bloaters

Son estructuras de código que complican el cambio por producir unidades demasiado grandes o por obligarnos a introducir mucho código auxiliar:

- **Data clump**: Grupos de datos que aparecen juntos repetidamente
- **Large class**: Clases que hacen demasiadas cosas
- **Long method**: Métodos excesivamente largos (ejercicio final recomendado)
- **Long parameter list**: Funciones con muchos parámetros
- **Primitive obsession**: Uso excesivo de tipos primitivos

#### Change Preventers

Se trata de estructuras, o la falta de ellas, que hacen que cualquier cambio sea costoso e incluso arriesgado al obligarnos a intervenir en muchos lugares del código a la vez debido a la duplicación literal o funcional de código:

- **Divergent change**: Una clase cambia frecuentemente por diferentes razones
- **Parallel inheritance hierarchy**: Jerarquías de herencia paralelas
- **Shotgun surgery**: Un cambio requiere modificar muchas clases diferentes

#### Couplers

Se trata de partes de código en los cuales los cambios en una unidad fuerzan cambios en otra, debido a que tienen un acoplamiento muy fuerte:

- **Feature envy**: Métodos más interesados en otras clases
- **Inappropriate intimacy**: Clases que conocen demasiado sobre la implementación de otras
- **Message chains**: Cadenas largas de llamadas (Ley de Demeter)
- **Middleman**: Clases que solo delegan a otras

#### Dispensables

En ocasiones, tenemos código innecesario, que introduce ruido dificultando la inteligibilidad del código:

- **Comments**: Comentarios que explican código mal escrito
- **Data class**: Clases que solo contienen datos sin comportamiento
- **Dead code**: Código que nunca se ejecuta
- **Duplicated code**: Código duplicado en múltiples lugares
- **Lazy class**: Clases que no hacen lo suficiente

#### OOP Abusers

La orientación a objetos mal aplicada puede dar lugar a varios Code smells:

- **Alternative classes with different interfaces**: Clases similares con interfaces diferentes
- **Refused bequest**: Cuando las clases no usan la herencia recibida
- **Switch statements**: Uso de switch/if-elseif en lugar de polimorfismo
- **Temporal instance variables**: Variables de instancia que solo se usan temporalmente

### Técnicas de Refactoring

Ejercicios prácticos de técnicas avanzadas de refactoring:

- **Golden Master**: Técnica para caracterizar el comportamiento de código legado sin tests
- **Parallel Change**: Técnicas para realizar cambios seguros en código en producción
  - Expand-Migrate-Contract
  - Sprout Change
  - Wrap Change

## Ejercicio final

La [QuoteBot kata](https://github.com/franiglesias/quotebot) es un ejercicio en el que aplicar todo lo aprendido en este curso. Se trata de una aplicación que ni siquiera se puede ejecutar en local.

## Estructura del Repositorio

Cada versión del curso está en su propia carpeta con su documentación específica. Consulta el README de cada versión para instrucciones detalladas de setup y ejecución.
