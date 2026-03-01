# Divergent Change - Ejercicio en TypeScript

## 📚 Documentación Completa

👉 **[Ver documentación completa de Divergent Change](../../../../docs/code-smells/change-preventers/divergent-change.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `divergent-change.ts`

**Tarea**: Añade un número de teléfono con validación, inclúyelo en las exportaciones y envía un SMS.

## Ejecutar tests

```bash
npm test -- divergent-change.test.ts
```

## Problema a experimentar

Tocarás validación, almacenamiento, exportAsJson/Csv y mensajería en un solo lugar, demostrando cómo un cambio fuerza ediciones en responsabilidades no relacionadas.
