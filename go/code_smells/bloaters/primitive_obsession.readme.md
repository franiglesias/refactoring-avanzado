# Primitive Obsession - Ejercicio en Go

## 📚 Documentación Completa

👉 **[Ver documentación completa de Primitive Obsession](../../../docs/code-smells/bloaters/primitive-obsession.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `primitive_obsession.go`

**Tarea**: Introduce soporte para diferentes monedas, para enviar la factura por email, y para formatear la dirección en función del país.

## Ejecutar tests

```bash
go test ./code_smells/bloaters/primitive_obsession_test.go
```

## Problema a experimentar

Dado que los primitivos no nos permiten garantizar la integridad de sus valores, tendrás que introducir validaciones en muchos lugares, incluso de forma repetida. Algunos datos siempre viajan juntos (Data Clump), por lo que tienes que asegurarte de que permanecen juntos.

Para formatear de forma diferente basándote en algún dato arbitrario tendrás que introducir lógica de decisión.
