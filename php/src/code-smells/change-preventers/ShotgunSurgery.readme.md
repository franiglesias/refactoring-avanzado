# Shotgun Surgery - Ejercicio en PHP

## 📚 Documentación Completa

👉 **[Ver documentación completa de Shotgun Surgery](../../../../docs/code-smells/change-preventers/shotgun-surgery.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `PriceCalculator.php`, `InvoiceService.php`, `SalesReport.php`, `LoyaltyPoints.php`

**Tarea**: Cambia el impuesto del 21% al 18.5% con redondeo a 2 decimales.

## Ejecutar tests

```bash
./vendor/bin/phpunit tests/CodeSmells/ChangePreventers/ShotgunSurgeryTest.php
```

## Problema a experimentar

Tendrás que buscar cada copia y asegurar un redondeo consistente en todas partes, destacando cómo la duplicación convierte un cambio pequeño en muchas ediciones arriesgadas.
