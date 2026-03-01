# Duplicated Code

Código duplicado.

## Definición

El código duplicado ocurre cuando la misma estructura de código o lógica aparece en más de un lugar. Es uno de los code smells más comunes y peligrosos, ya que cualquier cambio en la lógica debe replicarse en todas las copias, aumentando el riesgo de inconsistencias.

## Descripción

**Duplicated Code** es quizás el code smell más fundamental y pernicioso. Cuando el mismo código (o muy similar) aparece en múltiples lugares, se multiplica el esfuerzo de mantenimiento y el riesgo de bugs.

Tipos de duplicación:
- **Duplicación exacta**: El mismo código copiado y pegado
- **Duplicación estructural**: Misma lógica con diferentes variables o nombres
- **Duplicación semántica**: Diferentes implementaciones de la misma funcionalidad
- **Duplicación de concepto**: Múltiples formas de hacer lo mismo

El problema fundamental es que al cambiar la lógica, debes recordar actualizar todas las copias. Es inevitable que olvides alguna, creando inconsistencias y bugs sutiles.

## Síntomas

- El mismo bloque de código aparece en múltiples lugares
- Usar "copiar y pegar" es parte del workflow
- Mismo algoritmo con ligeras variaciones
- Cambios requieren tocar múltiples archivos o métodos
- Bugs que se arreglan en un lugar pero persisten en otro
- Funciones con nombres similares y código casi idéntico
- Sentimiento de "déjà vu" al leer el código

## Ejemplo

```pseudocode
function calculateOrderTotalWithTax(
  items: array of {price: number, qty: number},
  taxRate: number
): number {
  subtotal = 0
  for each item in items {
    subtotal = subtotal + item.price * item.qty
  }
  tax = subtotal * taxRate
  return subtotal + tax
}

function computeCartTotalIncludingTax(
  items: array of {price: number, quantity: number},
  taxRate: number
): number {
  subtotal = 0
  for each item in items {
    subtotal = subtotal + item.price * item.quantity
  }
  tax = subtotal * taxRate
  return subtotal + tax
}

// Mismo algoritmo, nombres ligeramente diferentes
// Cambiar la lógica requiere tocar ambos
```

## Ejercicio

Cambia la regla de impuestos para que sea escalonada (por ejemplo, 10% hasta $100 y 21% por encima).

## Problemas que encontrarás

Tendrás que actualizar múltiples implementaciones y recordar mantenerlas consistentes, lo que demuestra cómo la duplicación multiplica el esfuerzo y el riesgo de error humano.

## Proceso de Refactoring

### 1. Detectar duplicación
- Usa herramientas de detección de duplicación
- Busca visualmente patrones repetidos
- Identifica código con estructuras similares
- Marca todas las instancias del código duplicado

### 2. Analizar diferencias
- Compara las versiones duplicadas
- Identifica qué es idéntico y qué varía
- Las variaciones serán parámetros del método extraído

### 3. Extraer método común
- Crea un nuevo método que contenga la lógica común
- Parametriza las diferencias
- Dale un nombre que exprese claramente su propósito

### 4. Reemplazar duplicados con llamadas
- Reemplaza cada instancia duplicada con una llamada al método común
- Pasa las variaciones como argumentos
- Verifica que el comportamiento es idéntico después de cada cambio

### 5. Consolidar en clase o módulo apropiado
- Coloca el método extraído donde tiene más sentido
- Considera crear una clase utilidad si no hay un lugar obvio
- Asegúrate de que es fácilmente accesible donde se necesita

### 6. Extraer clase si hay múltiples métodos relacionados
- Si extraes varios métodos relacionados
- Considera crear una clase que los agrupe
- Esto previene crear muchas funciones utilidad dispersas

## Técnicas de Refactoring Aplicables

- **Extract Method**: Extraer código duplicado a un método
- **Extract Class**: Si hay múltiples métodos duplicados relacionados
- **Pull Up Method**: Si la duplicación está en subclases
- **Form Template Method**: Para algoritmos similares con variaciones
- **Substitute Algorithm**: Reemplazar con un algoritmo mejor
- **Extract Superclass**: Si clases diferentes tienen código duplicado

## Beneficios

- **DRY Principle**: Don't Repeat Yourself - una sola fuente de verdad
- **Cambios centralizados**: Modificar en un solo lugar
- **Consistencia garantizada**: Imposible tener versiones diferentes
- **Menos bugs**: No puedes olvidar actualizar una copia
- **Código más corto**: Menos líneas que mantener
- **Testing más fácil**: Testear la lógica una sola vez
- **Comprensión mejorada**: El código expresa conceptos claramente

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/dispensables/duplicated-code.ts) - [README](../../typescript/src/code-smells/dispensables/duplicated-code.readme.md)
- [Go](../../go/code_smells/dispensables/duplicated_code.go) - [README](../../go/code_smells/dispensables/duplicated_code.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/dispensables/DuplicatedCode.java) - [README](../../java/src/main/java/com/refactoring/codesmells/dispensables/DuplicatedCode.readme.md)
- [PHP](../../php/src/code-smells/dispensables/DuplicatedCode.php) - [README](../../php/src/code-smells/dispensables/DuplicatedCode.readme.md)
- [Python](../../python/src/code_smells/dispensables/duplicated_code.py) - [README](../../python/src/code_smells/dispensables/duplicated_code_readme.md)
- [C#](../../csharp/src/code-smells/dispensables/DuplicatedCode.cs) - [README](../../csharp/src/code-smells/dispensables/duplicated-code.readme.md)

## Referencias en Español

- [De abstracciones y duplicaciones (DRY)](https://franiglesias.github.io/dry-abstraction/) - Análisis profundo del principio DRY y cuándo aplicarlo correctamente

## Referencias

- [Refactoring Guru - Duplicate Code](https://refactoring.guru/smells/duplicate-code)
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
- Andrew Hunt, David Thomas - "The Pragmatic Programmer" - DRY Principle
