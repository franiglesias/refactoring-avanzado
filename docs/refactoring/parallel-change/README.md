# Parallel Change (Cambio Paralelo)

## Definición

**Parallel Change** (también conocido como Expand-Contract o Transitional Change) es un conjunto de técnicas que permiten realizar cambios en código en producción de forma incremental y segura, permitiendo que el código antiguo y el nuevo coexistan temporalmente durante un periodo de transición.

En lugar de hacer un cambio grande y arriesgado de una sola vez ("big bang deployment"), Parallel Change divide el cambio en pasos pequeños y reversibles, donde cada paso deja el sistema en un estado funcional.

## Cuándo Usar

- Cuando necesitas cambiar código que está en producción activamente
- Cuando múltiples equipos o sistemas dependen del código que vas a cambiar
- Cuando no puedes coordinar un despliegue simultáneo de todos los componentes
- Cuando necesitas mantener compatibilidad hacia atrás durante una migración
- Cuando quieres reducir el riesgo de un cambio grande
- Cuando trabajas en sistemas distribuidos o microservicios
- Cuando necesitas poder revertir cambios fácilmente

## Problema que Resuelve

Los cambios grandes en producción son arriesgados:

- **Despliegues "big bang"**: Todo o nada, si algo falla, todo se rompe
- **Incompatibilidad**: Clientes viejos dejan de funcionar inmediatamente
- **Coordinación compleja**: Múltiples equipos deben desplegar exactamente al mismo tiempo
- **Rollback difícil**: Revertir puede ser tan complejo como el despliegue original
- **Tiempo de inactividad**: Puede requerir apagar el sistema durante el cambio
- **Testing insuficiente**: Es difícil probar el cambio completo antes del despliegue

Parallel Change resuelve esto permitiendo:
1. Cambios incrementales y seguros
2. Coexistencia temporal de código viejo y nuevo
3. Despliegues independientes sin coordinación
4. Rollback trivial (solo deshacer el último paso)
5. Testing en producción con tráfico real
6. Zero downtime deployments

## Técnicas de Parallel Change

Existen tres técnicas principales dentro de Parallel Change:

### 1. Sprout Change (Técnica del Brote)

**Definición**: Añade nueva funcionalidad como código completamente nuevo y aislado, sin modificar el código existente. El código viejo llama al nuevo código en el punto necesario.

**Cuándo usar**: Para añadir funcionalidad nueva sin tocar código legacy.

**[Ver documentación completa de Sprout Change](./sprout-change.md)**

### 2. Wrap Change (Técnica del Envoltorio)

**Definición**: Envuelve el código existente en una nueva capa o interfaz, permitiendo interceptar, modificar o extender el comportamiento sin cambiar el código original.

**Cuándo usar**: Para cambiar o extender comportamiento existente sin modificar el código legacy.

**[Ver documentación completa de Wrap Change](./wrap-change.md)**

### 3. Expand-Migrate-Contract (Expandir-Migrar-Contraer)

**Definición**: Técnica en tres fases: primero expandes la interfaz/API para soportar tanto el uso viejo como el nuevo, luego migras gradualmente todos los clientes al nuevo uso, finalmente contraes eliminando el código viejo.

**Cuándo usar**: Para cambios estructurales grandes, refactorings de APIs, o migraciones de sistemas.

**[Ver documentación completa de Expand-Migrate-Contract](./expand-migrate-contract.md)**

## Comparación de Técnicas

| Aspecto | Sprout Change | Wrap Change | Expand-Migrate-Contract |
|---------|---------------|-------------|------------------------|
| **Complejidad** | Baja | Media | Alta |
| **Duración** | Horas | Días | Semanas/Meses |
| **Alcance** | Funcionalidad nueva localizada | Comportamiento de componente | APIs o sistema completo |
| **Modificación código viejo** | Mínima (1 llamada) | Ninguna (wrap externo) | Media (adaptación) |
| **Pasos** | 1-2 | 2-3 | 3+ |
| **Reversibilidad** | Muy fácil | Fácil | Moderada |
| **Testing** | Fácil (código aislado) | Moderado | Complejo (ambas versiones) |
| **Ejemplo** | Añadir validación nueva | Añadir logging a clase legacy | Cambiar API REST versión |

