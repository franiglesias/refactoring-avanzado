# Inappropriate Intimacy - Ejercicio en Python

## 📚 Documentación Completa

👉 **[Ver documentación completa de Inappropriate Intimacy](../../../../docs/code-smells/couplers/inappropriate-intimacy.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `inappropriate_intimacy.py`

**Tarea**: Añade una traza de auditoría cuando cambien los presupuestos y aplica reglas de presupuesto mínimo.

## Ejecutar tests

```bash
pytest src/code_smells/couplers/test_inappropriate_intimacy.py
```

## Problema a experimentar

Como Team y Manager tocan libremente los campos del otro, tendrás que esparcir comprobaciones y registros en muchos lugares, aumentando el acoplamiento y las regresiones.
