# Middleman

Intermediario.

## Definición

Una clase actúa como un simple intermediario, delegando todas las llamadas a otra clase sin agregar ningún valor. La clase intermediaria no tiene lógica propia, solo reenvía las llamadas.

## Ejemplo

```java
public static class Catalog {
    private Map<String, String> items = new HashMap<>();

    public void add(String id, String name) {
        items.put(id, name);
    }

    public Optional<String> find(String id) {
        return Optional.ofNullable(items.get(id));
    }

    public List<String> list() {
        return new ArrayList<>(items.values());
    }
}

public static class Shop {
    private Catalog catalog;

    public Shop(Catalog catalog) {
        this.catalog = catalog;
    }

    // Shop solo delega todas las llamadas a Catalog sin agregar valor
    public void add(String id, String name) {
        catalog.add(id, name);
    }

    public Optional<String> find(String id) {
        return catalog.find(id);
    }

    public List<String> list() {
        return catalog.list();
    }
}

public static List<String> demoMiddleman() {
    Catalog c = new Catalog();
    Shop s = new Shop(c);
    s.add("1", "Book");
    s.add("2", "Pen");
    return s.list();
}
```

## Ejercicio

Añade funcionalidad de búsqueda por nombre y filtrado por categoría.

## Problemas que encontrarás

Cada nueva funcionalidad en `Catalog` requerirá añadir un método de delegación correspondiente en `Shop`, sin aportar valor alguno.
