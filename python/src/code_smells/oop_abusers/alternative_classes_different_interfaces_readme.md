# Alternative Classes with Different Interfaces - Ejercicio en Python

## 📚 Documentación Completa

👉 **[Ver documentación completa de Alternative Classes with Different Interfaces](../../../../docs/code-smells/oop-abusers/alternative-classes-different-interfaces.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `alternative_classes_different_interfaces.py`

**Tarea**: Añade logging con marca de tiempo a ambas implementaciones y permite que el cliente pueda intercambiarlas en tiempo de ejecución sin usar condicionales.

## Ejecutar tests

```bash
pytest src/code_smells/oop_abusers/test_alternative_classes_different_interfaces.py
```

## Problema a experimentar

Al no compartir una interfaz común, te verás obligado a duplicar lógica en métodos con nombres distintos y a esparcir sentencias `if/else` o `switch` en los clientes, haciendo que cambios simples se vuelvan tediosos y propensos a errores.
