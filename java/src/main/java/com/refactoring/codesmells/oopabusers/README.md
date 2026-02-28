# OOP Abusers - Code Smells

Los **OOP Abusers** son code smells que representan un mal uso de los principios de programación orientada a objetos, especialmente herencia y polimorfismo.

## Ejercicios incluidos

### 1. Alternative Classes with Different Interfaces
[AlternativeClassesDifferentInterfaces.java](AlternativeClassesDifferentInterfaces.java)

Dos clases hacen lo mismo pero tienen interfaces diferentes, cuando deberían compartir una interfaz común.

### 2. Refused Bequest
[RefusedBequest.java](RefusedBequest.java)

Una subclase usa solo algunos de los métodos y propiedades heredados de su padre, indicando que la jerarquía de herencia está mal diseñada.

### 3. Switch Statements
[SwitchStatements.java](SwitchStatements.java)

Uso excesivo de switch/case o if-else encadenados que deberían ser reemplazados por polimorfismo.

### 4. Temporal Instance Variables
[TemporalInstanceVariables.java](TemporalInstanceVariables.java)

Variables de instancia que solo son válidas durante ciertos períodos del ciclo de vida del objeto.
