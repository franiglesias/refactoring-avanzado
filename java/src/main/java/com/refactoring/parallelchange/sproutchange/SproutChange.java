package com.refactoring.parallelchange.sproutchange;

import java.util.List;

public class SproutChange {

    public static class CartItem {
        public final String id;
        public final double price;
        public final int qty;
        public final String category;

        public CartItem(String id, double price, int qty, String category) {
            this.id = id;
            this.price = price;
            this.qty = qty;
            this.category = category;
        }
    }

    public enum Region {
        US, EU
    }

    public interface TaxPolicy {

    }

    // Regla existente: un único impuesto plano por región; los libros y la comida están exentos en la UE
    public static double calculateTotal(List<CartItem> cart, Region region) {
        double subtotal = cart.stream()
                .mapToDouble(it -> it.price * it.qty)
                .sum();

        double tax = 0;
        if (region == Region.US) {
            tax = subtotal * 0.07; // 7% plano
        } else if (region == Region.EU) {
            // exenciones ingenuas en línea
            double taxable = cart.stream()
                    .filter(it -> !"books".equals(it.category) && !"food".equals(it.category))
                    .mapToDouble(it -> it.price * it.qty)
                    .sum();
            tax = taxable * 0.2; // 20% plano solo sobre los ítems gravables
        }

        return roundCurrency(subtotal + tax);
    }

    public static double roundCurrency(double amount) {
        return Math.round(amount * 100.0) / 100.0;
    }

    // Uso de ejemplo, mantenido simple para estudiantes
    public static double demoSprout() {
        List<CartItem> cart = List.of(
                new CartItem("p1", 10, 2, "general"),
                new CartItem("b1", 20, 1, "books")
        );
        return calculateTotal(cart, Region.EU);
    }
}
