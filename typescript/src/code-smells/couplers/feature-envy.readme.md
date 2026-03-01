# Feature Envy - Ejercicio en TypeScript

## 📚 Documentación Completa

👉 **[Ver documentación completa de Feature Envy](../../../../docs/code-smells/couplers/feature-envy.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `feature-envy.ts`

**Tarea**: Añade envío gratis para clientes en ciertas ciudades y un recargo de fin de semana.

## Ejecutar tests

```bash
npm test -- feature-envy.test.ts
```

## Problema a experimentar

Probablemente, seguirás añadiendo condiciones dentro de `ShippingCalculator` que dependen de detalles internos de `Customer`, esparciendo reglas en el lugar equivocado y volviendo frágiles los cambios.
