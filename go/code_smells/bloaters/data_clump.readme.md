# Data Clump - Ejercicio en Go

## 📚 Documentación Completa

👉 **[Ver documentación completa de Data Clump](../../../docs/code-smells/bloaters/data-clump.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `data_clump.go`

**Tarea**: Añade país y provincia y reglas de formateo internacional de la dirección.

## Ejecutar tests

```bash
go test ./code_smells/bloaters/data_clump_test.go
```

## Problema a experimentar

Necesitarás modificar constructores, impresores y cualquier lugar que pase estos campos juntos, multiplicando la superficie de cambio.
