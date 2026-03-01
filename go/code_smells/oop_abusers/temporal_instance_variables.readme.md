# Temporal Instance Variables - Ejercicio en Go

## 📚 Documentación Completa

👉 **[Ver documentación completa de Temporal Instance Variables](../../../docs/code-smells/oop-abusers/temporal-instance-variables.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `temporal_instance_variables.go`

**Tarea**: Añade una validación para que no se pueda llamar a `Place()` si no se ha añadido al menos un ingrediente.

## Ejecutar tests

```bash
go test ./code_smells/oop_abusers/temporal_instance_variables_test.go
```

## Problema a experimentar

Te darás cuenta de que el objeto es una "máquina de estados" frágil. Si un cliente olvida llamar a `Start()` o intenta llamar a `AddTopping()` fuera de orden, el sistema puede fallar silenciosamente o requerir comprobaciones constantes de nulidad en cada método.
