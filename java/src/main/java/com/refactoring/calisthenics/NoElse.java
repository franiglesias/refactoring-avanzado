package com.refactoring.calisthenics;

public class NoElse {

    public static double calculateDiscount(String customerType, double amount) {
        if ("premium".equals(customerType)) {
            return amount * 0.20;
        } else if ("gold".equals(customerType)) {
            return amount * 0.15;
        } else if ("silver".equals(customerType)) {
            return amount * 0.10;
        } else {
            return amount * 0.05;
        }
    }
}
