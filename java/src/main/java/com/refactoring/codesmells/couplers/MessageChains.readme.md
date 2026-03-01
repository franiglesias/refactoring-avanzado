# Message Chains - Ejercicio en Java

## 📚 Documentación Completa

👉 **[Ver documentación completa de Message Chains](../../../../../../../docs/code-smells/couplers/message-chains.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `MessageChains.java`

**Tarea**: Añade un nivel adicional (Level3) y métodos para modificar valores en diferentes niveles.

## Ejecutar tests

```bash
mvn test -Dtest=MessageChainsTest
```

## Problema a experimentar

Cualquier cambio en la estructura intermedia (añadir, remover o reordenar niveles) rompe todos los lugares que navegan por la cadena.
