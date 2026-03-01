# Shotgun Surgery

Cirugía de escopeta.

## Definición

Cuando necesito hacer un cambio tengo que hacerlo en muchos lugares del código que pueden estar alejados entre sí, incluso en distintos módulos.

## Descripción

**Shotgun Surgery** es lo opuesto a **Divergent Change**. Mientras Divergent Change significa que una clase tiene múltiples razones para cambiar, Shotgun Surgery significa que un solo cambio lógico requiere modificaciones en múltiples clases dispersas por el código.

Este smell aparece cuando una responsabilidad o concepto está fragmentado en muchos lugares. Por ejemplo:
- La misma regla de negocio duplicada en múltiples clases
- Constantes mágicas repetidas en diferentes archivos
- Lógica de validación esparcida por toda la aplicación
- Algoritmos de cálculo replicados en varios lugares

El resultado es que hacer un cambio conceptualmente simple (como cambiar la tasa de impuestos) requiere tocar docenas de archivos, aumentando el riesgo de olvidar alguno y creando inconsistencias.

## Síntomas

- Un pequeño cambio requiere modificaciones en muchas clases diferentes
- La misma constante o valor literal aparece en múltiples lugares
- La misma lógica está duplicada en varios métodos o clases
- Necesitas usar "buscar y reemplazar" para hacer cambios
- Los cambios requieren modificar archivos en módulos no relacionados
- Es fácil olvidar actualizar uno de los lugares afectados
- Los bugs aparecen porque un lugar no se actualizó correctamente

## Ejemplo

```pseudocode
// La regla de impuestos (21%) está duplicada en múltiples clases

class PriceCalculator {
  function totalWithTax(items) {
    subtotal = sum(items, item => item.price * item.quantity)
    tax = subtotal * 0.21  // Tasa duplicada
    return subtotal + tax
  }
}

class InvoiceService {
  function createTotal(items) {
    base = sum(items, item => item.price * item.quantity)
    vat = base * 0.21  // Tasa duplicada
    return base + vat
  }
}

class SalesReport {
  function summarize(items) {
    sum = sum(items, item => item.price * item.quantity)
    tax = sum * 0.21  // Tasa duplicada
    total = sum + tax
    return "total=" + total.toFixed(2)
  }
}

class LoyaltyPoints {
  function points(items) {
    base = sum(items, item => item.price * item.quantity)
    withTax = base + base * 0.21  // Tasa duplicada
    return floor(withTax / 10)
  }
}

// Cambiar el impuesto al 18.5% requiere modificar 4 lugares
```

## Ejercicio

Cambia el impuesto del 21% al 18.5% con redondeo a 2 decimales.

## Problemas que encontrarás

Tendrás que buscar cada copia y asegurar un redondeo consistente en todas partes, destacando cómo la duplicación convierte un cambio pequeño en muchas ediciones arriesgadas.

## Proceso de Refactoring

### 1. Identificar la duplicación
- Busca el código o lógica que necesita cambiar
- Encuentra todas las copias dispersas por el código
- Usa herramientas de búsqueda para localizar duplicados
- Documenta todos los lugares afectados

### 2. Extraer a un lugar común
- Crea una clase, módulo o función que contenga la lógica
- Dale un nombre significativo que refleje su propósito
- Ejemplo: `TaxCalculator`, `Constants`, `BusinessRules`

### 3. Centralizar constantes
- Mueve valores literales a constantes nombradas
- Agrupa constantes relacionadas en un módulo o clase
- Usa enums o configuración para valores que pueden variar

### 4. Reemplazar duplicados con llamadas
- Reemplaza cada copia con una llamada al lugar centralizado
- Hazlo de manera incremental, verificando tests después de cada cambio
- Asegúrate de que el comportamiento es idéntico

### 5. Consolidar validaciones
- Si la validación está duplicada, extráela a un validador
- Crea métodos de validación reutilizables
- Ejemplo: `EmailValidator.validate()`, `PriceValidator.isPositive()`

### 6. Usar herencia o composición
- Si múltiples clases tienen el mismo comportamiento duplicado
- Extrae a una clase base o a un componente compartido
- Usa composición para compartir lógica sin herencia forzada

## Técnicas de Refactoring Aplicables

- **Extract Method**: Extraer lógica duplicada a un método
- **Extract Class**: Crear una clase para agrupar comportamiento relacionado
- **Move Method**: Mover comportamiento al lugar donde pertenece
- **Pull Up Method**: Si la duplicación está en subclases, subirla a la clase base
- **Inline Class**: Si una clase existe solo como fragmento, considerar integrarla
- **Replace Magic Number with Symbolic Constant**: Centralizar valores literales

## Beneficios

- **Cambios en un solo lugar**: Modificaciones centralizadas y seguras
- **Consistencia garantizada**: Imposible tener versiones diferentes de la lógica
- **Menos bugs**: No puedes olvidar actualizar una copia
- **Más fácil de entender**: El código expresa conceptos una sola vez
- **Tests más simples**: Testear la lógica en un solo lugar
- **Mantenimiento reducido**: Menos código que mantener sincronizado

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/change-preventers/shotgun-surgery.ts) - [README](../../typescript/src/code-smells/change-preventers/shotgun-surgery.readme.md)
- [Go](../../go/code_smells/change_preventers/shotgun_surgery.go) - [README](../../go/code_smells/change_preventers/shotgun_surgery.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/changepreventers/ShotgunSurgery.java) - [README](../../java/src/main/java/com/refactoring/codesmells/changepreventers/ShotgunSurgery.readme.md)
- [PHP](../../php/src/code-smells/change-preventers/ShotgunSurgery.php) - [README](../../php/src/code-smells/change-preventers/ShotgunSurgery.readme.md)
- [Python](../../python/src/code_smells/change_preventers/shotgun_surgery.py) - [README](../../python/src/code_smells/change_preventers/shotgun_surgery_readme.md)
- [C#](../../csharp/src/code-smells/change-preventers/ShotgunSurgery.cs) - [README](../../csharp/src/code-smells/change-preventers/shotgun-surgery.readme.md)

## Referencias en Español

- [De abstracciones y duplicaciones (DRY)](https://franiglesias.github.io/dry-abstraction/) - Análisis profundo del principio DRY y cómo aplicarlo correctamente
- [Refactor cotidiano (8). Dónde poner el conocimiento](https://franiglesias.github.io/everyday-refactor-8/) - Guía para centralizar conocimiento y evitar duplicación

## Referencias

- [Refactoring Guru - Shotgun Surgery](https://refactoring.guru/smells/shotgun-surgery)
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
- Andrew Hunt, David Thomas - "The Pragmatic Programmer" - DRY Principle
