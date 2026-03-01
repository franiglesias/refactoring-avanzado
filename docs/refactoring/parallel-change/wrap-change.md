# Wrap Change (Técnica del Envoltorio)

## Definición

**Wrap Change** (Cambio por Envoltorio o Cambio por Envoltura) es una técnica de refactoring que consiste en envolver código legacy existente con una nueva capa que añade, modifica o intercepta el comportamiento, sin necesidad de modificar el código original.

Es similar al patrón de diseño Decorator o Proxy: el código legacy queda intacto dentro de un "envoltorio" que controla cómo se usa.

## Cuándo Usar

- Necesitas modificar o extender el comportamiento de código legacy
- No puedes o no quieres modificar el código original (riesgo, permisos, o complejidad)
- Quieres añadir aspectos transversales (logging, validación, caching, etc.)
- Necesitas interceptar llamadas para cambiar entradas o salidas
- El código legacy no es testeable pero puedes hacer testeable el wrapper
- Quieres aplicar patrones Decorator, Proxy, o Adapter
- Necesitas aislar dependencias problemáticas

## Problema que Resuelve

Cuando necesitas cambiar comportamiento de código legacy, te enfrentas a:

- **Código frágil**: Tocar código viejo puede romper funcionalidad existente
- **Sin tests**: No hay red de seguridad para validar cambios
- **Alta complejidad**: El código es difícil de entender y modificar
- **Código bloqueado**: No puedes modificarlo (librería externa, código compartido, etc.)
- **Múltiples responsabilidades**: El código hace demasiadas cosas

Wrap Change resuelve esto:
1. Dejando el código legacy completamente intacto
2. Añadiendo comportamiento nuevo en una capa separada
3. Permitiendo testear el nuevo comportamiento aisladamente
4. Facilitando reversión (solo quitar el wrapper)

## Descripción Detallada

### Cómo Funciona

Wrap Change sigue este patrón:

1. **Identifica** el código legacy que necesitas cambiar
2. **Crea** un wrapper (clase o función) que encapsula el código legacy
3. **Añade** el comportamiento nuevo en el wrapper (antes, después, o alrededor)
4. **Reemplaza** las llamadas al código legacy por llamadas al wrapper
5. **Testea** el wrapper de forma aislada

### Tipos de Wrap

#### 1. Wrap Method (Envolver Método)

Creas una función nueva que llama a la función legacy.

```pseudocode
// Wrapper intercepta y delega
function newMethod(input) {
    // Comportamiento añadido ANTES
    preprocessedInput = preprocess(input)

    // Delegación al método legacy
    result = legacyMethod(preprocessedInput)

    // Comportamiento añadido DESPUÉS
    return postprocess(result)
}
```

#### 2. Wrap Class (Envolver Clase)

Creas una clase nueva que contiene la clase legacy.

```pseudocode
// Wrapper contiene instancia legacy
class NewClass {
    private legacyInstance

    constructor() {
        this.legacyInstance = new LegacyClass()
    }

    method(input) {
        // Nuevo comportamiento
        log("Calling method with: " + input)

        // Delegación
        result = this.legacyInstance.method(input)

        // Más comportamiento nuevo
        validateResult(result)
        return result
    }
}
```

### Diagrama Conceptual

