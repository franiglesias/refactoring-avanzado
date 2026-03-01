# Parallel Inheritance Hierarchy - Ejercicio en Go

## 📚 Documentación Completa

👉 **[Ver documentación completa de Parallel Inheritance Hierarchy](../../../docs/code-smells/change-preventers/parallel-inheritance-hierarchy.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `parallel_inheritance_hierarchy.go`

**Tarea**: Añade un componente `Image` que muestre una imagen.

## Ejecutar tests

```bash
go test ./code_smells/change_preventers/parallel_inheritance_hierarchy_test.go
```

## Problema a experimentar

Necesitarás añadir Image y renderImage a Renderer, e implementarlo en todos los renderers, mostrando cambios en paralelo.