## Diagrama General del Proceso

```
ESTADO INICIAL: Código Legacy
┌─────────────────────────────────────┐
│  Sistema Viejo                      │
│  (Funcionando en producción)       │
└─────────────────────────────────────┘


TÉCNICA 1: SPROUT CHANGE
┌─────────────────────────────────────┐
│  Sistema Viejo                      │
│         │                           │
│         ├──> [Nueva Funcionalidad]  │◄─── Código nuevo aislado
│         │                           │
│  (Sin modificar el código viejo)   │
└─────────────────────────────────────┘


TÉCNICA 2: WRAP CHANGE
┌─────────────────────────────────────┐
│  [Wrapper/Decorador]                │◄─── Capa nueva
│         │                           │
│         ▼                           │
│  ┌──────────────┐                  │
│  │Sistema Viejo │                  │◄─── Sin modificar
│  └──────────────┘                  │
└─────────────────────────────────────┘


TÉCNICA 3: EXPAND-MIGRATE-CONTRACT

Fase 1: EXPAND
┌─────────────────────────────────────┐
│  API Expandida                      │
│  ├──> [Interfaz Vieja] ─┐          │
│  └──> [Interfaz Nueva] ─┤          │
│                          ▼          │
│                  [Implementación]   │
└─────────────────────────────────────┘

Fase 2: MIGRATE
┌─────────────────────────────────────┐
│  Clientes migrando...               │
│  Cliente A ──> [Interfaz Nueva]     │
│  Cliente B ──> [Interfaz Nueva]     │
│  Cliente C ──> [Interfaz Vieja]     │◄─── En proceso
└─────────────────────────────────────┘

Fase 3: CONTRACT
┌─────────────────────────────────────┐
│  API Contraída                      │
│  [Interfaz Nueva]                   │◄─── Solo nueva
│         │                           │
│  (Código viejo eliminado)          │
└─────────────────────────────────────┘
```

## Principios Fundamentales

Todas las técnicas de Parallel Change comparten estos principios:

### 1. Compatibilidad Temporal

El sistema debe funcionar con el código viejo, el nuevo, o ambos simultáneamente durante la transición.

```pseudocode
// Mal: Cambio abrupto
function process(data) {
    return newImplementation(data)  // Rompe clientes viejos
}

// Bien: Compatibilidad durante transición
function process(data, useNewVersion = false) {
    if (useNewVersion) {
        return newImplementation(data)
    }
    return oldImplementation(data)  // Sigue funcionando
}
```

### 2. Pasos Pequeños y Verificables

Cada cambio debe ser lo suficientemente pequeño para:
- Ser revisado en minutos
- Ser testeado fácilmente
- Ser desplegado independientemente
- Ser revertido sin impacto

### 3. Estado Funcional Continuo

Después de cada paso, el sistema debe:
- Compilar correctamente
- Pasar todos los tests
- Poder desplegarse a producción
- Funcionar con tráfico real

### 4. Camino de Migración Claro

Debe estar claro:
- Qué código es viejo y qué es nuevo
- Qué clientes usan cada versión
- Cuándo completar cada fase
- Cómo medir el progreso

### 5. Estrategia de Rollback

En cualquier momento debe ser posible:
- Revertir al paso anterior
- Volver a usar código viejo
- Deshacer sin pérdida de datos

## Cuándo Usar Cada Técnica

### Usa Sprout Change si:
- Vas a AÑADIR funcionalidad nueva
- No necesitas cambiar comportamiento existente
- Puedes hacer el cambio de forma aislada
- El cambio es relativamente pequeño (horas/días)
- Quieres minimizar el riesgo

### Usa Wrap Change si:
- Vas a MODIFICAR o EXTENDER comportamiento existente
- No puedes o no quieres tocar el código legacy
- Necesitas interceptar llamadas (logging, validación, etc.)
- Quieres aplicar patrones Decorator o Proxy
- El cambio afecta a una clase o componente específico

