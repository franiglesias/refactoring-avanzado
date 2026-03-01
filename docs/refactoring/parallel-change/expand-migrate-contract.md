# Expand-Migrate-Contract (Expandir-Migrar-Contraer)

## Definición

**Expand-Migrate-Contract** es una técnica de refactoring en tres fases para realizar cambios estructurales grandes de forma segura y gradual. En lugar de un cambio "big bang", el proceso se divide en:

1. **Expand (Expandir)**: Ampliar la API/interfaz para soportar tanto el uso viejo como el nuevo simultáneamente
2. **Migrate (Migrar)**: Cambiar gradualmente todos los clientes del uso viejo al nuevo
3. **Contract (Contraer)**: Eliminar el código viejo una vez completada la migración

También conocido como **Parallel Change** o **Expand-Contract Pattern**.

## Cuándo Usar

- Necesitas cambiar interfaces públicas, APIs, o contratos
- Múltiples equipos o sistemas dependen del código que vas a cambiar
- El cambio afecta a muchos clientes que no puedes modificar simultáneamente
- Trabajas en sistemas distribuidos o microservicios
- Necesitas mantener compatibilidad hacia atrás durante semanas o meses
- No puedes coordinar un despliegue sincronizado de todos los componentes
- Quieres deployar cambios de forma incremental con capacidad de rollback
- El cambio es estructural y no puede hacerse con un simple refactor automático

## Problema que Resuelve

Los cambios grandes en sistemas en producción enfrentan desafíos críticos:

- **Despliegue "big bang"**: Todos deben cambiar al mismo tiempo o todo se rompe
- **Coordinación imposible**: 10 equipos no pueden desplegar exactamente a la misma hora
- **Rollback complejo**: Si algo falla, revertir es tan complejo como el cambio original
- **Clientes externos**: No puedes forzar a clientes externos a actualizar inmediatamente
- **Downtime**: Cambios abruptos requieren apagar sistemas
- **Testing limitado**: No puedes probar ambas versiones en producción

Expand-Migrate-Contract resuelve esto:
1. Permite coexistencia temporal de versiones vieja y nueva
2. Cada equipo migra a su propio ritmo
3. Rollback es trivial en cada fase
4. Zero downtime deployments
5. Testing incremental en producción

## Descripción Detallada

### Las Tres Fases

```
┌─────────────────────────────────────────────────────────────┐
│ ESTADO INICIAL: API Vieja                                  │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  interface OldAPI {                                         │
│      function doSomething(param1, param2)                   │
│  }                                                          │
│                                                             │
│  Clientes:                                                  │
│  ✓ Cliente A                                               │
│  ✓ Cliente B                                               │
│  ✓ Cliente C                                               │
│  ✓ Cliente D                                               │
│                                                             │
└─────────────────────────────────────────────────────────────┘


┌─────────────────────────────────────────────────────────────┐
│ FASE 1: EXPAND (Expandir) - Soportar ambas versiones       │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  interface ExpandedAPI {                                    │
│      // ⚠️  Método viejo - DEPRECATED                      │
│      @deprecated                                            │
│      function doSomething(param1, param2) {                │
│          // Adaptar a nueva interfaz internamente          │
│          return doSomethingNew({ param1, param2 })         │
│      }                                                      │
│                                                             │
│      // ✨ Método nuevo - RECOMENDADO                      │
│      function doSomethingNew(options)                      │
│  }                                                          │
│                                                             │
│  Clientes (sin cambios todavía):                           │
│  ✓ Cliente A → doSomething() [OLD]                        │
│  ✓ Cliente B → doSomething() [OLD]                        │
│  ✓ Cliente C → doSomething() [OLD]                        │
│  ✓ Cliente D → doSomething() [OLD]                        │
│                                                             │
└─────────────────────────────────────────────────────────────┘


┌─────────────────────────────────────────────────────────────┐
│ FASE 2: MIGRATE (Migrar) - Cambiar clientes gradualmente   │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  interface ExpandedAPI {                                    │
│      @deprecated                                            │
│      function doSomething(param1, param2)                  │
│                                                             │
│      function doSomethingNew(options)                      │
│  }                                                          │
│                                                             │
│  Clientes (migrando):                                       │
│  ✅ Cliente A → doSomethingNew() [NEW] ← Migrado          │
│  ✅ Cliente B → doSomethingNew() [NEW] ← Migrado          │
│  ⏳ Cliente C → doSomething() [OLD] ← En progreso         │
│  ⏳ Cliente D → doSomething() [OLD] ← Pendiente           │
│                                                             │
│  Métricas:                                                  │
│  - old_api_calls: 2000/min (bajando)                      │
│  - new_api_calls: 8000/min (subiendo)                     │
│  - migration_progress: 50%                                 │
│                                                             │
└─────────────────────────────────────────────────────────────┘


┌─────────────────────────────────────────────────────────────┐
│ FASE 2b: MIGRATE (Completada) - Todos migrados             │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  interface ExpandedAPI {                                    │
│      @deprecated                                            │
│      function doSomething(param1, param2)  ← Sin uso       │
│                                                             │
│      function doSomethingNew(options)                      │
│  }                                                          │
│                                                             │
│  Clientes (todos migrados):                                 │
│  ✅ Cliente A → doSomethingNew() [NEW]                    │
│  ✅ Cliente B → doSomethingNew() [NEW]                    │
│  ✅ Cliente C → doSomethingNew() [NEW]                    │
│  ✅ Cliente D → doSomethingNew() [NEW]                    │
│                                                             │
│  Métricas:                                                  │
│  - old_api_calls: 0/min ← ¡Sin uso!                       │
│  - new_api_calls: 10000/min                                │
│  - migration_progress: 100%                                │
│                                                             │
└─────────────────────────────────────────────────────────────┘


┌─────────────────────────────────────────────────────────────┐
│ FASE 3: CONTRACT (Contraer) - Eliminar código viejo        │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  interface NewAPI {                                         │
│      function doSomethingNew(options)                      │
│  }                                                          │
│                                                             │
│  // Código viejo eliminado                                 │
│  // Interfaz simplificada                                  │
│  // Sin deuda técnica                                      │
│                                                             │
│  Clientes:                                                  │
│  ✓ Cliente A                                               │
│  ✓ Cliente B                                               │
│  ✓ Cliente C                                               │
│  ✓ Cliente D                                               │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

## Ejemplo

### Contexto

Tenemos un servicio de envío de emails con interfaz antigua que necesitamos modernizar.

### Estado Inicial: API Vieja

```pseudocode
// ============================================
// ESTADO INICIAL: API que queremos cambiar
// ============================================

