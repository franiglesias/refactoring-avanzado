# Temporal Instance Variables - Ejercicio en Java

## 📚 Documentación Completa

👉 **[Ver documentación completa de Temporal Instance Variables](../../../../../../../docs/code-smells/oop-abusers/temporal-instance-variables.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `TemporalInstanceVariables.java`

**Tarea**: Añade validación de estado para prevenir el uso incorrecto (ej. llamar a `place()` sin haber llamado a `start()` primero).

## Ejecutar tests

```bash
mvn test -Dtest=TemporalInstanceVariablesTest
```

## Problema a experimentar

El estado temporal hace que el objeto sea frágil y propenso a errores. Es difícil razonar sobre el estado válido del objeto en diferentes puntos del flujo.
