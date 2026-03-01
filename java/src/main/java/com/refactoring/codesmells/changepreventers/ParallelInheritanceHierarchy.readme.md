# Parallel Inheritance Hierarchy - Ejercicio en Java

## 📚 Documentación Completa

👉 **[Ver documentación completa de Parallel Inheritance Hierarchy](../../../../../../../docs/code-smells/change-preventers/parallel-inheritance-hierarchy.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `ParallelInheritanceHierarchy.java`

**Tarea**: Añade un nuevo componente (Checkbox) y un nuevo renderizador (PlainText).

## Ejecutar tests

```bash
mvn test -Dtest=ParallelInheritanceHierarchyTest
```

## Problema a experimentar

Cada nuevo componente requiere añadir métodos a todos los renderizadores, y cada nuevo renderizador debe implementar métodos para todos los componentes existentes.
