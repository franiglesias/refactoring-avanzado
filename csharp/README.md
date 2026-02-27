# Curso Refactoring Avanzado - C#

Ejemplos y ejercicios del Curso de Refactoring Avanzado en C#.

## Preparación

### Requisitos previos

- [.NET 8.0 SDK](https://dotnet.microsoft.com/download/dotnet/8.0) o superior
- Un IDE compatible con C# (Visual Studio, Visual Studio Code con C# Dev Kit, o JetBrains Rider)

### Instalación

1. Verificar que tienes .NET instalado:
```bash
dotnet --version
```

2. Restaurar las dependencias del proyecto:
```bash
cd csharp
dotnet restore
```

3. Compilar el proyecto:
```bash
dotnet build
```

4. Ejecutar los tests:
```bash
dotnet test
```

## Contenido

### Mantenimiento diario de código: Calistenia

Un conjunto de reglas para escribir código nuevo o evaluar código existente y modificarlo para acercarlo a un mejor diseño.

[Ejercicios de calistenia](./src/calisthenics-exercises).

### Code Smells

En estos ejercicios de Code Smells se presenta cada _smell_ con un ejemplo de código y se propone un ejercicio.

Cada ejercicio presenta una dificultad debida al _code smell_, que deberías abordar primero con un refactor para reducir el coste de cambio.

Sugerencias para realizar los ejercicios:

1. Introduce tests para caracterizar el comportamiento actual del código
2. Intenta resolver el ejercicio sin refactorizar primero.
3. Realiza un refactor para reducir el coste del cambio.
4. Completa el ejercicio tras el refactor.

#### Bloaters

Code smells en los que se complica el cambio en el código por producir unidades demasiado grandes o por obligarnos a introducir mucho código auxiliar.

- [Data clump](src/code-smells/bloaters/data-clump.readme.md)
- [Large class](src/code-smells/bloaters/large-class.readme.md)
- [Long method](src/code-smells/bloaters/long-method.readme.md): Este es un ejemplo especialmente grande, y que en realidad incluye muchos otros code smells, por lo que se recomienda realizarlo como "gran ejercicio final"
- [Long parameter list](src/code-smells/bloaters/long-parameter-list.readme.md)
- [Primitive obsession](src/code-smells/bloaters/primitive-obsession.readme.md)

#### Change Preventers

Code smells que hacen que cualquier cambio sea costoso e incluso arriesgado al obligarnos a intervenir en muchos lugares del código a la vez.

- [Divergent change](src/code-smells/change-preventers/divergent-change.readme.md)
- [Parallel inheritance-hierarchy](src/code-smells/change-preventers/parallel-inheritance-hierarchy.readme.md)
- [Shotgun surgery](src/code-smells/change-preventers/shotgun-surgery.readme.md)

#### Couplers

Code smells en los que cambios en una unidad fuerzan cambios en otra que tiene un acoplamiento muy fuerte.

- [Feature envy](src/code-smells/couplers/feature-envy.readme.md)
- [Inappropriate intimacy](src/code-smells/couplers/inappropriate-intimacy.readme.md)
- [Message chains](src/code-smells/couplers/message-chains.readme.md)
- [Middleman](src/code-smells/couplers/middleman.readme.md)

#### Dispensables

Code smells debidos a código innecesario, que introduce ruido dificultando la inteligibilidad del código.

- [Comments](src/code-smells/dispensables/comments.readme.md)
- [Data class](src/code-smells/dispensables/data-class.readme.md)
- [Dead code](src/code-smells/dispensables/dead-code.readme.md)
- [Duplicated code](src/code-smells/dispensables/duplicated-code.readme.md)
- [Lazy class](src/code-smells/dispensables/lazy-class.readme.md)

#### OOP Abusers

Code smells debido a la aplicación inadecuada de la orientación a objetos.

- [Alternative classes different interfaces](src/code-smells/oop-abusers/alternative-classes-different-interfaces.readme.md)
- [Refused bequest](src/code-smells/oop-abusers/refused-bequest.readme.md)
- [Switch statements](src/code-smells/oop-abusers/switch-statements.readme.md)
- [Temporal instance variables](src/code-smells/oop-abusers/temporal-instance-variables.readme.md)

### Técnicas de Refactoring

Ejercicios prácticos de técnicas avanzadas de refactoring:

- [Golden Master](src/refactoring/golden-master/README.md): Técnica para crear tests de caracterización sobre código legado
- [Parallel Change](src/refactoring/parallel-change/): Técnicas para realizar cambios graduales manteniendo el sistema en funcionamiento
  - [Expand-Migrate-Contract](src/refactoring/parallel-change/expand-migrate-contract/README.md)
  - [Sprout Change](src/refactoring/parallel-change/sprout-change/README.md)
  - [Wrap Change](src/refactoring/parallel-change/wrap-change/README.md)

## Ejecutar tests

### Todos los tests
```bash
dotnet test
```

### Tests con output detallado
```bash
dotnet test --logger "console;verbosity=detailed"
```

### Tests con cobertura
```bash
dotnet test /p:CollectCoverage=true /p:CoverletOutputFormat=lcov
```

### Ejecutar un test específico
```bash
dotnet test --filter "FullyQualifiedName~NombreDelTest"
```

## Estructura del Proyecto

```
csharp/
├── src/
│   ├── calisthenics-exercises/     # 9 ejercicios de calistenia
│   ├── code-smells/
│   │   ├── bloaters/               # 5 smells
│   │   ├── change-preventers/      # 3 smells
│   │   ├── couplers/               # 4 smells
│   │   ├── dispensables/           # 5 smells
│   │   └── oop-abusers/            # 4 smells
│   └── refactoring/
│       ├── golden-master/
│       └── parallel-change/
│           ├── expand-migrate-contract/
│           ├── sprout-change/
│           └── wrap-change/
├── test/                           # Tests de verificación
├── RefactoringAvanzado.sln         # Solución
├── RefactoringAvanzado.csproj      # Proyecto
└── README.md
```

## Dependencias del Proyecto

- **xUnit 2.9.0**: Framework de testing para .NET
- **FluentAssertions 6.12.1**: Librería para aserciones expresivas en tests
- **Verify.Xunit 27.2.0**: Snapshot testing (equivalente a Vitest snapshots)
- **Microsoft.NET.Test.Sdk 17.11.0**: SDK de testing de .NET

## Notas sobre la traducción de TypeScript a C#

Esta versión en C# mantiene la misma estructura y contenido que la versión TypeScript, con las siguientes adaptaciones idiomáticas:

- **Tipos**: Los tipos TypeScript se traducen a clases, structs o records según corresponda
- **Interfaces**: Se mantienen como interfaces en C#
- **Funciones**: Las funciones de nivel superior se convierten en métodos estáticos de clases
- **Tests**: Vitest se traduce a xUnit con FluentAssertions
- **Snapshots**: Verify.Xunit proporciona funcionalidad similar a los snapshots de Vitest
- **Convenciones de nombres**: Se siguen las convenciones C# (PascalCase para clases y métodos públicos)
- **Nullability**: Se aprovecha el sistema de tipos nullable de C# 12