```
ANTES: Clientes usan código legacy directamente
┌──────────┐        ┌──────────────────────┐
│ Cliente  │───────>│   Código Legacy      │
│          │        │ (Difícil de cambiar) │
└──────────┘        └──────────────────────┘


DESPUÉS: Clientes usan wrapper que envuelve legacy
┌──────────┐        ┌────────────────────────────────┐
│ Cliente  │───────>│      Wrapper                   │
│          │        │  ┌──────────────────────────┐  │
└──────────┘        │  │  - Log antes             │  │
                    │  │  - Validación            │  │
                    │  │  - Cache                 │  │
                    │  └──────────────────────────┘  │
                    │               │                │
                    │               ▼                │
                    │  ┌──────────────────────────┐  │
                    │  │    Código Legacy         │  │
                    │  │    (Sin modificar)       │  │
                    │  └──────────────────────────┘  │
                    │               │                │
                    │               ▼                │
                    │  ┌──────────────────────────┐  │
                    │  │  - Transform result      │  │
                    │  │  - Handle errors         │  │
                    │  └──────────────────────────┘  │
                    └────────────────────────────────┘


PATRONES DE ENVOLTORIO:

1. DECORATOR PATTERN
   Wrapper → Legacy → Result
             ↓
         Add behavior

2. PROXY PATTERN
   Wrapper → Control access → Legacy
             Cache, lazy load

3. ADAPTER PATTERN
   Wrapper → Transform interface → Legacy
             New API → Old API

4. INTERCEPTOR PATTERN
   Wrapper → Before → Legacy → After
             Log       ↓       Validate
```

## Ejemplo

### ANTES: Código Legacy Problemático

```pseudocode
// Código legacy que nadie quiere tocar
class LegacyReportGenerator {
    function generateReport(data) {
        // 500 líneas de código complejo
        // - Lee directamente de base de datos
        // - Tiene lógica de negocio mezclada
        // - Escribe logs mezclados con lógica
        // - No es testeable
        // - No maneja errores correctamente

        // Conexión hardcodeada
        connection = connectToDatabase("production_db", "hardcoded_password")

        // Query complejo sin parametrizar
        query = "SELECT * FROM reports WHERE id = " + data.id

        // Sin manejo de errores
        results = connection.execute(query)

        // Lógica de negocio mezclada con formato
        report = ""
        for (row in results) {
            report += "<tr><td>" + row.name + "</td>"
            // ... 100 líneas más ...

            // Logs mezclados
            print("Processing row: " + row.id)
        }

        // Escribe directamente a archivo
        writeToFile("/tmp/report.html", report)

        return report
    }
}

// USO: Clientes llaman directamente
generator = new LegacyReportGenerator()
report = generator.generateReport(data)
```

**Problemas**:
- No podemos testear sin base de datos real
- No podemos controlar logging
- No hay validación de inputs
- No hay manejo de errores
- No podemos cachear resultados
- Passwords hardcodeados

**¿Qué necesitamos?**:
- Validación de entrada
- Logging estructurado
- Manejo de errores
- Cache de resultados
- Métricas de rendimiento

### DESPUÉS: Con Wrap Change

