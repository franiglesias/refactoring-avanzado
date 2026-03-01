# Técnicas de Refactoring Avanzado

## Introducción

Este módulo contiene técnicas avanzadas de refactoring diseñadas específicamente para trabajar con **código legacy** y realizar **cambios seguros** en sistemas en producción, especialmente cuando no existen tests o la cobertura es insuficiente.

A diferencia del refactoring tradicional que asume la existencia de tests, estas técnicas reconocen la realidad de muchos proyectos: código heredado sin documentación, sin tests, con alta complejidad y que debe mantenerse funcionando en producción mientras se mejora.

## ¿Cuándo Usar Estas Técnicas?

Estas técnicas son especialmente útiles cuando te enfrentas a:

- **Código legacy sin tests**: Sistema heredado sin cobertura de pruebas
- **Riesgo alto de ruptura**: Cambios que pueden afectar funcionalidad crítica en producción
- **Desconocimiento del comportamiento**: No está claro qué hace exactamente el código
- **Refactorings grandes**: Cambios que no pueden hacerse en un solo paso
- **Necesidad de despliegue continuo**: El código debe seguir funcionando mientras se refactoriza
- **Sistemas en producción**: No puedes detener el servicio para hacer cambios
- **Falta de documentación**: El código es la única fuente de verdad

## Técnicas Incluidas

### Golden Master

La técnica **Golden Master** (también conocida como Approval Testing o Characterization Testing) permite crear una red de seguridad de tests cuando no existen, capturando el comportamiento actual del sistema como "la verdad" y detectando automáticamente cualquier desviación.

Es la base para poder refactorizar código legacy de forma segura.

**[Ver documentación completa de Golden Master](./golden-master/README.md)**

### Parallel Change

**Parallel Change** es un conjunto de técnicas para realizar cambios en código en producción de forma incremental y segura, sin romper la funcionalidad existente. Permite que el código viejo y nuevo coexistan durante un periodo de transición.

Las tres técnicas principales son:

#### 1. Sprout Change (Técnica del Brote)

Añade nueva funcionalidad como código nuevo y aislado, sin modificar el código existente. Ideal para agregar comportamiento nuevo sin tocar el código legacy.

**[Ver documentación de Sprout Change](./parallel-change/sprout-change.md)**

#### 2. Wrap Change (Técnica del Envoltorio)

Envuelve el código existente en una nueva interfaz o capa, permitiendo interceptar, modificar o extender el comportamiento sin cambiar el código original.

**[Ver documentación de Wrap Change](./parallel-change/wrap-change.md)**

#### 3. Expand-Migrate-Contract (Expandir-Migrar-Contraer)

Técnica en tres fases para cambios estructurales grandes: primero se expande la API para soportar ambas versiones, luego se migran los clientes al nuevo código, finalmente se elimina el código viejo.

**[Ver documentación de Expand-Migrate-Contract](./parallel-change/expand-migrate-contract.md)**

**[Ver introducción a Parallel Change](./parallel-change/README.md)**

## Tabla Comparativa de Técnicas

| Técnica | Contexto Ideal | Tipo de Cambio | Duración | Requiere Tests Previos |
|---------|---------------|----------------|----------|----------------------|
| **Golden Master** | Sin tests, comportamiento desconocido | Preparación para refactoring | Una vez (preparación) | No (los crea) |
| **Sprout Change** | Añadir funcionalidad nueva | Adición de código | Corta (minutos/horas) | No (pero ayuda) |
| **Wrap Change** | Cambiar comportamiento existente | Envoltorio/decoración | Media (horas/días) | No (pero ayuda) |
| **Expand-Migrate-Contract** | Cambios estructurales grandes | Transformación gradual | Larga (días/semanas) | Sí (o Golden Master) |

## ¿Qué Técnica Usar?

### Usa Golden Master cuando:
- No tienes tests y necesitas crear una red de seguridad
- No sabes exactamente qué hace el código
- Quieres documentar el comportamiento actual
- Vas a hacer refactorings que cambien la estructura interna sin cambiar el comportamiento

### Usa Sprout Change cuando:
- Necesitas añadir funcionalidad nueva
- No quieres tocar el código legacy existente
- Puedes hacer el cambio de forma aislada
- La nueva funcionalidad es relativamente pequeña

### Usa Wrap Change cuando:
- Necesitas modificar o extender comportamiento existente
- No puedes o no quieres modificar el código original
- Quieres interceptar llamadas para añadir logging, validación, etc.
- Quieres aplicar el patrón Decorator o Proxy