class EmailService {
    // Interfaz vieja: parámetros separados, muchos opcionales
    function sendEmail(
        to,           // string
        subject,      // string
        body,         // string
        from = null,  // opcional
        cc = null,    // opcional
        bcc = null,   // opcional
        priority = "normal"  // opcional
    ) {
        // Implementación con muchos parámetros posicionales
        email = {
            to: to,
            subject: subject,
            body: body,
            from: from ?? "noreply@company.com",
            cc: cc ?? [],
            bcc: bcc ?? [],
            priority: priority
        }

        return smtpService.send(email)
    }
}

// CLIENTES: Múltiples equipos usando la API vieja
// Equipo A - Sistema de notificaciones
emailService.sendEmail("user@example.com", "Welcome", "Welcome message")

// Equipo B - Sistema de reportes
emailService.sendEmail(
    "admin@example.com",
    "Daily Report",
    reportBody,
    null,                    // from (default)
    ["manager@example.com"], // cc
    null,                    // bcc
    "high"                   // priority
)

// Equipo C - Sistema de alertas
emailService.sendEmail("ops@example.com", "ALERT", alertBody)

// Equipo D - Marketing (externo)
emailService.sendEmail("customer@example.com", "Newsletter", content)
```

**Problemas con la API vieja**:
- Demasiados parámetros posicionales (error-prone)
- Valores null para opcionales (confuso)
- Difícil añadir nuevos parámetros sin romper compatibilidad
- No es extensible

**Queremos cambiar a**: Un objeto de opciones con named parameters

### FASE 1: EXPAND - Soportar Ambas Interfaces

```pseudocode
// ============================================
// FASE 1: EXPAND
// Añadir nueva interfaz sin romper la vieja
// ============================================

class EmailService {
    // ⚠️  MÉTODO VIEJO - Marcado como deprecated
    @deprecated("Use sendEmailV2 instead. Will be removed in v3.0")
    function sendEmail(
        to,
        subject,
        body,
        from = null,
        cc = null,
        bcc = null,
        priority = "normal"
    ) {
        // ✨ CLAVE: Adaptar interfaz vieja a nueva internamente
        // El método viejo delega al nuevo
        options = {
            to: to,
            subject: subject,
            body: body
        }

        // Añadir opcionales solo si están presentes
        if (from != null) options.from = from
        if (cc != null) options.cc = cc
        if (bcc != null) options.bcc = bcc
        if (priority != "normal") options.priority = priority

        // Log para tracking de uso
        metrics.increment("email.api.v1.calls")
        logger.warn("Deprecated API used", {
            caller: getCallerInfo(),
            migrateUrl: "https://docs.company.com/email-api-migration"
        })

        // Delegar a la nueva implementación
        return this.sendEmailV2(options)
    }


    // ✨ MÉTODO NUEVO - Interfaz moderna
    function sendEmailV2(options) {
        // Validar opciones requeridas
        validateRequired(options, ["to", "subject", "body"])

        // Defaults para opcionales
        email = {
            to: options.to,
            subject: options.subject,
            body: options.body,
            from: options.from ?? "noreply@company.com",
            cc: options.cc ?? [],
            bcc: options.bcc ?? [],
            priority: options.priority ?? "normal",
            // ✨ Nuevas opciones posibles (extensible)
            attachments: options.attachments ?? [],
            replyTo: options.replyTo ?? null,
            templateId: options.templateId ?? null,
            metadata: options.metadata ?? {}
        }

        // Métricas
        metrics.increment("email.api.v2.calls")

        // Implementación real
        return smtpService.send(email)
    }
}


// ============================================
// DEPLOYMENT FASE 1
// ============================================

// 1. Deployar nueva versión del servicio
//    - Ambos métodos disponibles
//    - Método viejo delega a nuevo internamente

