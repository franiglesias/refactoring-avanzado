# Couplers - Code Smells

Los **Couplers** son code smells que representan acoplamiento excesivo entre clases. Esto hace que los cambios en una clase requieran cambios en otras clases relacionadas.

## Ejercicios incluidos

### 1. Feature Envy
[FeatureEnvy.java](FeatureEnvy.java)

Un método accede más a los datos de otra clase que a los de su propia clase.

### 2. Inappropriate Intimacy
[InappropriateIntimacy.java](InappropriateIntimacy.java)

Dos clases están demasiado acopladas, conociendo demasiado sobre la estructura interna de la otra.

### 3. Message Chains
[MessageChains.java](MessageChains.java)

El código navega a través de múltiples objetos para obtener un valor, creando dependencia con la estructura interna.

### 4. Middleman
[Middleman.java](Middleman.java)

Una clase simplemente delega todas sus llamadas a otra clase sin agregar ningún valor.
