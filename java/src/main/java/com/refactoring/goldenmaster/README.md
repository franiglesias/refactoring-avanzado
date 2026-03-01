# Golden Master - Ejercicio en Java

## 📚 Documentación Completa

👉 **[Ver documentación completa de Golden Master](../../../../../../docs/refactoring/golden-master/README.md)**

La documentación completa incluye:
- Definición y explicación detallada de la técnica
- Proceso paso a paso con ejemplos de código
- Problemas comunes y soluciones
- Referencias en español e inglés

## 🎯 Ejercicio

**Objetivo**: Crear una prueba de Golden Master para el código legado (`ReceiptPrinter`)

**Archivo a refactorizar**: `ReceiptPrinter.java` en `com.refactoring.refactoring`

## Contexto

La clase `ReceiptPrinter` tiene:
- Cálculos de totales con impuestos variables
- Descuentos aleatorios ("descuento de la suerte")
- Formato de recibos con fecha actual
- Comportamiento no determinista (Random, Date)

## Pasos recomendados

1. Identificar fuentes de no determinismo (Random, Date)
2. Usar costuras ya preparadas (constructor con Random, métodos protegidos)
3. Extender `ReceiptPrinter` para controlar comportamiento
4. Generar conjunto amplio de entradas con `OrderGenerator`
5. Capturar salida maestra (Approval Tests o snapshots)
6. Refactorizar con seguridad
7. Introducir tests unitarios para reemplazar Golden Master

## Ejecución

```bash
javac -d out src/main/java/com/refactoring/**/*.java
java -cp out com.refactoring.goldenmaster.GoldenMaster
```

## Criterios de aceptación

- ✅ Prueba captura múltiples casos de entrada
- ✅ Falla ante cambios en la salida
- ✅ Fuentes de no determinismo controladas mediante Seams
- ✅ Comportamiento original disponible (compatibilidad)