```pseudocode
// ============================================
// LEGACY CODE: Completamente sin modificar
// ============================================
class LegacyReportGenerator {
    function generateReport(data) {
        // 500 líneas de código complejo
        // (EXACTAMENTE IGUAL - sin ningún cambio)

        connection = connectToDatabase("production_db", "hardcoded_password")
        query = "SELECT * FROM reports WHERE id = " + data.id
        results = connection.execute(query)

        report = ""
        for (row in results) {
            report += "<tr><td>" + row.name + "</td>"
            print("Processing row: " + row.id)
        }

        writeToFile("/tmp/report.html", report)
        return report
    }
}


// ============================================
// WRAPPER: Comportamiento nuevo añadido
// ============================================
class ReportGeneratorWrapper {
    private legacyGenerator
    private logger
    private cache
    private metrics

    constructor(logger, cache, metrics) {
        this.legacyGenerator = new LegacyReportGenerator()
        this.logger = logger
        this.cache = cache
        this.metrics = metrics
    }

    function generateReport(data) {
        // ✨ NUEVO: Validación
        validateInput(data)

        // ✨ NUEVO: Logging estructurado
        this.logger.info("Starting report generation", {
            reportId: data.id,
            timestamp: now()
        })

        // ✨ NUEVO: Check cache
        cachedReport = this.cache.get("report_" + data.id)
        if (cachedReport != null) {
            this.logger.info("Returning cached report")
            return cachedReport
        }

        // ✨ NUEVO: Métricas de tiempo
        startTime = now()

        try {
            // 🔧 DELEGACIÓN: Llamada al código legacy
            // (El código legacy hace su magia sin saber que está envuelto)
            report = this.legacyGenerator.generateReport(data)

            // ✨ NUEVO: Cache del resultado
            this.cache.set("report_" + data.id, report, ttl: 3600)

            // ✨ NUEVO: Métricas de éxito
            duration = now() - startTime
            this.metrics.recordSuccess("report_generation", duration)

            // ✨ NUEVO: Logging de éxito
            this.logger.info("Report generated successfully", {
                reportId: data.id,
                duration: duration,
                size: report.length
            })

            return report

        } catch (error) {
            // ✨ NUEVO: Manejo de errores
            duration = now() - startTime
            this.metrics.recordFailure("report_generation", duration)

            this.logger.error("Report generation failed", {
                reportId: data.id,
                error: error.message,
                duration: duration
            })

            // Transformar error legacy en error estructurado
            throw new ReportGenerationError(
                "Failed to generate report " + data.id,
                originalError: error
            )
        }
    }

    private function validateInput(data) {
        if (data == null) {
            throw new ValidationError("Data cannot be null")
        }

        if (data.id == null || data.id <= 0) {
            throw new ValidationError("Invalid report ID: " + data.id)
        }

        // Más validaciones...
    }
}


// ============================================
// USO: Reemplazar llamadas al legacy por wrapper
// ============================================

// ANTES:
// generator = new LegacyReportGenerator()
// report = generator.generateReport(data)

// DESPUÉS:
logger = createLogger()
cache = createCache()
metrics = createMetrics()

generator = new ReportGeneratorWrapper(logger, cache, metrics)
report = generator.generateReport(data)  // Misma interfaz, comportamiento mejorado


// ============================================
// TESTS: Solo para el wrapper (legacy no se testea)
// ============================================

test "validates input before calling legacy" {
    wrapper = createWrapper()

    assertThrows(() => {
        wrapper.generateReport(null)
    }, ValidationError)

    // El código legacy nunca se llamó
    assert legacyGeneratorWasNotCalled()
}

test "returns cached report if available" {
    mockCache = createMockCache()
    mockCache.set("report_123", "cached report")

    wrapper = createWrapper(cache: mockCache)
    data = { id: 123 }

    result = wrapper.generateReport(data)

    assert result == "cached report"
    assert legacyGeneratorWasNotCalled()  // No llamó al legacy
}

test "calls legacy generator if cache miss" {
    mockCache = createEmptyCache()
    mockLegacy = createMockLegacyGenerator()
    mockLegacy.setReturnValue("generated report")

    wrapper = createWrapper(cache: mockCache, legacy: mockLegacy)
    data = { id: 123 }

    result = wrapper.generateReport(data)

    assert result == "generated report"
    assert mockLegacy.wasCalledWith(data)
}

test "records metrics on success" {
    mockMetrics = createMockMetrics()
    wrapper = createWrapper(metrics: mockMetrics)

    wrapper.generateReport({ id: 123 })

    assert mockMetrics.successWasRecorded()
    assert mockMetrics.getDuration() > 0
}

test "handles errors from legacy generator" {
    mockLegacy = createMockLegacyGenerator()
    mockLegacy.throwError("Database connection failed")

    wrapper = createWrapper(legacy: mockLegacy)

    assertThrows(() => {
        wrapper.generateReport({ id: 123 })
    }, ReportGenerationError)
}

test "logs all operations" {
    mockLogger = createMockLogger()
    wrapper = createWrapper(logger: mockLogger)

    wrapper.generateReport({ id: 123 })

    assert mockLogger.infoWasCalledWith("Starting report generation")
    assert mockLogger.infoWasCalledWith("Report generated successfully")
}
```

**Mejoras logradas**:
- Validación de entrada antes de tocar legacy
- Logging estructurado y controlado
- Cache de resultados
- Métricas de rendimiento
- Manejo robusto de errores
- 100% testeable (sin tocar base de datos)
- Código legacy completamente intacto

