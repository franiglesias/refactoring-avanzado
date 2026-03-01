# Mantener las Unidades Pequeñas

Regla 7 de Object Calisthenics

## Definición

Las clases deben tener un máximo de **50 líneas** y los métodos un máximo de **5 líneas** (según la regla original). Aunque estos números son arbitrarios, el principio es claro: mantener unidades de código pequeñas y enfocadas en una única responsabilidad.

## Descripción

**Unidades grandes** (clases y métodos) son un síntoma claro de violación del **Single Responsibility Principle**. Cuando una clase o método crece demasiado, típicamente está haciendo más de una cosa.

Problemas de unidades grandes:

1. **Múltiples responsabilidades**: Una clase con 500 líneas probablemente hace 5-10 cosas diferentes
2. **Alta complejidad ciclomática**: Más líneas = más caminos de ejecución = más complejidad
3. **Difícil de entender**: Hay que mantener demasiado contexto mental
4. **Testing complicado**: Muchos casos edge, configuración compleja
5. **Conflictos en control de versiones**: Más personas modifican el mismo archivo
6. **Violación de Open/Closed**: Añadir funcionalidad requiere modificar clases grandes
7. **Acoplamiento alto**: Clases grandes tienden a depender de muchas otras cosas

Números sugeridos (ajusta según contexto):
- **Métodos**: 5-15 líneas idealmente, máximo 20-30
- **Clases**: 100-200 líneas idealmente, máximo 300-400
- **Archivos**: Un archivo por clase/módulo

La regla original dice 5 líneas para métodos para forzar **composición extrema**. Es deliberadamente estricto para entrenar el instinto de descomposición. En código real, 10-15 líneas es razonable, pero si supera 20 líneas, probablemente necesita refactoring.

## Síntomas

### Para Clases

- Más de 200-300 líneas
- Scroll extensivo necesario para ver toda la clase
- Nombre que contiene "Manager", "Helper", "Util", "Service" genéricos
- Muchas dependencias (imports)
- Muchos métodos públicos (más de 7-10)
- Comentarios que separan "secciones" dentro de la clase
- Dificultad para nombrar la clase sin usar "and" o "or"
- Tests que requieren mucho setup

### Para Métodos

- Más de 20-30 líneas
- Múltiples niveles de indentación
- Bloques de código separados por comentarios
- Variables temporales que solo se usan en una sección
- Dificultad para nombrar sin usar "and"
- Necesidad de scroll para verlo completo

## Ejemplo

### Antes (Violación)

```pseudocode
// Clase de 200+ líneas con múltiples responsabilidades
class OrderProcessor {
  database db
  emailService email
  inventoryService inventory
  paymentGateway payment
  logger log
  config configuration

  // Método de 50+ líneas
  method processOrder(orderId) {
    // Validar pedido (10 líneas)
    order = db.findOrder(orderId)
    if (order == null) {
      log.error("Order not found: " + orderId)
      return false
    }
    if (order.status != "PENDING") {
      log.error("Order already processed: " + orderId)
      return false
    }
    if (order.items.length == 0) {
      log.error("Order has no items: " + orderId)
      return false
    }

    // Verificar inventario (15 líneas)
    for each item in order.items {
      stock = inventory.getStock(item.productId)
      if (stock < item.quantity) {
        log.error("Insufficient stock for: " + item.productId)
        sendOutOfStockEmail(order.customerEmail, item.productId)
        return false
      }
    }

    // Calcular total (20 líneas)
    subtotal = 0
    for each item in order.items {
      subtotal = subtotal + item.price * item.quantity
    }
    discount = 0
    if (order.couponCode != null) {
      coupon = db.findCoupon(order.couponCode)
      if (coupon != null and coupon.isValid()) {
        discount = subtotal * coupon.percentage / 100
      }
    }
    tax = (subtotal - discount) * configuration.getTaxRate()
    shipping = subtotal >= configuration.getFreeShippingThreshold() ? 0 : configuration.getShippingCost()
    total = subtotal - discount + tax + shipping

    // Procesar pago (10 líneas)
    paymentResult = payment.charge(order.customerId, total)
    if (not paymentResult.success) {
      log.error("Payment failed: " + paymentResult.error)
      sendPaymentFailedEmail(order.customerEmail, paymentResult.error)
      return false
    }

    // Actualizar inventario (10 líneas)
    for each item in order.items {
      inventory.reduceStock(item.productId, item.quantity)
    }

    // Actualizar pedido (5 líneas)
    order.status = "PROCESSED"
    order.total = total
    order.paymentId = paymentResult.transactionId
    db.updateOrder(order)

    // Enviar confirmación (5 líneas)
    emailBody = buildConfirmationEmail(order)
    email.send(order.customerEmail, "Order Confirmation", emailBody)

    log.info("Order processed successfully: " + orderId)
    return true
  }

  // Muchos otros métodos (100+ líneas más)
  method sendOutOfStockEmail(email, productId) { /* ... */ }
  method sendPaymentFailedEmail(email, error) { /* ... */ }
  method buildConfirmationEmail(order) { /* ... */ }
  method calculateDiscount(order) { /* ... */ }
  method validateCoupon(code) { /* ... */ }
  // etc...
}
```

