# Change Preventers - Code Smells

Los **Change Preventers** son code smells que hacen que el código sea difícil de cambiar. Cuando necesitas hacer una modificación, descubres que debes tocar múltiples lugares o que una clase cambia por muchas razones diferentes.

## Ejercicios incluidos

### 1. Divergent Change
[DivergentChange.java](DivergentChange.java)

Una clase cambia frecuentemente por diferentes razones. Indica que la clase tiene múltiples responsabilidades.

### 2. Shotgun Surgery
[ShotgunSurgery.java](ShotgunSurgery.java)

Un cambio requiere modificar múltiples clases. Lo opuesto a Divergent Change - la misma responsabilidad está dispersa en muchas clases.

### 3. Parallel Inheritance Hierarchy
[ParallelInheritanceHierarchy.java](ParallelInheritanceHierarchy.java)

Cada vez que añades una subclase a una jerarquía, debes añadir una subclase correspondiente a otra jerarquía.
