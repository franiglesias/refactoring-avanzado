# Expand-Migrate-Contract - Ejercicio en C#

## 📚 Documentación Completa

👉 **[Ver documentación completa de Expand-Migrate-Contract](../../../../../../docs/refactoring/parallel-change/expand-migrate-contract.md)**

La documentación completa incluye:
- Definición y explicación detallada de la técnica
- Proceso paso a paso con ejemplos de código
- Problemas comunes y soluciones
- Referencias en español e inglés

## 🎯 Ejercicio

**Objetivo**: Refactorizar el campo `Name` de `User` para que sea `FirstName` y `LastName` sin romper los tests

**Archivo a refactorizar**: `ExpandMigrateContract.cs`

## Ejecutar tests

```shell
dotnet test --filter "ExpandMigrateContract"
```

## Escenario

Tenemos clase `User` con campo `Name` (nombre completo) y cuatro funciones consumidoras:
- `FormatGreeting(user)` — saludo al usuario
- `FormatEmailHeader(user)` — cabecera de email
- `FormatDisplayName(user)` — nombre para mostrar con ID
- `BuildUserSummary(users)` — listado de nombres

## Pasos recomendados

### Fase 1: EXPAND (Expandir)
Añadir `FirstName` y `LastName` SIN eliminar `Name`

### Fase 2: MIGRATE (Migrar)
Migrar funciones una por una:
1. `FormatGreeting` → usa `FirstName`
2. `FormatEmailHeader` → usa `FirstName` + `LastName`
3. `FormatDisplayName` → usa `FirstName` + `LastName`
4. `BuildUserSummary` → usa `LastName`, `FirstName`

### Fase 3: CONTRACT (Contraer)
Eliminar campo `Name` cuando ningún consumidor lo use

## Criterios de aceptación

- ✅ Tests nunca fallan durante el proceso
- ✅ Commits pequeños entre cada paso
- ✅ Campo `Name` eliminado al final
- ✅ Todas las funciones migradas