// 2. Clientes NO cambian todavía (siguen usando API vieja)
//    Equipo A: emailService.sendEmail("user@...", "Welcome", "...")
//    Equipo B: emailService.sendEmail("admin@...", "Report", ...)
//    Equipo C: emailService.sendEmail("ops@...", "ALERT", ...)
//    Equipo D: emailService.sendEmail("customer@...", "Newsletter", ...)

// 3. Verificar que TODO sigue funcionando
//    - Tests: Ambas interfaces funcionan
//    - Producción: Sin errores
//    - Métricas: email.api.v1.calls = 100%

// ✅ Si algo falla, revertir es trivial (solo cambio en el servicio)
```

### FASE 2: MIGRATE - Migrar Clientes Gradualmente

```pseudocode
// ============================================
// FASE 2: MIGRATE
// Cambiar clientes uno por uno al nuevo API
// ============================================

// El servicio NO cambia (ya soporta ambas interfaces)
// Solo los CLIENTES cambian


// ────────────────────────────────────────────
// Semana 1: Equipo A migra (más simple)
// ────────────────────────────────────────────

// ANTES (API vieja)
emailService.sendEmail("user@example.com", "Welcome", "Welcome message")

// DESPUÉS (API nueva)
emailService.sendEmailV2({
    to: "user@example.com",
    subject: "Welcome",
    body: "Welcome message"
})

// Deploy: Solo sistema de notificaciones
// Métricas:
//   email.api.v1.calls: 75%
//   email.api.v2.calls: 25%


// ────────────────────────────────────────────
// Semana 2: Equipo C migra (alertas)
// ────────────────────────────────────────────

// ANTES
emailService.sendEmail("ops@example.com", "ALERT", alertBody)

// DESPUÉS (aprovecha nuevas features)
emailService.sendEmailV2({
    to: "ops@example.com",
    subject: "ALERT",
    body: alertBody,
    priority: "high",  // Más claro que antes
    metadata: {
        alertLevel: "critical",
        timestamp: now()
    }
})

// Deploy: Solo sistema de alertas
// Métricas:
//   email.api.v1.calls: 50%
//   email.api.v2.calls: 50%


// ────────────────────────────────────────────
// Semana 3: Equipo B migra (reportes)
// ────────────────────────────────────────────

// ANTES (parámetros confusos)
emailService.sendEmail(
    "admin@example.com",
    "Daily Report",
    reportBody,
    null,                    // ¿from?
    ["manager@example.com"], // cc
    null,                    // ¿bcc?
    "high"
)

// DESPUÉS (mucho más claro)
emailService.sendEmailV2({
    to: "admin@example.com",
    subject: "Daily Report",
    body: reportBody,
    cc: ["manager@example.com"],
    priority: "high",
    // Ahora podemos añadir attachments fácilmente
    attachments: [reportPDF]
})

// Deploy: Solo sistema de reportes
// Métricas:
//   email.api.v1.calls: 25%
//   email.api.v2.calls: 75%


// ────────────────────────────────────────────
// Semana 4-8: Equipo D migra (externo - más lento)
// ────────────────────────────────────────────

// Coordinación con equipo externo
// Documentación: https://docs.company.com/email-api-v2
// Soporte: email-migration@company.com
// Deadline: Final del trimestre

// ANTES
emailService.sendEmail("customer@example.com", "Newsletter", content)

// DESPUÉS
emailService.sendEmailV2({
    to: "customer@example.com",
    subject: "Newsletter",
    body: content,
    templateId: "newsletter_v2",  // Usan nueva feature
    metadata: {
        campaign: "spring_2024"
    }
})

// Deploy: Sistema de marketing
// Métricas:
//   email.api.v1.calls: 0% ← ¡Todos migrados!
//   email.api.v2.calls: 100%


// ────────────────────────────────────────────
// MONITORING durante migración
// ────────────────────────────────────────────

// Dashboard de migración
class EmailServiceMetrics {
    function trackUsage() {
        // Contadores por versión
        incrementCounter("email.api.v1.calls")
        incrementCounter("email.api.v2.calls")

        // Identificar quién usa API vieja
        if (usingOldAPI) {
            logger.warn("Old API usage detected", {
                service: callerServiceName,
                endpoint: callerEndpoint
            })
        }
    }

    function getMigrationProgress() {
        v1Calls = getCounter("email.api.v1.calls")
        v2Calls = getCounter("email.api.v2.calls")
        total = v1Calls + v2Calls

        return {
            v1Percentage: (v1Calls / total) * 100,
            v2Percentage: (v2Calls / total) * 100,
            servicesRemaining: listServicesUsingV1()
        }
    }
}

// Alertas automáticas
if (oldAPICallsLastWeek == 0) {
    notify("email-team", "All clients migrated! Ready for contract phase")
}
```

### FASE 3: CONTRACT - Eliminar Código Viejo

```pseudocode
// ============================================
// FASE 3: CONTRACT
// Eliminar API vieja una vez todos migraron
// ============================================

// VERIFICAR ANTES DE CONTRAER
// 1. Métricas muestran 0 llamadas a API vieja durante 2+ semanas
// 2. Todos los equipos confirmaron migración
// 3. Logs no muestran warnings de deprecated API

