# Comments

Comentarios.

## Definición

Los comentarios excesivos suelen ser una señal de que el código no es lo suficientemente claro por sí mismo. A menudo se utilizan para explicar código obvio o para compensar una mala elección de nombres, añadiendo ruido y riesgo de desactualización.

## Descripción

**Comments** como code smell no significa que todos los comentarios sean malos. Los comentarios útiles explican el "por qué" (decisiones de diseño, contexto del negocio, algoritmos complejos). El problema son los comentarios que explican el "qué" cuando el código debería ser autoexplicativo.

Comentarios problemáticos incluyen:
- Explicaciones de código obvio
- Comentarios que repiten lo que el código ya dice
- Comentarios que se vuelven obsoletos cuando el código cambia
- Comentarios que compensan nombres poco claros
- Bloques grandes de código comentado (dead code)

El código limpio debe leerse como prosa. Si necesitas comentarios extensivos para explicar qué hace el código, probablemente el código necesita refactorización, no comentarios.

## Síntomas

- Comentarios que explican cada línea de código
- Comentarios que repiten exactamente lo que el código hace
- Comentarios obsoletos que no coinciden con el código actual
- Variables o métodos con nombres crípticos explicados por comentarios
- Bloques grandes de código comentado sin razón clara
- Comentarios que explican condicionales complejos
- Comentarios que describen pasos de algoritmos simples

## Ejemplo

```pseudocode
// Esta función suma dos números y devuelve el resultado.
// Toma el parámetro a que es un número y el parámetro b que también es un número.
// Luego usa el operador más para calcular la suma de a y b.
// Finalmente, devuelve esa suma al invocador de esta función.
function add(a: number, b: number): number {
  // Declara una variable llamada result que contendrá la suma de a y b
  result = a + b  // calcula la suma agregando a y b
  // Devuelve el resultado a quien haya llamado a esta función
  return result  // fin de la función
}

// Ejemplo de uso de este código con mal olor
function demoCommentsSmell(): number {
  return add(2, 3)
}
```

## Ejercicio

Actualiza la función `add` para registrar (log) cuando la suma sea negativa.

## Problemas que encontrarás

Observa cómo los comentarios de alrededor se vuelven obsoletos o engañosos rápidamente al realizar cambios, obligándote a editar muchas líneas de comentario por un cambio diminuto en el código, aumentando el riesgo de desalineación entre lo que el código hace y lo que el comentario dice.

## Proceso de Refactoring

### 1. Identificar comentarios innecesarios
- Lee cada comentario y pregúntate si añade información valiosa
- Marca comentarios que solo repiten el código
- Identifica comentarios que explican código que debería ser obvio

### 2. Renombrar para claridad
- Si un comentario explica qué hace una variable o método
- Renombra el elemento para que sea autoexplicativo
- Ejemplo: `// días hasta caducidad` + `int x` → `daysUntilExpiration`

### 3. Extraer métodos para explicar lógica
- Si un comentario explica un bloque de código
- Extrae ese bloque a un método con nombre descriptivo
- Ejemplo: `// validar email` → `validateEmail()`

### 4. Simplificar condicionales complejos
- Si hay comentarios explicando condiciones
- Extrae la condición a una variable o método con nombre claro
- Ejemplo: `// verificar si el usuario es elegible` → `if (isUserEligible())`

### 5. Eliminar código comentado
- El código comentado es código muerto
- El control de versiones guarda el historial
- Elimina bloques de código comentado sin miedo

### 6. Preservar comentarios valiosos
- Mantén comentarios que explican "por qué", no "qué"
- Documenta decisiones de diseño no obvias
- Explica algoritmos complejos o fórmulas de negocio
- Advierte sobre efectos secundarios o comportamientos sorprendentes

## Técnicas de Refactoring Aplicables

- **Rename Variable/Method**: Hacer el código autoexplicativo
- **Extract Method**: Convertir bloques comentados en métodos con nombres claros
- **Introduce Explaining Variable**: Extraer expresiones complejas a variables nombradas
- **Replace Comment with Assertion**: Si el comentario describe una precondición
- **Remove Dead Code**: Eliminar código comentado

## Beneficios

- **Código autoexplicativo**: El código se lee como lenguaje natural
- **Sincronización garantizada**: El código y su explicación son lo mismo
- **Menos mantenimiento**: No hay que actualizar comentarios obsoletos
- **Mejor legibilidad**: Sin ruido de comentarios obvios
- **Refactoring más fácil**: Puedes cambiar código sin preocuparte por comentarios
- **Nombres mejores**: Fuerza a pensar en nombres descriptivos

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/dispensables/comments.ts) - [README](../../typescript/src/code-smells/dispensables/comments.readme.md)
- [Go](../../go/code_smells/dispensables/comments.go) - [README](../../go/code_smells/dispensables/comments.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/dispensables/Comments.java) - [README](../../java/src/main/java/com/refactoring/codesmells/dispensables/Comments.readme.md)
- [PHP](../../php/src/code-smells/dispensables/Comments.php) - [README](../../php/src/code-smells/dispensables/Comments.readme.md)
- [Python](../../python/src/code_smells/dispensables/comments.py) - [README](../../python/src/code_smells/dispensables/comments_readme.md)
- [C#](../../csharp/src/code-smells/dispensables/Comments.cs) - [README](../../csharp/src/code-smells/dispensables/comments.readme.md)

## Referencias en Español

- [Refactor cotidiano (1). Fuera comentarios](https://franiglesias.github.io/everyday-refactor-1/) - Guía práctica para eliminar comentarios innecesarios mediante refactorización
- [Sobre la expresividad del código](https://franiglesias.github.io/codigo-expresivo/) - Cómo hacer que el código sea autoexplicativo sin comentarios
- [Cómo poner nombres](https://franiglesias.github.io/naming-things/) - Técnicas para crear nombres claros que eliminen la necesidad de comentarios

## Referencias

- [Refactoring Guru - Comments](https://refactoring.guru/smells/comments)
- Robert C. Martin - "Clean Code" - Comments chapter
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
