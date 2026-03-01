# Switch Statements - Ejercicio en Java

## 📚 Documentación Completa

👉 **[Ver documentación completa de Switch Statements](../../../../../../../docs/code-smells/oop-abusers/switch-statements.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `SwitchStatements.java`

**Tarea**: Añade un nuevo método de pago (CRYPTOCURRENCY) con reglas especiales de procesamiento y comisiones.

## Ejecutar tests

```bash
mvn test -Dtest=SwitchStatementsTest
```

## Problema a experimentar

Tendrás que modificar múltiples `switch` statements en diferentes métodos. A medida que el sistema crece, olvidar actualizar uno de estos puntos genera bugs difíciles de detectar.
