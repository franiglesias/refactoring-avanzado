# Comments - Ejercicio en Go

## 📚 Documentación Completa

👉 **[Ver documentación completa de Comments](../../../docs/code-smells/dispensables/comments.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `comments.go`

**Tarea**: Actualiza la función `Add` para registrar (log) cuando la suma sea negativa.

## Ejecutar tests

```bash
go test ./code_smells/dispensables/comments_test.go
```

## Problema a experimentar

Observa cómo los comentarios de alrededor se vuelven obsoletos o engañosos rápidamente al realizar cambios, obligándote a editar muchas líneas de comentario por un cambio diminuto en el código, aumentando el riesgo de desalineación entre lo que el código hace y lo que el comentario dice.
