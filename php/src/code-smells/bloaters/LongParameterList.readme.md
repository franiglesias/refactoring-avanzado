# Long Parameter List - Ejercicio en PHP

## 📚 Documentación Completa

👉 **[Ver documentación completa de Long Parameter List](../../../../docs/code-smells/bloaters/long-parameter-list.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `ReportGenerator.php`

**Tarea**: Añade dos opciones más (p. ej., locale y pageSize) al reporte.

## Ejecutar tests

```bash
./vendor/bin/phpunit tests/CodeSmells/Bloaters/LongParameterListTest.php
```

## Problema a experimentar

Con más de tres parámetros es difícil recordar con exactitud cuáles son, el orden o el tipo de cada uno. Añadir parámetros no hace más que aumentar la dificultad de uso y mantenimiento.
