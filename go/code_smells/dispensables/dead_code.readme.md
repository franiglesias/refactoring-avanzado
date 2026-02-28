# Dead Code

Código muerto.

## Definición

El código muerto incluye declaraciones de variables, funciones o bloques de código que nunca se utilizan o son inalcanzables. Estos fragmentos añaden ruido, aumentan el coste de mantenimiento y pueden confundir a los desarrolladores sobre el flujo real de la aplicación.

## Ejemplo

Variables constantes no utilizadas y código después de un `return` son ejemplos claros de código que no aporta nada a la ejecución.

```go
const THE_ANSWER_TO_EVERYTHING = 42

func formatCurrency(amount float64) string {
	return fmt.Sprintf("$%.2f", amount)
}

func ActiveFunction(value int) int {
	if value < 0 {
		return 0
		neverRuns := value * -1
		fmt.Println("This will never be printed", neverRuns)
	}

	temp := value * 2

	return value + 1
}

func DemoDeadCode() string {
	result := ActiveFunction(5)
	return formatCurrency(float64(result))
}
```

## Ejercicio

Arregla un bug en `ActiveFunction` (por ejemplo, cambia el manejo de valores negativos).

## Problemas que encontrarás

Observa cómo el código muerto cercano dificulta razonar sobre lo que realmente se ejecuta, lo que puede invitar a errores o a olvidar la limpieza necesaria durante el proceso de refactorización.
