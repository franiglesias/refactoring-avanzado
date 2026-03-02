# Sprout Change (Técnica del Brote)

## Definición

**Sprout Change** (Cambio por Brote o Cambio por Germinación) es una técnica de refactoring que consiste en añadir nueva funcionalidad como código completamente nuevo y aislado, sin modificar el código legacy existente. El código viejo simplemente llama al nuevo código en el punto apropiado.

Como un brote que crece desde una planta establecida, el nuevo código emerge del viejo sin alterar su estructura interna.

## Cuándo Usar

- Necesitas añadir funcionalidad nueva a código legacy
- El código existente es complejo y no quieres tocarlo
- No tienes tests para el código viejo
- El cambio puede expresarse como una operación separada
- Quieres minimizar el riesgo de romper funcionalidad existente
- Puedes hacer el nuevo código testeable independientemente
- El código viejo tiene un punto claro donde insertar la llamada

## Problema que Resuelve

Cuando necesitas añadir funcionalidad a código legacy, enfrentas un dilema:

- **Modificar el código viejo**: Arriesgado, puede introducir bugs, difícil de testear
- **No hacer nada**: La deuda técnica crece, el código se vuelve más complejo
- **Refactorizar todo primero**: Muy lento, puede tomar semanas

Sprout Change resuelve esto ofreciendo un camino intermedio:
1. Escribe el nuevo código de forma limpia y testeable
2. Haz una modificación mínima al código viejo (una llamada)
3. La funcionalidad nueva queda aislada y fácil de mantener
4. Reduces el riesgo al mínimo

## Descripción Detallada

### Cómo Funciona

Sprout Change sigue un patrón simple:

1. **Identifica** dónde necesitas añadir la funcionalidad nueva
2. **Crea** un método/función/clase nueva con el comportamiento nuevo
3. **Testea** el código nuevo de forma aislada
4. **Llama** al código nuevo desde el código viejo con una modificación mínima
5. **Verifica** que todo funciona correctamente

### Diagrama Conceptual

```
ANTES: Código Legacy sin el cambio
┌──────────────────────────────────────┐
│ function processOrder(order) {       │
│                                      │
│     validateOrder(order)             │
│     calculateTotal(order)            │
│     saveToDatabase(order)            │
│                                      │
│     // Necesitamos añadir notificación aquí
│                                      │
│     return order                     │
│ }                                    │
└──────────────────────────────────────┘


DESPUÉS: Con Sprout Change
┌──────────────────────────────────────┐
│ function processOrder(order) {       │
│                                      │
│     validateOrder(order)             │
│     calculateTotal(order)            │
│     saveToDatabase(order)            │
│                                      │
│     sendOrderNotification(order) ◄───┼─── Una línea añadida
│                                      │
│     return order                     │
│ }                                    │
└──────────────────────────────────────┘
                │
                └─────────────────┐
                                  ▼
        ┌─────────────────────────────────────────┐
        │ // CÓDIGO NUEVO: Separado y testeable   │
        │                                         │
        │ function sendOrderNotification(order) { │
        │     recipient = order.customer.email    │
        │     subject = "Order Confirmation"      │
        │     body = formatOrderEmail(order)      │
        │     emailService.send(recipient,        │
        │                       subject, body)    │
        │ }                                       │
        │                                         │
        │ function formatOrderEmail(order) {      │
        │     // Lógica de formateo               │
        │     return template                     │
        │ }                                       │
        └─────────────────────────────────────────┘
                          │
                          ▼
        ┌─────────────────────────────────────────┐
        │ // TESTS: Solo para código nuevo        │
        │                                         │
        │ test "sends notification email" {       │
        │     mockEmailService = createMock()     │
        │     order = createTestOrder()           │
        │                                         │
        │     sendOrderNotification(order)        │
        │                                         │
        │     assert mockEmailService.wasCalled() │
        │     assert recipient == order.customer  │
        │ }                                       │
        └─────────────────────────────────────────┘
```

