# Middleman - Ejercicio en Java

## 📚 Documentación Completa

👉 **[Ver documentación completa de Middleman](../../../../../../../docs/code-smells/couplers/middleman.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `Middleman.java`

**Tarea**: Añade funcionalidad de búsqueda por nombre y filtrado por categoría.

## Ejecutar tests

```bash
mvn test -Dtest=MiddlemanTest
```

## Problema a experimentar

Cada nueva funcionalidad en `Catalog` requerirá añadir un método de delegación correspondiente en `Shop`, sin aportar valor alguno.
