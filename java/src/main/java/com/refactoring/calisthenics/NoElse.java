package com.refactoring.calisthenics;

/**
 * Example calculating discount using if-else.
 * Rule: Don't use the ELSE keyword.
 */
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

    /*
     * Exercise: Refactor this code to eliminate all else keywords
     * Techniques to consider:
     * 1. Early returns
     * 2. Guard clauses
     * 3. Polymorphism (strategy pattern)
     * 4. Look-up tables or maps
     */
}
