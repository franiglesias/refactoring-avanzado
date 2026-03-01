# Duplicated Code - Ejercicio en Go

## 📚 Documentación Completa

👉 **[Ver documentación completa de Duplicated Code](../../../docs/code-smells/dispensables/duplicated-code.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `duplicated_code.go`

**Tarea**: Añade un nuevo tipo de email (p. ej., notificación de promoción) con las mismas validaciones y lógica de envío.

## Ejecutar tests

```bash
go test ./code_smells/dispensables/duplicated_code_test.go
```

## Problema a experimentar

Tendrás que duplicar nuevamente toda la lógica de validación y envío, lo que demuestra cómo la duplicación multiplica el esfuerzo y el riesgo de error humano.