class EmailService {
    // ❌ ELIMINADO: Método viejo removido completamente
    // function sendEmail(...) { ... }  ← BORRADO

    // ✨ ÚNICO MÉTODO: Interfaz nueva (renombrada)
    function sendEmail(options) {  // Ahora es el método principal
        // Validar opciones
        validateRequired(options, ["to", "subject", "body"])

        // Construir email
        email = {
            to: options.to,
            subject: options.subject,
            body: options.body,
            from: options.from ?? "noreply@company.com",
            cc: options.cc ?? [],
            bcc: options.bcc ?? [],
            priority: options.priority ?? "normal",
            attachments: options.attachments ?? [],
            replyTo: options.replyTo ?? null,
            templateId: options.templateId ?? null,
            metadata: options.metadata ?? {}
        }

        // Métricas simplificadas (sin versioning)
        metrics.increment("email.sent")

        // Enviar
        return smtpService.send(email)
    }
}

// CLIENTES: Todos usando interfaz nueva
// Equipo A
emailService.sendEmail({
    to: "user@example.com",
    subject: "Welcome",
    body: "Welcome message"
})

// Equipo B
emailService.sendEmail({
    to: "admin@example.com",
    subject: "Daily Report",
    body: reportBody,
    cc: ["manager@example.com"],
    priority: "high",
    attachments: [reportPDF]
})

// Equipo C
emailService.sendEmail({
    to: "ops@example.com",
    subject: "ALERT",
    body: alertBody,
    priority: "high",
    metadata: { alertLevel: "critical" }
})

// Equipo D
emailService.sendEmail({
    to: "customer@example.com",
    subject: "Newsletter",
    body: content,
    templateId: "newsletter_v2"
})


// ============================================
// DEPLOYMENT FASE 3
// ============================================

// 1. Crear rama con contracción
git checkout -b contract/remove-old-email-api

// 2. Eliminar código deprecated
//    - Borrar método sendEmailV2 (o renombrar a sendEmail)
//    - Eliminar adaptador de API vieja
//    - Limpiar métricas de versioning
//    - Actualizar tests

// 3. Actualizar documentación
//    - Marcar API v1 como removida
//    - Actualizar ejemplos
//    - Changelog: "BREAKING: Removed deprecated sendEmail(to, subject, body...)"

// 4. Deploy con anuncio
//    Announcement: "EmailService v3.0 released - deprecated API removed"
//    Rollback plan: Revertir a v2.x si problemas (aunque no debería haberlos)

// 5. Monitorear
//    - Sin errores esperados (todos ya migraron)
//    - Métricas normales
//    - Performance igual o mejor (menos código)

// ✅ MIGRACIÓN COMPLETA
//    - Código más limpio
//    - Sin deuda técnica
//    - API moderna y extensible
```

## Proceso Paso a Paso

### Paso 1: Planificar el Cambio

Antes de empezar, documenta:

```pseudocode
// PLAN DE MIGRACIÓN

// 1. Estado actual
CurrentAPI:
  - Método: processOrder(orderId, userId, options)
  - Usado por: 15 servicios, 3 equipos, 2 sistemas externos
  - Tráfico: 10,000 calls/min

// 2. Estado deseado
NewAPI:
  - Método: processOrder({ orderId, userId, context })
  - Razón: Más extensible, mejor manejo de errores, metadata adicional

// 3. Cronograma
Week 1: EXPAND - Implementar nueva interfaz
Week 2-4: MIGRATE - Migrar servicios internos (10 servicios)
Week 5-8: MIGRATE - Migrar servicios de otros equipos (5 servicios)
Week 9-12: MIGRATE - Coordinar con sistemas externos (2 sistemas)
Week 13: Verificación - Confirmar 0 uso de API vieja
Week 14: CONTRACT - Eliminar código deprecated

// 4. Métricas de éxito
- 100% clientes migrados
- 0 llamadas a API vieja durante 2 semanas
- Sin incidentes de producción
- Performance igual o mejor

// 5. Riesgos y mitigaciones
Riesgo: Sistemas externos no migran a tiempo
Mitigación: Deadline flexible, soporte extendido

Riesgo: Breaking change no detectado
Mitigación: Tests de compatibilidad, feature flag
```

### Paso 2: Implementar Fase EXPAND

Crea la interfaz expandida que soporta ambos usos:

```pseudocode
class ServiceAPI {
    // ESTRATEGIA 1: Método deprecated + método nuevo
    @deprecated("Use processV2")
    function process(oldParams) {
        // Adaptar a nueva interfaz
        newParams = adaptOldToNew(oldParams)
        return this.processV2(newParams)
    }

    function processV2(newParams) {
        // Implementación real
    }


    // ESTRATEGIA 2: Overloading (si el lenguaje lo soporta)
    function process(params) {
        if (isOldFormat(params)) {
            metrics.increment("api.old")
            params = adaptOldToNew(params)
        } else {
            metrics.increment("api.new")
        }

        return executeProcess(params)
    }


    // ESTRATEGIA 3: Parámetro opcional de versión
    function process(params, version = "v1") {
        if (version == "v1") {
            metrics.increment("api.v1")
            params = adaptV1ToV2(params)
        } else {
            metrics.increment("api.v2")
        }

        return executeProcess(params)
    }
}


