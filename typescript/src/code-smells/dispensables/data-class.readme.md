# Data Class - Ejercicio en TypeScript

## 📚 Documentación Completa

👉 **[Ver documentación completa de Data Class](../../../../docs/code-smells/dispensables/data-class.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `data-class.ts`

**Tarea**: Implementa reglas de dominio adicionales, como requerir verificación de email o restringir el registro a ciertos dominios (ej. `company.com`).

## Ejecutar tests

```bash
npm test -- data-class.test.ts
```

## Problema a experimentar

Tendrás que modificar múltiples servicios y lugares que manipulan `UserRecord`. Esto demuestra cómo separar el comportamiento de los datos provoca que cambios simples se dispersen ampliamente por el código (Shotgun Surgery).
