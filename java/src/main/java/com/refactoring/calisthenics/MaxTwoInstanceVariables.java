package com.refactoring.calisthenics;

import java.util.*;

public class MaxTwoInstanceVariables {

    public static void main(String[] args) {
        CheckoutSession session = new CheckoutSession();
        session.addItem("1", 100.0, 2);
        session.addItem("2", 50.0, 1);
        System.out.println("Total: " + session.total());
    }

    public static class CheckoutSession {
        private List<CartItem> cartItems = new ArrayList<>();
        private String customerId = null;
        private String shippingAddress = null;
        private String billingAddress = null;
        private String couponCode = null;
        private String paymentMethod = null;
        private String currency = "USD";
        private double taxRate = 0.21;

        public void addItem(String id, double price, int qty) {
            cartItems.add(new CartItem(id, price, qty));
        }

        public double total() {
            double subtotal = cartItems.stream()
                .mapToDouble(i -> i.price * i.qty)
                .sum();

            double discount = couponCode != null ? 10 : 0;
            double taxed = (subtotal - discount) * (1 + taxRate);
            return "USD".equals(currency) ? taxed : taxed * 0.9;
        }

        private static class CartItem {
            String id;
            double price;
            int qty;

            CartItem(String id, double price, int qty) {
                this.id = id;
                this.price = price;
                this.qty = qty;
            }
        }
    }
}
