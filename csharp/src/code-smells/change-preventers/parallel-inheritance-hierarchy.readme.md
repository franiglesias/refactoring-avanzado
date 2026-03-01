# Parallel Inheritance Hierarchy - Ejercicio en C#

## 📚 Documentación Completa

👉 **[Ver documentación completa de Parallel Inheritance Hierarchy](../../../../docs/code-smells/change-preventers/parallel-inheritance-hierarchy.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `ParallelInheritanceHierarchy.cs`

**Tarea**: Añade un componente `Image` que muestre una imagen.

## Ejecutar tests

```bash
dotnet test --filter "ParallelInheritanceHierarchy"
```

## Problema a experimentar

Necesitarás añadir Image y renderImage a Renderer, e implementarlo en todos los renderers, mostrando cambios en paralelo.
