# Golden Master (Approval Testing)

## Definición

**Golden Master** (también conocido como Approval Testing, Characterization Testing o Snapshot Testing) es una técnica de testing que captura el comportamiento actual de un sistema como "la verdad" (el master), y lo usa como referencia para detectar automáticamente cualquier cambio en el comportamiento futuro.

En lugar de escribir assertions específicas sobre lo que el código debería hacer, capturamos lo que hace actualmente y verificamos que siga haciéndolo igual después de los cambios.

## Cuándo Usar

- Cuando heredas código sin tests y necesitas crear una red de seguridad rápidamente
- Cuando no sabes exactamente qué hace el código pero funciona en producción
- Antes de empezar un refactoring grande en código legacy
- Cuando el comportamiento del sistema es complejo y difícil de especificar
- Para documentar el comportamiento actual de un sistema
- Como paso previo a escribir tests unitarios más específicos
- Cuando tienes código con múltiples outputs complejos (HTML, JSON, XML, reportes, etc.)

## Problema que Resuelve

El mayor obstáculo para refactorizar código legacy es la falta de tests. Sin tests, no puedes saber si tus cambios rompieron algo. Pero escribir tests para código complejo y mal diseñado es:

- **Lento**: Puede tomar días o semanas
- **Difícil**: El código no es testeable sin refactorizar
- **Paradójico**: Necesitas refactorizar para poder testear, pero necesitas tests para refactorizar con seguridad
- **Intimidante**: No sabes por dónde empezar

Golden Master resuelve esta paradoja permitiéndote:
1. Crear tests automáticamente sin entender todo el código
2. Tener una red de seguridad en minutos u horas (no días)
3. Detectar cambios de comportamiento inmediatamente
4. Refactorizar con confianza

## Descripción Detallada

### Cómo Funciona

El proceso es simple pero poderoso:

1. **Captura**: Ejecutas el código con múltiples inputs y guardas todos los outputs
2. **Aprobación**: Revisas que los outputs son correctos (comportamiento actual de producción)
3. **Referencia**: Ese output aprobado se convierte en el "Golden Master"
4. **Verificación**: Cada vez que cambias el código, comparas el nuevo output con el master
5. **Detección**: Cualquier diferencia es reportada como fallo
6. **Decisión**: Si el cambio es intencional, actualizas el master; si no, corriges el bug

### Diagrama de Flujo

```
┌─────────────────────────────────────────────────────────────┐
│ FASE 1: CAPTURA DEL COMPORTAMIENTO ACTUAL                  │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
        ┌──────────────────────────────────┐
        │  Ejecutar código con inputs      │
        │  representativos                 │
        └──────────────────────────────────┘
                            │
                            ▼
        ┌──────────────────────────────────┐
        │  Capturar outputs completos      │
        │  (archivos, stdout, logs, etc.)  │
        └──────────────────────────────────┘
                            │
                            ▼
        ┌──────────────────────────────────┐
        │  Guardar como Golden Master      │
        │  (archivo de referencia)         │
        └──────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ FASE 2: REFACTORING CON RED DE SEGURIDAD                   │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
        ┌──────────────────────────────────┐
        │  Hacer cambios en el código      │
        └──────────────────────────────────┘
                            │
                            ▼
        ┌──────────────────────────────────┐
        │  Ejecutar con mismos inputs      │
        └──────────────────────────────────┘
                            │
                            ▼
        ┌──────────────────────────────────┐
        │  Comparar output nuevo con       │
        │  Golden Master                   │
        └──────────────────────────────────┘
                            │
                ┌───────────┴───────────┐
                ▼                       ▼
        ┌──────────────┐        ┌──────────────┐
        │  Idénticos   │        │  Diferentes  │
        │  ✓ Test OK   │        │  ✗ Test FAIL │
        └──────────────┘        └──────────────┘
                                        │
                        ┌───────────────┴────────────────┐
                        ▼                                ▼
            ┌───────────────────┐          ┌─────────────────────┐
            │ Cambio intencional│          │ Bug introducido     │
            │ Actualizar master │          │ Corregir código     │
            └───────────────────┘          └─────────────────────┘
```

## Ejemplo

### ANTES: Código Legacy Sin Tests

