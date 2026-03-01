# Code Smells - Índice de Ejercicios

Los code smells son indicadores de problemas potenciales en el diseño del código. Esta sección contiene documentación detallada de 21 code smells diferentes, organizados en 5 categorías.

## Categorías

### [Bloaters](bloaters/) (5 ejercicios)

Code smells en los que se complica el cambio en el código por producir unidades demasiado grandes o por obligarnos a introducir mucho código auxiliar.

| Code Smell | Descripción | Dificultad |
|------------|-------------|------------|
| [Data Clump](bloaters/data-clump.md) | Grupo de datos que siempre viajan juntos | ⭐⭐ |
| [Large Class](bloaters/large-class.md) | Clase con demasiadas responsabilidades | ⭐⭐⭐ |
| [Long Method](bloaters/long-method.md) | Método excesivamente largo | ⭐⭐⭐⭐ |
| [Long Parameter List](bloaters/long-parameter-list.md) | Método con muchos parámetros | ⭐⭐ |
| [Primitive Obsession](bloaters/primitive-obsession.md) | Uso excesivo de tipos primitivos | ⭐⭐ |

### [Change Preventers](change-preventers/) (3 ejercicios)

Code smells que hacen que cualquier cambio sea costoso e incluso arriesgado al obligarnos a intervenir en muchos lugares del código a la vez.

| Code Smell | Descripción | Dificultad |
|------------|-------------|------------|
| [Divergent Change](change-preventers/divergent-change.md) | Clase que cambia por múltiples razones | ⭐⭐⭐ |
| [Parallel Inheritance Hierarchy](change-preventers/parallel-inheritance-hierarchy.md) | Jerarquías que crecen en paralelo | ⭐⭐⭐ |
| [Shotgun Surgery](change-preventers/shotgun-surgery.md) | Cambio que requiere modificar múltiples clases | ⭐⭐⭐ |

### [Couplers](couplers/) (4 ejercicios)

Code smells en los que cambios en una unidad fuerzan cambios en otra que tiene un acoplamiento muy fuerte.

| Code Smell | Descripción | Dificultad |
|------------|-------------|------------|
| [Feature Envy](couplers/feature-envy.md) | Método que usa más datos de otra clase | ⭐⭐ |
| [Inappropriate Intimacy](couplers/inappropriate-intimacy.md) | Clases demasiado acopladas | ⭐⭐⭐ |
| [Message Chains](couplers/message-chains.md) | Cadenas largas de llamadas | ⭐⭐ |
| [Middleman](couplers/middleman.md) | Clase que solo delega | ⭐ |

### [Dispensables](dispensables/) (5 ejercicios)

Code smells debidos a código innecesario, que introduce ruido dificultando la inteligibilidad del código.

| Code Smell | Descripción | Dificultad |
|------------|-------------|------------|
| [Comments](dispensables/comments.md) | Comentarios que compensan código poco claro | ⭐ |
| [Data Class](dispensables/data-class.md) | Clase sin comportamiento | ⭐⭐ |
| [Dead Code](dispensables/dead-code.md) | Código que nunca se ejecuta | ⭐ |
| [Duplicated Code](dispensables/duplicated-code.md) | Código repetido | ⭐⭐ |
| [Lazy Class](dispensables/lazy-class.md) | Clase que hace muy poco | ⭐ |

### [OOP Abusers](oop-abusers/) (4 ejercicios)

Code smells debido a la aplicación inadecuada de la orientación a objetos.

| Code Smell | Descripción | Dificultad |
|------------|-------------|------------|
| [Alternative Classes with Different Interfaces](oop-abusers/alternative-classes-different-interfaces.md) | Clases similares con interfaces diferentes | ⭐⭐ |
| [Refused Bequest](oop-abusers/refused-bequest.md) | Subclase que no usa la herencia recibida | ⭐⭐⭐ |
| [Switch Statements](oop-abusers/switch-statements.md) | Switch sobre tipos en lugar de polimorfismo | ⭐⭐ |
| [Temporal Instance Variables](oop-abusers/temporal-instance-variables.md) | Variables de instancia válidas solo en ciertas fases | ⭐⭐ |

## Leyenda de Dificultad

- ⭐ **Fácil**: Refactoring directo, cambios localizados
- ⭐⭐ **Moderado**: Requiere planificación, múltiples pasos
- ⭐⭐⭐ **Avanzado**: Cambios estructurales, impacto en múltiples componentes
- ⭐⭐⭐⭐ **Experto**: Refactoring complejo, requiere técnicas avanzadas

## Cómo Usar Este Índice

### Por Categoría
Navega por categoría según el tipo de problema que observes en tu código:
- ¿Clases o métodos muy grandes? → **Bloaters**
- ¿Cambios que afectan muchos lugares? → **Change Preventers**
- ¿Dependencias muy fuertes entre clases? → **Couplers**
- ¿Código innecesario o redundante? → **Dispensables**
- ¿Problemas con herencia o polimorfismo? → **OOP Abusers**

### Por Síntomas
Identifica síntomas específicos y busca el code smell correspondiente:
- "Siempre paso los mismos 4 parámetros juntos" → Data Clump
- "Esta clase tiene más de 500 líneas" → Large Class
- "Este método tiene más de 50 líneas" → Long Method
- "Tengo que cambiar 10 archivos para añadir un nuevo tipo" → Switch Statements o Shotgun Surgery
- "Estos comentarios explican qué hace el código" → Comments
- "Esta clase solo tiene getters y setters" → Data Class

### Ruta de Aprendizaje Sugerida

#### Principiantes
1. Comments
2. Dead Code
3. Duplicated Code
4. Lazy Class
5. Middleman

#### Intermedios
6. Data Clump
7. Long Parameter List
8. Primitive Obsession
9. Feature Envy
10. Message Chains
11. Data Class
12. Switch Statements
13. Alternative Classes

#### Avanzados
14. Large Class
15. Inappropriate Intimacy
16. Temporal Instance Variables
17. Divergent Change
18. Shotgun Surgery
19. Parallel Inheritance Hierarchy
20. Refused Bequest

#### Ejercicio Final
21. Long Method (integra múltiples smells)

## Recursos Adicionales

- [Volver al índice principal](../README.md)
- [Refactoring Guru - Code Smells](https://refactoring.guru/refactoring/smells)
- [Martin Fowler - Catalog of Refactorings](https://refactoring.com/catalog/)

## Implementaciones

Cada ejercicio está implementado en 6 lenguajes:
- [TypeScript](../../typescript/)
- [Go](../../go/)
- [Java](../../java/)
- [PHP](../../php/)
- [Python](../../python/)
- [C#](../../csharp/)
