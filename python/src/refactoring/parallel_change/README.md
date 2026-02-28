# Parallel Change (Cambio en Paralelo)

El **Parallel Change** (también conocido como **Expand-Contract**) es una técnica de refactorización que permite realizar cambios significativos en un sistema de forma segura y gradual, sin romper el código existente.

## Concepto

En lugar de hacer un cambio brusco que podría romper dependencias, el parallel change se realiza en fases:

1. **EXPAND**: Añadir la nueva funcionalidad sin eliminar la antigua
2. **MIGRATE**: Migrar gradualmente los consumidores a la nueva funcionalidad
3. **CONTRACT**: Eliminar la funcionalidad antigua una vez que nadie la use

Esta técnica es especialmente útil cuando:
- Tienes múltiples consumidores de una API
- No puedes cambiar todos los puntos de uso de una vez
- Necesitas mantener el sistema funcionando durante el refactoring
- Trabajas en un equipo y otros desarrolladores dependen de tu código

## Ejercicios

### 1. Expand-Migrate-Contract

El ejemplo clásico de parallel change: cambiar la estructura de datos de una clase mientras mantienes compatibilidad.

[Ver ejercicio de Expand-Migrate-Contract](expand_migrate_contract_readme.md)

**Archivos:**
- `expand_migrate_contract.py`
- `test_expand_migrate_contract.py`

### 2. Sprout Change

Hacer "brotar" (sprout) nueva funcionalidad sin modificar el código existente, permitiendo migración gradual.

[Ver ejercicio de Sprout Change](sprout_change_readme.md)

**Archivos:**
- `sprout_change.py`

### 3. Wrap Change

Envolver dependencias problemáticas con una capa que añade funcionalidad sin cambiar la interfaz.

[Ver ejercicio de Wrap Change](wrap_change_readme.md)

**Archivos:**
- `wrap_change.py`

## Beneficios

- ✅ **Seguridad**: Los tests siguen pasando en cada paso
- ✅ **Reversibilidad**: Puedes revertir cambios sin perder trabajo
- ✅ **Colaboración**: Otros desarrolladores pueden seguir trabajando
- ✅ **Incremental**: Puedes hacer commits pequeños y frecuentes
- ✅ **Sin "Big Bang"**: Evitas cambios masivos que puedan fallar

## Patrón TCR (Test && Commit || Revert)

Estos ejercicios se benefician del patrón TCR:

```bash
pytest path/to/test.py && git commit -am "refactor: mensaje" || git reset --hard
```

Si los tests pasan → commit automático
Si los tests fallan → revertir cambios

Esto te fuerza a hacer cambios pequeños y seguros.

## Referencias

- [Parallel Change - Martin Fowler](https://martinfowler.com/bliki/ParallelChange.html)
- [Working Effectively with Legacy Code - Michael Feathers](https://www.amazon.com/Working-Effectively-Legacy-Michael-Feathers/dp/0131177052)
