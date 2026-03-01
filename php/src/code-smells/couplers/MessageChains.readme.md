# Message Chains - Ejercicio en PHP

## 📚 Documentación Completa

👉 **[Ver documentación completa de Message Chains](../../../../docs/code-smells/couplers/message-chains.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `Root.php`, `Level1.php` y `Level2.php`

**Tarea**: Inserta un nuevo `Level` entre `Root` y `Level1`, o reubica `getValue`.

## Ejecutar tests

```bash
./vendor/bin/phpunit tests/CodeSmells/Couplers/MessageChainsTest.php
```

## Problema a experimentar

Observa cómo cada cliente que usa `$root->getNext()->getNext()->getValue()` debe cambiar, revelando cómo las cadenas de mensajes vuelven costosas refactorizaciones simples.
