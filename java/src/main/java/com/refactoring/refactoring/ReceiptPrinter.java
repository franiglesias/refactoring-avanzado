package com.refactoring.refactoring;

import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.Random;

/**
 * Prints receipts for orders.
 * Do not change this class at the beginning of the exercise; first create the Golden Master.
 */
public class ReceiptPrinter {
    private final Random random;

    public ReceiptPrinter() {
        this.random = new Random();
    }

    // Constructor for testing with seeded random
    public ReceiptPrinter(Random random) {
        this.random = random;
    }

    public String print(Order order) {
        Date now = getCurrentDate();
        SimpleDateFormat dateFormat = new SimpleDateFormat("MM/dd/yy HH:mm:ss");

        String header = String.format("Recibo %s - %s", order.getId(), dateFormat.format(now));

        double subtotal = 0.0;
        List<String> lines = new ArrayList<>();
        int index = 1;
        for (OrderItem item : order.getItems()) {
            double lineTotal = roundMoney(item.getUnitPrice() * item.getQuantity());
            subtotal = roundMoney(subtotal + lineTotal);
            lines.add(String.format("%d. %s (%s) x%d = $%.2f",
                    index++, item.getDescription(), item.getSku(), item.getQuantity(), lineTotal));
        }

        double luckyDiscountPct = discount();
        double luckyDiscount = roundMoney(subtotal * luckyDiscountPct);

        double taxableGeneral = 0.0;
        double foodTax = 0.0;
        for (OrderItem item : order.getItems()) {
            if (!"books".equals(item.getCategory()) && !"food".equals(item.getCategory())) {
                taxableGeneral += item.getUnitPrice() * item.getQuantity();
            }
            if ("food".equals(item.getCategory())) {
                foodTax += item.getUnitPrice() * item.getQuantity() * 0.03;
            }
        }
        double generalTax = taxableGeneral * 0.07;
        double taxes = roundMoney(generalTax + foodTax);

        double total = roundMoney(subtotal - luckyDiscount + taxes);

        String discountLine;
        if (luckyDiscount > 0) {
            discountLine = String.format("Descuento de la suerte: -$%.2f (%.2f%%)",
                    luckyDiscount, luckyDiscountPct * 100);
        } else {
            discountLine = "Descuento de la suerte: $0.00 (0.00%)";
        }

        List<String> summary = List.of(
                String.format("Subtotal: $%.2f", subtotal),
                discountLine,
                String.format("Impuestos: $%.2f", taxes),
                String.format("TOTAL: $%.2f", total)
        );

        List<String> parts = new ArrayList<>();
        parts.add(header);
        parts.addAll(lines);
        parts.add("---");
        parts.addAll(summary);

        return String.join("\n", parts);
    }

    protected double discount() {
        double luckyDiscountPct = 0.0;
        if (random.nextDouble() < 0.1) {
            luckyDiscountPct = random.nextDouble() * 0.05;
        }
        return luckyDiscountPct;
    }

    protected Date getCurrentDate() {
        return new Date();
    }

    private double roundMoney(double n) {
        return Math.round(n * 100.0) / 100.0;
    }
}
