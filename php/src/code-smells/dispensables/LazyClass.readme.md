# Lazy Class - Ejercicio en PHP

## 📚 Documentación Completa

👉 **[Ver documentación completa de Lazy Class](../../../../docs/code-smells/dispensables/lazy-class.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `ShippingLabelBuilder.php` y `Address.php`

**Tarea**: Reescribe el código para eliminar la necesidad de la clase `ShippingLabelBuilder`.

## Ejecutar tests

```bash
./vendor/bin/phpunit tests/CodeSmells/Dispensables/LazyClassTest.php
```

## Problema a experimentar

Mantener una estructura de clase para una lógica tan simple te obliga a instanciar objetos innecesariamente y añade capas de abstracción que dificultan la legibilidad del código sin ofrecer beneficios a cambio.
