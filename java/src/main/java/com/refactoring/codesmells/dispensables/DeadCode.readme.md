# Dead Code - Ejercicio en Java

## 📚 Documentación Completa

👉 **[Ver documentación completa de Dead Code](../../../../../../../docs/code-smells/dispensables/dead-code.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `DeadCode.java`

**Tarea**: Identifica y elimina todo el código muerto. Añade un nuevo método que use la constante previamente no utilizada.

## Ejecutar tests

```bash
mvn test -Dtest=DeadCodeTest
```

## Problema a experimentar

Es fácil que el código muerto se acumule con el tiempo, especialmente durante refactorizaciones. El código no utilizado confunde a los desarrolladores y hace más difícil encontrar el código realmente relevante.