### Usa Expand-Migrate-Contract cuando:
- Necesitas cambiar interfaces públicas o contratos
- Tienes múltiples clientes que dependen del código actual
- El cambio es demasiado grande para hacerlo de una vez
- Necesitas mantener compatibilidad durante la migración
- Trabajas en un sistema distribuido o con múltiples equipos

## Rutas de Aprendizaje Sugeridas

### Ruta 1: Principiante en Legacy Code
1. **Golden Master**: Aprende a crear tests de caracterización
2. **Sprout Change**: Practica añadir funcionalidad sin tocar legacy
3. **Wrap Change**: Practica envolver código legacy
4. **Expand-Migrate-Contract**: Aplica a cambios más grandes

### Ruta 2: Cambio Urgente en Producción
1. **Sprout Change**: Añade el cambio de forma aislada
2. **Golden Master**: Si hay tiempo, crea tests después
3. **Expand-Migrate-Contract**: Para migración posterior

### Ruta 3: Refactoring Profundo
1. **Golden Master**: Crea red de seguridad primero
2. **Expand-Migrate-Contract**: Planifica la migración
3. **Sprout/Wrap**: Usa según necesites durante el proceso

## Principios Fundamentales

Todas estas técnicas se basan en principios comunes:

1. **Seguridad primero**: No rompas lo que funciona
2. **Cambios incrementales**: Pasos pequeños y verificables
3. **Reversibilidad**: Siempre debe ser posible volver atrás
4. **Coexistencia**: Código viejo y nuevo pueden convivir temporalmente
5. **Medición**: Captura el comportamiento antes de cambiar
6. **Verificación continua**: Tests en cada paso

## Implementaciones por Lenguaje

Cada técnica tiene implementaciones prácticas en 6 lenguajes:

- [TypeScript](../../typescript/src/refactoring/)
- [Go](../../go/refactoring/)
- [Java](../../java/src/main/java/com/refactoring/refactoring/)
- [PHP](../../php/src/refactoring/)
- [Python](../../python/src/refactoring/)
- [C#](../../csharp/src/refactoring/)

## Referencias en Español

### Artículos de Fran Iglesias

**Sobre Legacy Code y Refactoring:**
- [El código en producción es legacy](https://franiglesias.github.io/prod-code-is-legacy/)
- [Refactor rompiendo cosas](https://franiglesias.github.io/refactor-by-breaking/)
- [Modernizando el legacy](https://franiglesias.github.io/modernizando-el-legacy/)
- [Ejercicio de refactor](https://franiglesias.github.io/ejercicio-de-refactor-1/)

**Sobre Golden Master:**
- [Approval Testing](https://franiglesias.github.io/approval_testing/)
- [Golden Master - Cookbook](https://franiglesias.github.io/golden-cookbook-master-approval/)
- [Quotebot Kata](https://franiglesias.github.io/quotebot-kata/)

**Sobre Técnicas de Cambio Incremental:**
- [Introducción al Refactor](https://franiglesias.github.io/intro_refactor_1/)
- [Refactoring - Camp Rule](https://franiglesias.github.io/refactoring-camp-rule/)

## Referencias en Inglés

### Libros Fundamentales
- **Working Effectively with Legacy Code** - Michael Feathers (2004)
  - Fuente original de las técnicas Sprout y Wrap
  - Capítulos 6-10 sobre Breaking Dependencies

- **Refactoring: Improving the Design of Existing Code** - Martin Fowler (1999, 2nd ed. 2018)
  - Catálogo de refactorings seguros

- **Growing Object-Oriented Software, Guided by Tests** - Steve Freeman & Nat Pryce (2009)
  - Aproximación a legacy code con testing

### Artículos y Blogs
- [Parallel Change Pattern](https://martinfowler.com/bliki/ParallelChange.html) - Martin Fowler
- [Expand-Contract Pattern](https://martinfowler.com/bliki/ParallelChange.html) - Danilo Sato
- [Branch by Abstraction](https://www.branchbyabstraction.com/) - Paul Hammant
- [ApprovalTests.com](https://approvaltests.com/) - Llewellyn Falco

### Videos y Charlas
- [Testing Legacy Code](https://www.youtube.com/watch?v=_NnElPO5BU0) - Sandro Mancuso
- [Get Your Legacy Code Under Test](https://www.youtube.com/watch?v=LDKJRjKPYZ8) - Michael Feathers

## Recursos Adicionales

- **Katas recomendadas**: Gilded Rose, Tennis Refactoring, Trivia
- **Herramientas**: ApprovalTests, TextTest, Jest Snapshots
- **Comunidad**: Software Crafters, Legacy Code Rocks podcast

---

**Nota**: Este es material educativo para aprender técnicas profesionales de refactoring. Cada técnica incluye ejemplos prácticos y ejercicios en múltiples lenguajes de programación.
