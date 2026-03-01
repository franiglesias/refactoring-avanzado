# Long Parameter List

Lista larga de parámetros.

## Definición

Una función recibe más de tres o cuatro parámetros, haciendo difícil recordar el orden, el tipo y el propósito de cada uno.

## Descripción

Un **Long Parameter List** ocurre cuando un método o función acepta demasiados parámetros. Aunque algunos lenguajes modernos mitigan este problema con parámetros nombrados, las listas largas de parámetros siguen siendo problemáticas porque:
- Son difíciles de recordar
- Son propensas a errores de orden
- Dificultan añadir nuevos parámetros
- Sugieren que el método hace demasiado
- Hacen el código más rígido y difícil de cambiar

Este smell a menudo aparece cuando se intenta evitar dependencias entre objetos pasando todos los datos necesarios como parámetros individuales. Irónicamente, esto crea un acoplamiento más fuerte porque los clientes deben conocer exactamente qué datos necesita el método.

## Síntomas

- Métodos con más de 3-4 parámetros
- Parámetros que siempre se pasan juntos (Data Clump)
- Parámetros booleanos que controlan el flujo del método
- Dificultad para recordar el orden de los parámetros
- Necesidad de consultar la firma del método frecuentemente
- Llamadas al método que ocupan múltiples líneas
- Errores frecuentes al pasar argumentos en el orden incorrecto

## Ejemplo

```pseudocode
class ReportGenerator {
  function generateReport(
    title,
    startDate,
    endDate,
    includeCharts,
    includeSummary,
    authorName,
    authorEmail
  ) {
    print "Generando reporte: " + title
    print "Desde " + startDate + " hasta " + endDate
    print "Autor: " + authorName + " (" + authorEmail + ")"

    if (includeCharts) {
      print "Incluyendo gráficos..."
    }

    if (includeSummary) {
      print "Incluyendo resumen..."
    }

    print "Reporte generado exitosamente."
  }
}

// Uso - difícil de entender qué significa cada valor
generator.generateReport(
  "Ventas Q1",
  date("2025-01-01"),
  date("2025-03-31"),
  true,
  false,
  "Pat Smith",
  "pat@example.com"
)
```

## Ejercicio

Añade dos opciones más (por ejemplo, locale y pageSize) al reporte.

## Problemas que encontrarás

Con más de tres parámetros es difícil recordar con exactitud cuáles son, el orden o el tipo de cada uno. Añadir parámetros no hace más que aumentar la dificultad de uso y mantenimiento.

## Proceso de Refactoring

### 1. Identificar parámetros relacionados
- Busca grupos de parámetros que conceptualmente van juntos
- Ejemplo: startDate y endDate representan un "período"
- Ejemplo: authorName y authorEmail representan un "autor"

### 2. Introducir Parameter Objects
- Crea una clase o estructura para cada grupo de parámetros relacionados
- Dale un nombre significativo (DateRange, Author, ReportOptions)
- Reemplaza los parámetros individuales con el objeto

### 3. Preservar objetos completos
- Si ya tienes un objeto con los datos necesarios, pásalo directamente
- En lugar de: `method(user.name, user.email, user.id)`
- Usa: `method(user)`
- El método extrae lo que necesita del objeto

### 4. Usar Builder Pattern para construcciones complejas
- Si el objeto tiene muchas opciones configurables
- Implementa un builder que permita construir el objeto paso a paso
- Proporciona valores por defecto sensatos

### 5. Extraer métodos especializados
- Si diferentes combinaciones de parámetros indican diferentes operaciones
- Crea métodos específicos en lugar de uno general con muchos parámetros
- Ejemplo: en lugar de `generate(type, includeX, includeY)` → `generateSummary()`, `generateDetailed()`

### 6. Considerar métodos de configuración
- Para objetos que necesitan muchos parámetros de inicialización
- Usa el patrón de construcción fluida o métodos setter encadenables
- Permite configurar el objeto paso a paso

## Técnicas de Refactoring Aplicables

- **Introduce Parameter Object**: Agrupar parámetros relacionados en un objeto
- **Preserve Whole Object**: Pasar el objeto completo en lugar de partes
- **Replace Parameter with Method Call**: Si el parámetro puede calcularse internamente
- **Replace Parameter with Explicit Methods**: Crear métodos separados para diferentes casos
- **Builder Pattern**: Para objetos complejos con muchas opciones

## Beneficios

- **Código más legible**: Los objetos con nombres claros explican su propósito
- **Menos errores**: Imposible equivocarse con el orden de los parámetros
- **Más fácil de extender**: Añadir campos al objeto no rompe las llamadas existentes
- **Mejor encapsulación**: Los objetos pueden contener lógica de validación
- **Firmas más estables**: Cambios en parámetros individuales no afectan la firma del método
- **Reutilización**: Los Parameter Objects pueden usarse en múltiples métodos

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/bloaters/long-parameter-list.ts) - [README](../../typescript/src/code-smells/bloaters/long-parameter-list.readme.md)
- [Go](../../go/code_smells/bloaters/long_parameter_list.go) - [README](../../go/code_smells/bloaters/long_parameter_list.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/bloaters/LongParameterList.java) - [README](../../java/src/main/java/com/refactoring/codesmells/bloaters/LongParameterList.readme.md)
- [PHP](../../php/src/code-smells/bloaters/LongParameterList.php) - [README](../../php/src/code-smells/bloaters/LongParameterList.readme.md)
- [Python](../../python/src/code_smells/bloaters/long_parameter_list.py) - [README](../../python/src/code_smells/bloaters/long_parameter_list_readme.md)
- [C#](../../csharp/src/code-smells/bloaters/LongParameterList.cs) - [README](../../csharp/src/code-smells/bloaters/long-parameter-list.readme.md)

## Referencias en Español

- [Parameter (Introduce Parameter Object)](https://franiglesias.github.io/blogtober19-parameter/) - Patrón para reducir listas largas de parámetros agrupándolos en objetos
- [Refactor cotidiano (4). Sustituye escalares por objetos](https://franiglesias.github.io/everyday-refactor-4/) - Guía práctica para reemplazar parámetros primitivos con objetos significativos

## Referencias

- [Refactoring Guru - Long Parameter List](https://refactoring.guru/smells/long-parameter-list)
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
- Joshua Bloch - "Effective Java" - Builder Pattern
