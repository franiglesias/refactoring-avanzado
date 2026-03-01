# Alternative Classes with Different Interfaces

Clases alternativas con diferentes interfaces.

## Definición

Dos o más clases hacen lo mismo pero tienen interfaces diferentes (nombres de métodos distintos, parámetros diferentes). Esto dificulta el uso intercambiable y genera duplicación de lógica cliente.

## Ejemplo

```java
public static class TextLogger {
    public void log(String message) {
        System.out.println("[text] " + message);
    }
}

public static class MessageWriter {
    public void write(String entry) {
        System.out.println("[text] " + entry);
    }
}

public static void useAltClasses(String choice, String msg) {
    if ("logger".equals(choice)) {
        new TextLogger().log(msg);
    } else {
        new MessageWriter().write(msg);
    }
}
```

## Ejercicio

Añade soporte para niveles de log (INFO, WARNING, ERROR) en ambas clases.

## Problemas que encontrarás

Tendrás que duplicar la lógica de niveles en ambas clases y en todo el código que las use, en lugar de tener una interfaz común que permita el polimorfismo.
