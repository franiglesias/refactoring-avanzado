# Duplicated Code - Ejercicio en Java

## 📚 Documentación Completa

👉 **[Ver documentación completa de Duplicated Code](../../../../../../../docs/code-smells/dispensables/duplicated-code.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `DuplicatedCode.java`

**Tarea**: Añade validación adicional de dominio de email y soporte para CC/BCC en todos los métodos de envío.

## Ejecutar tests

```bash
mvn test -Dtest=DuplicatedCodeTest
```

## Problema a experimentar

Tendrás que replicar los mismos cambios en múltiples lugares, con alto riesgo de olvidar actualizar alguno y generar comportamiento inconsistente.