## Ejemplo

### ANTES: Código Legacy Sin el Cambio

```pseudocode
class OrderProcessor {
    function processOrder(order) {
        // 200 líneas de código legacy complicado
        // Sin tests, difícil de entender

        // Validar datos básicos
        if (order.items.isEmpty()) {
            throw new Error("Order must have items")
        }

        // Calcular total con lógica complicada
        total = 0
        for (item in order.items) {
            price = item.basePrice
            // 50 líneas de lógica de descuentos, impuestos, etc.
            if (item.isOnSale && order.customer.isPremium) {
                price = applyComplexDiscountLogic(item, order)
            }
            total += price * item.quantity
        }
        order.total = total

        // Guardar en base de datos
        database.save(order)

        // NECESITAMOS AÑADIR: Enviar notificación al cliente
        // Pero no queremos tocar todo este código legacy

        return order
    }

    // ... 300 líneas más de código complicado
}
```

**Problemas**:
- El código es complejo y sin tests
- Tocar este código es arriesgado
- Mezcla múltiples responsabilidades
- Añadir la notificación aquí complicaría más el código

### DESPUÉS: Con Sprout Change

```pseudocode
// ============================================
// CÓDIGO LEGACY: Modificación mínima
// ============================================
class OrderProcessor {
    function processOrder(order) {
        // 200 líneas de código legacy complicado
        // (SIN CAMBIOS - dejamos como está)

        if (order.items.isEmpty()) {
            throw new Error("Order must have items")
        }

        total = 0
        for (item in order.items) {
            price = item.basePrice
            if (item.isOnSale && order.customer.isPremium) {
                price = applyComplexDiscountLogic(item, order)
            }
            total += price * item.quantity
        }
        order.total = total

        database.save(order)

        // ✨ ÚNICA MODIFICACIÓN: Una línea
        notifyOrderProcessed(order)

        return order
    }
}


// ============================================
// CÓDIGO NUEVO: Separado, limpio y testeable
// ============================================

// Sprout Method: Nueva función aislada
function notifyOrderProcessed(order) {
    notification = createOrderNotification(order)
    sendNotification(notification)
}

function createOrderNotification(order) {
    return {
        recipient: order.customer.email,
        subject: "Your Order #" + order.id + " is Confirmed",
        body: formatOrderConfirmationEmail(order),
        template: "order_confirmation",
        metadata: {
            orderId: order.id,
            customerId: order.customer.id,
            total: order.total
        }
    }
}

function formatOrderConfirmationEmail(order) {
    template = loadTemplate("order_confirmation")

    data = {
        customerName: order.customer.name,
        orderId: order.id,
        items: formatOrderItems(order.items),
        total: formatCurrency(order.total),
        estimatedDelivery: calculateDeliveryDate(order)
    }

    return template.render(data)
}

function formatOrderItems(items) {
    return items.map(item => {
        return {
            name: item.name,
            quantity: item.quantity,
            price: formatCurrency(item.price * item.quantity)
        }
    })
}

function sendNotification(notification) {
    emailService = getEmailService()
    emailService.send(
        to: notification.recipient,
        subject: notification.subject,
        body: notification.body,
        template: notification.template
    )
}


// ============================================
// TESTS: Solo para el código nuevo
// ============================================

test "creates order notification with correct data" {
    order = createTestOrder({
        id: 123,
        customer: { email: "customer@example.com", name: "John" },
        items: [{ name: "Book", price: 10, quantity: 2 }],
        total: 20
    })

    notification = createOrderNotification(order)

    assert notification.recipient == "customer@example.com"
    assert notification.subject.contains("123")
    assert notification.metadata.orderId == 123
}

test "formats order items correctly" {
    items = [
        { name: "Book", price: 10, quantity: 2 },
        { name: "Pen", price: 1, quantity: 5 }
    ]

    formatted = formatOrderItems(items)

    assert formatted.length == 2
    assert formatted[0].name == "Book"
    assert formatted[0].quantity == 2
}

test "sends notification via email service" {
    mockEmailService = createMockEmailService()
    injectDependency(mockEmailService)

    notification = {
        recipient: "test@example.com",
        subject: "Test",
        body: "Test body",
        template: "order_confirmation"
    }

    sendNotification(notification)

    assert mockEmailService.sendWasCalledWith({
        to: "test@example.com",
        subject: "Test",
        body: "Test body"
    })
}

test "notifyOrderProcessed integrates correctly" {
    mockEmailService = createMockEmailService()
    injectDependency(mockEmailService)

    order = createTestOrder({ id: 123 })

    notifyOrderProcessed(order)

    assert mockEmailService.sendWasCalled()
    call = mockEmailService.getLastCall()
    assert call.to == order.customer.email
    assert call.subject.contains("123")
}
```

