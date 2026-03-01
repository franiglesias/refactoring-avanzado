# Feature Envy - Ejercicio en Java

## 📚 Documentación Completa

👉 **[Ver documentación completa de Feature Envy](../../../../../../../docs/code-smells/couplers/feature-envy.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `FeatureEnvy.java`

**Tarea**: Añade validación de número de teléfono y formateo especial para clientes VIP.

## Ejecutar tests

```bash
mvn test -Dtest=FeatureEnvyTest
```

## Problema a experimentar

Probablemente, seguirás añadiendo lógica en `InvoiceService` que depende de detalles internos de `Customer`, esparciendo reglas en el lugar equivocado y volviendo frágiles los cambios.
