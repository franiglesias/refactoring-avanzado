# Message Chains - Ejercicio en C#

## 📚 Documentación Completa

👉 **[Ver documentación completa de Message Chains](../../../../docs/code-smells/couplers/message-chains.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `MessageChains.cs`

**Tarea**: Inserta un nuevo `Level` entre `Root` y `Level1`, o reubica `getValue`.

## Ejecutar tests

```bash
dotnet test --filter "MessageChains"
```

## Problema a experimentar

Observa cómo cada cliente que usa root.getNext().getNext().getValue() debe cambiar, revelando cómo las cadenas de mensajes vuelven costosas refactorizaciones simples.