```pseudocode
class InvoiceGenerator {
    function generateInvoice(order) {
        // 500 líneas de código complejo
        // Lógica de negocio mezclada con formato
        // Sin tests, sin documentación

        html = "<html><body>"
        html += "<h1>Invoice #" + order.id + "</h1>"

        // Cálculos complejos de impuestos, descuentos, etc.
        total = 0
        for item in order.items {
            price = item.price
            if (order.customer.isVIP) {
                price = price * 0.9  // 10% descuento
            }
            if (item.category == "BOOKS") {
                price = price * 1.04  // 4% tax
            } else {
                price = price * 1.21  // 21% tax
            }
            total += price * item.quantity
            html += "<tr><td>" + item.name + "</td>..."
        }

        html += "<p>Total: " + total + "</p>"
        html += "</body></html>"
        return html
    }
}

// ¿Cómo probamos esto? ¿Cómo refactorizamos con seguridad?
```

### DESPUÉS: Con Golden Master

```pseudocode
// TEST 1: Capturar comportamiento actual
test "golden master for invoice generation" {
    // Preparar casos representativos
    testCases = [
        createOrder(id: 1, customer: regularCustomer, items: [book, electronic]),
        createOrder(id: 2, customer: vipCustomer, items: [book]),
        createOrder(id: 3, customer: regularCustomer, items: [multipleBooks]),
        createOrder(id: 4, customer: vipCustomer, items: [mixedItems]),
        // Más casos que cubran edge cases
    ]

    generator = new InvoiceGenerator()

    for testCase in testCases {
        // Generar output
        result = generator.generateInvoice(testCase)

        // Comparar con golden master guardado
        approvalFile = "approved/invoice_" + testCase.id + ".html"

        if (not fileExists(approvalFile)) {
            // Primera vez: guardar para revisión manual
            saveForApproval(result, approvalFile)
            fail("New output captured. Please review and approve.")
        } else {
            // Comparar con master aprobado
            expected = readFile(approvalFile)
            if (result != expected) {
                // Mostrar diferencias
                showDiff(expected, result)
                fail("Output differs from approved golden master!")
            }
        }
    }
}

// Ahora podemos refactorizar con confianza:

class InvoiceGenerator {
    function generateInvoice(order) {
        // REFACTORIZADO: Lógica separada de presentación
        data = calculateInvoiceData(order)
        return formatAsHTML(data)
    }

    private function calculateInvoiceData(order) {
        // Lógica de negocio pura
        calculator = new PriceCalculator()
        items = []
        total = 0

        for item in order.items {
            price = calculator.calculateItemPrice(item, order.customer)
            items.add({
                name: item.name,
                price: price,
                quantity: item.quantity,
                subtotal: price * item.quantity
            })
            total += price * item.quantity
        }

        return {
            orderId: order.id,
            items: items,
            total: total
        }
    }

    private function formatAsHTML(data) {
        // Formato puro
        renderer = new HTMLRenderer()
        return renderer.renderInvoice(data)
    }
}

// El test de Golden Master sigue pasando
// El output HTML es idéntico
// Pero el código es mucho mejor
```

## Proceso Paso a Paso

### Paso 1: Identificar el Punto de Entrada y Output Observable

Encuentra dónde ejecutar el código y qué capturar:
- **Punto de entrada**: Función, método o endpoint que quieres testear
- **Output observable**: Resultado que puedes capturar (archivo, string, JSON, stdout, logs)
- Si no hay output claro, puede que necesites refactorizar primero para extraer uno (Sprout)

**Ejemplo**:
```pseudocode
// Punto de entrada
result = generateReport(input_data)

// Output observable
file_written = "report.pdf"
or
string_returned = result
or
console_output = captured_stdout
```

### Paso 2: Crear Casos de Test Representativos

Identifica inputs que ejerciten diferentes caminos del código:
- **Casos típicos**: Los que ocurren el 80% del tiempo en producción
- **Edge cases**: Valores límite, listas vacías, valores máximos
- **Casos especiales**: Condiciones de negocio específicas
- **Casos históricos**: Si tienes logs de producción, úsalos

**Tip**: No necesitas cobertura del 100% inicialmente. Empieza con 3-5 casos clave y añade más según necesites.

```pseudocode
testInputs = [
    { type: "TYPICAL", data: {...} },        // Caso normal
    { type: "EMPTY", data: {...} },          // Lista vacía
    { type: "BOUNDARY", data: {...} },       // Valor límite
    { type: "COMPLEX", data: {...} },        // Caso complejo real
    { type: "EDGE_CASE", data: {...} },      // Condición rara
]
```

### Paso 3: Ejecutar el Código y Capturar el Output

Ejecuta el código con tus inputs de test y guarda los outputs:

