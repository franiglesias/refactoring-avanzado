package com.refactoring.codesmells.couplers;

import java.util.*;

/**
 * Code Smell: Middleman [Intermediario]
 * Shop actúa como un simple intermediario, delegando todas las llamadas a Catalog
 * sin agregar ningún valor.
 *
 * La clase Shop no tiene lógica propia, solo reenvía las llamadas.
 */
public class Middleman {

    public static void main(String[] args) {
        List<String> items = demoMiddleman();
        System.out.println("Items: " + items);
    }

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
}