**Mejoras**:
- Código nuevo está separado y es limpio
- Es fácil de testear independientemente
- El código legacy no se tocó (casi)
- Riesgo minimizado (solo una línea modificada)
- Funcionalidad nueva es mantenible

## Proceso Paso a Paso

### Paso 1: Identificar el Punto de Inserción

Encuentra exactamente dónde en el código legacy necesitas añadir la funcionalidad nueva.

```pseudocode
function legacyProcess() {
    step1()
    step2()
    step3()  // ← Después de este paso necesito nueva funcionalidad
    step4()
}
```

**Preguntas clave**:
- ¿Dónde exactamente debe ejecutarse el nuevo código?
- ¿Qué datos necesita el nuevo código?
- ¿Puede fallar el nuevo código sin afectar al flujo principal?
- ¿Hay efectos secundarios del nuevo código?

### Paso 2: Determinar los Datos Necesarios

Identifica qué información del código legacy necesita pasar al código nuevo.

```pseudocode
// ¿Qué datos están disponibles en el punto de inserción?
function legacyProcess(order, user, context) {
    // ... código ...

    // En este punto tenemos disponibles:
    // - order (con order.total ya calculado)
    // - user
    // - context (con context.timestamp)

    // ¿Qué necesitamos pasar al nuevo código?
    // Solo 'order' parece suficiente
}
```

**Tip**: Intenta minimizar los datos que pasas. Si necesitas muchos parámetros, considera crear un objeto de datos específico.

### Paso 3: Diseñar la Interfaz del Código Nuevo

Decide la firma de la nueva función/método/clase.

```pseudocode
// Opción 1: Función simple
function sendNotification(order) { ... }

// Opción 2: Clase si es más complejo
class OrderNotifier {
    function notify(order) { ... }
}

// Opción 3: Método en clase nueva
class NotificationService {
    function notifyOrderProcessed(order) { ... }
}
```

**Criterios de diseño**:
- **Simple**: Una función si es lógica simple
- **Clase**: Si necesitas mantener estado o configuración
- **Servicio**: Si la funcionalidad puede crecer o necesita dependencias

### Paso 4: Implementar el Código Nuevo

Escribe el código nuevo de forma limpia, siguiendo buenas prácticas.

```pseudocode
// Implementación clara y testeable
function sendNotification(order) {
    // Validar input
    if (order == null || order.customer == null) {
        throw new Error("Invalid order")
    }

    // Separar responsabilidades
    emailData = buildEmailData(order)
    deliveryResult = deliverEmail(emailData)

    // Manejo de errores
    if (!deliveryResult.success) {
        logError("Failed to send notification", deliveryResult.error)
        // Decidir si lanzar excepción o solo loggear
    }

    return deliveryResult
}

function buildEmailData(order) {
    // Lógica pura de construcción de datos
    return {
        to: order.customer.email,
        subject: buildSubject(order),
        body: buildBody(order)
    }
}

function deliverEmail(emailData) {
    // Interacción con servicio externo
    return emailService.send(emailData)
}
```

