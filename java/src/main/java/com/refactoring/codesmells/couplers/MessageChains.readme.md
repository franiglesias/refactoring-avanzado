# Message Chains

Cadenas de mensajes.

## Definición

El cliente debe conocer toda la cadena de navegación para obtener un valor simple. Esto crea acoplamiento a la estructura interna y cualquier cambio en la cadena rompe el código cliente.

## Ejemplo

```java
public static class Level2 {
    private int value;

    public Level2(int value) {
        this.value = value;
    }

    public int getValue() {
        return value;
    }
}

public static class Level1 {
    private Level2 next;

    public Level1(Level2 next) {
        this.next = next;
    }

    public Level2 getNext() {
        return next;
    }
}

public static class Root {
    private Level1 next;

    public Root(Level1 next) {
        this.next = next;
    }

    public Level1 getNext() {
        return next;
    }
}

public static int readDeep(Root root) {
    // Cliente debe conocer toda la cadena de navegación
    return root.getNext().getNext().getValue();
}
```

## Ejercicio

Añade un nivel adicional (Level3) y métodos para modificar valores en diferentes niveles.

## Problemas que encontrarás

Cualquier cambio en la estructura intermedia (añadir, remover o reordenar niveles) rompe todos los lugares que navegan por la cadena.
