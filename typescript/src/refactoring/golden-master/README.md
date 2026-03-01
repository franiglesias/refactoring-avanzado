# Golden Master - Ejercicio en TypeScript

## 📚 Documentación Completa

👉 **[Ver documentación completa de Golden Master](../../../../docs/refactoring/golden-master/README.md)**

La documentación completa incluye:
- Definición y explicación detallada de la técnica
- Proceso paso a paso con ejemplos de código
- Problemas comunes y soluciones
- Referencias en español e inglés

## 🎯 Ejercicio

**Objetivo**: Crear una prueba de Golden Master para el código legado (`ReceiptPrinter`)

**Archivo a refactorizar**: `receipt-printer.ts`

## Ejecutar tests

```shell
npm run test -- src/refactoring/golden-master/golden-master.test.ts
```

## Pasos recomendados

1. Identificar fuentes de no determinismo
2. Introducir costuras (SEAMS) mínimas sin cambiar el comportamiento
3. Generar conjunto amplio de entradas estables
4. Capturar la salida maestra con snapshots
5. Refactorizar con seguridad para invertir dependencias
6. Introducir tests unitarios para reemplazar el Golden Master

## Criterios de aceptación

- ✅ La prueba captura múltiples casos de entrada
- ✅ Falla ante cambios en la salida
- ✅ Fuentes de no determinismo controladas mediante Seams
- ✅ Comportamiento original disponible (compatibilidad)