// TESTING: Verificar compatibilidad
test "old interface still works" {
    service = new ServiceAPI()

    // Llamada vieja debe funcionar
    result = service.process(oldStyleParams)

    assert result.success
}

test "new interface works" {
    service = new ServiceAPI()

    // Llamada nueva debe funcionar
    result = service.processV2(newStyleParams)

    assert result.success
}

test "both interfaces produce same result" {
    service = new ServiceAPI()

    oldResult = service.process(oldParams)
    newResult = service.processV2(adaptedParams)

    // Resultado equivalente
    assert oldResult.data == newResult.data
}
```

### Paso 3: Añadir Observabilidad

Instrumenta el código para tracking de migración:

```pseudocode
class ServiceAPI {
    function process(params) {
        // Detectar qué versión se usa
        version = detectVersion(params)

        // Métricas
        metrics.increment("api.calls.total")
        metrics.increment("api.calls." + version)

        // Logging para deprecation
        if (version == "old") {
            logger.warn("Deprecated API used", {
                version: version,
                caller: getCallerIdentity(),
                migrateGuide: "https://docs.../migration"
            })
        }

        // Timestamp para analytics
        recordUsage(version, getCallerIdentity(), now())

        // Ejecutar
        return execute(params)
    }
}


// Dashboard de migración
class MigrationDashboard {
    function getProgress() {
        last7Days = getUsageStats(since: now() - 7days)

        return {
            oldAPIPercentage: calculatePercentage(last7Days, "old"),
            newAPIPercentage: calculatePercentage(last7Days, "new"),
            clientsRemaining: listClientsUsingOld(),
            estimatedCompletionDate: extrapolateCompletion(),
            blockers: identifyStuckClients()
        }
    }
}
```

### Paso 4: Comunicar y Documentar

Informa a todos los stakeholders:

```pseudocode
// ANUNCIO DE DEPRECATION

Subject: EmailService API v1 Deprecated - Migration Required

Dear Team,

We've released EmailService v2 with improved interface.
The old API (v1) is now DEPRECATED and will be removed in Q2 2024.

OLD API (deprecated):
  emailService.sendEmail(to, subject, body, from, cc, bcc, priority)

NEW API (recommended):
  emailService.sendEmailV2({
    to: "user@example.com",
    subject: "Subject",
    body: "Body",
    // ... optional parameters
  })

Benefits:
- Clearer named parameters
- Easier to extend
- Support for attachments and templates

Migration Guide: https://docs.company.com/email-migration
Support Channel: #email-api-migration
Deadline: June 30, 2024

Your service currently uses OLD API. Please migrate by the deadline.

Need help? Contact: api-team@company.com

Thanks,
API Team
```

### Paso 5: Ejecutar Fase MIGRATE

Migra clientes en orden de prioridad:

```pseudocode
// PRIORIZACIÓN DE MIGRACIÓN

// 1. Servicios que TÚ controlas (fácil)
Priority 1: Internal services (Week 2-4)
  - Service A: Owned by our team
  - Service B: Owned by our team
  - Service C: Owned by our team

// 2. Servicios de otros equipos internos (coordinación media)
Priority 2: Other teams (Week 5-8)
  - Service D: Team X (contacto: john@company)
  - Service E: Team Y (contacto: jane@company)

// 3. Sistemas externos (difícil, largo)
Priority 3: External systems (Week 9-12)
  - Partner API: External company (contacto: partner@external.com)
  - Legacy integration: Vendor system (limited control)


// ESTRATEGIA POR CLIENTE

// Para servicios propios: Migrar directamente
function migrateOwnService(service) {
    // 1. Crear rama
    git checkout -b migrate/email-api-v2

    // 2. Cambiar código
    replaceOldAPICallsWithNew(service)

    // 3. Tests
    runTests()

    // 4. Deploy
    deployToProduction()

    // 5. Verificar métricas
    confirmNoErrors()
}

// Para otros equipos: Crear tickets y dar soporte
function migrateOtherTeamService(service, team) {
    // 1. Crear ticket en su sistema
    ticket = createJiraTicket({
        team: team,
        title: "Migrate to EmailService v2",
        description: "...",
        guide: "https://docs.../migration",
        deadline: "June 30, 2024"
    })

    // 2. Ofrecer soporte
    offerPairProgramming(team)

    // 3. Tracking
    trackProgress(ticket)

    // 4. Reminder si no progresan
    if (notProgressingAfter2Weeks()) {
        sendReminder(team)
    }
}

