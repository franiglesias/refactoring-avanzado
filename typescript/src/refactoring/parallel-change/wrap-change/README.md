# Wrap Change - Ejercicio en TypeScript

## 📚 Documentación Completa

👉 **[Ver documentación completa de Wrap Change](../../../../../docs/refactoring/parallel-change/wrap-change.md)**

La documentación completa incluye:
- Definición y explicación detallada de la técnica
- Proceso paso a paso con ejemplos de código
- Problemas comunes y soluciones
- Referencias en español e inglés

## 🎯 Ejercicio

**Objetivo**: Envolver `LegacyEmailService` para añadir funcionalidad sin cambiar su interfaz

**Archivo a refactorizar**: `wrap-change.ts`

## Ejecutar tests

```shell
npm run test -- src/refactoring/parallel-change/wrap-change/wrap-change.test.ts
```

## Escenario

Tenemos `LegacyEmailService` (dependencia externa que NO podemos modificar) que se usa directamente. El servicio es limitado: no valida emails, no tiene retry, no tiene logging.

Queremos mejorar la funcionalidad SIN cambiar todas las llamadas existentes.

## Pasos recomendados

1. Crear wrapper con la misma interfaz `sendEmail(to, subject, body)`
2. Instanciar wrapper y reemplazarlo en el código
3. Añadir validación dentro del wrapper (con tests)
4. Añadir logging
5. Migrar puntos de uso uno por uno
6. Añadir más funcionalidad según necesites (retry, sanitización, plantillas)

## Criterios de aceptación

- ✅ Interfaz pública NO cambia
- ✅ Clientes no necesitan modificarse
- ✅ Wrapper añade funcionalidad (validación, logging, etc.)
- ✅ Servicio legacy usado internamente
- ✅ Migración punto por punto sin romper nada
