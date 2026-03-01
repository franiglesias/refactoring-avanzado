# Alternative Classes with Different Interfaces - Ejercicio en Java

## 📚 Documentación Completa

👉 **[Ver documentación completa de Alternative Classes with Different Interfaces](../../../../../../../docs/code-smells/oop-abusers/alternative-classes-different-interfaces.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `AlternativeClassesDifferentInterfaces.java`

**Tarea**: Añade soporte para niveles de log (INFO, WARNING, ERROR) en ambas clases.

## Ejecutar tests

```bash
mvn test -Dtest=AlternativeClassesDifferentInterfacesTest
```

## Problema a experimentar

Tendrás que duplicar la lógica de niveles en ambas clases y en todo el código que las use, en lugar de tener una interfaz común que permita el polimorfismo.