**Principios**:
- Código claro y legible
- Funciones pequeñas y enfocadas
- Fácil de testear (inyección de dependencias)
- Manejo explícito de errores

### Paso 5: Escribir Tests para el Código Nuevo

Antes de integrarlo, testea el código nuevo exhaustivamente.

```pseudocode
describe "sendNotification" {
    test "sends email with correct recipient" {
        order = createTestOrder()

        sendNotification(order)

        assert lastEmailSent().to == order.customer.email
    }

    test "throws error for null order" {
        assertThrows(() => sendNotification(null))
    }

    test "throws error for order without customer" {
        order = { customer: null }

        assertThrows(() => sendNotification(order))
    }

    test "logs error if email fails to send" {
        mockEmailService.setNextResultToFail()
        order = createTestOrder()

        sendNotification(order)

        assert errorWasLogged()
    }

    test "includes order details in email body" {
        order = createTestOrder({ id: 123, total: 99.99 })

        sendNotification(order)

        emailBody = lastEmailSent().body
        assert emailBody.contains("123")
        assert emailBody.contains("99.99")
    }
}
```

**Cobertura importante**:
- Casos felices (happy path)
- Casos de error
- Casos límite (boundary cases)
- Integración con dependencias

### Paso 6: Insertar la Llamada en el Código Legacy

Ahora, con confianza, añade la llamada mínima en el código viejo.

```pseudocode
function legacyProcess(order) {
    step1()
    step2()
    step3()

    // ✨ INSERCIÓN: Una línea
    sendNotification(order)

    step4()
}
```

**Consideraciones**:
- **Manejo de errores**: ¿Debe propagarse la excepción o solo loggearse?
- **Transacción**: ¿Debe ejecutarse dentro de la misma transacción?
- **Sincronía**: ¿Debe ser síncrono o puede ser asíncrono?

```pseudocode
// Opción 1: Sincrónico con propagación de errores
sendNotification(order)  // Si falla, falla todo

// Opción 2: Sincrónico sin propagación
try {
    sendNotification(order)
} catch (error) {
    logError("Notification failed", error)
    // Continúa el flujo normal
}

// Opción 3: Asíncrono
scheduleAsyncTask(() => sendNotification(order))
```

### Paso 7: Testear la Integración

Verifica que el código nuevo se ejecuta correctamente desde el código legacy.

```pseudocode
test "processOrder sends notification" {
    mockEmailService = createMockEmailService()
    order = createTestOrder()

    // Ejecutar el flujo completo (legacy + nuevo)
    processOrder(order)

    // Verificar que se llamó al código nuevo
    assert mockEmailService.sendWasCalled()
}

test "processOrder continues if notification fails" {
    mockEmailService = createMockThatFails()
    order = createTestOrder()

    // No debe lanzar excepción
    result = processOrder(order)

    // Debe completarse el procesamiento
    assert result.status == "processed"
    assert orderWasSaved(order)
}
```

### Paso 8: Documentar y Limpiar

Documenta qué hiciste y por qué, y limpia si es necesario.

```pseudocode
// Legacy code (legado - no modificar innecesariamente)
function processOrder(order) {
    step1()
    step2()
    step3()

    // SPROUT: Notificación añadida el 2024-01-15
    // TODO: Refactorizar processOrder para separar responsabilidades
    sendNotification(order)

    step4()
}


// ====================================
// Order Notification Module
// ====================================
// Añadido: 2024-01-15
// Autor: Tu Nombre
// Propósito: Enviar notificación por email cuando se procesa una orden
//
// Este código fue añadido usando Sprout Change para evitar
// modificar el código legacy complejo de OrderProcessor.
//
// TODO Futuro: Cuando OrderProcessor sea refactorizado,
// integrar esto como parte de un pipeline de eventos.
// ====================================

function sendNotification(order) {
    // ...
}
```

