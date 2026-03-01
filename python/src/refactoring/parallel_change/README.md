# Parallel Change - Ejercicios en Python

## 📚 Documentación Completa

👉 **[Ver documentación completa de Parallel Change](../../../../../docs/refactoring/parallel-change/README.md)**

## 🎯 Técnicas

- **[Sprout Change](sprout_change_readme.md)** - Hacer brotar nuevo código
- **[Wrap Change](wrap_change_readme.md)** - Envolver dependencias
- **[Expand-Migrate-Contract](expand_migrate_contract_readme.md)** - Cambios estructurales en 3 fases

## Ejecutar tests

```bash
pytest src/refactoring/parallel_change/ -v
```

## Patrón TCR (Test && Commit || Revert)

```bash
pytest src/refactoring/parallel_change/ && git commit -am "refactor: mensaje" || git reset --hard
```

Si los tests pasan → commit automático. Si fallan → revertir cambios.