**Problemas**:
- Clase con múltiples responsabilidades (validación, inventario, pago, email, cálculo)
- Método gigante que hace todo
- Imposible de entender sin leer todo
- Testing requiere configurar todo (db, email, inventory, payment)
- Alto acoplamiento

### Después (Cumplimiento)

```pseudocode
// Clase enfocada: solo orquestación (30 líneas)
class OrderProcessor {
  OrderValidator validator
  OrderPricing pricing
  PaymentProcessor paymentProcessor
  InventoryManager inventoryManager
  OrderNotifier notifier
  OrderRepository repository

  method process(orderId) {
    order = this.loadAndValidate(orderId)
    this.checkInventory(order)
    total = this.calculateAndSetTotal(order)
    this.processPayment(order, total)
    this.updateInventoryAndOrder(order)
    this.notifyCustomer(order)
  }

  private method loadAndValidate(orderId) {
    order = repository.find(orderId)
    validator.ensureCanProcess(order)
    return order
  }

  private method checkInventory(order) {
    inventoryManager.ensureAvailability(order.items)
  }

  private method calculateAndSetTotal(order) {
    total = pricing.calculate(order)
    order.setTotal(total)
    return total
  }

  private method processPayment(order, total) {
    paymentProcessor.charge(order, total)
  }

  private method updateInventoryAndOrder(order) {
    inventoryManager.reduceStock(order.items)
    repository.save(order)
  }

  private method notifyCustomer(order) {
    notifier.sendConfirmation(order)
  }
}

// Clase enfocada: validación (20 líneas)
class OrderValidator {
  method ensureCanProcess(order) {
    this.ensureExists(order)
    this.ensureIsPending(order)
    this.ensureHasItems(order)
  }

  private method ensureExists(order) {
    if (order == null) throw "Order not found"
  }

  private method ensureIsPending(order) {
    if (order.status != "PENDING") throw "Order already processed"
  }

  private method ensureHasItems(order) {
    if (order.items.isEmpty()) throw "Order has no items"
  }
}

// Clase enfocada: cálculo de precios (40 líneas)
class OrderPricing {
  DiscountCalculator discountCalculator
  TaxCalculator taxCalculator
  ShippingCalculator shippingCalculator

  method calculate(order) {
    subtotal = this.calculateSubtotal(order)
    discount = discountCalculator.calculate(order, subtotal)
    tax = taxCalculator.calculate(subtotal - discount)
    shipping = shippingCalculator.calculate(subtotal)
    return new OrderTotal(subtotal, discount, tax, shipping)
  }

  private method calculateSubtotal(order) {
    total = 0
    for each item in order.items {
      total = total + item.calculateTotal()
    }
    return total
  }
}

// Clase enfocada: gestión de inventario (30 líneas)
class InventoryManager {
  InventoryRepository repository
  OutOfStockNotifier notifier

  method ensureAvailability(items) {
    for each item in items {
      this.checkStock(item)
    }
  }

  private method checkStock(item) {
    stock = repository.getStock(item.productId)
    if (stock < item.quantity) {
      this.handleOutOfStock(item)
    }
  }

  private method handleOutOfStock(item) {
    notifier.notify(item)
    throw "Insufficient stock for: " + item.productId
  }

  method reduceStock(items) {
    for each item in items {
      repository.reduce(item.productId, item.quantity)
    }
  }
}

// Clase enfocada: procesamiento de pagos (25 líneas)
class PaymentProcessor {
  PaymentGateway gateway
  PaymentFailureHandler failureHandler

  method charge(order, amount) {
    result = gateway.charge(order.customerId, amount)
    this.handleResult(result, order)
    order.setPaymentId(result.transactionId)
  }

  private method handleResult(result, order) {
    if (not result.success) {
      failureHandler.handle(order, result.error)
      throw "Payment failed"
    }
  }
}

// Clase enfocada: notificaciones (20 líneas)
class OrderNotifier {
  EmailService emailService
  EmailTemplateBuilder templateBuilder

  method sendConfirmation(order) {
    email = this.buildEmail(order)
    emailService.send(email)
  }

  private method buildEmail(order) {
    body = templateBuilder.buildConfirmation(order)
    return new Email(
      order.customerEmail,
      "Order Confirmation",
      body
    )
  }
}
```