```pseudocode
for testCase in testInputs {
    // Ejecutar
    result = systemUnderTest.execute(testCase.data)

    // Capturar output completo
    outputFile = "golden_masters/test_" + testCase.type + ".approved"

    // Serializar si es necesario
    if (result is Object) {
        content = toJSON(result)  // o toXML, toString, etc.
    } else {
        content = result
    }

    // Guardar
    writeFile(outputFile, content)
}
```

### Paso 4: Revisar y Aprobar el Golden Master

**CRÍTICO**: No asumas que el output capturado es correcto automáticamente.

1. **Revisar cada archivo capturado manualmente**
2. **Verificar contra producción** si es posible
3. **Consultar con expertos de negocio** si hay dudas
4. **Documentar comportamientos extraños** que encuentres
5. **Solo aprobar cuando estés seguro** que refleja el comportamiento correcto actual

```pseudocode
// Revisión manual
for approvedFile in directory("golden_masters/*.approved") {
    print("Reviewing: " + approvedFile)
    content = readFile(approvedFile)

    // Mostrar de forma legible
    displayForReview(content)

    // Decisión humana
    answer = ask("Is this output correct? (yes/no/fix)")

    if (answer == "no") {
        // Investigar por qué el comportamiento actual es incorrecto
        // Puede que necesites fix el bug antes de capturar el master
    } else if (answer == "fix") {
        // Editar manualmente el output esperado
        editedContent = manualEdit(content)
        writeFile(approvedFile, editedContent)
    }
    // else: aprobado implícitamente
}
```

### Paso 5: Crear el Test Automatizado

Ahora escribe el test que compare outputs futuros con el master:

```pseudocode
test "golden master regression test" {
    for testCase in testInputs {
        // 1. Ejecutar código actual
        actualOutput = systemUnderTest.execute(testCase.data)

        // 2. Cargar golden master
        masterFile = "golden_masters/test_" + testCase.type + ".approved"
        expectedOutput = readFile(masterFile)

        // 3. Normalizar si es necesario (timestamps, IDs aleatorios, etc.)
        actualNormalized = normalize(actualOutput)
        expectedNormalized = normalize(expectedOutput)

        // 4. Comparar
        if (actualNormalized != expectedNormalized) {
            // 5. Mostrar diferencias útiles
            diff = computeDiff(expectedNormalized, actualNormalized)
            saveDiffReport("diffs/test_" + testCase.type + ".diff", diff)

            fail(
                "Golden Master mismatch for " + testCase.type + "\n" +
                "See diff at: diffs/test_" + testCase.type + ".diff"
            )
        }
    }
}
```

### Paso 6: Normalizar Datos No Deterministas

Algunos outputs contienen datos que cambian en cada ejecución y deben normalizarse:

```pseudocode
function normalize(content) {
    result = content

    // Reemplazar timestamps
    result = replaceRegex(result, /\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}/, "{{TIMESTAMP}}")

    // Reemplazar UUIDs
    result = replaceRegex(result, /[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/, "{{UUID}}")

    // Reemplazar IDs autoincrementales
    result = replaceRegex(result, /"id":\s*\d+/, '"id": {{ID}}')

    // Normalizar espacios en blanco
    result = normalizeWhitespace(result)

    // Ordenar arrays si el orden no importa
    if (isJSON(result)) {
        obj = parseJSON(result)
        sortArraysRecursively(obj)
        result = toJSON(obj)
    }

    return result
}
```

### Paso 7: Ejecutar y Validar

Ejecuta el test varias veces para confirmar que es estable:

```bash
# Debe pasar consistentemente
run_tests golden_master_test
run_tests golden_master_test  # Segunda vez
run_tests golden_master_test  # Tercera vez

# Si alguna vez falla, investiga:
# - ¿Hay algo no determinista sin normalizar?
# - ¿Hay dependencias externas (DB, API, tiempo)?
# - ¿El código tiene efectos secundarios indeseados?
```

### Paso 8: Refactorizar con Confianza

Ahora que tienes la red de seguridad:

1. **Hacer cambio pequeño** en el código (renombrar variable, extraer método, etc.)
2. **Ejecutar golden master test** inmediatamente
3. **Si pasa**: El comportamiento se mantiene, continúa
4. **Si falla**: Revisar si el cambio es intencional o un bug
5. **Repetir** con el siguiente cambio pequeño

