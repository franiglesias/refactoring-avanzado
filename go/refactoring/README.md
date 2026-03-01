# Refactoring Avanzado - Ejercicios en Go

## 📚 Documentación Completa

👉 **[Ver documentación completa de Técnicas de Refactoring](../../../docs/refactoring/README.md)**

La documentación completa incluye explicaciones detalladas, ejemplos de código y referencias para cada técnica.

## 🎯 Técnicas Disponibles

### Golden Master
- **[Documentación](../../../docs/refactoring/golden-master/README.md)** - Tests de caracterización para código legacy
- **Implementación**: `golden_master.go` y `golden_master_test.go`

### Parallel Change
- **[Documentación general](../../../docs/refactoring/parallel-change/README.md)** - Visión general de cambios en paralelo
- **[Sprout Change](parallel_change/sprout_change/README.md)** - Hacer brotar nuevo código
- **[Wrap Change](parallel_change/wrap_change/README.md)** - Envolver dependencias
- **[Expand-Migrate-Contract](parallel_change/expand_migrate_contract/README.md)** - Cambios estructurales en 3 fases

## Ejecutar tests

```bash
go test ./refactoring/...
```
