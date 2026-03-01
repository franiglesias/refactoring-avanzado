# Shotgun Surgery - Ejercicio en C#

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

**Archivo**: `ShotgunSurgery.cs`

**Tarea**: Cambia el impuesto del 21% al 18.5% con redondeo a 2 decimales.

## Ejecutar tests

```bash
dotnet test --filter "ShotgunSurgery"
```

## Problema a experimentar

Tendrás que buscar cada copia y asegurar un redondeo consistente en todas partes, destacando cómo la duplicación convierte un cambio pequeño en muchas ediciones arriesgadas.
