# Dead Code

Código muerto.

## Definición

El código muerto incluye declaraciones de variables, funciones o bloques de código que nunca se utilizan o son inalcanzables. Estos fragmentos añaden ruido, aumentan el coste de mantenimiento y pueden confundir a los desarrolladores sobre el flujo real de la aplicación.

## Ejemplo

Variables constantes no utilizadas y código después de un `return` son ejemplos claros de código que no aporta nada a la ejecución.

```typescript
const THE_ANSWER_TO_EVERYTHING = 42

function formatCurrency(amount: number): string {
  return `$${amount.toFixed(2)}`
}

export function activeFunction(value: number): number {
  if (value < 0) {
    return 0
    const neverRuns = value * -1
    console.log('This will never be printed', neverRuns)
  }

  const temp = value * 2

  return value + 1
}

export function demoDeadCode(): string {
  const result = activeFunction(5)
  return formatCurrency(result)
}
```

## Ejercicio

Arregla un bug en `activeFunction` (por ejemplo, cambia el manejo de valores negativos).

## Problemas que encontrarás

Observa cómo el código muerto cercano dificulta razonar sobre lo que realmente se ejecuta, lo que puede invitar a errores o a olvidar la limpieza necesaria durante el proceso de refactorización.
