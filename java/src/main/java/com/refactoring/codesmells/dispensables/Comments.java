package com.refactoring.codesmells.dispensables;

public class Comments {

    public static void main(String[] args) {
        int result = demoCommentsSmell();
        System.out.println("Resultado: " + result);
    }

    /**
     * Esta función suma dos números y devuelve el resultado.
     * Toma el parámetro a que es un número y el parámetro b que también es un número.
     * Luego usa el operador más para calcular la suma de a y b.
     * Finalmente, devuelve esa suma al invocador de esta función.
     */
    public static int add(int a, int b) {
        // Declara una variable llamada result que contendrá la suma de a y b
        int result = a + b; // calcula la suma agregando a y b
        // Devuelve el resultado a quien haya llamado a esta función
        return result; // fin de la función
    }

    /**
     * Ejemplo de uso de este código con mal olor: llamar a una función trivial
     * que no debería necesitar comentarios
     */
    public static int demoCommentsSmell() {
        return add(2, 3);
    }
}
