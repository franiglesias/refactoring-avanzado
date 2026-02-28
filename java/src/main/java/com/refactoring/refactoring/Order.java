package com.refactoring.refactoring;

import java.util.List;
import java.util.Objects;

/**
 * Represents a customer order.
 */
public class Order {
    private final String id;
    private final String customerName;
    private final List<OrderItem> items;

    public Order(String id, String customerName, List<OrderItem> items) {
        this.id = id;
        this.customerName = customerName;
        this.items = items;
    }

    public String getId() {
        return id;
    }

    public String getCustomerName() {
        return customerName;
    }

    public List<OrderItem> getItems() {
        return items;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Order order = (Order) o;
        return Objects.equals(id, order.id) &&
                Objects.equals(customerName, order.customerName) &&
                Objects.equals(items, order.items);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id, customerName, items);
    }
}
