package com.refactoring.refactoring;

import java.util.Objects;

/**
 * Represents an item in an order.
 */
public class OrderItem {
    private final String sku;
    private final String description;
    private final double unitPrice;
    private final int quantity;
    private final String category; // "general", "food", "books", or null

    public OrderItem(String sku, String description, double unitPrice, int quantity, String category) {
        this.sku = sku;
        this.description = description;
        this.unitPrice = unitPrice;
        this.quantity = quantity;
        this.category = category;
    }

    public String getSku() {
        return sku;
    }

    public String getDescription() {
        return description;
    }

    public double getUnitPrice() {
        return unitPrice;
    }

    public int getQuantity() {
        return quantity;
    }

    public String getCategory() {
        return category;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        OrderItem orderItem = (OrderItem) o;
        return Double.compare(orderItem.unitPrice, unitPrice) == 0 &&
                quantity == orderItem.quantity &&
                Objects.equals(sku, orderItem.sku) &&
                Objects.equals(description, orderItem.description) &&
                Objects.equals(category, orderItem.category);
    }

    @Override
    public int hashCode() {
        return Objects.hash(sku, description, unitPrice, quantity, category);
    }
}