## Problemas Comunes

### 1. Demasiadas Dependencias del Código Legacy

**Problema**: El código nuevo necesita 10 parámetros del código legacy.

```pseudocode
// Mal: Firma compleja
function newFeature(param1, param2, param3, param4, param5, param6) {
    // ...
}

// Llamada fea
newFeature(a, b, c, d, e, f)
```

**Solución**: Crear un objeto de contexto o pasar el objeto completo.

```pseudocode
// Bien: Objeto de contexto
class OperationContext {
    order, user, settings, timestamp, ...
}

function newFeature(context) {
    // Extraer lo que necesites
    order = context.order
    user = context.user
}

// O aún mejor: pasar solo el objeto principal
function newFeature(order) {
    // Acceder a través del orden
    user = order.customer
    settings = order.settings
}
```

### 2. El Código Nuevo Necesita Modificar Estado del Legacy

**Problema**: El código nuevo necesita cambiar variables del código legacy.

```pseudocode
// Mal: No puedes modificar variables locales del legacy
function legacyCode() {
    status = "pending"
    processStuff()

    newFeature()  // ¿Cómo modifico 'status' desde aquí?
}
```

**Solución**: Devolver valores o pasar objetos mutables.

```pseudocode
// Opción 1: Devolver nuevo estado
function legacyCode() {
    status = "pending"
    processStuff()

    status = newFeature(status)  // Retorna nuevo estado
}

// Opción 2: Pasar objeto mutable
function legacyCode() {
    state = { status: "pending" }
    processStuff()

    newFeature(state)  // Modifica state.status internamente
}
```

### 3. Punto de Inserción No Claro

**Problema**: No hay un lugar obvio donde insertar la llamada.

```pseudocode
// El código está todo enredado
function messyCode() {
    doSomethingA()
    x = calculateX()
    doSomethingB()
    y = calculateY(x)
    doSomethingC()
    // ¿Dónde inserto mi código?
}
```

**Solución**: Primero hacer un mini-refactor para crear el punto de inserción.

```pseudocode
// Paso 1: Extraer sección como método (sin cambiar comportamiento)
function messyCode() {
    doSomethingA()
    x = calculateX()
    doSomethingB()
    processCalculations(x)
}

function processCalculations(x) {
    y = calculateY(x)
    doSomethingC()

    // Ahora hay un punto claro
    newFeature(x, y)
}
```

### 4. El Código Nuevo También se Vuelve Legacy

**Problema**: El código "nuevo" también se complica con el tiempo.

**Solución**: Aplicar Sprout Change recursivamente.

```pseudocode
// Primera iteración: Sprout simple
function feature() {
    doStuff()
}

// Después necesitas añadir más
function feature() {
    doStuff()

    // Segundo Sprout dentro del primero
    additionalFeature()  // ← Nuevo brote desde el brote anterior
}

function additionalFeature() {
    // Código limpio y separado
}
```

### 5. Tests Frágiles por Acoplamiento

**Problema**: Los tests del código nuevo dependen demasiado del legacy.

```pseudocode
// Mal: Test acoplado al legacy
test "new feature works" {
    // Necesito crear todo el contexto legacy
    setupComplexLegacyState()
    legacyObject = createLegacyObject()
    legacyObject.setupMoreStuff()

    result = legacyObject.newFeature()

    assert result == expected
}
```

**Solución**: Testear el código nuevo de forma aislada.

```pseudocode
// Bien: Test independiente
test "new feature works" {
    // Crear solo lo mínimo necesario
    input = createSimpleInput()

    result = newFeature(input)  // Llamada directa

    assert result == expected
}

// Test de integración separado
test "new feature integrates with legacy" {
    // Este puede ser más complejo
    legacyObject = createLegacyObject()

    legacyObject.process()  // Internamente llama newFeature

    assert legacyObject.state == expected
}
```