```pseudocode
// Ciclo de refactoring seguro
while (code needs improvement) {
    // 1. Identificar smell
    smell = identifyNextSmell(code)

    // 2. Aplicar refactoring pequeño
    refactor(smell)

    // 3. Verificar con golden master
    run_tests(golden_master_test)

    // 4. Si falla inesperadamente, revertir
    if (test_failed && not_intentional_change) {
        revert_changes()
        investigate()
    }

    // 5. Commit de cambio seguro
    if (test_passes) {
        commit("Refactor: " + description)
    }
}
```

## Problemas Comunes

### 1. Output No Determinista

**Problema**: El test falla aleatoriamente porque el output contiene timestamps, IDs aleatorios, o datos dependientes del momento de ejecución.

**Solución**:
- Normalizar datos variables con funciones de limpieza
- Usar mocks para fuentes de no-determinismo (reloj, random)
- Inyectar dependencias controlables (generadores de IDs fijos)

```pseudocode
// Mal: timestamp real
timestamp = Date.now()

// Bien: timestamp inyectado
class SystemUnderTest {
    constructor(clockProvider) {
        this.clock = clockProvider
    }

    process() {
        timestamp = this.clock.now()  // Controlable en tests
    }
}

// En test
fixedClock = { now: () => "2024-01-01T00:00:00Z" }
system = new SystemUnderTest(fixedClock)
```

### 2. Outputs Demasiado Grandes

**Problema**: El golden master ocupa megabytes o es difícil de revisar.

**Solución**:
- Dividir en múltiples tests más pequeños
- Capturar solo las partes relevantes del output
- Usar resúmenes o hashes para outputs muy grandes

```pseudocode
// Mal: capturar reporte completo de 10MB
result = generateHugeReport()

// Bien: capturar secciones específicas
result = {
    summary: generateHugeReport().getSummary(),
    firstPage: generateHugeReport().getPage(1),
    totals: generateHugeReport().getTotals(),
    // Solo lo esencial
}
```

### 3. Demasiados Casos de Test

**Problema**: Crear golden masters para 100 casos diferentes es inmanejable.

**Solución**:
- Empezar con pocos casos representativos (3-5)
- Añadir casos solo cuando encuentres bugs
- Usar técnicas de partición de equivalencia
- Priorizar casos reales de producción

### 4. Comportamiento Actual es Incorrecto

**Problema**: El golden master captura un bug que existe en producción.

**Solución**:
- **Opción A**: Capturar el bug en el master, refactorizar, luego arreglar el bug con nuevo test
- **Opción B**: Arreglar el bug primero, luego capturar el comportamiento correcto
- **Documentar** el bug conocido en el test

```pseudocode
test "golden master - includes known bug #1234" {
    // NOTA: Este test captura el bug conocido en ticket #1234
    // donde las fechas en febrero están mal calculadas.
    // Una vez refactorizado, arreglaremos el bug y actualizaremos el master.

    result = systemUnderTest.execute(input)
    assertMatchesGoldenMaster(result, "master_with_known_bug.approved")
}
```

### 5. Diferencias en Formato pero No en Contenido

**Problema**: El test falla porque cambió el formato (espacios, orden) pero el contenido es equivalente.

**Solución**:
- Comparar estructuras parseadas en vez de strings
- Normalizar formatos antes de comparar
- Usar comparadores semánticos

```pseudocode
// Mal: comparar strings de JSON
expected = '{"name":"John","age":30}'
actual = '{"age":30,"name":"John"}'  // Orden diferente, falla

// Bien: comparar objetos
expectedObj = parseJSON(expected)
actualObj = parseJSON(actual)
assertDeepEquals(expectedObj, actualObj)  // Pasa
```

### 6. Test Difícil de Mantener

**Problema**: Cada pequeño cambio intencional rompe decenas de golden masters.

**Solución**:
- Capturar en nivel de abstracción correcto (no muy bajo)
- Usar formato estable (JSON mejor que HTML)
- Actualizar masters en lote cuando sea necesario
- Considerar hacer el golden master más específico (testar menos cosas por archivo)

## Criterios de Aceptación

Sabes que has aplicado Golden Master correctamente cuando:

1. **Tests estables**: Pasan consistentemente sin cambios en el código
2. **Detectan cambios**: Cualquier modificación de comportamiento hace fallar el test
3. **Output verificado**: Has revisado manualmente que los masters son correctos
4. **Sin falsos positivos**: No fallan por razones aleatorias o no relacionadas
5. **Cobertura adecuada**: Los casos de test cubren los caminos principales del código
6. **Fácil de entender**: Otro desarrollador puede ver qué se está testeando
7. **Rápido**: Los tests ejecutan en segundos, no minutos
8. **Determinista**: No necesitas ejecutar varias veces para confirmar

