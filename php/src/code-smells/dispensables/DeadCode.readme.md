# Dead Code - Ejercicio en PHP

## 📚 Documentación Completa

👉 **[Ver documentación completa de Dead Code](../../../../docs/code-smells/dispensables/dead-code.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `functions.php`

**Tarea**: Arregla un bug en `activeFunction` (por ejemplo, cambia el manejo de valores negativos).

## Ejecutar tests

```bash
./vendor/bin/phpunit tests/CodeSmells/Dispensables/DeadCodeTest.php
```

## Problema a experimentar

Observa cómo el código muerto cercano dificulta razonar sobre lo que realmente se ejecuta, lo que puede invitar a errores o a olvidar la limpieza necesaria durante el proceso de refactorización.
