# Expand-Migrate-Contract - Ejercicio en Python

## 📚 Documentación Completa

👉 **[Ver documentación completa de Expand-Migrate-Contract](../../../../../docs/refactoring/parallel-change/expand-migrate-contract.md)**

La documentación completa incluye:
- Definición y explicación detallada de la técnica
- Proceso paso a paso con ejemplos de código
- Problemas comunes y soluciones
- Referencias en español e inglés

## 🎯 Ejercicio

**Objetivo**: Refactorizar el campo `name` de `User` para que sea `first_name` y `last_name` sin romper los tests

**Archivo a refactorizar**: `expand_migrate_contract.py`

## Ejecutar tests

```bash
pytest src/refactoring/parallel_change/ -k expand -v
```

## Escenario

Tenemos clase `User` con campo `name` (nombre completo) y cuatro funciones consumidoras:
- `format_greeting(user)` — saludo al usuario
- `format_email_header(user)` — cabecera de email
- `format_display_name(user)` — nombre para mostrar con ID
- `build_user_summary(users)` — listado de nombres

## Pasos recomendados

### Fase 1: EXPAND (Expandir)
Añadir `first_name` y `last_name` SIN eliminar `name`

### Fase 2: MIGRATE (Migrar)
Migrar funciones una por una:
1. `format_greeting` → usa `first_name`
2. `format_email_header` → usa `first_name` + `last_name`
3. `format_display_name` → usa `first_name` + `last_name`
4. `build_user_summary` → usa `last_name`, `first_name`

### Fase 3: CONTRACT (Contraer)
Eliminar campo `name` cuando ningún consumidor lo use

## Criterios de aceptación

- ✅ Tests nunca fallan durante el proceso
- ✅ Commits pequeños entre cada paso
- ✅ Campo `name` eliminado al final
- ✅ Todas las funciones migradas
