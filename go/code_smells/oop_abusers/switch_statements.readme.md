# Switch Statements - Ejercicio en Go

## 📚 Documentación Completa

👉 **[Ver documentación completa de Switch Statements](../../../docs/code-smells/oop-abusers/switch-statements.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `switch_statements.go`

**Tarea**: Añade un nuevo tipo de método de pago (`cryptocurrency`) con reglas de procesamiento, comisión y nombre especiales.

## Ejecutar tests

```bash
go test ./code_smells/oop_abusers/switch_statements_test.go
```

## Problema a experimentar

Tendrás que modificar todos los `switch` dispersos por el código y cualquier otro lugar que dependa del tipo de método de pago. A medida que el sistema crece, olvidar actualizar uno de estos puntos genera bugs difíciles de detectar.
