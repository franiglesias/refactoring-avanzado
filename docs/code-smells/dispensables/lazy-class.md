# Lazy Class

Clase perezosa.

## Definición

Una clase perezosa es aquella que no aporta suficiente valor para justificar su existencia. Suelen ser clases que solo envuelven una operación trivial o que tienen muy poca responsabilidad, añadiendo una complejidad innecesaria al sistema.

## Descripción

**Lazy Class** es una clase que no hace lo suficiente para justificar su existencia. Cada clase en un sistema tiene un costo: debe ser entendida, mantenida, testeada y navegada. Si una clase no proporciona suficiente valor para compensar ese costo, es una Lazy Class.

Causas comunes:
- Clases que alguna vez hicieron algo útil pero se fueron vaciando
- Over-engineering donde se creó una abstracción innecesaria
- Clases creadas anticipando funcionalidad futura que nunca llegó
- Wrappers triviales que solo delegan sin añadir valor
- Clases que solo tienen un método con lógica trivial

El problema es que estas clases añaden indirección innecesaria, haciendo el código más difícil de navegar y entender sin proporcionar beneficios a cambio.

## Síntomas

- La clase tiene solo uno o dos métodos simples
- La clase solo envuelve una operación trivial
- La clase podría ser reemplazada por una función
- La clase tiene muy poco o ningún estado
- La clase fue creada "por si acaso" pero no se usa mucho
- Eliminar la clase simplificaría el código
- La clase no tiene tests significativos porque no hay lógica que testear

## Ejemplo

```pseudocode
type Address = {
  name: string
  line1: string
  city: string
}

class ShippingLabelBuilder {
  function build(address: Address): string {
    // Solo una concatenación simple
    return address.name + " — " + address.line1 +
           (address.city ? ", " + address.city : "")
  }
}

function printShippingLabel() {
  address = {
    name: "John Doe",
    line1: "123 Main St",
    city: "New York"
  }

  // Necesitamos instanciar una clase para una operación trivial
  labelBuilder = new ShippingLabelBuilder()
  label = labelBuilder.build(address)
  print(label)
}
```

## Ejercicio

Reescribe el código para eliminar la necesidad de la clase `ShippingLabelBuilder`.

## Problemas que encontrarás

Mantener una estructura de clase para una lógica tan simple te obliga a instanciar objetos innecesariamente y añade capas de abstracción que dificultan la legibilidad del código sin ofrecer beneficios a cambio.

## Proceso de Refactoring

### 1. Identificar la Lazy Class
- Busca clases con muy pocos métodos o responsabilidades
- Identifica clases que solo envuelven operaciones simples
- Verifica el ratio valor/complejidad de la clase

### 2. Evaluar si aporta valor
- ¿La clase encapsula lógica compleja?
- ¿Proporciona una abstracción útil?
- ¿Facilita testing o reuso?
- ¿Es un punto de extensión necesario?
- Si no aporta nada de esto, considérala lazy

### 3. Inline la clase
- Si la clase tiene un solo método, conviértelo en función
- Si tiene pocos métodos, muévelos a clases relacionadas
- Elimina la clase completamente

### 4. Colapsar jerarquías innecesarias
- Si una subclase hace muy poco
- Combínala con la clase padre
- Usa **Collapse Hierarchy**

### 5. Convertir a función o método
- Si la clase solo tiene un método estático o simple
- Reemplázala con una función pura
- Ejemplo: `ShippingLabelBuilder.build()` → `formatShippingLabel()`

### 6. Integrar en clase cliente
- Si la lazy class solo la usa una clase
- Considera mover su funcionalidad a la clase cliente
- Reduce indirección innecesaria

## Técnicas de Refactoring Aplicables

- **Inline Class**: Eliminar la clase y mover su funcionalidad a donde se usa
- **Collapse Hierarchy**: Si es una subclase trivial, fusionarla con la padre
- **Inline Method**: Si tiene un solo método, hacerlo inline
- **Remove Middle Man**: Si solo delega sin añadir valor

## Beneficios

- **Menos complejidad**: Menos clases que entender y mantener
- **Mejor navegación**: Menos saltos entre archivos
- **Código más directo**: Sin indirección innecesaria
- **Performance**: Menos instanciaciones y llamadas
- **Legibilidad**: La lógica está donde se espera
- **Menos tests**: No necesitas testear wrappers triviales

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/dispensables/lazy-class.ts) - [README](../../typescript/src/code-smells/dispensables/lazy-class.readme.md)
- [Go](../../go/code_smells/dispensables/lazy_class.go) - [README](../../go/code_smells/dispensables/lazy_class.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/dispensables/LazyClass.java) - [README](../../java/src/main/java/com/refactoring/codesmells/dispensables/LazyClass.readme.md)
- [PHP](../../php/src/code-smells/dispensables/LazyClass.php) - [README](../../php/src/code-smells/dispensables/LazyClass.readme.md)
- [Python](../../python/src/code_smells/dispensables/lazy_class.py) - [README](../../python/src/code_smells/dispensables/lazy_class_readme.md)
- [C#](../../csharp/src/code-smells/dispensables/LazyClass.cs) - [README](../../csharp/src/code-smells/dispensables/lazy-class.readme.md)

## Referencias en Español

- [Object Calisthenics para adelgazar las clases](https://franiglesias.github.io/calistenics-and-small-classes/) - Balance entre clases pequeñas útiles y clases innecesarias

## Referencias

- [Refactoring Guru - Lazy Class](https://refactoring.guru/smells/lazy-class)
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
- YAGNI Principle - You Aren't Gonna Need It
