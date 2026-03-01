# Long Method - Ejercicio en Go

## 📚 Documentación Completa

👉 **[Ver documentación completa de Long Method](../../../docs/code-smells/bloaters/long-method.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `long_method.go`

**Tarea**: Añade soporte de cupones con expiración y multi-moneda (USD/EUR) con reglas de redondeo distintas.

## Ejecutar tests

```bash
go test ./code_smells/bloaters/long_method_test.go
```

## Problema a experimentar

Tienes que tocar diferentes secciones dentro del método, lo que genera riesgo de cambios indeseados
y aumenta el esfuerzo de mantenimiento.