// Para externos: Coordinación y soporte extendido
function migrateExternalSystem(system, contact) {
    // 1. Email formal
    sendMigrationNotice(contact, {
        deadline: "June 30, 2024",
        support: "api-support@company.com"
    })

    // 2. Documentación detallada
    provideDetailedDocs(system)

    // 3. Ambiente de staging
    provideTestingEnvironment()

    // 4. Deadline flexible
    if (needsExtension()) {
        grantExtension("September 30, 2024")
    }
}
```

### Paso 6: Verificar Migración Completa

Antes de contraer, confirma que todos migraron:

```pseudocode
function verifyMigrationComplete() {
    // 1. Métricas: 0 llamadas a API vieja
    oldAPICalls = metrics.get("api.v1.calls", last: 14days)
    if (oldAPICalls > 0) {
        return {
            ready: false,
            reason: "Old API still receiving " + oldAPICalls + " calls",
            action: "Identify remaining clients and migrate"
        }
    }

    // 2. Logs: Sin warnings de deprecated
    deprecationWarnings = logs.count("Deprecated API used", last: 14days)
    if (deprecationWarnings > 0) {
        return {
            ready: false,
            reason: "Still seeing deprecation warnings",
            clients: identifyClientsFromLogs()
        }
    }

    // 3. Confirmación de equipos
    pendingTeams = teams.filter(t => !t.confirmedMigration)
    if (pendingTeams.length > 0) {
        return {
            ready: false,
            reason: "Teams pending confirmation: " + pendingTeams,
            action: "Request explicit confirmation"
        }
    }

    // 4. Tests: API vieja puede ser removida sin romper tests
    testResults = runTestsWithoutOldAPI()
    if (testResults.failed > 0) {
        return {
            ready: false,
            reason: "Tests still depend on old API",
            failingTests: testResults.failures
        }
    }

    // ✅ LISTO PARA CONTRAER
    return {
        ready: true,
        message: "All clients migrated. Safe to remove old API.",
        stats: {
            migratedClients: countMigratedClients(),
            timelineWeeks: calculateMigrationDuration(),
            incidentsDuringMigration: 0
        }
    }
}
```

### Paso 7: Ejecutar Fase CONTRACT

Elimina el código viejo:

```pseudocode
// CONTRACCIÓN SEGURA

// 1. Crear rama
git checkout -b contract/remove-email-api-v1

// 2. Eliminar código deprecated
class EmailService {
    // ❌ BORRAR método viejo
    // @deprecated
    // function sendEmail(to, subject, body, ...) { ... }

    // ✨ MANTENER solo método nuevo (renombrar si quieres)
    function sendEmail(options) {
        // Implementación nueva
    }
}

// 3. Eliminar adaptadores
// ❌ function adaptOldToNew(oldParams) { ... }  // Ya no necesario

// 4. Simplificar métricas
// ANTES:
metrics.increment("api.v1.calls")
metrics.increment("api.v2.calls")

// DESPUÉS:
metrics.increment("api.calls")  // Solo una métrica


// 5. Limpiar tests
// ❌ Borrar tests del API vieja
// test "old API works" { ... }

// ✅ Mantener tests del API nueva
test "email service sends email" {
    service.sendEmail({ to: "...", subject: "...", body: "..." })
    assert emailWasSent()
}

// 6. Actualizar documentación
// - Remover referencias a API v1
// - Actualizar changelog
// - Marcar v1 como "removed"

// 7. Commit y PR
git commit -m "Remove deprecated EmailService v1 API

All clients have been migrated to v2.
Metrics show 0 usage of v1 for 2+ weeks.

BREAKING CHANGE: EmailService.sendEmail(to, subject, body, ...) removed.
Use EmailService.sendEmail({ to, subject, body, ... }) instead.

Migration guide: https://docs.../email-migration"

// 8. Deploy con monitoring
deploy()
monitorForErrors()

// 9. Si todo bien, celebrar 🎉
notifyTeam("Email API migration complete! Old API removed. 🎉")
```

### Paso 8: Retrospectiva y Documentación

Aprende de la experiencia:

```pseudocode
// POST-MORTEM DE MIGRACIÓN

Migration: EmailService v1 → v2
Duration: 12 weeks (planned: 14 weeks)
Teams involved: 4
Services migrated: 15
External systems: 2

✅ What went well:
- Clear communication and documentation
- Metrics dashboard helpful for tracking progress
- No production incidents during migration
- All teams cooperated

⚠️  Challenges:
- External partner took longer than expected (8 weeks vs 4 planned)
- One team forgot about migration until week 10
- Initial metrics had bug (undercounting v1 calls)

📚 Learnings:
- Start communication 2 weeks before deprecation
- Add automated reminders for teams
- Test metrics implementation thoroughly
- Give external partners 2x time buffer

📝 Template for next migration:
- Use this same process
- Improve automation of tracking
- Create reusable dashboard template
```

## Problemas Comunes

### 1. Cliente No Puede Migrar a Tiempo

**Problema**: Un cliente crítico no puede migrar antes del deadline.

**Solución**: Extender deadline solo para ese cliente o mantener API vieja más tiempo.

```pseudocode
// Opción 1: Extension selective por cliente
class ServiceAPI {
    function process(params, clientId) {
        // Clientes con extensión pueden seguir usando API vieja
        if (hasExtension(clientId)) {
            logger.warn("Client using extended deprecation", { clientId })
            return processOld(params)
        }

        // Otros clientes: Error si usan API vieja
        if (isOldFormat(params)) {
            throw new DeprecatedAPIError(
                "API v1 removed. Please migrate to v2"
            )
        }

        return processNew(params)
    }
}

