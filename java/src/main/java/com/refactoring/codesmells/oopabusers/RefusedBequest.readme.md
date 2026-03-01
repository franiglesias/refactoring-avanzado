# Refused Bequest - Ejercicio en Java

## 📚 Documentación Completa

👉 **[Ver documentación completa de Refused Bequest](../../../../../../../docs/code-smells/oop-abusers/refused-bequest.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `RefusedBequest.java`

**Tarea**: Añade métodos de configuración (configure, validate) en BaseController que también sean rechazados por ReadOnlyController.

## Ejecutar tests

```bash
mvn test -Dtest=RefusedBequestTest
```

## Problema a experimentar

La herencia se vuelve cada vez más inapropiada, con más métodos vacíos o que lanzan excepciones, violando el Principio de Sustitución de Liskov.
