# Sprout Change - Ejercicio en Java

## 📚 Documentación Completa

👉 **[Ver documentación completa de Sprout Change](../../../../../../../docs/refactoring/parallel-change/sprout-change.md)**

La documentación completa incluye:
- Definición y explicación detallada de la técnica
- Proceso paso a paso con ejemplos de código
- Problemas comunes y soluciones
- Referencias en español e inglés

## 🎯 Ejercicio

**Objetivo**: Introducir estrategias de política de impuestos sin romper el comportamiento actual

**Archivo a refactorizar**: Crear clase con `calculateTotal(cart, region)`

## Escenario

Función con lógica de impuestos embebida:
- Región `US`: 7% plano sobre el subtotal
- Región `EU`: 20% plano solo sobre ítems gravables (libros y comida exentos)

Queremos añadir nuevas regiones con políticas diferentes sin romper el código existente.

## Pasos recomendados

1. Hacer brotar interfaz `TaxPolicy` con método `compute(cart): double`
2. Crear implementaciones `USTaxPolicy` y `EUTaxPolicy`
3. Añadir parámetro opcional a `calculateTotal` para `TaxPolicy`
4. Crear política adaptadora `LegacyInlineTaxPolicy` para reproducir comportamiento actual
5. Migrar puntos de llamada para pasar una política
6. Eliminar ramas de impuesto en línea cuando todo use políticas

## Criterios de aceptación

- ✅ Totales permanecen numéricamente idénticos durante la migración
- ✅ `calculateTotal` delega completamente a `TaxPolicy`
- ✅ Lógica antigua eliminada
- ✅ Tests siguen pasando en cada paso
