# Dead Code

Código muerto.

## Definición

Código que nunca se ejecuta o nunca se usa. Variables, constantes, métodos o clases no utilizados que deben eliminarse. Aumentan la deuda técnica y dificultan la comprensión del código.

## Ejemplo

```java
public class DeadCode {

    // Constante no utilizada
    private static final int THE_ANSWER_TO_EVERYTHING = 42;

    public static void main(String[] args) {
        String result = demoDeadCode();
        System.out.println(result);
    }

    // Método no utilizado
    private static String formatCurrency(double amount) {
        return String.format("$%.2f", amount);
    }

    public static int activeFunction(int value) {
        if (value < 0) {
            return 0;
            // Código inalcanzable - nunca se ejecuta
            // int neverRuns = value * -1;
            // System.out.println("This will never be printed: " + neverRuns);
        }

        // Variable no utilizada
        int temp = value * 2;

        return value + 1;
    }

    public static String demoDeadCode() {
        int result = activeFunction(5);
        return formatCurrency(result);
    }
}
```

## Ejercicio

Identifica y elimina todo el código muerto. Añade un nuevo método que use la constante previamente no utilizada.

## Problemas que encontrarás

Es fácil que el código muerto se acumule con el tiempo, especialmente durante refactorizaciones. El código no utilizado confunde a los desarrolladores y hace más difícil encontrar el código realmente relevante.
