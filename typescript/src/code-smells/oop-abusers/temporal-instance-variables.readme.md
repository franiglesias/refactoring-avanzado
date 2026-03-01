# Temporal Instance Variables - Ejercicio en TypeScript

## 📚 Documentación Completa

👉 **[Ver documentación completa de Temporal Instance Variables](../../../../docs/code-smells/oop-abusers/temporal-instance-variables.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `temporal-instance-variables.ts`

**Tarea**: Añade una validación para que no se pueda llamar a `place()` si no se ha añadido al menos un ingrediente.

## Ejecutar tests

```bash
npm test -- temporal-instance-variables.test.ts
```

## Problema a experimentar

Te darás cuenta de que el objeto es una "máquina de estados" frágil. Si un cliente olvida llamar a `start()` o intenta llamar a `addTopping()` fuera de orden, el sistema puede fallar silenciosamente o requerir comprobaciones constantes de nulidad en cada método.
