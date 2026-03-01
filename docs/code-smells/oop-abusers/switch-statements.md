# Switch Statements

Sentencias switch.

## Definición

El uso excesivo de `switch` o múltiples `if/else` basados en un código de tipo suele ser una señal de que falta polimorfismo. El problema principal es que cada vez que se añade una nueva variante (un nuevo tipo), hay que buscar y modificar todos los bloques `switch` dispersos por la aplicación.

## Descripción

**Switch Statements** como code smell no significa que todo `switch` sea malo. El problema aparece cuando usas `switch` o cadenas largas de `if/else` basadas en códigos de tipo, y este patrón se repite en múltiples lugares del código.

Este smell es problemático porque:
- Cada nueva variante requiere modificar todos los switches
- Es fácil olvidar actualizar algún switch
- Viola el principio Open/Closed (abierto a extensión, cerrado a modificación)
- Duplica la lógica de discriminación de tipos
- No aprovecha el polimorfismo orientado a objetos

En programación orientada a objetos, el polimorfismo permite que diferentes clases proporcionen diferentes implementaciones del mismo método, eliminando la necesidad de switches basados en tipo.

## Síntomas

- Múltiples `switch` o `if/else` basados en el mismo código de tipo
- Los switches aparecen en diferentes métodos o clases
- Añadir un nuevo tipo requiere modificar múltiples switches
- Los cases del switch contienen lógica significativa
- El mismo tipo de discriminación aparece repetidamente
- Tests que verifican todos los casos de cada switch

## Ejemplo

```pseudocode
type EmployeeKind = "engineer" | "manager" | "sales"

type EmployeeRecord = {
  kind: EmployeeKind
  base: number
  bonus: number
  commission: number
}

function calculatePay(employee: EmployeeRecord): number {
  switch (employee.kind) {
    case "engineer":
      return employee.base
    case "manager":
      return employee.base + employee.bonus
    case "sales":
      return employee.base + employee.commission
    default:
      throw error "Unknown employee type"
  }
}

// Si añades "contractor", debes modificar este switch
// y cualquier otro switch que discrimine por tipo de empleado
```

## Ejercicio

Añade un nuevo tipo de empleado (`contractor`) con una regla de pago especial (por ejemplo, tarifa por horas).

## Problemas que encontrarás

Tendrás que modificar el `switch` y cualquier otro código que dependa del tipo de empleado. A medida que el sistema crece, olvidar actualizar uno de estos puntos genera bugs difíciles de detectar.

## Proceso de Refactoring

### 1. Identificar switches relacionados
- Encuentra todos los switches que discriminan sobre el mismo tipo
- Mapea qué operaciones se realizan para cada tipo
- Identifica el comportamiento que varía según el tipo

### 2. Crear jerarquía de clases
- Define una clase base o interfaz para el concepto
- Crea subclases para cada tipo/variante
- Ejemplo: `Employee` (base), `Engineer`, `Manager`, `SalesRep` (subclases)

### 3. Mover comportamiento a las subclases
- Cada caso del switch se convierte en un método en la subclase correspondiente
- Ejemplo: `calculatePay()` se implementa diferente en cada subclase
- Usa **Replace Type Code with Subclasses**

### 4. Reemplazar switch con llamada polimórfica
- En lugar de: `switch(employee.kind) { ... }`
- Usa: `employee.calculatePay()`
- El polimorfismo selecciona automáticamente la implementación correcta

### 5. Eliminar el campo de tipo si es posible
- Si el único propósito del campo era discriminar en switches
- Considera eliminarlo una vez que uses polimorfismo
- El tipo de la clase misma es la discriminación

### 6. Usar patrón Strategy si las variantes son configurables
- Si las variantes no son tipos inherentes del objeto
- Usa el patrón Strategy en lugar de subclases
- Permite cambiar comportamiento en tiempo de ejecución

## Técnicas de Refactoring Aplicables

- **Replace Type Code with Subclasses**: Crear jerarquía de clases
- **Replace Type Code with State/Strategy**: Usar patrones de comportamiento
- **Replace Conditional with Polymorphism**: Eliminar switches con polimorfismo
- **Extract Method**: Extraer cada caso a un método antes de mover a subclase
- **Move Method**: Mover comportamiento a la clase apropiada

## Beneficios

- **Open/Closed Principle**: Añadir tipos sin modificar código existente
- **Un lugar por operación**: Cada comportamiento en su subclase
- **Imposible olvidar casos**: El compilador ayuda con métodos abstractos
- **Código más limpio**: Sin switches largos y complejos
- **Mejor extensibilidad**: Nuevos tipos son nuevas clases
- **Polimorfismo natural**: Aprovecha características OOP
- **Testing más fácil**: Testear cada tipo independientemente

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/oop-abusers/switch-statements.ts) - [README](../../typescript/src/code-smells/oop-abusers/switch-statements.readme.md)
- [Go](../../go/code_smells/oop_abusers/switch_statements.go) - [README](../../go/code_smells/oop_abusers/switch_statements.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/oopabusers/SwitchStatements.java) - [README](../../java/src/main/java/com/refactoring/codesmells/oopabusers/SwitchStatements.readme.md)
- [PHP](../../php/src/code-smells/oop-abusers/SwitchStatements.php) - [README](../../php/src/code-smells/oop-abusers/SwitchStatements.readme.md)
- [Python](../../python/src/code_smells/oop_abusers/switch_statements.py) - [README](../../python/src/code_smells/oop_abusers/switch_statements_readme.md)
- [C#](../../csharp/src/code-smells/oop-abusers/SwitchStatements.cs) - [README](../../csharp/src/code-smells/oop-abusers/switch-statements.readme.md)

## Referencias en Español

- [Polimorfismo y extensibilidad de objetos](https://franiglesias.github.io/polimorfismo-y-extensibilidad-de-objetos/) - Uso de polimorfismo para reemplazar switches y mejorar extensibilidad
- [Los principios SOLID (Open/Closed)](https://franiglesias.github.io/principios-solid/) - Principio Open/Closed para código extensible sin modificación

## Referencias

- [Refactoring Guru - Switch Statements](https://refactoring.guru/smells/switch-statements)
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
- Robert C. Martin - "Agile Software Development" - Open/Closed Principle
- Gang of Four - "Design Patterns" - Strategy, State patterns
