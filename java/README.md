# Curso Refactoring Avanzado (Java)

Ejemplos y ejercicios del Curso de Refactoring Avanzado convertidos a Java.

## Preparación

### Requisitos

- Java 11 o superior
- Maven 3.6+ (o usar el wrapper incluido `mvnw`)

### Instalar dependencias

```bash
mvn clean install
```

### Ejecutar tests

```bash
# Ejecutar todos los tests
mvn test

# Ejecutar tests con output detallado
mvn test -X

# Ejecutar tests con cobertura
mvn test jacoco:report

# Ver reporte de cobertura (después de ejecutar el comando anterior)
open target/site/jacoco/index.html  # macOS
xdg-open target/site/jacoco/index.html  # Linux

# Ejecutar un test específico
mvn test -Dtest=ReceiptPrinterTest

# Ejecutar tests en modo watch (requiere herramienta externa)
# Instalar: https://github.com/davidB/watchexec
watchexec -e java,xml mvn test
```

## Contenido

### Técnicas de Refactoring

Ejercicios prácticos de técnicas avanzadas de refactoring:

#### Golden Master

Técnica para caracterizar el comportamiento de código legado sin tests.

- Ejercicio: [com.refactoring.refactoring](./src/main/java/com/refactoring/refactoring/)

```bash
mvn test -Dtest=ReceiptPrinterTest
```

#### Parallel Change

Técnicas para realizar cambios seguros en código en producción:

- Expand-Migrate-Contract
- Sprout Change
- Wrap Change

### Mantenimiento diario de código: Calistenia

Un conjunto de reglas para escribir código nuevo o evaluar código existente y modificarlo para acercarlo a un mejor diseño.

Ejercicios en [com.refactoring.calisthenics](./src/main/java/com/refactoring/calisthenics/):

1. Un nivel de indentación por método - [OneLevelIndentation.java](./src/main/java/com/refactoring/calisthenics/OneLevelIndentation.java)
2. No uses ELSE - [NoElse.java](./src/main/java/com/refactoring/calisthenics/NoElse.java)
3. Envuelve primitivos - [WrapPrimitives.java](./src/main/java/com/refactoring/calisthenics/WrapPrimitives.java)
4. Colecciones de primera clase
5. Un punto por línea
6. No uses abreviaciones
7. Mantén las entidades pequeñas
8. No más de 2 variables de instancia
9. Sin getters ni setters

### Code Smells

En estos ejercicios de Code Smells se presenta cada _smell_ con un ejemplo de código y se propone un ejercicio.

Cada ejercicio presenta una dificultad debida al _code smell_, que deberías abordar primero con un refactor para reducir el coste de cambio.

Sugerencias para realizar los ejercicios:

1. Introduce tests para caracterizar el comportamiento actual del código
2. Intenta resolver el ejercicio sin refactorizar primero
3. Realiza un refactor para reducir el coste del cambio
4. Completa el ejercicio tras el refactor

#### Bloaters

Code smells en los que se complica el cambio en el código por producir unidades demasiado grandes o por obligarnos a introducir mucho código auxiliar.

- [Data clump](src/main/java/com/refactoring/codesmells/bloaters/DataClump.java) - Grupos de datos que aparecen juntos repetidamente
- [Long parameter list](src/main/java/com/refactoring/codesmells/bloaters/LongParameterList.java) - Funciones con muchos parámetros
- [Primitive obsession](src/main/java/com/refactoring/codesmells/bloaters/PrimitiveObsession.java) - Uso excesivo de tipos primitivos

#### Couplers

Code smells en los que cambios en una unidad fuerzan cambios en otra que tiene un acoplamiento muy fuerte.

- [Feature envy](src/main/java/com/refactoring/codesmells/couplers/FeatureEnvy.java) - Métodos más interesados en otras clases

#### Dispensables

Code smells debidos a código innecesario, que introduce ruido dificultando la inteligibilidad del código.

- [Duplicated code](src/main/java/com/refactoring/codesmells/dispensables/DuplicatedCode.java) - Código duplicado en múltiples lugares

#### OOP Abusers

Code smells debido a la aplicación inadecuada de la orientación a objetos.

- [Switch statements](src/main/java/com/refactoring/codesmells/oopabusers/SwitchStatements.java) - Uso de switch en lugar de polimorfismo

## Estructura del Proyecto

```
java/
├── src/
│   ├── main/java/com/refactoring/
│   │   ├── refactoring/          # Técnicas de refactoring
│   │   │   ├── Order.java
│   │   │   ├── OrderItem.java
│   │   │   ├── OrderGenerator.java
│   │   │   └── ReceiptPrinter.java
│   │   ├── calisthenics/         # Ejercicios de Object Calisthenics
│   │   │   ├── OneLevelIndentation.java
│   │   │   ├── NoElse.java
│   │   │   └── WrapPrimitives.java
│   │   └── codesmells/           # Ejemplos de Code Smells
│   │       ├── bloaters/
│   │       ├── couplers/
│   │       ├── dispensables/
│   │       └── oopabusers/
│   └── test/java/com/refactoring/
│       └── refactoring/
│           └── ReceiptPrinterTest.java
├── pom.xml
└── README.md
```

## Comandos Útiles

```bash
# Compilar el proyecto
mvn compile

# Limpiar y compilar
mvn clean compile

# Ejecutar todos los tests
mvn test

# Ejecutar tests con cobertura
mvn test jacoco:report

# Ver reporte de cobertura
open target/site/jacoco/index.html

# Empaquetar el proyecto
mvn package

# Verificar el proyecto (compile, test, package)
mvn verify

# Formatear código (requiere plugin adicional)
mvn fmt:format

# Verificar estilo de código
mvn checkstyle:check

# Limpiar todo
mvn clean
```

## Versión de Java

Este proyecto requiere Java 11 o superior.

## Diferencias con otras versiones

- Se usan **clases** en lugar de estructuras más simples
- **PascalCase** para nombres de clases e interfaces
- **camelCase** para nombres de métodos y variables
- **Paquetes** (packages) para organizar el código
- **JUnit 5** con **AssertJ** para assertions más expresivas
- **JaCoCo** para análisis de cobertura de código
- **Maven** como herramienta de build (alternativa: Gradle)

## Convenciones de Java

- Los archivos de test están en `src/test/java/` con el mismo paquete que la clase bajo test
- Los nombres de clases de test terminan en `Test`
- Los métodos de test empiezan con `test` o usan la anotación `@Test`
- Se usa JUnit 5 (Jupiter) como framework de testing
- AssertJ proporciona assertions fluidas y legibles
- JaCoCo genera reportes de cobertura en `target/site/jacoco/`

## Dependencias del Proyecto

- **JUnit 5 (5.10.1)**: Framework de testing para Java
- **AssertJ (3.24.2)**: Librería para aserciones expresivas y fluidas
- **JaCoCo (0.8.11)**: Herramienta para análisis de cobertura de código