**Mejoras**:
- Cada clase tiene una única responsabilidad clara
- Métodos pequeños (5-10 líneas) y enfocados
- Nombres descriptivos cuentan la historia
- Testing simplificado (cada clase se testea independientemente)
- Bajo acoplamiento, alta cohesión
- Fácil de extender sin modificar (Open/Closed)

## Ejercicio

**Tarea**: En el código proporcionado en tu lenguaje, identifica clases y métodos grandes y descompónlos en unidades más pequeñas y enfocadas.

**Criterios de éxito**:
1. Ningún método supera 15-20 líneas
2. Ninguna clase supera 100-150 líneas
3. Cada clase tiene una única responsabilidad
4. Cada método hace una cosa
5. Los nombres son claros y específicos

## Problemas que Encontrarás

### 1. "Ahora tengo demasiadas clases pequeñas"

Esto es correcto. Muchas clases pequeñas y enfocadas son mejores que pocas clases grandes y complejas. Es más fácil:
- Entender una clase pequeña
- Testear una clase pequeña
- Reutilizar una clase pequeña
- Modificar una clase pequeña

### 2. "Los números son arbitrarios"

Sí, lo son. No es que 49 líneas esté bien y 51 esté mal. Los límites son **guías** para detectar problemas. Si una clase tiene 150 líneas pero es cohesiva y clara, probablemente está bien. Si tiene 60 líneas pero hace 3 cosas diferentes, necesita refactoring.

### 3. "Performance por muchas clases pequeñas"

La sobrecarga es despreciable:
- Los compiladores modernos optimizan agresivamente
- El impacto en runtime es insignificante
- Los beneficios de mantenibilidad superan cualquier costo
- Si tienes un cuello de botella real, optimiza ese punto específico

### 4. "Dónde poner las clases pequeñas"

Opciones de organización:
- **Por feature**: Clases relacionadas en el mismo módulo/package
- **Por capa**: Separar por responsabilidad arquitectónica
- **Por dominio**: Siguiendo DDD, agrupar por bounded context

No pongas todo en un solo archivo/directorio. Usa estructura de carpetas significativa.

### 5. "Tests se complican con muchas clases"

Al contrario, se simplifican:
- Cada clase se testea independientemente con menos casos
- Usa mocks/stubs para dependencias
- Tests unitarios más simples y rápidos
- Tests de integración solo para flujos completos

## Proceso de Aplicación

### 1. Identificar unidades grandes

- Busca clases con más de 150-200 líneas
- Busca métodos con más de 20-30 líneas
- Marca las que tienen comentarios separando "secciones"

