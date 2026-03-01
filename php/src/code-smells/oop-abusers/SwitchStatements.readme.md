# Switch Statements - Ejercicio en PHP

## 📚 Documentación Completa

👉 **[Ver documentación completa de Switch Statements](../../../../docs/code-smells/oop-abusers/switch-statements.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `EmployeeRecord.php` y `functions.php`

**Tarea**: Añade un nuevo tipo de empleado (`contractor`) con una regla de pago especial (ej. tarifa por horas).

## Ejecutar tests

```bash
./vendor/bin/phpunit tests/CodeSmells/OopAbusers/SwitchStatementsTest.php
```

## Problema a experimentar

Tendrás que modificar el `switch` (o `match`) y cualquier otro código que dependa del tipo de empleado. A medida que el sistema crece, olvidar actualizar uno de estos puntos genera bugs difíciles de detectar.
