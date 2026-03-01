package com.refactoring.codesmells.changepreventers;

import java.util.List;

public class ShotgunSurgery {

    public static void main(String[] args) {
        List<LineItem> items = List.of(
            new LineItem("Producto A", 100.0, 2),
            new LineItem("Producto B", 50.0, 1)
        );

        double[] result = demoShotgun(items);
        System.out.println("Calc: " + result[0] + ", Invoice: " + result[1]);
    }

    public static class LineItem {
        String name;
        double price;
        int qty;

        public LineItem(String name, double price, int qty) {
            this.name = name;
            this.price = price;
            this.qty = qty;
        }
    }

    public static class PriceCalculator {
        public double totalWithTax(List<LineItem> items) {
            double subtotal = items.stream()
                .mapToDouble(i -> i.price * i.qty)
                .sum();
            double tax = subtotal * 0.21;
            return subtotal + tax;
        }
    }

    public static class InvoiceService {
        public double createTotal(List<LineItem> items) {
            double base = items.stream()
                .mapToDouble(i -> i.price * i.qty)
                .sum();
            double vat = base * 0.21;
            return base + vat;
        }
    }

    public static class SalesReport {
        public String summarize(List<LineItem> items) {
            double sum = items.stream()
                .mapToDouble(i -> i.price * i.qty)
                .sum();
            double tax = sum * 0.21;
            double total = sum + tax;
            return String.format("total=%.2f", total);
        }
    }

    public static class LoyaltyPoints {
        public int points(List<LineItem> items) {
            double base = items.stream()
                .mapToDouble(i -> i.price * i.qty)
                .sum();
            double withTax = base + base * 0.21;
            return (int) Math.floor(withTax / 10);
        }
    }

    public static double[] demoShotgun(List<LineItem> items) {
        double calc = new PriceCalculator().totalWithTax(items);
        double inv = new InvoiceService().createTotal(items);
        return new double[]{calc, inv};
    }
}