### 2. Identificar responsabilidades

Para cada unidad grande:
- Lista todas las "cosas" que hace
- Agrupa operaciones relacionadas
- Identifica cohesión natural

### 3. Extraer clases

Por cada responsabilidad identificada:
- Crea una nueva clase con nombre descriptivo
- Mueve los métodos relacionados
- Inyecta dependencias necesarias

### 4. Descomponer métodos grandes

Para cada método grande:
- Identifica bloques lógicos (a menudo separados por comentarios)
- Extrae cada bloque a un método con nombre descriptivo
- El método original queda como orquestador

### 5. Aplicar Single Responsibility Principle

Pregúntate: "¿Esta clase tiene solo una razón para cambiar?"
- **No** → Identifica las múltiples razones y separa

### 6. Usar inyección de dependencias

Las clases pequeñas colaboran a través de interfaces:
- Define interfaces claras
- Inyecta dependencias en constructor
- Facilita testing y reemplazo

### 7. Reorganizar estructura

- Agrupa clases relacionadas en módulos/packages
- Usa nombres de carpetas significativos
- Mantén cohesión en la organización

## Técnicas de Refactoring Aplicables

- **Extract Method**: Descomponer métodos grandes
- **Extract Class**: Separar responsabilidades en nuevas clases
- **Move Method**: Mover métodos a la clase apropiada
- **Inline Class**: Si una clase es demasiado pequeña y trivial, considera combinarla
- **Replace Method with Method Object**: Para métodos muy complejos
- **Introduce Parameter Object**: Agrupar parámetros relacionados
- **Compose Method**: Hacer métodos que lean como narrativa de alto nivel

## Beneficios

### 1. Single Responsibility Principle

Cada unidad hace una cosa y la hace bien.

### 2. Comprensibilidad

Unidades pequeñas son fáciles de entender completamente.

### 3. Testing Simplificado

Clases y métodos pequeños son fáciles de testear en aislamiento.

### 4. Reusabilidad

Unidades pequeñas y enfocadas pueden reutilizarse en diferentes contextos.

### 5. Mantenibilidad

Cambios están localizados en unidades pequeñas específicas.

### 6. Menos Conflictos en VCS

Archivos pequeños = menos personas trabajando en el mismo archivo = menos conflictos.

### 7. Open/Closed Principle

Más fácil extender funcionalidad sin modificar código existente.

### 8. Onboarding Más Rápido

Nuevos desarrolladores pueden entender unidades pequeñas gradualmente.

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/calisthenics-exercises/small-entities.ts)
- [Go](../../go/calisthenics_exercises/07_small_entities.go)
- [Java](../../java/src/main/java/com/refactoring/calisthenics/SmallEntities.java)
- [PHP](../../php/src/calisthenics-exercises/SmallEntities.php)
- [Python](../../python/src/calisthenics_exercises/small_entities.py)
- [C#](../../csharp/src/calisthenics-exercises/SmallEntities.cs)

## Referencias en Español

- [Métodos largos](https://franiglesias.github.io/long-method/) - Problema de métodos largos y soluciones
- [Object Calisthenics para adelgazar las clases](https://franiglesias.github.io/calistenics-and-small-classes/) - Técnicas para mantener clases pequeñas
- [Ejercicio de refactor (2) Extraer hasta la última gota](https://franiglesias.github.io/ejercicio-de-refactor-2/) - Práctica de extracción

## Referencias

- **"Clean Code"** - Robert C. Martin - Capítulos sobre funciones y clases pequeñas
- **"Refactoring: Improving the Design of Existing Code"** - Martin Fowler - Extract Method, Extract Class
- [Object Calisthenics - William Durand](https://williamdurand.fr/2013/06/03/object-calisthenics/) - Regla #7
- **"The Art of Readable Code"** - Boswell & Foucher - Técnicas para código simple
- [Single Responsibility Principle](https://en.wikipedia.org/wiki/Single-responsibility_principle) - Principio SOLID relacionado
- **"Smalltalk Best Practice Patterns"** - Kent Beck - Composed Method pattern
