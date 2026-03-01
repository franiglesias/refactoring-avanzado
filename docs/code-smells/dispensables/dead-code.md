# Dead Code

Código muerto.

## Definición

El código muerto incluye declaraciones de variables, funciones o bloques de código que nunca se utilizan o son inalcanzables. Estos fragmentos añaden ruido, aumentan el coste de mantenimiento y pueden confundir a los desarrolladores sobre el flujo real de la aplicación.

## Descripción

**Dead Code** es código que existe en el sistema pero que nunca se ejecuta o nunca se usa. Puede tomar varias formas:
- Variables declaradas pero nunca usadas
- Funciones o métodos que nunca se llaman
- Parámetros que nunca se usan
- Código después de `return`, `break`, o `throw` (inalcanzable)
- Ramas de condicionales que nunca se ejecutan
- Clases o módulos enteros que no se importan

El código muerto aparece por:
- Refactorizaciones incompletas donde no se eliminó el código viejo
- Funcionalidades deprecadas que nunca se removieron
- Código experimental que se dejó comentado o sin usar
- Miedo a eliminar código "por si acaso"

El problema es que este código confunde, aumenta la superficie de mantenimiento, y puede llevar a modificar código que realmente no afecta nada.

## Síntomas

- Variables declaradas pero nunca referenciadas
- Funciones o métodos sin llamadas en todo el código
- Parámetros que nunca se usan dentro de una función
- Código después de sentencias de retorno o excepciones
- Imports de módulos que nunca se usan
- Constantes definidas pero nunca referenciadas
- Condicionales con ramas que nunca se pueden alcanzar
- Clases o módulos completos que no se usan

## Ejemplo

```pseudocode
// Constante nunca usada
THE_ANSWER_TO_EVERYTHING = 42

// Función nunca llamada
function formatCurrency(amount: number): string {
  return "$" + amount.toFixed(2)
}

function activeFunction(value: number): number {
  if (value < 0) {
    return 0
    // Código inalcanzable después del return
    neverRuns = value * -1
    print "This will never be printed: " + neverRuns
  }

  // Variable declarada pero nunca usada
  temp = value * 2

  return value + 1
}

function demoDeadCode(): string {
  result = activeFunction(5)
  return formatCurrency(result)
}
```

## Ejercicio

Arregla un bug en `activeFunction` (por ejemplo, cambia el manejo de valores negativos).

## Problemas que encontrarás

Observa cómo el código muerto cercano dificulta razonar sobre lo que realmente se ejecuta, lo que puede invitar a errores o a olvidar la limpieza necesaria durante el proceso de refactorización.

## Proceso de Refactoring

### 1. Identificar código muerto
- Usa herramientas de análisis estático del IDE
- Busca warnings sobre código no usado
- Analiza cobertura de tests para encontrar código no ejecutado
- Busca imports, variables y funciones sin referencias

### 2. Verificar que realmente está muerto
- Busca todas las referencias en el código
- Verifica que no se usa mediante reflexión o dinámicamente
- Comprueba si se usa en tests
- Confirma que no es parte de una API pública

### 3. Eliminar código inalcanzable
- Código después de `return`, `break`, `throw`
- Ramas de condicionales imposibles (`if (true)`, `if (false)`)
- Elimina inmediatamente, estos nunca se ejecutan

### 4. Remover variables y parámetros no usados
- Elimina variables declaradas pero no referenciadas
- Elimina parámetros que no se usan en el método
- Si un parámetro es parte de una interfaz, considera marcarlo explícitamente

### 5. Eliminar funciones y clases sin usar
- Elimina métodos que nunca se llaman
- Elimina clases que nunca se instancian
- Usa el control de versiones si necesitas recuperarlas

### 6. Limpiar código comentado
- El código comentado es código muerto
- El control de versiones preserva el historial
- Elimina bloques de código comentado sin miedo

## Técnicas de Refactoring Aplicables

- **Remove Dead Code**: Eliminar código que no se usa
- **Remove Unused Parameter**: Eliminar parámetros no usados
- **Inline Method/Variable**: Si solo se usa una vez, considera inline
- **Simplify Conditional**: Eliminar ramas imposibles

## Beneficios

- **Código más limpio**: Menos ruido y confusión
- **Mejor legibilidad**: Solo código que realmente importa
- **Mantenimiento reducido**: Menos código que mantener
- **Performance**: Menos código que cargar y compilar
- **Búsquedas más precisas**: No encontrarás código irrelevante
- **Refactoring más seguro**: No modificarás código que no se usa
- **Onboarding más fácil**: Nuevos desarrolladores no se confunden

## Versiones por Lenguaje

- [TypeScript](../../typescript/src/code-smells/dispensables/dead-code.ts) - [README](../../typescript/src/code-smells/dispensables/dead-code.readme.md)
- [Go](../../go/code_smells/dispensables/dead_code.go) - [README](../../go/code_smells/dispensables/dead_code.readme.md)
- [Java](../../java/src/main/java/com/refactoring/codesmells/dispensables/DeadCode.java) - [README](../../java/src/main/java/com/refactoring/codesmells/dispensables/DeadCode.readme.md)
- [PHP](../../php/src/code-smells/dispensables/DeadCode.php) - [README](../../php/src/code-smells/dispensables/DeadCode.readme.md)
- [Python](../../python/src/code_smells/dispensables/dead_code.py) - [README](../../python/src/code_smells/dispensables/dead_code_readme.md)
- [C#](../../csharp/src/code-smells/dispensables/DeadCode.cs) - [README](../../csharp/src/code-smells/dispensables/dead-code.readme.md)

## Referencias en Español

- [Refactor cotidiano (1). Fuera comentarios](https://franiglesias.github.io/everyday-refactor-1/) - Incluye eliminación de código muerto y código comentado

## Referencias

- [Refactoring Guru - Dead Code](https://refactoring.guru/smells/dead-code)
- Martin Fowler - "Refactoring: Improving the Design of Existing Code"
- Robert C. Martin - "Clean Code"
