# Expand-Migrate-Contract - Ejercicio en TypeScript

## 📚 Documentación Completa

👉 **[Ver documentación completa de Expand-Migrate-Contract](../../../../../docs/refactoring/parallel-change/expand-migrate-contract.md)**

La documentación completa incluye:
- Definición y explicación detallada de la técnica
- Proceso paso a paso con ejemplos de código
- Problemas comunes y soluciones
- Referencias en español e inglés

## 🎯 Ejercicio

**Objetivo**: Refactorizar el campo `name` de `User` para que sea `firstName` y `lastName` sin romper los tests

**Archivo a refactorizar**: `user.ts`

## Ejecutar tests

```shell
npm run test -- src/refactoring/parallel-change/expand-migrate-contract/user.test.ts
```

## Escenario

Tenemos clase `User` con campo `name` (nombre completo) y cuatro funciones consumidoras:
- `formatGreeting(user)` — saludo al usuario
- `formatEmailHeader(user)` — cabecera de email
- `formatDisplayName(user)` — nombre para mostrar con ID
- `buildUserSummary(users)` — listado de nombres

## Pasos recomendados

### Fase 1: EXPAND (Expandir)
Añadir `firstName` y `lastName` SIN eliminar `name`

### Fase 2: MIGRATE (Migrar)
Migrar funciones una por una:
1. `formatGreeting` → usa `firstName`
2. `formatEmailHeader` → usa `firstName` + `lastName`
3. `formatDisplayName` → usa `firstName` + `lastName`
4. `buildUserSummary` → usa `lastName`, `firstName`

### Fase 3: CONTRACT (Contraer)
Eliminar campo `name` cuando ningún consumidor lo use

## Uso del script TCR

```bash
npm run tcr
```

Tests pasan → commit automático. Tests fallan → revert automático.

## Criterios de aceptación

- ✅ Tests nunca fallan durante el proceso
- ✅ Commits pequeños entre cada paso
- ✅ Campo `name` eliminado al final
- ✅ Todas las funciones migradas
