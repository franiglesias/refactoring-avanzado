# Primitive Obsession - Ejercicio en Java

## 📚 Documentación Completa

👉 **[Ver documentación completa de Primitive Obsession](../../../../../../../docs/code-smells/bloaters/primitive-obsession.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `PrimitiveObsession.java`

**Tarea**: Añade soporte para múltiples monedas con conversión y validación de códigos de moneda ISO.

## Ejecutar tests

```bash
mvn test -Dtest=PrimitiveObsessionTest
```

## Problema a experimentar

La validación y las reglas de negocio seguirán esparcidas por múltiples lugares, haciendo el código frágil y difícil de mantener.
