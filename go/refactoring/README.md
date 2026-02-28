# Ejercicios de Refactoring

Este directorio contiene ejercicios prácticos de refactoring en Go.

## Contenido

### 1. Golden Master

Técnica para caracterizar el comportamiento de código legado antes de refactorizarlo.

Los archivos ya existentes son:
- `golden_master.go` - Implementación de ejemplo
- `golden_master_test.go` - Tests usando la técnica Golden Master

### 2. Parallel Change (Cambio en Paralelo)

Conjunto de técnicas para refactorizar código de forma segura manteniendo el sistema funcionando en todo momento.

Ver [directorio parallel_change](./parallel_change/) para más detalles.

Incluye:
- **Sprout Change**: Introducir nuevo comportamiento en paralelo
- **Wrap Change**: Envolver dependencias para mejorarlas
- **Expand-Migrate-Contract**: Refactorizar interfaces en tres fases

## Ejecutar tests

Para ejecutar todos los tests de refactoring:

```shell
go test ./refactoring/...
```

Para ejecutar tests específicos:

```shell
# Golden Master
go test ./refactoring -run TestGoldenMaster

# Parallel Change
go test ./refactoring/parallel_change/...
```

## Recursos adicionales

Estos ejercicios están basados en técnicas descritas en:
- "Working Effectively with Legacy Code" de Michael Feathers
- "Refactoring" de Martin Fowler
- "Object Calisthenics" de Jeff Bay
