# Feature Envy - Ejercicio en Go

## 📚 Documentación Completa

👉 **[Ver documentación completa de Feature Envy](../../../docs/code-smells/couplers/feature-envy.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `feature_envy.go`

**Tarea**: Añade validación de formato de email y validación de dirección completa (calle, ciudad, código postal).

## Ejecutar tests

```bash
go test ./code_smells/couplers/feature_envy_test.go
```

## Problema a experimentar

Probablemente, seguirás añadiendo condiciones dentro de `InvoiceService` que dependen de detalles internos de `Customer`, esparciendo reglas en el lugar equivocado y volviendo frágiles los cambios.
