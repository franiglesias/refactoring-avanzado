# Long Method

Método largo.

## Definición

Un método en una clase es muy largo, conteniendo demasiadas líneas de código y realizando múltiples operaciones que deberían estar separadas en métodos más pequeños y cohesivos.

## Descripción

Un **Long Method** es un método que hace demasiado. Aunque no hay un número mágico de líneas que defina un método largo, generalmente cualquier método con más de 10-20 líneas merece ser examinado. El problema real no es la longitud per se, sino la cantidad de responsabilidades que acumula.

Los métodos largos son difíciles de:
- Entender completamente
- Modificar sin introducir bugs
- Testear exhaustivamente
- Reutilizar parcialmente
- Nombrar apropiadamente

El código tiende a crecer cuando es más fácil añadir líneas a un método existente que crear una nueva abstracción. Con el tiempo, estos métodos se vuelven "procedimientos gigantes" que requieren mantener mucho contexto mental.

## Síntomas

- El método tiene más de 20-30 líneas de código
- Necesitas scroll para ver el método completo
- El método tiene múltiples niveles de indentación (más de 3-4)
- Bloques de código separados por comentarios que explican "secciones"
- Variables temporales que solo se usan en una parte del método
- Dificultad para nombrar el método sin usar "and" o "or"
- Tests que requieren múltiples asserts para verificar diferentes aspectos

## Ejemplo

```pseudocode
function processOrder(order) {
  // Validar el pedido
  if (order.items is empty) {
    print "El pedido no tiene productos"
    return
  }

  for each item in order.items {
    if (item.price < 0 or item.quantity <= 0) {
      print "Producto inválido"
      return
    }
  }

  // Calcular subtotal
  subtotal = 0
  for each item in order.items {
    subtotal = subtotal + item.price * item.quantity
  }

  // Aplicar descuento VIP
  discount = 0
  if (order.customerType == "VIP") {
    discount = subtotal * 0.1
  }

  // Calcular impuestos
  taxableAmount = subtotal - discount
  tax = taxableAmount * 0.21

  // Calcular envío
  shipping = taxableAmount >= 50 ? 0 : 5

  // Calcular total
  total = taxableAmount + tax + shipping

  // Actualizar pedido
  order.subtotal = subtotal
  order.discount = discount
  order.tax = tax
  order.shipping = shipping
  order.total = total

  // Guardar en base de datos (muchas líneas)
  connectionString = "Server=fake.db.local;..."
  record = createDatabaseRecord(order)
  serialized = toJson(record)
  saveToDatabase(serialized)

  // Enviar email de confirmación (muchas líneas)
  emailTemplate = buildEmailTemplate(order)
  emailBody = buildHtmlBody(order)
  attachments = buildAttachments(order)
  sendEmail(order.customerEmail, emailBody, attachments)

  // Imprimir recibo (muchas líneas)
  receipt = buildReceipt(order)
  printToThermalPrinter(receipt)
}
```

## Ejercicio

Añade soporte de cupones con expiración y multi-moneda (USD/EUR) con reglas de redondeo distintas.

## Problemas que encontrarás

Tienes que tocar diferentes secciones dentro del método, lo que genera riesgo de cambios indeseados y aumenta el esfuerzo de mantenimiento. Cada modificación requiere entender todo el contexto del método completo.

## Proceso de Refactoring

### 1. Identificar bloques lógicos
- Lee el método completo y marca las diferentes "secciones"
- Identifica bloques de código que realizan una tarea cohesiva
- Busca comentarios que separan secciones (suelen indicar candidatos para extracción)

### 2. Extraer métodos auxiliares
- Comienza con los bloques más independientes
- Crea un nuevo método con un nombre descriptivo
- Mueve el código del bloque al nuevo método
- Identifica qué variables necesita como parámetros y qué retorna

### 3. Simplificar condicionales complejos
- Extrae condiciones complejas a métodos con nombres significativos
- Ejemplo: `if (item.price < 0 or item.quantity <= 0)` → `if (isInvalidItem(item))`
- Esto hace el flujo principal más legible

### 4. Eliminar variables temporales innecesarias
- Reemplaza variables temporales con llamadas a métodos
- Si una variable se usa solo una vez, considera eliminarla
- Usa **Replace Temp with Query** cuando sea apropiado

### 5. Aplicar **Compose Method**
- El método principal debe leerse como un resumen de pasos de alto nivel
- Cada paso debe ser una llamada a un método bien nombrado
- Objetivo: método principal de 5-10 líneas que cuenta una historia clara

### 6. Mover comportamiento a clases apropiadas
- Si los métodos extraídos usan muchos datos de otro objeto, muévelos allí
- Ejemplo: cálculos de impuestos pueden ir a una clase `TaxCalculator`
- Sigue el principio **Tell, Don't Ask**

## Técnicas de Refactoring Aplicables

- **Extract Method**: Extraer bloques de código a métodos separados
- **Replace Temp with Query**: Eliminar variables temporales usando métodos
- **Introduce Parameter Object**: Si el método tiene muchos parámetros relacionados
- **Decompose Conditional**: Extraer condiciones complejas a métodos
- **Replace Method with Method Object**: Para métodos muy complejos con muchas variables locales
- **Move Method**: Mover comportamiento a la clase apropiada

## Beneficios

- **Legibilidad mejorada**: Código que se lee como prosa, no como código de máquina
- **Reutilización**: Métodos pequeños pueden ser reutilizados en otros contextos
- **Testing más fácil**: Métodos pequeños son más fáciles de testear en aislamiento
- **Menos bugs**: Métodos simples tienen menos lugares donde esconder errores
- **Mantenimiento simplificado**: Cambios localizados en métodos específicos
- **Mejor navegación**: Nombres descriptivos facilitan encontrar el código relevante

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/bloaters/long-method.ts) - [README](../../typescript/src/code-smells/bloaters/long-method.readme.md)
- [Go](../../go/code_smells/bloaters/long_method.go) - [README](../../go/code_smells/bloaters/long_method.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/bloaters/LongMethod.java) - [README](../../java/src/main/java/com/refactoring/codesmells/bloaters/LongMethod.readme.md)
- [PHP](../../php/src/code-smells/bloaters/LongMethod.php) - [README](../../php/src/code-smells/bloaters/LongMethod.readme.md)
- [Python](../../python/src/code_smells/bloaters/long_method.py) - [README](../../python/src/code_smells/bloaters/long_method_readme.md)
- [C#](../../csharp/src/code-smells/bloaters/LongMethod.cs) - [README](../../csharp/src/code-smells/bloaters/long-method.readme.md)

## Referencias en Español

- [Métodos largos](https://franiglesias.github.io/long-method/) - Análisis detallado del problema de métodos largos y cómo refactorizarlos
- [Ejercicio de refactor (2) Extraer hasta la última gota](https://franiglesias.github.io/ejercicio-de-refactor-2/) - Ejercicio práctico de extracción de métodos
- [Object Calisthenics para adelgazar las clases](https://franiglesias.github.io/calistenics-and-small-classes/) - Incluye técnicas para mantener métodos cortos y cohesivos

## Referencias

- [Refactoring Guru - Long Method](https://refactoring.guru/smells/long-method)
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
- Kent Beck - "Smalltalk Best Practice Patterns" - Compose Method pattern
