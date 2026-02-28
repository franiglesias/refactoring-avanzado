# Técnica de refactorización: Cambio en Paralelo (Expand–Migrate–Contract)

Este ejercicio te ayuda a practicar la técnica de **Parallel Change** en su forma más pura:
expandir la interfaz, migrar los consumidores y contraer la interfaz antigua.

## Escenario

Tenemos una clase `User` con un campo `name` (nombre completo como string único) y **cuatro
funciones consumidoras** que dependen de ese campo:

- `format_greeting(user)` — saludo al usuario.
- `format_email_header(user)` — cabecera de email con nombre y dirección.
- `format_display_name(user)` — nombre para mostrar con ID.
- `build_user_summary(users)` — listado de nombres de usuarios.

**Objetivo**: Refactorizar `name` para que sea `first_name` y `last_name`, aplicando cambio en
paralelo para no romper nunca los tests.

## Ejecutar tests

```shell
pytest src/refactoring/parallel_change/test_expand_migrate_contract.py -v
```

## Ejercicio: Expand–Migrate–Contract

### Fase 1: EXPAND (Expandir)

Añade los campos `first_name` y `last_name` a la clase `User` **sin eliminar** el campo `name`.
Puedes hacer que `name` se calcule a partir de los nuevos campos, o aceptar los tres en el
constructor.

- Los tests deben seguir pasando sin cambios.
- Haz commit (o usa el script TCR).

### Fase 2: MIGRATE (Migrar)

Migra las funciones consumidoras **una por una** para que usen `first_name` y/o `last_name` en lugar
de `name`. Después de migrar cada función:

1. Actualiza el test correspondiente si es necesario.
2. Ejecuta los tests — deben pasar.
3. Haz commit (o usa el script TCR).

Orden sugerido de migración:

1. `format_greeting` — usa solo `first_name`.
2. `format_email_header` — usa `first_name` + `last_name`.
3. `format_display_name` — usa `first_name` + `last_name`.
4. `build_user_summary` — usa `last_name`, `first_name`.

### Fase 3: CONTRACT (Contraer)

Una vez que **ningún consumidor** use `name`:

1. Elimina el campo `name` de la clase `User`.
2. Ajusta el constructor para que solo reciba `first_name` y `last_name`.
3. Los tests deben pasar.
4. Haz commit (o usa el script TCR).

## Uso del script TCR

El patrón TCR (Test && Commit || Revert) puede ayudarte a mantener pasos pequeños:

```bash
pytest src/refactoring/parallel_change/test_expand_migrate_contract.py && git commit -am "refactor: ..." || git reset --hard
```

Si los tests pasan, se hace commit automáticamente. Si fallan, se revierten los cambios al último
commit. Esto fuerza pasos pequeños y seguros.
