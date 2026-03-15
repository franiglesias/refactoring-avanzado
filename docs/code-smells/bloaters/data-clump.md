# Data Clump

Grupo de datos.

## Definición

El mismo grupo de campos de datos viaja junto por muchos lugares, lo que sugiere un Value Object faltante y duplicación.

## Descripción

El **Data Clump** ocurre cuando el mismo conjunto de variables aparece repetidamente en múltiples lugares del código:

- Como parámetros en varios métodos
- Como campos en diferentes clases
- Como variables locales con los mismos nombres

Este patrón sugiere que existe un concepto del dominio que no ha sido modelado explícitamente. Los datos viajan juntos porque representan una entidad cohesiva que debería ser un objeto por derecho propio.

## Síntomas

- Varios métodos aceptan los mismos 3-4 parámetros en el mismo orden
- Múltiples clases tienen los mismos campos
- Cambiar la estructura de estos datos requiere modificaciones en muchos lugares para mantener la coherencia
- Los mismos grupos de validaciones se repiten en diferentes ubicaciones

## Ejemplo

El código muestra dos clases (`Invoice` y `ShippingLabel`) que repiten los mismos cuatro campos relacionados con la dirección del cliente:

```pseudocode
class Invoice {
  private customerName: string
  private customerStreet: string
  private customerCity: string
  private customerZip: string

  constructor(customerName, customerStreet, customerCity, customerZip) {
    // inicialización
  }
}

class ShippingLabel {
  private customerName: string
  private customerStreet: string
  private customerCity: string
  private customerZip: string

  constructor(customerName, customerStreet, customerCity, customerZip) {
    // inicialización
  }
}
```

## Ejercicio

Añade país y provincia y reglas de formateo internacional de la dirección.

## Problemas que encontrarás

Necesitarás modificar constructores, impresores y cualquier lugar que pase estos campos juntos, multiplicando la superficie de cambio.

## Proceso de Refactoring

### 1. Identificar el Data Clump

- Busca grupos de 3 o más campos/parámetros que siempre aparecen juntos
- Identifica el concepto del dominio que representan (ej: "Address", "Money", "DateRange")

### 2. Extraer Value Object
- 
- Crea una nueva clase/struct que encapsule estos datos
- Dale un nombre significativo que represente el concepto del dominio
- Mueve los campos relacionados a esta nueva clase

### 3. Reemplazar progresivamente
- 
- Comienza por una clase o método
- Reemplaza los campos individuales con una instancia del Value Object
- Actualiza los constructores/métodos para aceptar el Value Object
- Ejecuta tests después de cada cambio

### 4. Añadir comportamiento

- Mueve la lógica relacionada con estos datos al Value Object
- Ejemplos: validación, formateo, conversiones, comparaciones
- Esto sigue el principio "Tell, Don't Ask"

### 5. Consolidar duplicación

- Una vez que todos los lugares usan el Value Object
- La lógica duplicada desaparece naturalmente
- Los cambios futuros se hacen en un solo lugar

## Técnicas de Refactoring Aplicables

- **Extract Class**: Crear la nueva clase de Value Object
- **Introduce Parameter Object**: Reemplazar lista de parámetros con el objeto
- **Preserve Whole Object**: Pasar el objeto completo en lugar de campos individuales
- **Move Method**: Mover comportamiento relacionado al Value Object

## Beneficios

- **Menos parámetros**: Métodos más fáciles de entender y llamar. El objeto introducido representa un concepto que da sentido al conjunto de datos que agrupa.
- **Cambios centralizados**: Modificar la estructura una sola vez o añadirle comportamiento.
- **Tipo más seguro**: El compilador ayuda a prevenir errores, el código del objeto se preocupa por mantener sus invariantes.
- **Comportamiento cohesivo**: La lógica relacionada vive junta.
- **Mejor semántica**: El código expresa conceptos del dominio.

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/bloaters/data-clump.ts) - [README](../../typescript/src/code-smells/bloaters/data-clump.readme.md)
- [Go](../../go/code_smells/bloaters/data_clump.go) - [README](../../go/code_smells/bloaters/data_clump.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/bloaters/DataClump.java) - [README](../../java/src/main/java/com/refactoring/codesmells/bloaters/DataClump.readme.md)
- [PHP](../../php/src/code-smells/bloaters/DataClump.php) - [README](../../php/src/code-smells/bloaters/DataClump.readme.md)
- [Python](../../python/src/code_smells/bloaters/data_clump.py) - [README](../../python/src/code_smells/bloaters/data_clump_readme.md)
- [C#](../../csharp/src/code-smells/bloaters/DataClump.cs) - [README](../../csharp/src/code-smells/bloaters/data-clump.readme.md)

## Referencias en Español

- [Primitive Obsession](https://franiglesias.github.io/primitive-obsession/) - Explicación detallada sobre el problema de usar tipos primitivos para conceptos del dominio
- [Encapsular primitivos y colecciones](https://franiglesias.github.io/encapsulate/) - Técnicas para encapsular tipos primitivos y colecciones en objetos del dominio
- [Parameter Object](https://franiglesias.github.io/blogtober19-parameter/) - Patrón para agrupar parámetros relacionados en un objeto
- [Refactor cotidiano (4). Sustituye escalares por objetos](https://franiglesias.github.io/everyday-refactor-4/) - Guía práctica para reemplazar tipos primitivos con objetos

## Referencias

- [Refactoring Guru - Data Clumps](https://refactoring.guru/smells/data-clumps)
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
