# Inappropriate Intimacy - Ejercicio en Java

## 📚 Documentación Completa

👉 **[Ver documentación completa de Inappropriate Intimacy](../../../../../../../docs/code-smells/couplers/inappropriate-intimacy.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `InappropriateIntimacy.java`

**Tarea**: Añade gestión de miembros del equipo y límites de presupuesto por categoría.

## Ejecutar tests

```bash
mvn test -Dtest=InappropriateIntimacyTest
```

## Problema a experimentar

El acoplamiento entre `Manager` y `Team` te obligará a modificar ambas clases simultáneamente, aumentando la complejidad y el riesgo de errores.
