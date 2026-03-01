# Middleman - Ejercicio en TypeScript

## 📚 Documentación Completa

👉 **[Ver documentación completa de Middleman](../../../../docs/code-smells/couplers/middleman.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `middleman.ts`

**Tarea**: Añade una funcionalidad `searchByPrefix` en `Catalog` y propágala a través de `Shop`.

## Ejecutar tests

```bash
npm test -- middleman.test.ts
```

## Problema a experimentar

Añadirás métodos a `Shop` que solo pasan a través hacia `Catalog`, fomentando la duplicación accidental y ocultando dónde vive el comportamiento real cuando necesites cambiarlo después.
