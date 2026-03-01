# Refused Bequest

Herencia rechazada.

## Definición

Una subclase hereda de una clase base pero no usa (o anula con implementaciones vacías) muchos de los métodos heredados. Esto indica que la herencia no es apropiada - la subclase no cumple con la relación "es-un".

## Ejemplo

```java
public static class BaseController {
    public void start() {
        System.out.println("starting");
    }

    public void stop() {
        System.out.println("stopping");
    }

    public void reset() {
        System.out.println("resetting");
    }
}

public static class ReadOnlyController extends BaseController {
    @Override
    public void start() {
        // No hace nada - rechaza la herencia
    }

    @Override
    public void stop() {
        // No hace nada - rechaza la herencia
    }
}

public static void demoRefusedBequest(boolean readonly) {
    BaseController controller = readonly ? new ReadOnlyController() : new BaseController();
    controller.start();
    controller.stop();
}
```

## Ejercicio

Añade métodos de configuración (configure, validate) en BaseController que también sean rechazados por ReadOnlyController.

## Problemas que encontrarás

La herencia se vuelve cada vez más inapropiada, con más métodos vacíos o que lanzan excepciones, violando el Principio de Sustitución de Liskov.
