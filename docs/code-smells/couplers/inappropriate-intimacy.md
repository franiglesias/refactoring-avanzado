# Inappropriate Intimacy

Intimidad inapropiada.

## Definición

Dos clases tienen intimidad inapropiada cuando pueden acceder y manipular el estado interno de la otra, creando un acoplamiento fuerte y diseños frágiles.

## Descripción

**Inappropriate Intimacy** ocurre cuando dos clases conocen demasiado sobre los detalles internos de la otra. Están excesivamente acopladas, accediendo directamente a los campos privados (o deberían ser privados), llamando a métodos internos, y generalmente violando la encapsulación.

Este smell se manifiesta cuando:
- Las clases se modifican mutuamente el estado
- Una clase depende de la implementación interna de otra
- Ambas clases tienen referencias circulares
- Cambiar una clase frecuentemente requiere cambiar la otra
- Las clases parecen estar divididas artificialmente

Es diferente de **Feature Envy** porque aquí el acoplamiento es bidireccional: ambas clases están íntimamente involucradas en los asuntos de la otra.

## Síntomas

- Referencias bidireccionales entre clases
- Una clase modifica directamente los campos de otra
- Métodos que exponen detalles internos solo para otra clase
- Las clases frecuentemente cambian juntas
- Dificultad para entender o usar una clase sin la otra
- Conversaciones complejas entre dos clases
- Tests que requieren configurar ambas clases juntas

## Ejemplo

```pseudocode
class Team {
  name: string
  budget: Budget
  manager: Manager

  function assignManager(m: Manager) {
    manager = m
    m.assignTeam(this)  // Bidireccional
  }

  function raiseBudget(amount: number) {
    budget.amount = budget.amount + amount  // Acceso directo
  }

  function rename(newName: string) {
    name = newName
  }
}

class Manager {
  name: string
  team: Team

  function assignTeam(t: Team) {
    if (team is not null) {
      throw error "Team already assigned"
    }
    team = t
  }

  function raiseTeamBudget(amount: number) {
    if (team is not null) {
      team.raiseBudget(amount)  // Manager manipula Team
    }
  }

  function renameTeam(newName: string) {
    if (team is not null) {
      team.rename(newName)  // Manager manipula Team
    }
  }
}

class Budget {
  amount: number  // Público, no encapsulado

  function raise(amount: number) {
    if (amount + this.amount > 3000) {
      throw error "Budget exceeded"
    }
    amount = amount + this.amount
  }
}
```

## Ejercicio

Añade una traza de auditoría cuando cambien los presupuestos y aplica reglas de presupuesto mínimo.

## Problemas que encontrarás

Como Team y Manager tocan libremente los campos del otro, tendrás que esparcir comprobaciones y registros en muchos lugares, aumentando el acoplamiento y las regresiones.

## Proceso de Refactoring

### 1. Identificar la intimidad excesiva
- Mapea las interacciones entre las clases
- Identifica qué información comparten
- Verifica si la relación es realmente bidireccional necesaria

### 2. Encapsular estado interno
- Hacer privados los campos que no deberían ser públicos
- Eliminar setters que permiten modificación directa
- Exponer solo la interfaz pública necesaria

### 3. Extraer clase para responsabilidades compartidas
- Si ambas clases comparten lógica relacionada
- Crea una tercera clase que contenga esa responsabilidad
- Ejemplo: `TeamAssignment` que gestiona la relación Team-Manager

### 4. Eliminar referencias bidireccionales
- Analiza si realmente necesitas navegación en ambas direcciones
- Considera hacer la relación unidireccional
- Si es necesario bidireccional, centraliza la gestión de la relación

### 5. Aplicar Law of Demeter
- Un método debe llamar solo a métodos de:
  - El objeto mismo
  - Parámetros del método
  - Objetos que crea
  - Campos directos
- No debe navegar a través de objetos retornados

### 6. Mover métodos que manipulan otra clase
- Si Manager tiene métodos que principalmente manipulan Team
- Considera si esos métodos deberían estar en Team
- Usa **Move Method** para reequilibrar responsabilidades

## Técnicas de Refactoring Aplicables

- **Extract Class**: Crear una clase para gestionar la relación
- **Move Method**: Mover comportamiento a la clase apropiada
- **Move Field**: Mover datos a donde se usan
- **Change Bidirectional Association to Unidirectional**: Simplificar la relación
- **Hide Delegate**: Ocultar la navegación entre objetos
- **Encapsulate Field**: Proteger acceso directo a campos

## Beneficios

- **Mejor encapsulación**: Cada clase protege su estado interno
- **Menos acoplamiento**: Las clases pueden evolucionar independientemente
- **Código más robusto**: Cambios en una clase no rompen la otra
- **Testing más fácil**: Puedes testear clases en aislamiento
- **Mantenimiento simplificado**: Entender y modificar cada clase es más fácil
- **Menos efectos secundarios**: Cambios localizados y predecibles

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/couplers/inappropriate-intimacy.ts) - [README](../../typescript/src/code-smells/couplers/inappropriate-intimacy.readme.md)
- [Go](../../go/code_smells/couplers/inappropriate_intimacy.go) - [README](../../go/code_smells/couplers/inappropriate_intimacy.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/couplers/InappropriateIntimacy.java) - [README](../../java/src/main/java/com/refactoring/codesmells/couplers/InappropriateIntimacy.readme.md)
- [PHP](../../php/src/code-smells/couplers/InappropriateIntimacy.php) - [README](../../php/src/code-smells/couplers/InappropriateIntimacy.readme.md)
- [Python](../../python/src/code_smells/couplers/inappropriate_intimacy.py) - [README](../../python/src/code_smells/couplers/inappropriate_intimacy_readme.md)
- [C#](../../csharp/src/code-smells/couplers/InappropriateIntimacy.cs) - [README](../../csharp/src/code-smells/couplers/inappropriate-intimacy.readme.md)

## Referencias en Español

- [Dependencias y acoplamiento](https://franiglesias.github.io/dependencias-acoplamiento/) - Análisis detallado de tipos de acoplamiento y cómo gestionarlo
- [Evita el acoplamiento fuerte con configurable dependency](https://franiglesias.github.io/configurable_dependency/) - Técnica para reducir el acoplamiento entre clases

## Referencias

- [Refactoring Guru - Inappropriate Intimacy](https://refactoring.guru/smells/inappropriate-intimacy)
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
- The Pragmatic Programmer - Law of Demeter
