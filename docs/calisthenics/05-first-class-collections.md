# Colecciones de Primera Clase

Regla 5 de Object Calisthenics

## Definición

Cualquier clase que contenga una colección (array, lista, set, map, etc.) no debe contener ninguna otra propiedad. La colección debe ser el único campo de instancia, y la clase debe encapsular todo el comportamiento relacionado con esa colección.

## Descripción

Una **First-Class Collection** es un objeto cuyo único propósito es envolver una colección y proporcionar comportamiento específico del dominio relacionado con ella.

Los tipos de colección primitivos (arrays, listas, sets, maps) son estructuras genéricas que:
- No expresan el concepto del dominio que representan
- No protegen invariantes de la colección
- No encapsulan operaciones específicas del dominio
- Fuerzan código cliente a trabajar con estructuras genéricas

Problemas de usar colecciones primitivas directamente:

1. **Sin invariantes**: Una lista de usuarios puede estar vacía cuando debería tener al menos uno
2. **Validación duplicada**: Cada lugar que usa la colección debe validar
3. **Operaciones esparcidas**: Lógica de filtrado, búsqueda, agregación duplicada
4. **Sin expresividad**: `array` no comunica que es una lista de pedidos, carrito de compra, o historial
5. **Comportamiento de dominio perdido**: Operaciones como "calcular total", "encontrar por criterio", "validar capacidad" están en el código cliente
6. **Testing complicado**: Hay que crear manualmente colecciones válidas en cada test

La solución es crear objetos que:
- Encapsulan la colección como único campo privado
- Exponen operaciones del dominio, no operaciones genéricas de colección
- Protegen invariantes (tamaño mínimo/máximo, elementos únicos, ordenamiento)
- Expresan conceptos del dominio explícitamente

## Síntomas

- Clases con una colección y otros campos
- Arrays/listas pasados como parámetros directamente
- Métodos que retornan colecciones primitivas
- Loops que filtran/transforman colecciones repetidos en varios lugares
- Validaciones de colección (`if (list.size() > 0)`) duplicadas
- Operaciones matemáticas sobre colecciones (`sum`, `average`) en código cliente
- Comentarios explicando qué representa una colección
- Getters que retornan colecciones mutables
- Lógica de negocio que opera directamente sobre colecciones

## Ejemplo

### Antes (Violación)

```pseudocode
class ShoppingCart {
  array items
  string userId
  decimal discount

  method addItem(item) {
    this.items.add(item)
  }

  method getItems() {
    return this.items  // Expone colección mutable
  }
}

class OrderService {
  method calculateTotal(cart) {
    // Lógica duplicada: sumar precios
    total = 0
    for each item in cart.items {
      total = total + item.price * item.quantity
    }
    return total
  }

  method hasExpensiveItems(cart) {
    // Lógica duplicada: filtrar por criterio
    for each item in cart.items {
      if (item.price > 100) {
        return true
      }
    }
    return false
  }

  method getItemCount(cart) {
    // Lógica duplicada: contar
    count = 0
    for each item in cart.items {
      count = count + item.quantity
    }
    return count
  }
}

class InventoryService {
  method checkStock(items) {
    // Validación duplicada
    if (items.size() == 0) {
      throw "Cart is empty"
    }

    // Misma lógica de iteración en otro servicio
    for each item in items {
      stock = getStock(item.productId)
      if (stock < item.quantity) {
        return false
      }
    }
    return true
  }
}
```

**Problemas**:
- `items` expuesta directamente (puede mutarse externamente)
- Operaciones sobre la colección duplicadas en múltiples servicios
- Sin validación de invariantes (carrito puede estar vacío)
- Sin expresividad del dominio (`array` no dice que es un carrito)
- Acoplamiento a estructura interna

### Después (Cumplimiento)

