package com.refactoring.calisthenics;

import java.util.ArrayList;
import java.util.List;

/**
 * Object Calisthenics: First Class Collections
 * Cualquier clase que contenga una colección no debe contener otras propiedades.
 *
 * Las funciones addProduct, totalPrice y removeProduct operan sobre una lista
 * sin encapsulación. Deberían estar en una clase que envuelva la colección.
 */
public class FirstClassCollections {

    public static void main(String[] args) {
        List<Product> products = new ArrayList<>();
        addProduct(products, new Product("1", 100.0));
        addProduct(products, new Product("2", 50.0));
        System.out.println("Total: " + totalPrice(products));

        products = removeProduct(products, "1");
        System.out.println("Total después de remover: " + totalPrice(products));
    }

    public static class Product {
        String id;
        double price;

        public Product(String id, double price) {
            this.id = id;
            this.price = price;
        }
    }

    public static void addProduct(List<Product> products, Product product) {
        boolean exists = products.stream()
            .anyMatch(p -> p.id.equals(product.id));
        if (!exists) {
            products.add(product);
        }
    }

    public static double totalPrice(List<Product> products) {
        return products.stream()
            .mapToDouble(p -> p.price)
            .sum();
    }

    public static List<Product> removeProduct(List<Product> products, String productId) {
        List<Product> result = new ArrayList<>();
        for (Product p : products) {
            if (!p.id.equals(productId)) {
                result.add(p);
            }
        }
        return result;
    }
}
