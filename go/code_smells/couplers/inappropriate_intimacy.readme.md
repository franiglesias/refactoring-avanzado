# Inappropriate Intimacy - Ejercicio en Go

## 📚 Documentación Completa

👉 **[Ver documentación completa de Inappropriate Intimacy](../../../docs/code-smells/couplers/inappropriate-intimacy.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `inappropriate_intimacy.go`

**Tarea**: Añade una traza de auditoría cuando cambien los presupuestos y aplica reglas de presupuesto mínimo.

## Ejecutar tests

```bash
go test ./code_smells/couplers/inappropriate_intimacy_test.go
```

## Problema a experimentar

Como Team y Manager tocan libremente los campos del otro, tendrás que esparcir comprobaciones y registros en muchos lugares, aumentando el acoplamiento y las regresiones.