### Usa Expand-Migrate-Contract si:
- Vas a cambiar INTERFACES PÚBLICAS o CONTRATOS
- Múltiples equipos o sistemas dependen del código
- El cambio es estructural y grande (semanas/meses)
- Necesitas mantener compatibilidad durante migración
- Trabajas en microservicios o sistemas distribuidos

## Ejemplo Comparativo

Imaginemos que tenemos una función de cálculo de precio que necesita cambios:

```pseudocode
// CÓDIGO ORIGINAL
function calculatePrice(item) {
    return item.basePrice * 1.21  // IVA fijo 21%
}
```

### Escenario 1: Añadir descuento por volumen (Sprout)

```pseudocode
// SPROUT: Añadir nueva función aislada
function calculateVolumeDiscount(quantity, basePrice) {
    if (quantity >= 100) return basePrice * 0.9
    if (quantity >= 50) return basePrice * 0.95
    return basePrice
}

function calculatePrice(item) {
    discountedPrice = calculateVolumeDiscount(item.quantity, item.basePrice)
    return discountedPrice * 1.21
}
```

### Escenario 2: Añadir logging sin tocar código (Wrap)

```pseudocode
// WRAP: Envolver función original
function calculatePriceWithLogging(item) {
    log("Calculating price for item: " + item.id)
    startTime = now()

    result = calculatePrice(item)  // Código original sin tocar

    log("Price calculated: " + result + " in " + (now() - startTime) + "ms")
    return result
}

// Usar el wrapper
price = calculatePriceWithLogging(item)
```

### Escenario 3: Cambiar IVA variable por país (Expand-Migrate-Contract)

```pseudocode
// FASE 1: EXPAND - Soportar ambas interfaces
function calculatePrice(item, country = null) {
    basePrice = item.basePrice

    // Nuevo comportamiento
    if (country != null) {
        taxRate = getTaxRateForCountry(country)
        return basePrice * (1 + taxRate)
    }

    // Comportamiento viejo (compatibilidad)
    return basePrice * 1.21
}

// FASE 2: MIGRATE - Cambiar clientes uno a uno
// Cliente 1: calculatePrice(item, "ES")  ✓ Migrado
// Cliente 2: calculatePrice(item, "FR")  ✓ Migrado
// Cliente 3: calculatePrice(item)        ⏳ Pendiente

// FASE 3: CONTRACT - Eliminar código viejo
function calculatePrice(item, country) {
    taxRate = getTaxRateForCountry(country)
    return item.basePrice * (1 + taxRate)
}
```

## Errores Comunes

### 1. No Planificar la Eliminación del Código Viejo

**Problema**: El código viejo queda indefinidamente, acumulando deuda técnica.

**Solución**:
- Establecer fecha límite para eliminar código viejo
- Documentar qué código es temporal
- Usar feature flags con expiración
- Crear tickets para cleanup

### 2. Hacer Pasos Demasiado Grandes

**Problema**: "Un paso más" que en realidad son múltiples cambios complejos.

**Solución**:
- Si el paso toma más de un día, dividirlo
- Cada commit debe ser desplegable independientemente
- Si no puedes explicar el paso en una frase, es muy grande

### 3. No Testear Ambas Versiones

**Problema**: Solo se testea el código nuevo, el viejo se rompe.

**Solución**:
- Mantener tests para código viejo
- Añadir tests para código nuevo
- Tests de integración que verifiquen compatibilidad

### 4. No Medir el Progreso de Migración

**Problema**: No sabes cuántos clientes siguen usando código viejo.

**Solución**:
- Añadir métricas/logging
- Dashboard de adopción
- Alertas cuando código viejo sigue en uso pasado deadline

### 5. No Comunicar el Cambio

**Problema**: Otros equipos no saben que hay que migrar.

**Solución**:
- Documentar plan de migración
- Comunicar deprecaciones claramente
- Ofrecer soporte durante migración

## Criterios de Éxito

Has aplicado Parallel Change correctamente cuando:

1. **Zero downtime**: El sistema nunca dejó de funcionar durante el cambio
2. **Rollback fácil**: Pudiste revertir cualquier paso sin problemas
3. **Sin big bang**: No hubo un único punto crítico de despliegue
4. **Clientes felices**: Los consumidores del código no experimentaron rupturas
5. **Código limpio al final**: No queda código muerto o deuda técnica
6. **Métricas claras**: Tienes visibilidad del progreso de migración
7. **Testing completo**: Ambas versiones estuvieron testeadas durante transición

