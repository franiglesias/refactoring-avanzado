# Inappropriate Intimacy - Ejercicio en PHP

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

**Archivo**: `Team.php`, `Manager.php` y `Budget.php`

**Tarea**: Añade una traza de auditoría cuando cambien los presupuestos y aplica reglas de presupuesto mínimo.

## Ejecutar tests

```bash
./vendor/bin/phpunit tests/CodeSmells/Couplers/InappropriateIntimacyTest.php
```

## Problema a experimentar

Como Team y Manager tocan libremente los campos del otro, tendrás que esparcir comprobaciones y registros en muchos lugares, aumentando el acoplamiento y las regresiones.
