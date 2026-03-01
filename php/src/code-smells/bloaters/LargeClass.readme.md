# Large Class - Ejercicio en PHP

## 📚 Documentación Completa

👉 **[Ver documentación completa de Large Class](../../../../docs/code-smells/bloaters/large-class.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `UserAccount.php`

**Tarea**: Añade soporte para autenticación de dos factores (2FA) y preferencias de notificación.

## Ejecutar tests

```bash
./vendor/bin/phpunit tests/CodeSmells/Bloaters/LargeClassTest.php
```

## Problema a experimentar

Tocarás autenticación, estado y notificaciones en una clase inflada, aumentando la probabilidad de romper comportamiento no relacionado.