## Beneficios

### Técnicos
- Reduce riesgo de cambios grandes
- Permite rollback trivial
- Facilita testing en producción
- Mantiene sistema funcional continuamente

### De Negocio
- Zero downtime deployments
- Despliegues independientes por equipo
- Migración gradual y controlada
- Menos incidentes de producción

### De Equipo
- Menos estrés en despliegues
- Más confianza en los cambios
- Mejor colaboración entre equipos
- Aprendizaje continuo

## Técnicas Relacionadas

- **Golden Master**: Úsalo antes de Parallel Change para crear red de seguridad
- **Feature Flags**: Complementa Expand-Migrate-Contract para controlar adopción
- **Branch by Abstraction**: Variante de Expand-Migrate-Contract para cambios en arquitectura
- **Strangler Fig Pattern**: Aplicación de Parallel Change a nivel de sistema completo
- **Blue-Green Deployment**: Estrategia de despliegue compatible con Parallel Change

## Versiones por Lenguaje

- [TypeScript](../../../typescript/src/refactoring/parallel-change/) - [README](../../../typescript/src/refactoring/parallel-change/README.md)
- [Go](../../../go/refactoring/parallel-change/) - [README](../../../go/refactoring/parallel-change/README.md)
- [Java](../../../java/src/main/java/com/refactoring/refactoring/parallel_change/) - [README](../../../java/src/main/java/com/refactoring/refactoring/parallel_change/README.md)
- [PHP](../../../php/src/refactoring/parallel-change/) - [README](../../../php/src/refactoring/parallel-change/README.md)
- [Python](../../../python/src/refactoring/parallel_change/) - [README](../../../python/src/refactoring/parallel_change/README.md)
- [C#](../../../csharp/src/refactoring/parallel-change/) - [README](../../../csharp/src/refactoring/parallel-change/README.md)

## Referencias en Español

### Artículos de Fran Iglesias

- [Refactor rompiendo cosas](https://franiglesias.github.io/refactor-by-breaking/) - Cuándo y cómo romper compatibilidad de forma segura
- [Modernizando el legacy](https://franiglesias.github.io/modernizando-el-legacy/) - Estrategias de migración gradual
- [Introducción al Refactor](https://franiglesias.github.io/intro_refactor_1/) - Técnicas Sprout y Wrap explicadas
- [Refactoring - Camp Rule](https://franiglesias.github.io/refactoring-camp-rule/) - Mejora incremental del código

## Referencias en Inglés

### Libros
- **Working Effectively with Legacy Code** - Michael Feathers (2004)
  - Capítulo 6: "I Don't Have Much Time and I Have to Change It" (Sprout)
  - Capítulo 7: "It Takes Forever to Make a Change" (Wrap)

- **Refactoring: Improving the Design of Existing Code** - Martin Fowler (2018)
  - Parallel Change pattern

### Artículos
- [Parallel Change](https://martinfowler.com/bliki/ParallelChange.html) - Martin Fowler
- [Expand-Contract Pattern](https://www.tim-wellhausen.de/papers/ExpandContract.pdf) - Tim Wellhausen
- [Branch by Abstraction](https://www.branchbyabstraction.com/) - Paul Hammant
- [Strangler Fig Application](https://martinfowler.com/bliki/StranglerFigApplication.html) - Martin Fowler

### Videos
- [Advanced Testing Patterns](https://www.youtube.com/watch?v=_NnElPO5BU0) - Sandro Mancuso
- [Refactoring to Microservices](https://www.youtube.com/watch?v=64NgHs0T0uA) - Sam Newman

---

**Próximos Pasos**: Profundiza en cada técnica específica:
1. [Sprout Change](./sprout-change.md) - Para añadir funcionalidad nueva
2. [Wrap Change](./wrap-change.md) - Para modificar comportamiento existente
3. [Expand-Migrate-Contract](./expand-migrate-contract.md) - Para cambios estructurales grandes