## Proceso Paso a Paso

### Paso 1: Identificar el Código a Envolver

Determina qué método, función o clase legacy necesitas envolver.

```pseudocode
// Código target
class LegacyService {
    function processData(input) {
        // Código complicado
        return result
    }
}
```

**Criterios**:
- ¿Es un punto de entrada claro?
- ¿Puedes interceptar las llamadas?
- ¿La interfaz es relativamente estable?

### Paso 2: Analizar la Interfaz

Documenta la interfaz exacta del código legacy:

```pseudocode
// Análisis de interfaz
function legacyMethod(input: DataType): ReturnType
    - Parameters: input (type: DataType, constraints: ...)
    - Returns: ReturnType
    - Throws: Exception1, Exception2
    - Side effects: Writes to DB, calls external API
    - Preconditions: input must not be null
    - Postconditions: result is always non-null
```

### Paso 3: Decidir el Tipo de Wrapper

Elige entre diferentes patrones de wrapper:

```pseudocode
// Opción 1: DECORATOR - Añadir comportamiento
class DecoratorWrapper {
    function method(input) {
        logBefore()
        result = legacy.method(input)  // Delegación simple
        logAfter()
        return result
    }
}

// Opción 2: ADAPTER - Cambiar interfaz
class AdapterWrapper {
    function newMethodName(newInput) {
        oldInput = adaptInput(newInput)  // Transformar
        result = legacy.oldMethodName(oldInput)
        return adaptOutput(result)  // Transformar
    }
}

// Opción 3: PROXY - Control de acceso
class ProxyWrapper {
    function method(input) {
        if (!hasPermission(input)) {
            throw UnauthorizedError()
        }
        return legacy.method(input)  // Delegación condicional
    }
}

// Opción 4: CIRCUIT BREAKER - Protección
class CircuitBreakerWrapper {
    function method(input) {
        if (circuitIsOpen()) {
            return fallbackResponse()
        }
        try {
            result = legacy.method(input)
            recordSuccess()
            return result
        } catch (error) {
            recordFailure()
            openCircuitIfNeeded()
            throw error
        }
    }
}
```

### Paso 4: Crear la Estructura del Wrapper

Implementa el wrapper básico que delega al legacy:

```pseudocode
class ServiceWrapper {
    private legacyService

    constructor(legacyService = null) {
        // Permitir inyección para testing
        if (legacyService == null) {
            this.legacyService = new LegacyService()
        } else {
            this.legacyService = legacyService
        }
    }

    function processData(input) {
        // Por ahora solo delegar (verificar que funciona)
        return this.legacyService.processData(input)
    }
}
```

**Test inicial**:
```pseudocode
test "wrapper delegates to legacy correctly" {
    mockLegacy = createMock()
    mockLegacy.setReturnValue("expected result")

    wrapper = new ServiceWrapper(mockLegacy)
    result = wrapper.processData("test input")

    assert result == "expected result"
    assert mockLegacy.wasCalledWith("test input")
}
```

### Paso 5: Añadir Comportamiento Incremental

Añade el nuevo comportamiento paso a paso:

```pseudocode
// Paso 5a: Añadir validación
class ServiceWrapper {
    function processData(input) {
        // ✨ Nuevo: Validación
        if (input == null) {
            throw ValidationError("Input cannot be null")
        }

        return this.legacyService.processData(input)
    }
}

// Test
test "validates input" {
    wrapper = new ServiceWrapper()
    assertThrows(() => wrapper.processData(null))
}


// Paso 5b: Añadir logging
class ServiceWrapper {
    constructor(legacyService = null, logger = null) {
        this.legacyService = legacyService ?? new LegacyService()
        this.logger = logger ?? createLogger()
    }

    function processData(input) {
        if (input == null) {
            throw ValidationError("Input cannot be null")
        }

        // ✨ Nuevo: Logging
        this.logger.info("Processing data", { input: input })

        result = this.legacyService.processData(input)

        this.logger.info("Data processed", { result: result })

        return result
    }
}


// Paso 5c: Añadir manejo de errores
class ServiceWrapper {
    function processData(input) {
        if (input == null) {
            throw ValidationError("Input cannot be null")
        }

        this.logger.info("Processing data", { input: input })

        try {
            result = this.legacyService.processData(input)
            this.logger.info("Data processed", { result: result })
            return result

        } catch (error) {
            // ✨ Nuevo: Manejo de errores
            this.logger.error("Processing failed", { error: error })

            // Transformar error legacy en error estructurado
            throw new ProcessingError(
                "Failed to process data",
                cause: error,
                input: input
            )
        }
    }
}
```