```pseudocode
// First-Class Collection: CartItems
class CartItems {
  private array items  // Única propiedad

  constructor(items) {
    this.items = items
  }

  method add(item) {
    // Validación de negocio
    if (this.contains(item.productId)) {
      this.increaseQuantity(item.productId, item.quantity)
    } else {
      this.items.add(item)
    }
  }

  method remove(productId) {
    this.items = filter(this.items, (item) => item.productId != productId)
  }

  method calculateTotal() {
    total = 0
    for each item in this.items {
      total = total + item.price * item.quantity
    }
    return new Money(total, "USD")
  }

  method hasExpensiveItems() {
    for each item in this.items {
      if (item.price > 100) {
        return true
      }
    }
    return false
  }

  method getTotalItemCount() {
    count = 0
    for each item in this.items {
      count = count + item.quantity
    }
    return count
  }

  method isEmpty() {
    return this.items.size() == 0
  }

  method contains(productId) {
    for each item in this.items {
      if (item.productId == productId) {
        return true
      }
    }
    return false
  }

  method getByProductId(productId) {
    for each item in this.items {
      if (item.productId == productId) {
        return item
      }
    }
    return null
  }

  private method increaseQuantity(productId, quantity) {
    item = this.getByProductId(productId)
    item.quantity = item.quantity + quantity
  }

  method getProductIds() {
    return map(this.items, (item) => item.productId)
  }

  // Iterator pattern para poder iterar sin exponer colección interna
  method forEach(callback) {
    for each item in this.items {
      callback(item)
    }
  }

  method count() {
    return this.items.size()
  }
}

// Uso mejorado
class ShoppingCart {
  CartItems items  // No otros campos en esta versión simplificada
  // Si necesitas userId y discount, crea otra First-Class Collection
  // o extrae a otra clase (ej: ShoppingSession con cart, user, discount)

  constructor() {
    this.items = new CartItems([])
  }

  method addItem(item) {
    this.items.add(item)
  }

  method removeItem(productId) {
    this.items.remove(productId)
  }

  method getTotal() {
    return this.items.calculateTotal()
  }

  method isEmpty() {
    return this.items.isEmpty()
  }

  method hasExpensiveItems() {
    return this.items.hasExpensiveItems()
  }
}

class OrderService {
  method createOrder(cart) {
    if (cart.isEmpty()) {
      throw "Cannot create order from empty cart"
    }

    // Comportamiento encapsulado en CartItems
    total = cart.getTotal()
    return new Order(cart.items, total)
  }
}

class InventoryService {
  method checkStock(cartItems) {
    if (cartItems.isEmpty()) {
      throw "Cart is empty"
    }

    allInStock = true
    cartItems.forEach((item) => {
      stock = this.getStock(item.productId)
      if (stock < item.quantity) {
        allInStock = false
      }
    })
    return allInStock
  }
}
```

**Mejoras**:
- Colección encapsulada en objeto dedicado (`CartItems`)
- Operaciones de dominio centralizadas
- Validaciones e invariantes protegidos
- Sin duplicación de lógica
- Tipo explícito del dominio
- Inmutabilidad controlada (puedes hacer la colección interna inmutable)
- Testing simplificado (CartItems es testeable aisladamente)

## Ejercicio

**Tarea**: En el código proporcionado en tu lenguaje, identifica todas las colecciones y envuélvelas en First-Class Collections que encapsulen el comportamiento específico del dominio.

**Criterios de éxito**:
1. Cada colección está envuelta en una clase dedicada
2. La clase de colección solo tiene la colección como campo (sin otros campos)
3. Operaciones relacionadas con la colección están en la clase
4. No hay código cliente que opere directamente sobre la colección
5. La colección interna no se expone (o se retorna copia inmutable)

## Problemas que Encontrarás

### 1. Clases con colección + otros campos

Si tu clase tiene una colección y otros campos, no puedes aplicar esta regla directamente porque dice "no otros campos". Opciones:

**Opción A**: Extrae la colección a su propia clase First-Class Collection, y úsala desde la clase original
**Opción B**: Si la colección y los otros campos están muy relacionados, quizás necesitas reorganizar el diseño

### 2. "Es solo una lista, ¿por qué una clase?"

Porque la lista representa un concepto del dominio con operaciones específicas. `CartItems` no es "solo una lista", es un concepto de negocio con reglas.

### 3. Exponer la colección para iteración

No expongas la colección directamente. Alternativas:
- Proporciona métodos específicos (`calculateTotal()`, `hasExpensiveItems()`)
- Implementa Iterator pattern (`forEach(callback)`)
- Retorna copias inmutables si necesitas acceso read-only
- Usa métodos de transformación (`map()`, `filter()`) que retornan nuevas First-Class Collections

### 4. Performance de encapsulación

Sí, hay un pequeño overhead, pero:
- Es despreciable comparado con operaciones reales (DB, network)
- Los beneficios de mantenibilidad superan cualquier costo
- Si tienes un cuello de botella real, optimiza ese punto específico

### 5. Colecciones vacías vs null

Usa el patrón **Null Object**:
- Nunca uses `null` para colecciones
- Usa colecciones vacías como valor por defecto
- Proporciona métodos como `isEmpty()` para verificar

