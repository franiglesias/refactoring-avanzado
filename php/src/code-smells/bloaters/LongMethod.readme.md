# Long Method - Ejercicio en PHP

## 📚 Documentación Completa

👉 **[Ver documentación completa de Long Method](../../../../docs/code-smells/bloaters/long-method.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `OrderService.php`

**Tarea**: Añade soporte de cupones con expiración y multi-moneda (USD/EUR) con reglas de redondeo distintas.

## Ejecutar tests

```bash
./vendor/bin/phpunit tests/CodeSmells/Bloaters/LongMethodTest.php
```

## Problema a experimentar

Tienes que tocar diferentes secciones dentro del método, lo que genera riesgo de cambios indeseados y aumenta el esfuerzo de mantenimiento.
