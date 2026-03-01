# Refused Bequest - Ejercicio en Go

## 📚 Documentación Completa

👉 **[Ver documentación completa de Refused Bequest](../../../docs/code-smells/oop-abusers/refused-bequest.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `refused_bequest.go`

**Tarea**: Añade un método de ciclo de vida `Pause` a la interfaz `Controller` y haz que `Start` y `Stop` sean obligatorios con lógica real.

## Ejecutar tests

```bash
go test ./code_smells/oop_abusers/refused_bequest_test.go
```

## Problema a experimentar

`ReadOnlyController` se verá forzado a implementar métodos que no tienen sentido para su propósito, lo que te obligará a lanzar excepciones o dejar implementaciones vacías que violan el Principio de Sustitución de Liskov.
