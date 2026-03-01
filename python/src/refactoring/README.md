# Refactoring Avanzado - Ejercicios en Python

## 📚 Documentación Completa

👉 **[Ver documentación completa de Técnicas de Refactoring](../../../../docs/refactoring/README.md)**

La documentación completa incluye explicaciones detalladas, ejemplos de código y referencias para cada técnica.

## 🎯 Técnicas Disponibles

### Golden Master
- **[Documentación](../../../../docs/refactoring/golden-master/README.md)** - Tests de caracterización para código legacy
- **Implementación**: `golden_master.py` y `test_golden_master.py`

### Parallel Change
- **[Documentación general](../../../../docs/refactoring/parallel-change/README.md)** - Visión general de cambios en paralelo
- **[Detalles](parallel_change/README.md)** - Ejercicios en Python

## Ejecutar tests

```bash
# Todos los tests de refactoring
pytest src/refactoring/ -v

# Golden Master
pytest src/refactoring/test_golden_master.py -v

# Parallel Change
pytest src/refactoring/parallel_change/ -v
```
