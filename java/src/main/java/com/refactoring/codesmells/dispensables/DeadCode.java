package com.refactoring.codesmells.dispensables;

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
