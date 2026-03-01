# Divergent Change - Ejercicio en C#

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

**Archivo**: `DivergentChange.cs`

**Tarea**: Añade un número de teléfono con validación, inclúyelo en las exportaciones y envía un SMS.

## Ejecutar tests

```bash
dotnet test --filter "DivergentChange"
```

## Problema a experimentar

Tocarás validación, almacenamiento, exportAsJson/Csv y mensajería en un solo lugar, demostrando cómo un cambio fuerza ediciones en responsabilidades no relacionadas.