## Proceso de Aplicación

### 1. Identificar colecciones en el dominio

Busca:
- Campos que son arrays, listas, sets, maps
- Parámetros de métodos que son colecciones
- Retornos de métodos que son colecciones
- Colecciones que representan conceptos del dominio

### 2. Identificar operaciones duplicadas

Para cada colección, busca:
- ¿Qué operaciones se hacen sobre ella en múltiples lugares?
- ¿Qué validaciones se repiten?
- ¿Qué cálculos se realizan?
- ¿Qué búsquedas/filtrados se hacen?

### 3. Crear la First-Class Collection

```pseudocode
class ConceptCollection {
  private collectionType items

  constructor(items) {
    this.validate(items)
    this.items = items
  }

  // Métodos de dominio específicos
  method operationSpecificToThisConcept() {
    // ...
  }

  // Métodos de acceso controlado
  method forEach(callback) {
    for each item in this.items {
      callback(item)
    }
  }
}
```

### 4. Mover operaciones a la colección

- Identifica lógica duplicada que opera sobre la colección
- Muévela a métodos en la First-Class Collection
- Elimina duplicación en código cliente

### 5. Proteger invariantes

Añade validaciones en:
- Constructor (tamaño inicial, elementos válidos)
- Métodos de modificación (add, remove)
- Asegura que la colección nunca esté en estado inválido

### 6. Encapsular mutabilidad

- Haz la colección interna privada
- No expongas referencias mutables
- Proporciona métodos específicos para modificación controlada
- Considera hacer la colección inmutable (métodos retornan nuevas instancias)

### 7. Reemplazar uso directo de colecciones

- Cambia parámetros de métodos a usar First-Class Collections
- Cambia retornos de métodos a usar First-Class Collections
- Actualiza código cliente para usar métodos de la colección, no operaciones directas

## Técnicas de Refactoring Aplicables

- **Extract Class**: Crear la First-Class Collection
- **Move Method**: Mover operaciones sobre colección a la nueva clase
- **Encapsulate Collection**: Proteger la colección interna
- **Replace Array with Object**: Reemplazar array primitivo con objeto
- **Introduce Parameter Object**: Si colección y metadatos se pasan juntos
- **Remove Middle Man**: Si la clase original solo delega a la colección

## Beneficios

### 1. Encapsulación de Comportamiento

Lógica relacionada con la colección está centralizada, no duplicada.

### 2. Protección de Invariantes

La colección garantiza sus propias reglas de validación y consistencia.

### 3. Expresividad del Dominio

`CartItems` comunica mucho más que `array` sobre el propósito y uso.

### 4. Eliminación de Duplicación

Operaciones sobre la colección se implementan una vez, se usan muchas veces.

### 5. Testing Simplificado

La First-Class Collection es testeable independientemente del resto del sistema.

### 6. Inmutabilidad Controlada

Puedes hacer la colección inmutable fácilmente, retornando nuevas instancias en modificaciones.

### 7. Evolución del Código

Cambios en cómo se maneja la colección están localizados en un solo lugar.

### 8. Mejor API

Métodos como `hasExpensiveItems()` son más expresivos que `cart.items.filter(...)`.

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/calisthenics-exercises/first-class-collections.ts)
- [Go](../../go/calisthenics_exercises/04_first_class_collections.go)
- [Java](../../java/src/main/java/com/refactoring/calisthenics/FirstClassCollections.java)
- [PHP](../../php/src/calisthenics-exercises/FirstClassCollections.php)
- [Python](../../python/src/calisthenics_exercises/first_class_collections.py)
- [C#](../../csharp/src/calisthenics-exercises/FirstClassCollections.cs)

## Referencias en Español

- [Calistenias para objetos de valor](https://franiglesias.github.io/calistenics-and-value-objects/) - Incluye discusión sobre colecciones
- [Object Calisthenics para adelgazar las clases](https://franiglesias.github.io/calistenics-and-small-classes/) - Técnicas para encapsular colecciones

## Referencias

- **"Domain-Driven Design"** - Eric Evans - Collections como parte del modelo de dominio
- **"Refactoring: Improving the Design of Existing Code"** - Martin Fowler - Encapsulate Collection
- [Object Calisthenics - William Durand](https://williamdurand.fr/2013/06/03/object-calisthenics/) - Regla #5
- [First-Class Collection Pattern](https://refactoring.com/catalog/encapsulateCollection.html) - Martin Fowler
- **"Smalltalk Best Practice Patterns"** - Kent Beck - Collection protocols
