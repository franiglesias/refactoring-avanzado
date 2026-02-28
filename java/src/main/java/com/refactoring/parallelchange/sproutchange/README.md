# Técnica de refactorización: Cambio en Paralelo usando Sprout

Este ejercicio te ayuda a practicar cómo introducir nuevo comportamiento "haciendo brotar" (sprout) código nuevo, manteniendo el código antiguo funcionando para poder migrar los puntos de llamada de forma gradual y segura.

## Escenario

Tenemos una función de total de checkout con reglas de impuestos embebidas en línea. El producto quiere introducir políticas de impuestos por región (estándar y reducida), pero no debemos romper el comportamiento existente. Practicaremos haciendo brotar una nueva abstracción (`TaxPolicy`) y movernos hacia ella de forma incremental.

## Implementación ingenua actual (intencionalmente rígida)

Necesitarás crear una clase con una función `calculateTotal(cart, region)` con lógica de impuestos en línea:

- Región `US`: 7% plano sobre el subtotal
- Región `EU`: 20% plano solo sobre los ítems gravables (libros y comida exentos)

Queremos introducir el producto en nuevas regiones:
- Región `RU`: 10% plano sobre items gravables, excepto comida que es el 5%
- Región `UK`: 10% hasta un importe de 150 y 12% para items de más de 150 sobre los ítems gravables. Libros exentos. Comida 2% plano.

## Ejercicio: Cambio en Paralelo usando SPROUT

**Objetivo**: Introducir estrategias de política de impuestos sin romper el comportamiento actual.

### Pasos (idealmente con un commit entre cada paso)

1. **Haz brotar** un nuevo concepto `TaxPolicy` (interfaz) con un método `compute(cart): double`.
   NO cambies aún `calculateTotal`.
   - Crea implementaciones `USTaxPolicy` y `EUTaxPolicy`
   - Mantenlas sin usar al principio (build en verde)

2. **Añade un parámetro opcional** a `calculateTotal`: usar un patrón como Builder o parámetros opcionales.
   Por defecto, usa el comportamiento actual si no se proporciona.
   - Cuando `policy` esté presente, delega el cálculo de impuestos en él
   - De lo contrario, conserva la lógica embebida

3. **Crea una política adaptadora** que reproduzca el comportamiento actual (`LegacyInlineTaxPolicy`)
   para demostrar paridad.
   - Úsala para validar que no hay cambio de comportamiento

4. **Migra los puntos de llamada** para pasar una política.
   - Primero pasa `LegacyInlineTaxPolicy` para mantener el comportamiento
   - Luego cambia a las políticas `EU`/`US` según convenga

5. **Finalmente, elimina las ramas** de impuesto en línea de `calculateTotal` una vez que todos los puntos de llamada usen una política.

### Criterios de aceptación

- Todos los totales permanecen numéricamente idénticos hasta que la migración los cambie intencionalmente
- Los nombres y responsabilidades son claros
- El código documenta los pasos de sprout mediante commits o comentarios

## Estructura sugerida

```java
public interface TaxPolicy {
    double compute(Cart cart);
}

public class USTaxPolicy implements TaxPolicy {
    @Override
    public double compute(Cart cart) {
        double subtotal = cart.getSubtotal();
        return subtotal * 0.07; // 7% plano
    }
}

public class EUTaxPolicy implements TaxPolicy {
    @Override
    public double compute(Cart cart) {
        double taxable = cart.getTaxableAmount(List.of("books", "food"));
        return taxable * 0.20; // 20% sobre gravables
    }
}

public class Calculator {
    public double calculateTotal(Cart cart, String region, TaxPolicy policy) {
        double subtotal = cart.getSubtotal();
        double tax;

        if (policy != null) {
            tax = policy.compute(cart);
        } else {
            // Lógica legacy embebida
            if ("US".equals(region)) {
                tax = subtotal * 0.07;
            } else if ("EU".equals(region)) {
                double taxable = cart.getTaxableAmount(List.of("books", "food"));
                tax = taxable * 0.20;
            } else {
                tax = 0;
            }
        }

        return subtotal + tax;
    }
}
```

## Ejecución

1. Implementa la lógica inicial con impuestos embebidos
2. Crea tests que verifiquen el comportamiento actual
3. Aplica la técnica Sprout paso a paso
4. Verifica que los tests siguen pasando en cada paso
5. Finalmente elimina el código legacy

## Recursos adicionales

- Ver ejercicio equivalente en TypeScript: `typescript/src/refactoring/parallel-change/sprout-change/`
