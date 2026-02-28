# Ejercicios de Refactoring

Este directorio contiene ejercicios prácticos de técnicas de refactoring para trabajar con código legado y realizar cambios seguros.

## Ejercicios

### 1. Golden Master

Crear una prueba de Golden Master para código legado sin tests.

**Objetivo:** Caracterizar el comportamiento actual y establecer una red de seguridad para refactorizar.

**Archivos:**
- `golden_master.py` - Código legado (`ReceiptPrinter`)
- `test_golden_master.py` - Tests

**Ejecutar tests:**
```shell
pytest src/refactoring/test_golden_master.py -v
```

#### Pasos recomendados (haz commits entre pasos)

1. Identificar fuentes de no determinismo que rompen cualquier intento de escribir tests.
2. Introducir costuras (SEAMS) mínimas sin cambiar el comportamiento.
   1. Aisla las fuentes no deterministas en métodos protegidos (seams).
   2. Extiende la clase y sobreescribe las costuras para controlar su comportamiento.
3. Generar un conjunto amplio y estable de entradas. Puedes usar esta forma para que el test se ejecute con un valor nuevo cada vez:

```python
@pytest.mark.parametrize('customer', ['Ana', 'Luis', 'Mar', 'Iván', 'Sofía'])
def test_with_different_customers(customer, snapshot):
    # Aquí el test
    pass
```

4. Capturar la salida maestra.
5. Escribir la prueba usando como SUT la clase derivada, mediante `assert result == snapshot`.
6. Refactorizar con seguridad para invertir las dependencias.
7. Introducir tests unitarios/integración para reemplazar el Golden Master y poder introducir nueva funcionalidad.

#### Criterios de aceptación

- La prueba de Golden Master captura múltiples casos de entrada y falla ante cambios en la salida.
- Las fuentes de no determinismo están controladas mediante Seams o, finalmente, inversión de dependencias.
- El comportamiento original de `ReceiptPrinter` sigue disponible para el resto del código (compatibilidad), incluyendo el uso de comportamientos no deterministas.

### 2. Parallel Change

Técnicas para realizar cambios significativos de forma gradual y segura.

**Objetivo:** Aprender a refactorizar interfaces y comportamientos sin romper el código existente.

Ver [parallel_change/README.md](parallel_change/README.md) para más detalles.

**Ejercicios incluidos:**
- **Expand-Migrate-Contract**: Cambiar estructura de datos gradualmente
- **Sprout Change**: Introducir nueva funcionalidad sin modificar código existente
- **Wrap Change**: Envolver dependencias problemáticas añadiendo funcionalidad

**Ejecutar tests:**
```shell
pytest src/refactoring/parallel_change/ -v
```

## Técnicas de Refactoring

### Golden Master Testing

Técnica para caracterizar comportamiento de código legado sin tests:
1. Capturar salidas actuales como "golden" (referencia)
2. Ejecutar tests comparando con golden
3. Refactorizar con confianza
4. Reemplazar con tests unitarios

### Parallel Change (Expand-Contract)

Técnica para cambios incrementales en APIs:
1. **Expand**: Añadir nueva funcionalidad (convive con la antigua)
2. **Migrate**: Migrar consumidores gradualmente
3. **Contract**: Eliminar funcionalidad antigua

### Seams

Puntos en el código donde puedes inyectar comportamiento diferente sin modificar el código original. Útiles para testing de código legado.

## Referencias

- [Working Effectively with Legacy Code - Michael Feathers](https://www.amazon.com/Working-Effectively-Legacy-Michael-Feathers/dp/0131177052)
- [Refactoring - Martin Fowler](https://refactoring.com/)
- [Golden Master Testing](https://softwareengineering.stackexchange.com/questions/358589/what-is-golden-master-testing)
