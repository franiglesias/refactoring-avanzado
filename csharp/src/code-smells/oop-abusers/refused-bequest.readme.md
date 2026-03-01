# Refused Bequest - Ejercicio en C#

## 📚 Documentación Completa

👉 **[Ver documentación completa de Refused Bequest](../../../../docs/code-smells/oop-abusers/refused-bequest.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `RefusedBequest.cs`

**Tarea**: Añade un método de ciclo de vida `pause` a la interfaz `Controller` y haz que `start` y `stop` sean obligatorios con lógica real.

## Ejecutar tests

```bash
dotnet test --filter "RefusedBequest"
```

## Problema a experimentar

`ReadOnlyController` se verá forzado a implementar métodos que no tienen sentido para su propósito, lo que te obligará a lanzar excepciones o dejar implementaciones vacías que violan el Principio de Sustitución de Liskov.
