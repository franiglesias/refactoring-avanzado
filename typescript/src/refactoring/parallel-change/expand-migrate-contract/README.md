# Técnica de refactorización: Cambio en Paralelo (Expand–Migrate–Contract)

Este ejercicio te ayuda a practicar la técnica de **Parallel Change** en su forma más pura:
expandir la interfaz, migrar los consumidores y contraer la interfaz antigua.


## Escenario

Tenemos una clase `User` con un campo `name` (nombre completo como string único) y **cuatro
funciones consumidoras** que dependen de ese campo:

- `formatGreeting(user)` — saludo al usuario.
- `formatEmailHeader(user)` — cabecera de email con nombre y dirección.
- `formatDisplayName(user)` — nombre para mostrar con ID.
- `buildUserSummary(users)` — listado de nombres de usuarios.

**Objetivo**: Refactorizar `name` para que sea `firstName` y `lastName`, aplicando cambio en
paralelo para no romper nunca los tests.

## Ejecutar tests

```shell
npm run test -- src/refactoring/parallel-change/expand-migrate-contract/user.test.ts
```

## Ejercicio: Expand–Migrate–Contract

### Fase 1: EXPAND (Expandir)

Añade los campos `firstName` y `lastName` a la clase `User` **sin eliminar** el campo `name`.
Puedes hacer que `name` se calcule a partir de los nuevos campos, o aceptar los tres en el
constructor.

- Los tests deben seguir pasando sin cambios.
- Haz commit (o usa el script TCR).

### Fase 2: MIGRATE (Migrar)

Migra las funciones consumidoras **una por una** para que usen `firstName` y/o `lastName` en lugar
de `name`. Después de migrar cada función:

1. Actualiza el test correspondiente si es necesario.
2. Ejecuta los tests — deben pasar.
3. Haz commit (o usa el script TCR).

Orden sugerido de migración:

1. `formatGreeting` — usa solo `firstName`.
2. `formatEmailHeader` — usa `firstName` + `lastName`.
3. `formatDisplayName` — usa `firstName` + `lastName`.
4. `buildUserSummary` — usa `lastName`, `firstName`.

### Fase 3: CONTRACT (Contraer)

Una vez que **ningún consumidor** use `name`:

1. Elimina el campo `name` de la clase `User`.
2. Ajusta el constructor para que solo reciba `firstName` y `lastName`.
3. Los tests deben pasar.
4. Haz commit (o usa el script TCR).

## Uso del script TCR

Ejecuta el script TCR después de cada cambio pequeño:

```bash
npm run tcr
```

Si los tests pasan, se hace commit automáticamente. Si fallan, se revierten los cambios al último
commit. Esto fuerza pasos pequeños y seguros.
