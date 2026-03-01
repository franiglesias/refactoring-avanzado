# Long Parameter List - Ejercicio en Java

## 📚 Documentación Completa

👉 **[Ver documentación completa de Long Parameter List](../../../../../../../docs/code-smells/bloaters/long-parameter-list.md)**

La documentación completa incluye:
- Definición y descripción detallada
- Síntomas para identificarlo
- Ejemplo en pseudocódigo
- Proceso de refactoring paso a paso
- Técnicas aplicables
- Referencias en español e inglés

## 🎯 Ejercicio

**Archivo**: `LongParameterList.java`

**Tarea**: Añade más campos relacionados con el usuario (fecha de nacimiento, número de identificación fiscal, preferencias de idioma).

## Ejecutar tests

```bash
mvn test -Dtest=LongParameterListTest
```

## Problema a experimentar

Cada nuevo parámetro alarga aún más la firma del método, haciendo el código cada vez más difícil de leer y mantener. Los cambios en el orden de los parámetros pueden generar bugs sutiles.
