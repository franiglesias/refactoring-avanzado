# Técnica de refactorización: Cambio en Paralelo (Expand–Migrate–Contract)

Este ejercicio te ayuda a practicar la técnica de **Parallel Change** en su forma más pura: expandir la interfaz, migrar los consumidores y contraer la interfaz antigua.

## Escenario

Tenemos una estructura `User` con un campo `Name` (nombre completo como string único) y **cuatro funciones consumidoras** que dependen de ese campo:

- `FormatGreeting(user)` — saludo al usuario.
- `FormatEmailHeader(user)` — cabecera de email con nombre y dirección.
- `FormatDisplayName(user)` — nombre para mostrar con ID.
- `BuildUserSummary(users)` — listado de nombres de usuarios.

**Objetivo**: Refactorizar `Name` para que sea `FirstName` y `LastName`, aplicando cambio en paralelo para no romper nunca los tests.

## Ejecutar tests

```shell
go test ./refactoring/parallel_change/expand_migrate_contract
```

## Ejercicio: Expand–Migrate–Contract

### Fase 1: EXPAND (Expandir)

Añade los campos `FirstName` y `LastName` a la estructura `User` **sin eliminar** el campo `Name`. Puedes hacer que `Name` se calcule a partir de los nuevos campos, o aceptar los tres en la función constructora.

- Los tests deben seguir pasando sin cambios.
- Haz commit.

### Fase 2: MIGRATE (Migrar)

Migra las funciones consumidoras **una por una** para que usen `FirstName` y/o `LastName` en lugar de `Name`. Después de migrar cada función:

1. Actualiza el test correspondiente si es necesario.
2. Ejecuta los tests — deben pasar.
3. Haz commit.

Orden sugerido de migración:

1. `FormatGreeting` — usa solo `FirstName`.
2. `FormatEmailHeader` — usa `FirstName` + `LastName`.
3. `FormatDisplayName` — usa `FirstName` + `LastName`.
4. `BuildUserSummary` — usa `LastName`, `FirstName`.

### Fase 3: CONTRACT (Contraer)

Una vez que **ningún consumidor** use `Name`:

1. Elimina el campo `Name` de la estructura `User`.
2. Ajusta la función constructora para que solo reciba `FirstName` y `LastName`.
3. Los tests deben pasar.
4. Haz commit.

## Uso de TCR (Test && Commit || Revert)

Puedes ejecutar los tests después de cada cambio pequeño:

```bash
go test ./refactoring/parallel_change/expand_migrate_contract && git commit -am "paso X" || git reset --hard
```

Si los tests pasan, se hace commit automáticamente. Si fallan, se revierten los cambios al último commit. Esto fuerza pasos pequeños y seguros.