// Opción 2: Mantener API vieja más tiempo pero con overhead
class ServiceAPI {
    function process(params) {
        if (isOldFormat(params)) {
            // Añadir latencia artificial para incentivar migración
            sleep(100ms)

            logger.error("DEPRECATED API USAGE - Performance degraded", {
                caller: getCallerIdentity(),
                message: "Migrate immediately to avoid performance issues"
            })

            // Cobrar más por uso de API vieja (si es API pública)
            chargeDeprecationFee()
        }

        return execute(params)
    }
}
```

### 2. Cambio No Es Backward Compatible

**Problema**: Descubres que el cambio rompe compatibilidad de forma no obvia.

**Solución**: Fase EXPAND debe manejar ambos casos correctamente.

```pseudocode
// Ejemplo: Cambio en validación que puede romper clientes
// OLD API: Acepta emails sin validar
// NEW API: Valida formato de email

// ❌ MAL: Rompe clientes con datos inválidos
function sendEmailV2(options) {
    if (!isValidEmail(options.to)) {
        throw ValidationError("Invalid email")  // Rompe clientes existentes
    }
    // ...
}

// ✅ BIEN: Comportamiento gradual
function sendEmailV2(options, strictValidation = false) {
    if (!isValidEmail(options.to)) {
        if (strictValidation) {
            throw ValidationError("Invalid email")
        } else {
            // Log pero permite continuar (compatibilidad)
            logger.warn("Invalid email format", { email: options.to })
            // Intentar enviar de todas formas (como API vieja)
        }
    }
    // ...
}

// Migración en dos fases:
// Fase 1: strictValidation=false (default) - Compatible
// Fase 2: Después que todos migraron, cambiar a strictValidation=true
```

### 3. Difícil Detectar Qué Clientes Usan API Vieja

**Problema**: Métricas no identifican claramente quién debe migrar.

**Solución**: Mejorar observabilidad y logging.

```pseudocode
class ServiceAPI {
    @deprecated
    function oldMethod(params) {
        // Capturar información del caller
        callerInfo = {
            serviceId: getCallerServiceId(),
            serviceName: getCallerServiceName(),
            endpoint: getCallerEndpoint(),
            ip: getCallerIP(),
            userAgent: getCallerUserAgent(),
            timestamp: now()
        }

        // Log estructurado
        logger.warn("Deprecated API usage", callerInfo)

        // Métrica con tags
        metrics.increment("api.deprecated.calls", tags: {
            service: callerInfo.serviceName
        })

        // Persistir para análisis
        database.insert("deprecated_api_usage", callerInfo)

        // Ejecutar
        return execute(params)
    }
}

// Query para identificar clientes
SELECT service_name, COUNT(*) as call_count, MAX(timestamp) as last_used
FROM deprecated_api_usage
WHERE timestamp > NOW() - INTERVAL '7 days'
GROUP BY service_name
ORDER BY call_count DESC

// Resultados:
// service_name      | call_count | last_used
// ------------------+------------+-------------------
// marketing-service | 50000      | 2024-01-15 14:30
// alerts-service    | 1000       | 2024-01-14 09:15
// legacy-batch      | 10         | 2024-01-10 03:00
```

### 4. Fase MIGRATE Toma Demasiado Tiempo

**Problema**: Después de 3 meses, todavía hay clientes sin migrar.

**Solución**: Estrategias de aceleración.

```pseudocode
// ESTRATEGIA 1: Deadline duro con breaking change
// Comunicar claramente:
"API v1 will be REMOVED on March 31st.
After this date, calls to v1 will return HTTP 410 Gone."

// Enforcement automático
function oldMethod(params) {
    if (now() > REMOVAL_DATE) {
        throw new APIRemovedError(
            "API v1 was removed on March 31st. Use v2.",
            migrationGuide: "https://docs.../migration"
        )
    }
    // ...
}


// ESTRATEGIA 2: Degradación gradual
// Mes 1-2: Funciona normal
// Mes 3: Añadir 100ms latencia
// Mes 4: Añadir 500ms latencia
// Mes 5: Añadir 2s latencia
// Mes 6: Devolver error

function oldMethod(params) {
    monthsSinceDeprecation = calculateMonthsSince(DEPRECATION_DATE)

    if (monthsSinceDeprecation >= 6) {
        throw APIRemovedError()
    } else if (monthsSinceDeprecation >= 3) {
        // Latencia progresiva
        latency = monthsSinceDeprecation * 500  // 500ms, 1s, 1.5s, 2s...
        sleep(latency)

        logger.error("DEPRECATED API - SEVERE PERFORMANCE DEGRADATION", {
            latency: latency,
            message: "Migrate immediately"
        })
    }

    return execute(params)
}


// ESTRATEGIA 3: Ofrecer migración automática
// Script que migra automáticamente código de clientes
npm install -g email-api-migration-tool

emailApiMigrate --path ./src --dry-run
emailApiMigrate --path ./src --apply

// El tool hace find & replace inteligente del código
```

### 5. Rollback Durante MIGRATE

**Problema**: Durante la migración, necesitas hacer rollback.

**Solución**: Rollback es simple en cada fase.

```pseudocode
// ESCENARIO: Cliente migró a v2 pero tiene problemas

// Opción 1: Rollback del cliente (su deploy)
// Cliente simplemente revierte su código a usar v1
// API sigue soportando ambas versiones
git revert <migration-commit>
deploy()

