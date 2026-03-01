# Switch Statements - Ejercicio en TypeScript

## 📚 Documentación Completa

👉 **[Ver documentación completa de Switch Statements](../../../../docs/code-smells/oop-abusers/switch-statements.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `switch-statements.ts`

**Tarea**: Añade un nuevo tipo de empleado (`contractor`) con una regla de pago especial (ej. tarifa por horas).

## Ejecutar tests

```bash
npm test -- switch-statements.test.ts
```

## Problema a experimentar

Tendrás que modificar el `switch` y cualquier otro código que dependa del tipo de empleado. A medida que el sistema crece, olvidar actualizar uno de estos puntos genera bugs difíciles de detectar.
