# Comments

Comentarios.

## Definición

Los comentarios excesivos suelen ser una señal de que el código no es lo suficientemente claro por sí mismo. A menudo se utilizan para explicar código obvio o para compensar una mala elección de nombres, añadiendo ruido y riesgo de desactualización.

## Ejemplo

En este ejemplo, los comentarios explican paso a paso operaciones matemáticas básicas y declaraciones de variables, algo que el código ya expresa claramente.

```typescript
// Esta función suma dos números y devuelve el resultado.
// Toma el parámetro a que es un número y el parámetro b que también es un número.
// Luego usa el operador más para calcular la suma de a y b.
// Finalmente, devuelve esa suma al invocador de esta función.
export function add(a: number, b: number): number {
  // Declara una variable llamada result que contendrá la suma de a y b
  const result = a + b // calcula la suma agregando a y b
  // Devuelve el resultado a quien haya llamado a esta función
  return result // fin de la función
}

// Ejemplo de uso de este código con mal olor: llamar a una función trivial que no debería necesitar comentarios
export function demoCommentsSmell(): number {
  return add(2, 3)
}
```

## Ejercicio

Actualiza la función `add` para registrar (log) cuando la suma sea negativa.

## Problemas que encontrarás

Observa cómo los comentarios de alrededor se vuelven obsoletos o engañosos rápidamente al realizar cambios, obligándote a editar muchas líneas de comentario por un cambio diminuto en el código, aumentando el riesgo de desalineación entre lo que el código hace y lo que el comentario dice.
