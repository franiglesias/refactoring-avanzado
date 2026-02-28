# Técnica de refactorización: Cambio en Paralelo usando Sprout

Este ejercicio te ayuda a practicar cómo introducir nuevo comportamiento "haciendo brotar" (sprout) código nuevo, manteniendo el código antiguo funcionando para poder migrar los puntos de llamada de forma gradual y segura.

## Escenario

Tenemos una función de total de checkout con reglas de impuestos embebidas en línea. El producto quiere introducir políticas de impuestos por región (estándar y reducida), pero no debemos romper el comportamiento existente. Practicaremos haciendo brotar una nueva abstracción (`TaxPolicy`) y movernos hacia ella de forma incremental.

## Ejecutar tests

```shell
go test ./refactoring/parallel_change/sprout_change
```

## Implementación ingenua actual (intencionalmente rígida)

En `sprout_change.go` existe la función `CalculateTotal(cart, region)` con lógica de impuestos en línea:

- Región `US`: 7% plano sobre el subtotal.
- Región `EU`: 20% plano solo sobre los ítems gravables (libros y comida exentos).

Queremos introducir el producto en nuevas regiones, por lo que necesitamos una forma de configurar estas reglas.

- Región `RU`: 10% plano sobre items gravables, excepto comida que es el 5%.
- Región `UK`: 10% hasta un importe de 150 y 12% para items de más de 150 sobre los ítems gravables. Libros exentos. Comida 2% plano.

También hay funciones auxiliares como `RoundCurrency` y un uso de ejemplo en `DemoSprout()`.

## Ejercicio: Cambio en Paralelo usando SPROUT

Objetivo: Introducir estrategias de política de impuestos sin romper el comportamiento actual.

### Pasos (idealmente con un commit entre cada paso)

1. Haz _brotar_ un nuevo concepto `TaxPolicy` (interfaz) con un método `Compute(cart Cart) float64`. NO cambies aún `CalculateTotal`.

- Crea ejemplos `USTaxPolicy` y `EUTaxPolicy`.
- Mantenlos sin usar al principio (build en verde).

2. Añade un parámetro opcional a `CalculateTotal`: `opts *Options` donde `Options` tiene un campo `Policy TaxPolicy`. Por defecto, usa el comportamiento actual si no se proporciona. A la larga, dejaremos de pasar `region` a `CalculateTotal`.

- Cuando `opts.Policy` esté presente, delega el cálculo de impuestos en él; de lo contrario, conserva la lógica embebida.

3. Crea una política adaptadora que reproduzca el comportamiento actual (`LegacyInlineTaxPolicy`) para demostrar paridad.

- Úsala desde `DemoSprout` para validar que no hay cambio de comportamiento.

4. Migra los puntos de llamada (aquí solo `DemoSprout`) para pasar una política.

- Primero pasa `LegacyInlineTaxPolicy` para mantener el comportamiento.
- Luego cambia a las políticas `EU`/`US` según convenga.

5. Finalmente, elimina las ramas de impuesto en línea de `CalculateTotal` una vez que todos los puntos de llamada usen una política.

- Aceptación: `CalculateTotal` delega completamente a una `TaxPolicy`; la lógica antigua se elimina.

### Criterios de aceptación

- Todos los totales permanecen numéricamente idénticos hasta que la migración (paso 4, segundo punto) los cambie intencionalmente.
- No deben fallar las pruebas; los nombres y responsabilidades son claros.
- El código (o los commits) documenta(n) los pasos de sprout mediante commits o comentarios.