## Criterios de Aceptación

Has aplicado Sprout Change correctamente cuando:

1. **Modificación mínima**: Solo añadiste 1-3 líneas al código legacy
2. **Código aislado**: El nuevo código está completamente separado
3. **Tests independientes**: Puedes testear el código nuevo sin el legacy
4. **Sin side effects**: No modificaste el comportamiento existente
5. **Código limpio**: El nuevo código sigue buenas prácticas
6. **Documentado**: Está claro qué es nuevo y qué es legacy
7. **Fácil de revertir**: Puedes quitar la funcionalidad fácilmente

## Beneficios

### Inmediatos
- Riesgo minimizado (casi no tocas código legacy)
- Funcionalidad nueva es testeable desde el inicio
- Desarrollo rápido (no necesitas entender todo el legacy)

### A Mediano Plazo
- Código nuevo es fácil de mantener y evolucionar
- Puedes refactorizar el código nuevo independientemente
- Creas "islas de código limpio" en un mar de legacy

### A Largo Plazo
- Gradualmente el sistema tiene más código limpio que legacy
- Facilitas refactorings futuros (puedes extraer el código nuevo primero)
- Mejoras la arquitectura sin reescribir todo

## Técnicas Relacionadas

- **Wrap Change**: Úsalo cuando necesites modificar comportamiento existente en vez de añadir nuevo
- **Golden Master**: Combínalo con Sprout para tener red de seguridad antes del cambio
- **Extract Method**: Usa refactorings seguros para crear el punto de inserción si no existe
- **Dependency Injection**: Hace el código sprouted más testeable
- **Expand-Migrate-Contract**: Evolución natural de múltiples Sprouts hacia refactor completo

## Versiones por Lenguaje

- [TypeScript](../../../../typescript/src/refactoring/parallel-change/sprout-change/) - [README](../../../../typescript/src/refactoring/parallel-change/sprout-change/README.md)
- [Go](../../../../go/refactoring/parallel-change/sprout-change/) - [README](../../../../go/refactoring/parallel-change/sprout-change/README.md)
- [Java](../../../../java/src/main/java/com/refactoring/refactoring/parallel_change/sprout_change/) - [README](../../../../java/src/main/java/com/refactoring/refactoring/parallel_change/sprout_change/README.md)
- [PHP](../../../../php/src/refactoring/parallel-change/sprout-change/) - [README](../../../../php/src/refactoring/parallel-change/sprout-change/README.md)
- [Python](../../../../python/src/refactoring/parallel_change/sprout_change/) - [README](../../../../python/src/refactoring/parallel_change/sprout_change/README.md)
- [C#](../../../../csharp/src/refactoring/parallel-change/sprout-change/) - [README](../../../../csharp/src/refactoring/parallel-change/sprout-change/README.md)

## Referencias en Español

### Artículos de Fran Iglesias

- [Introducción al Refactor](https://franiglesias.github.io/intro_refactor_1/) - Explicación detallada de Sprout y Wrap
- [Refactoring - Camp Rule](https://franiglesias.github.io/refactoring-camp-rule/) - Mejora incremental con Sprout
- [Modernizando el legacy](https://franiglesias.github.io/modernizando-el-legacy/) - Estrategias incluyendo Sprout

## Referencias en Inglés

### Libros
- **Working Effectively with Legacy Code** - Michael Feathers (2004)
  - Capítulo 6: "I Don't Have Much Time and I Have to Change It"
  - Sección "Sprout Method" y "Sprout Class"

### Artículos
- [Sprout Method](https://refactoring.guru/smells/refused-bequest) - Refactoring Guru
- [Legacy Code Techniques](https://www.sitepoint.com/working-effectively-legacy-code/) - Practical examples

---

**Próxima Técnica**: Si Sprout Change no es suficiente porque necesitas modificar comportamiento existente (no solo añadir), considera [Wrap Change](./wrap-change.md).
