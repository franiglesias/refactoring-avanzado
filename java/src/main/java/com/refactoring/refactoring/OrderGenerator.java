package com.refactoring.refactoring;

import java.util.ArrayList;
import java.util.List;

/**
 * Utility class to generate test orders.
 */
public class OrderGenerator {
    private static final List<Product> PRODUCTS = List.of(
            new Product("BK-001", "Libro: Clean Code", 30.0, "books"),
            new Product("FD-010", "Café en grano 1kg", 12.5, "food"),
            new Product("GN-777", "Cuaderno A5", 5.2, "general"),
            new Product("GN-123", "Bolígrafos (pack 10)", 3.9, "general"),
            new Product("FD-222", "Té verde 200g", 6.75, "food")
    );

    public static final List<String> CUSTOMERS = List.of("Ana", "Luis", "Mar", "Iván", "Sofía");

    public static Order generateOrder(String id, String customerName, int numItems, int quantity) {
        List<OrderItem> items = new ArrayList<>();
        for (int i = 0; i < numItems && i < PRODUCTS.size(); i++) {
            Product p = PRODUCTS.get(i);
            items.add(new OrderItem(p.sku, p.description, p.unitPrice, quantity, p.category));
        }
        return new Order(id, customerName, items);
    }

    private static class Product {
        final String sku;
        final String description;
        final double unitPrice;
        final String category;

        Product(String sku, String description, double unitPrice, String category) {
            this.sku = sku;
            this.description = description;
            this.unitPrice = unitPrice;
            this.category = category;
        }
    }
}