### Paso 6: Añadir Tests Exhaustivos

Testea todos los caminos del wrapper:

```pseudocode
describe "ServiceWrapper" {
    test "successful processing" {
        mockLegacy = createMock()
        mockLegacy.setReturnValue("success")
        wrapper = new ServiceWrapper(mockLegacy)

        result = wrapper.processData("input")

        assert result == "success"
    }

    test "validation error" {
        wrapper = new ServiceWrapper()
        assertThrows(() => wrapper.processData(null), ValidationError)
    }

    test "legacy error is wrapped" {
        mockLegacy = createMock()
        mockLegacy.throwError("Legacy error")
        wrapper = new ServiceWrapper(mockLegacy)

        error = assertThrows(() => wrapper.processData("input"))

        assert error instanceof ProcessingError
        assert error.cause.message == "Legacy error"
    }

    test "logs all operations" {
        mockLogger = createMock()
        wrapper = new ServiceWrapper(logger: mockLogger)

        wrapper.processData("input")

        assert mockLogger.infoWasCalledTimes(2)
    }

    test "logs errors" {
        mockLogger = createMock()
        mockLegacy = createMockThatFails()
        wrapper = new ServiceWrapper(mockLegacy, mockLogger)

        try {
            wrapper.processData("input")
        } catch {}

        assert mockLogger.errorWasCalled()
    }
}
```

### Paso 7: Reemplazar Llamadas al Legacy

Identifica todos los lugares que usan el código legacy y reemplázalos:

```pseudocode
// Antes
service = new LegacyService()
result = service.processData(input)

// Después
service = new ServiceWrapper()  // Con el wrapper
result = service.processData(input)  // Misma interfaz
```

**Estrategia de migración**:
```pseudocode
// Opción 1: Feature flag
if (config.useNewWrapper) {
    service = new ServiceWrapper()
} else {
    service = new LegacyService()
}

// Opción 2: Reemplazo gradual por módulo
// Módulo A: Ya usa wrapper
// Módulo B: Ya usa wrapper
// Módulo C: Todavía usa legacy (migrar pronto)

// Opción 3: Alias temporal
// Crear alias para migración gradual
const ServiceProcessor = ServiceWrapper  // Nuevo nombre

// Clientes migran a nuevo nombre
service = new ServiceProcessor()
```

### Paso 8: Monitorear y Limpiar

Monitorea el comportamiento en producción:

```pseudocode
// Añadir métricas para comparar
class ServiceWrapper {
    function processData(input) {
        metrics.increment("wrapper.calls")
        startTime = now()

        try {
            result = this.legacyService.processData(input)
            duration = now() - startTime

            metrics.timing("wrapper.duration", duration)
            metrics.increment("wrapper.success")

            return result
        } catch (error) {
            metrics.increment("wrapper.errors")
            throw error
        }
    }
}

// Dashboard muestra:
// - wrapper.calls: 10,000/min
// - wrapper.duration: p50=50ms, p99=200ms
// - wrapper.success: 99.9%
// - wrapper.errors: 0.1%
```

Una vez estable, considera refactorizar el legacy interno o eliminar el wrapper si ya no es necesario.

## Problemas Comunes

### 1. Wrapper Demasiado Complejo

**Problema**: El wrapper tiene cientos de líneas, casi tanto como el legacy.

