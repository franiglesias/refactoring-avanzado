# Calistenia de Objetos (Object Calisthenics)

Un conjunto de nueve reglas propuestas por Jeff Bay que, aplicadas a nuestro código, pueden mejorar su calidad y aproximarlo a un mejor diseño.

Se basan en fijarnos en ciertas características aplicables a cualquier código.

Si observamos que el código incumple alguna de esas reglas, intentamos arreglarlo para que se ajuste a ella.

## Las 9 reglas

1. **No usar abreviaturas** - [NoAbbreviations.java](NoAbbreviations.java)
2. **No usar ELSE** - [NoElse.java](NoElse.java)
3. **Un solo nivel de indentación** - [OneLevelIndentation.java](OneLevelIndentation.java)
4. **Encapsular primitivos** - [WrapPrimitives.java](WrapPrimitives.java)
5. **Colecciones de primera clase** - [FirstClassCollections.java](FirstClassCollections.java)
6. **No usar getters y setters** - [NoGettersSetters.java](NoGettersSetters.java)
7. **Mantener las unidades de código pequeñas** - [SmallEntities.java](SmallEntities.java)
8. **Máximo de dos variables de instancia por clase** - [MaxTwoInstanceVariables.java](MaxTwoInstanceVariables.java)
9. **No más de un punto por línea** - [OneDotPerLine.java](OneDotPerLine.java)

## Descripción de cada regla

### 1. No usar abreviaturas
Los identificadores abreviados oscurecen la intención y hacen que sea más costoso comprender los conceptos que maneja el código.

**Refactoring:** Renombrar identificadores.

### 2. No usar ELSE
Las cláusulas ELSE ocultan reglas importantes y dificultan la comprensión del flujo del código.

**Refactoring:** Aplicar retorno temprano, cláusulas de guarda.

### 3. Un solo nivel de indentación
Múltiples niveles de indentación revelan mezcla de niveles de abstracción.

**Refactoring:** Extraer métodos, descomponer condicionales.

### 4. Encapsular primitivos
Los tipos primitivos no protegen las invariantes del dominio.

**Refactoring:** Introducir value objects.

### 5. Colecciones de primera clase
Encapsular colecciones en objetos que representen conceptos del dominio.

**Refactoring:** Envolver colecciones en clases.

### 6. No usar getters y setters
Exponer la estructura interna genera acoplamiento.

**Refactoring:** Encapsular campos, Tell Don't Ask.

### 7. Mantener las unidades de código pequeñas
Clases y métodos grandes tienen múltiples responsabilidades.

**Refactoring:** Extraer clases, extraer métodos.

### 8. Máximo de dos variables de instancia por clase
Muchas variables de instancia indican responsabilidades múltiples.

**Refactoring:** Introducir value objects, extraer clases.

### 9. No más de un punto por línea
Evitar exponer la estructura interna (Ley de Demeter).

**Refactoring:** Esconder delegación, extraer métodos.

## Cómo usar estos ejercicios

Cada archivo Java contiene un ejemplo de violación de una regla de calistenia. Tu tarea es:

1. Identificar la violación
2. Aplicar el refactoring correspondiente
3. Verificar que el comportamiento se mantiene
4. Observar cómo mejora el diseño

## Ejecutar los ejemplos

Cada archivo tiene un método `main()` que puedes ejecutar directamente:

```bash
javac com/refactoring/calisthenics/NoAbbreviations.java
java com.refactoring.calisthenics.NoAbbreviations
```

O desde tu IDE favorito.