// Opción 2: Feature flag por cliente
function process(params, clientId) {
    // Forzar uso de API vieja para cliente específico
    if (shouldUseOldAPI(clientId)) {
        return processV1(params)
    }

    return processV2(params)
}

// Configuración dinámica
config.set("client_123_use_old_api", true)  // Rollback inmediato

// Opción 3: A/B testing gradual
// Migrar solo 10% del tráfico primero
function process(params) {
    if (random() < 0.1) {  // 10% usa nuevo
        return processV2(params)
    } else {  // 90% usa viejo
        return processV1(params)
    }
}

// Si funciona bien, incrementar gradualmente
// 10% → 25% → 50% → 75% → 100%
```

## Criterios de Aceptación

Has aplicado Expand-Migrate-Contract correctamente cuando:

1. **Fase Expand**: Ambas interfaces funcionan simultáneamente sin interferencia
2. **Compatibilidad**: Clientes existentes siguen funcionando sin cambios
3. **Observabilidad**: Puedes medir qué clientes usan cada versión
4. **Comunicación**: Todos los stakeholders están informados y tienen guías
5. **Migración**: Cada cliente migra independientemente sin coordinación
6. **Verificación**: 0 uso de API vieja durante período de verificación (2+ semanas)
7. **Contracción**: Código viejo eliminado sin incidentes
8. **Zero downtime**: Sistema funcionó continuamente durante todo el proceso

## Beneficios

### Inmediatos
- Cambios grandes sin "big bang deployment"
- Rollback trivial en cada fase
- Zero downtime

### A Mediano Plazo
- Equipos migran a su propio ritmo
- Menos coordinación y menos estrés
- Testing incremental en producción

### A Largo Plazo
- Patrón reusable para futuras migraciones
- Cultura de cambios incrementales seguros
- Menos deuda técnica acumulada

## Técnicas Relacionadas

- **Golden Master**: Úsalo para validar que ambas APIs producen resultados equivalentes
- **Feature Flags**: Complementa E-M-C para controlar rollout gradual
- **Branch by Abstraction**: Variante de E-M-C para cambios en arquitectura interna
- **Strangler Fig Pattern**: E-M-C aplicado a nivel de sistema completo
- **Blue-Green Deployment**: Estrategia de deployment que complementa E-M-C
- **Canary Releases**: Testing incremental durante fase MIGRATE

## Versiones por Lenguaje

- [TypeScript](../../../../typescript/src/refactoring/parallel-change/expand-migrate-contract/) - [README](../../../../typescript/src/refactoring/parallel-change/expand-migrate-contract/README.md)
- [Go](../../../../go/refactoring/parallel-change/expand-migrate-contract/) - [README](../../../../go/refactoring/parallel-change/expand-migrate-contract/README.md)
- [Java](../../../../java/src/main/java/com/refactoring/refactoring/parallel_change/expand_migrate_contract/) - [README](../../../../java/src/main/java/com/refactoring/refactoring/parallel_change/expand_migrate_contract/README.md)
- [PHP](../../../../php/src/refactoring/parallel-change/expand-migrate-contract/) - [README](../../../../php/src/refactoring/parallel-change/expand-migrate-contract/README.md)
- [Python](../../../../python/src/refactoring/parallel_change/expand_migrate_contract/) - [README](../../../../python/src/refactoring/parallel_change/expand_migrate_contract/README.md)
- [C#](../../../../csharp/src/refactoring/parallel-change/expand-migrate-contract/) - [README](../../../../csharp/src/refactoring/parallel-change/expand-migrate-contract/README.md)

## Referencias en Español

### Artículos de Fran Iglesias

- [Refactor rompiendo cosas](https://franiglesias.github.io/refactor-by-breaking/) - Cuándo y cómo aplicar E-M-C
- [Modernizando el legacy](https://franiglesias.github.io/modernizando-el-legacy/) - Estrategia de migración con E-M-C

## Referencias en Inglés

### Artículos
- [Parallel Change](https://martinfowler.com/bliki/ParallelChange.html) - Martin Fowler
- [Expand-Contract Pattern](https://www.tim-wellhausen.de/papers/ExpandContract.pdf) - Tim Wellhausen (paper original)
- [Branch by Abstraction](https://www.branchbyabstraction.com/) - Paul Hammant
- [Transitional Change](https://www.industriallogic.com/blog/parallel-change/) - Joshua Kerievsky

### Libros
- **Working Effectively with Legacy Code** - Michael Feathers (2004)
  - Conceptos fundamentales de cambio incremental

- **Refactoring: Improving the Design of Existing Code** - Martin Fowler (2018)
  - Parallel Change pattern

### Videos y Talks
- [Workflows of Refactoring](https://www.youtube.com/watch?v=vqEg37e4Mkw) - Martin Fowler
- [Parallel Change](https://www.youtube.com/watch?v=J5SN7E_wYD8) - Danilo Sato

---

**Resumen**: Expand-Migrate-Contract es la técnica definitiva para cambios estructurales grandes en producción. Requiere más planificación que Sprout o Wrap, pero permite cambios que serían imposibles o demasiado arriesgados de otra forma.

**Volver a**: [Parallel Change - Introducción](./README.md) | [Documentación Principal](../README.md)