## Beneficios

### Inmediatos
- Red de seguridad en horas, no días
- No necesitas entender todo el código para empezar
- Detecta regresiones automáticamente
- Documentación viva del comportamiento actual

### A Mediano Plazo
- Permite refactorizar código legacy con confianza
- Base para escribir tests unitarios más específicos después
- Identifica comportamientos inesperados o bugs
- Reduce el miedo a tocar código viejo

### A Largo Plazo
- Facilita la transición de legacy a código limpio
- Mejora la calidad del código progresivamente
- Reduce el tiempo de debugging
- Aumenta la velocidad de desarrollo

## Técnicas Relacionadas

- **Characterization Tests**: Mismo concepto, nombre usado por Michael Feathers
- **Snapshot Testing**: Variante usada en frameworks como Jest
- **Sprout Change**: Úsala junto con Golden Master para añadir funcionalidad nueva
- **Wrap Change**: Combínala con Golden Master para cambiar comportamiento existente
- **Expand-Migrate-Contract**: Usa Golden Master para validar la migración

## Versiones por Lenguaje

- [TypeScript](../../../typescript/src/refactoring/golden-master/) - [README](../../../typescript/src/refactoring/golden-master/README.md)
- [Go](../../../go/refactoring/golden-master/) - [README](../../../go/refactoring/golden-master/README.md)
- [Java](../../../java/src/main/java/com/refactoring/refactoring/golden_master/) - [README](../../../java/src/main/java/com/refactoring/refactoring/golden_master/README.md)
- [PHP](../../../php/src/refactoring/golden-master/) - [README](../../../php/src/refactoring/golden-master/README.md)
- [Python](../../../python/src/refactoring/golden_master/) - [README](../../../python/src/refactoring/golden_master/README.md)
- [C#](../../../csharp/src/refactoring/golden-master/) - [README](../../../csharp/src/refactoring/golden-master/README.md)

## Referencias en Español

### Artículos de Fran Iglesias

- [Approval Testing](https://franiglesias.github.io/approval_testing/) - Introducción completa a la técnica
- [Golden Master - Cookbook](https://franiglesias.github.io/golden-cookbook-master-approval/) - Receta práctica paso a paso
- [El código en producción es legacy](https://franiglesias.github.io/prod-code-is-legacy/) - Por qué necesitas esta técnica
- [Quotebot Kata](https://franiglesias.github.io/quotebot-kata/) - Ejercicio práctico guiado
- [Ejercicio de refactor](https://franiglesias.github.io/ejercicio-de-refactor-1/) - Caso de uso real

## Referencias en Inglés

### Libros
- **Working Effectively with Legacy Code** - Michael Feathers (2004)
  - Capítulo 13: "I Need to Make a Change, but I Don't Know What Tests to Write"
  - Introduce el concepto de Characterization Tests

### Artículos y Recursos
- [ApprovalTests.com](https://approvaltests.com/) - Framework y documentación oficial
- [Approval Testing](https://blog.thecodewhisperer.com/permalink/approval-testing) - J.B. Rainsberger
- [Characterization Testing](https://michaelfeathers.silvrback.com/characterization-testing) - Michael Feathers
- [Golden Master Testing](https://methodpoet.com/golden-master-testing/) - Emily Bache

### Videos
- [Approval Testing in Practice](https://www.youtube.com/watch?v=GJ5PxkN3-Vs) - Llewellyn Falco
- [Testing Legacy Code](https://www.youtube.com/watch?v=_NnElPO5BU0) - Sandro Mancuso
- [Surviving Legacy Code](https://www.youtube.com/watch?v=9OUz4yYM9SE) - J.B. Rainsberger

### Herramientas por Lenguaje
- **Java**: ApprovalTests.Java
- **JavaScript/TypeScript**: Jest Snapshots, jest-approval-tests
- **Python**: approvaltests, pytest-snapshot
- **C#**: ApprovalTests.Net
- **Go**: go-approval-tests
- **PHP**: PHPUnit Snapshot Assertions

---

**Próximos Pasos**: Una vez que tengas tu Golden Master funcionando, puedes proceder con técnicas de Parallel Change para refactorizar de forma segura:
- [Parallel Change - Introducción](../parallel-change/README.md)
- [Sprout Change](../parallel-change/sprout-change.md)
- [Wrap Change](../parallel-change/wrap-change.md)
