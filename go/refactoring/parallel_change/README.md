# Técnicas de Refactorización: Cambio en Paralelo

Este directorio contiene ejercicios prácticos para aplicar diferentes técnicas de refactorización segura en código legado.

## Técnicas incluidas

### 1. Sprout Change (Cambio Brote)

Técnica para introducir nuevo comportamiento creando código nuevo en paralelo con el antiguo, permitiendo migración gradual.

[Ver ejercicio](./sprout_change/)

### 2. Wrap Change (Cambio Envolvente)

Técnica para mejorar una dependencia problemática envolviéndola sin cambiar su interfaz pública.

[Ver ejercicio](./wrap_change/)

### 3. Expand-Migrate-Contract

Técnica para refactorizar interfaces existentes de forma segura en tres fases:
- **Expand**: Añadir nueva interfaz sin eliminar la antigua
- **Migrate**: Migrar consumidores uno por uno
- **Contract**: Eliminar interfaz antigua una vez migrados todos los consumidores

[Ver ejercicio](./expand_migrate_contract/)

## Principios generales

Todas estas técnicas comparten principios comunes:

1. **Nunca romper los tests**: Cada paso debe mantener el sistema funcionando
2. **Cambios incrementales**: Pequeños pasos que pueden comitearse individualmente
3. **Paralelismo**: Mantener código antiguo y nuevo funcionando simultáneamente durante la transición
4. **Migración gradual**: Mover consumidores uno por uno en lugar de cambios masivos

## Ejecutar todos los tests

```shell
go test ./refactoring/parallel_change/...
```