**Solución**: Dividir en múltiples wrappers con responsabilidades únicas.

```pseudocode
// Mal: Wrapper monolítico
class ServiceWrapper {
    // 500 líneas: validación, logging, cache, retry, circuit breaker...
}

// Bien: Chain of wrappers
service = new LegacyService()
service = new ValidationWrapper(service)
service = new LoggingWrapper(service)
service = new CacheWrapper(service)
service = new CircuitBreakerWrapper(service)

// Cada wrapper es simple y enfocado
```

### 2. Interfaz del Legacy Cambia

**Problema**: El código legacy cambia su interfaz y rompe el wrapper.

**Solución**: Tests de integración y versioning.

```pseudocode
// Test de contrato
test "wrapper is compatible with legacy interface" {
    realLegacy = new LegacyService()  // Real, no mock
    wrapper = new ServiceWrapper(realLegacy)

    // Verificar que la interfaz es compatible
    result = wrapper.processData(testInput)

    assert result != null  // Contrato básico
}

// Si legacy cambia, el test falla y sabes que debes actualizar wrapper
```

### 3. Performance Overhead

**Problema**: El wrapper añade latencia significativa.

**Solución**: Optimizar o hacer el wrapper opcional.

```pseudocode
// Medir overhead
test "wrapper overhead is acceptable" {
    legacy = new LegacyService()
    wrapper = new ServiceWrapper(legacy)

    // Benchmark legacy directo
    legacyTime = benchmark(() => legacy.processData(input))

    // Benchmark con wrapper
    wrapperTime = benchmark(() => wrapper.processData(input))

    overhead = wrapperTime - legacyTime

    // Overhead debe ser < 10% del tiempo total
    assert overhead < legacyTime * 0.1
}

// Si overhead es alto, optimizar o hacer wrapper opcional
class ServiceWrapper {
    constructor(options) {
        this.enableLogging = options.logging ?? true
        this.enableCache = options.cache ?? true
        // Permitir desactivar features costosas
    }
}
```

### 4. Dificultad para Testear Legacy Envuelto

**Problema**: Necesitas testear el comportamiento del legacy pero está envuelto.

**Solución**: Inyectar el legacy como dependencia.

```pseudocode
// Mal: Legacy hardcodeado
class ServiceWrapper {
    constructor() {
        this.legacy = new LegacyService()  // No testeable
    }
}

// Bien: Legacy inyectable
class ServiceWrapper {
    constructor(legacyService = null) {
        this.legacy = legacyService ?? new LegacyService()
    }
}

// En tests
mockLegacy = createMock()
wrapper = new ServiceWrapper(mockLegacy)

// Verificar interacción con legacy
wrapper.processData(input)
assert mockLegacy.wasCalledWith(expectedParams)
```

### 5. Wrapper Modifica Estado del Legacy Inesperadamente

**Problema**: El wrapper cambia el input antes de pasarlo al legacy, causando bugs sutiles.

**Solución**: Hacer copias y documentar claramente.

```pseudocode
// Mal: Modificar input directamente
class ServiceWrapper {
    function processData(input) {
        input.timestamp = now()  // Modifica el input original
        return this.legacy.processData(input)
    }
}

// Bien: Copiar antes de modificar
class ServiceWrapper {
    function processData(input) {
        // Crear copia para no afectar al original
        enrichedInput = cloneDeep(input)
        enrichedInput.timestamp = now()

        return this.legacy.processData(enrichedInput)
    }
}

// O documentar explícitamente
class ServiceWrapper {
    /**
     * NOTA: Este wrapper modifica el input añadiendo timestamp.
     * Si necesitas el input original, hacer copia antes de llamar.
     */
    function processData(input) {
        input.timestamp = now()
        return this.legacy.processData(input)
    }
}
```

## Criterios de Aceptación

Has aplicado Wrap Change correctamente cuando:

