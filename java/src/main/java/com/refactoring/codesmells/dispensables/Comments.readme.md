# Comments - Ejercicio en Java

## 📚 Documentación Completa

👉 **[Ver documentación completa de Comments](../../../../../../../docs/code-smells/dispensables/comments.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `Comments.java`

**Tarea**: Actualiza la función `add` para registrar (log) cuando la suma sea negativa.

## Ejecutar tests

```bash
mvn test -Dtest=CommentsTest
```

## Problema a experimentar

Observa cómo los comentarios de alrededor se vuelven obsoletos o engañosos rápidamente al realizar cambios, obligándote a editar muchas líneas de comentario por un cambio diminuto en el código, aumentando el riesgo de desalineación entre lo que el código hace y lo que el comentario dice.
