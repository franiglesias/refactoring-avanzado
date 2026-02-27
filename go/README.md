# Curso Refactoring Avanzado (Go)

Ejemplos y ejercicios del Curso de Refactoring Avanzado convertidos a Go.

## Preparación

### Requisitos

- Go 1.21 o superior

### Instalar dependencias

```bash
go mod download
```

### Ejecutar tests

```bash
# Ejecutar todos los tests
go test ./...

# Ejecutar tests con verbose output
go test -v ./...

# Ejecutar tests con cobertura
go test -cover ./...

# Ver reporte detallado de cobertura
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Ejecutar tests específicos
go test -v ./refactoring -run TestReceiptPrinter
```

## Contenido

### Técnicas de Refactoring

#### Golden Master

Técnica para caracterizar el comportamiento de código legado sin tests.

- Ejercicio: [refactoring/golden_master.go](./refactoring/golden_master.go)

```bash
go test -v ./refactoring -run TestReceiptPrinter
```

#### Parallel Change

Técnicas para realizar cambios seguros en código en producción:

- Expand-Migrate-Contract
- Sprout Change
- Wrap Change

### Mantenimiento diario de código: Calistenia

Un conjunto de reglas para escribir código nuevo o evaluar código existente y modificarlo para acercarlo a un mejor diseño.

Ejercicios en [calisthenics_exercises/](./calisthenics_exercises):

1. Un nivel de indentación por método - [01_one_level_indentation.go](./calisthenics_exercises/01_one_level_indentation.go)
2. No uses ELSE - [02_no_else.go](./calisthenics_exercises/02_no_else.go)
3. Envuelve primitivos - [03_wrap_primitives.go](./calisthenics_exercises/03_wrap_primitives.go)
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

- [Data clump](code_smells/bloaters/data_clump.go) - Grupos de datos que aparecen juntos repetidamente
- [Long parameter list](code_smells/bloaters/long_parameter_list.go) - Funciones con muchos parámetros
- [Primitive obsession](code_smells/bloaters/primitive_obsession.go) - Uso excesivo de tipos primitivos

#### Couplers

Code smells en los que cambios en una unidad fuerzan cambios en otra que tiene un acoplamiento muy fuerte.

- [Feature envy](code_smells/couplers/feature_envy.go) - Métodos más interesados en otras clases

#### Dispensables

Code smells debidos a código innecesario, que introduce ruido dificultando la inteligibilidad del código.

- [Duplicated code](code_smells/dispensables/duplicated_code.go) - Código duplicado en múltiples lugares

#### OOP Abusers

Code smells debido a la aplicación inadecuada de la orientación a objetos.

- [Switch statements](code_smells/oop_abusers/switch_statements.go) - Uso de switch en lugar de polimorfismo

## Estructura del Proyecto

```
go/
├── refactoring/              # Técnicas de refactoring
│   ├── golden_master.go
│   └── golden_master_test.go
├── calisthenics_exercises/   # Ejercicios de Object Calisthenics
│   ├── 01_one_level_indentation.go
│   ├── 02_no_else.go
│   └── 03_wrap_primitives.go
└── code_smells/              # Ejemplos de Code Smells
    ├── bloaters/
    ├── couplers/
    ├── dispensables/
    └── oop_abusers/
```

## Comandos Útiles

```bash
# Descargar dependencias
go mod download

# Ejecutar todos los tests
go test ./...

# Ejecutar tests en modo watch (requiere herramienta externa)
# Instalar: go install github.com/cespare/reflex@latest
reflex -r '\.go$' -s -- go test ./...

# Ver cobertura de tests
go test -cover ./...

# Generar reporte HTML de cobertura
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Formatear código
go fmt ./...

# Verificar código con el linter
go vet ./...

# Instalar y usar golangci-lint (recomendado)
# https://golangci-lint.run/usage/install/
golangci-lint run

# Ejecutar tests de un paquete específico
go test -v ./refactoring

# Ejecutar un test específico
go test -v ./refactoring -run TestReceiptPrinter_WithFixedRandomness
```

## Versión de Go

Este proyecto requiere Go 1.21 o superior.

## Diferencias con otras versiones

- Se usan `struct` en lugar de clases
- `PascalCase` para nombres exportados y `camelCase` para nombres no exportados
- Paquetes en lugar de módulos de ES6
- `go test` con el paquete `testing` en lugar de otros frameworks
- Interfaces implícitas en lugar de explícitas
- Composición sobre herencia (Go no tiene herencia de clases)

## Convenciones de Go

- Los archivos de test terminan en `_test.go`
- Los tests son funciones que empiezan con `Test`
- Se usa el paquete `testing` de la librería estándar
- Se recomienda usar `github.com/stretchr/testify` para aserciones más expresivas
- La carpeta de cada paquete contiene su código y sus tests