1. **Legacy intacto**: No modificaste ni una línea del código legacy
2. **Interfaz compatible**: El wrapper expone la misma interfaz (o mejor) que el legacy
3. **Delegación clara**: Es obvio que el wrapper delega al legacy
4. **Tests independientes**: Puedes testear el wrapper sin ejecutar el legacy real
5. **Comportamiento preservado**: La funcionalidad original sigue funcionando
6. **Nuevo comportamiento añadido**: El wrapper añade valor (logging, validación, etc.)
7. **Fácil de remover**: Puedes quitar el wrapper y volver al legacy si es necesario

## Beneficios

### Inmediatos
- Zero modificación del código legacy (riesgo minimizado)
- Añade funcionalidad sin tocar código frágil
- Testeable independientemente del legacy

### A Mediano Plazo
- Permite interceptar y controlar comportamiento
- Facilita migración gradual a nuevo código
- Crea abstracción sobre legacy

### A Largo Plazo
- Base para reemplazar legacy completamente
- Patrón reusable en múltiples partes del sistema
- Mejora calidad sin reescritura completa

## Técnicas Relacionadas

- **Sprout Change**: Úsalo cuando necesites añadir funcionalidad completamente nueva (no modificar existente)
- **Decorator Pattern**: Wrap Change implementa este patrón de diseño
- **Proxy Pattern**: Variante de Wrap para control de acceso
- **Adapter Pattern**: Variante de Wrap para transformar interfaces
- **Strangler Fig Pattern**: Wrapping a nivel de sistema para migración completa
- **Golden Master**: Úsalo para validar que el wrapper preserva comportamiento

## Versiones por Lenguaje

- [TypeScript](../../../../typescript/src/refactoring/parallel-change/wrap-change/) - [README](../../../../typescript/src/refactoring/parallel-change/wrap-change/README.md)
- [Go](../../../../go/refactoring/parallel-change/wrap-change/) - [README](../../../../go/refactoring/parallel-change/wrap-change/README.md)
- [Java](../../../../java/src/main/java/com/refactoring/refactoring/parallel_change/wrap_change/) - [README](../../../../java/src/main/java/com/refactoring/refactoring/parallel_change/wrap_change/README.md)
- [PHP](../../../../php/src/refactoring/parallel-change/wrap-change/) - [README](../../../../php/src/refactoring/parallel-change/wrap-change/README.md)
- [Python](../../../../python/src/refactoring/parallel_change/wrap_change/) - [README](../../../../python/src/refactoring/parallel_change/wrap_change/README.md)
- [C#](../../../../csharp/src/refactoring/parallel-change/wrap-change/) - [README](../../../../csharp/src/refactoring/parallel-change/wrap-change/README.md)

## Referencias en Español

### Artículos de Fran Iglesias

- [Introducción al Refactor](https://franiglesias.github.io/intro_refactor_1/) - Explicación detallada de Wrap y Sprout
- [Refactor rompiendo cosas](https://franiglesias.github.io/refactor-by-breaking/) - Cuándo y cómo usar wrappers
- [Modernizando el legacy](https://franiglesias.github.io/modernizando-el-legacy/) - Wrappers como estrategia de migración

## Referencias en Inglés

### Libros
- **Working Effectively with Legacy Code** - Michael Feathers (2004)
  - Capítulo 7: "It Takes Forever to Make a Change"
  - Sección "Wrap Method" y "Wrap Class"

- **Design Patterns: Elements of Reusable Object-Oriented Software** - Gang of Four (1994)
  - Decorator Pattern
  - Proxy Pattern

### Artículos
- [Decorator Pattern](https://refactoring.guru/design-patterns/decorator) - Refactoring Guru
- [Proxy Pattern](https://refactoring.guru/design-patterns/proxy) - Refactoring Guru
- [Strangler Fig Application](https://martinfowler.com/bliki/StranglerFigApplication.html) - Martin Fowler

---

**Próxima Técnica**: Para cambios más grandes que requieren múltiples fases, ve [Expand-Migrate-Contract](./expand-migrate-contract.md).
