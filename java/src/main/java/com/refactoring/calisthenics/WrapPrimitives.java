package com.refactoring.calisthenics;

/**
 * Example demonstrating using primitives directly.
 * Rule: Wrap all primitives and strings in their own types.
 */
public class WrapPrimitives {

    public static class Order {
        private final String id;
        private final String customerName;
        private final String email;
        private final double amount;
        private final String currency;

        public Order(String id, String customerName, String email, double amount, String currency) {
            this.id = id;
            this.customerName = customerName;
            this.email = email;
            this.amount = amount;
            this.currency = currency;
        }

        public String getId() {
            return id;
        }

        public String getCustomerName() {
            return customerName;
        }

        public String getEmail() {
            return email;
        }

        public double getAmount() {
            return amount;
        }

        public String getCurrency() {
            return currency;
        }
    }

    public static void validateAndProcess(Order order) {
        // Email validation
        if (order.getEmail() == null || order.getEmail().isEmpty() || !order.getEmail().contains("@")) {
            throw new IllegalArgumentException("invalid email");
        }

        // Amount validation
        if (order.getAmount() <= 0) {
            throw new IllegalArgumentException("amount must be positive");
        }

        System.out.printf("Processing order %s for %s (%.2f %s)%n",
                order.getId(), order.getCustomerName(), order.getAmount(), order.getCurrency());
    }

    /*
     * Exercise: Refactor this code to wrap primitives in domain types
     * Create types like:
     * - Email (with validation)
     * - Money (with amount and currency)
     * - OrderId
     * - CustomerName
     * This makes the domain model explicit and validation automatic
     */
}
