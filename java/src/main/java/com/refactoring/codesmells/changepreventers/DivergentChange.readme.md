# Divergent Change - Ejercicio en Java

## 📚 Documentación Completa

👉 **[Ver documentación completa de Divergent Change](../../../../../../../docs/code-smells/change-preventers/divergent-change.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `DivergentChange.java`

**Tarea**: Añade soporte para exportar a XML y enviar notificaciones SMS.

## Ejecutar tests

```bash
mvn test -Dtest=DivergentChangeTest
```

## Problema a experimentar

Cada nuevo formato de exportación o canal de comunicación requiere modificar esta clase, acumulando responsabilidades no relacionadas.
