# Sprout Change - Ejercicio en TypeScript

## 📚 Documentación Completa

👉 **[Ver documentación completa de Sprout Change](../../../../../docs/refactoring/parallel-change/sprout-change.md)**

La documentación completa incluye:
- Definición y explicación detallada de la técnica
- Proceso paso a paso con ejemplos de código
- Problemas comunes y soluciones
- Referencias en español e inglés

## 🎯 Ejercicio

**Objetivo**: Introducir estrategias de política de impuestos sin romper el comportamiento actual

**Archivo a refactorizar**: `sprout-change.ts`

## Ejecutar tests

```shell
npm run test -- src/refactoring/parallel-change/sprout-change/sprout-change.test.ts
```

## Escenario

Tenemos `calculateTotal(cart, region)` con lógica de impuestos embebida:
- Región `US`: 7% plano sobre el subtotal
- Región `EU`: 20% plano solo sobre ítems gravables (libros y comida exentos)

Queremos añadir nuevas regiones con políticas diferentes sin romper el código existente.

## Pasos recomendados

1. Hacer brotar el concepto `TaxPolicy` (interfaz) con método `compute(cart): number`
2. Añadir parámetro opcional a `calculateTotal`: `opts?: { policy?: TaxPolicy }`
3. Crear política adaptadora `LegacyInlineTaxPolicy` para reproducir comportamiento actual
4. Migrar puntos de llamada para pasar una política
5. Eliminar ramas de impuesto en línea cuando todo use políticas

## Criterios de aceptación

- ✅ Totales permanecen numéricamente idénticos durante la migración
- ✅ No fallan los tipos de TypeScript
- ✅ `calculateTotal` delega completamente a `TaxPolicy`
- ✅ Lógica antigua eliminada
