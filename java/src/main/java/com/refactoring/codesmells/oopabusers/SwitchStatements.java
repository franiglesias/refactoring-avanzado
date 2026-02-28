package com.refactoring.codesmells.oopabusers;

/**
 * Example of Switch Statements code smell.
 */
public class SwitchStatements {

    public enum PaymentMethod {
        CREDIT_CARD,
        PAYPAL,
        BANK_TRANSFER,
        CASH
    }

    public static class PaymentProcessor {

        public void processPayment(PaymentMethod method, double amount) {
            switch (method) {
                case CREDIT_CARD:
                    System.out.printf("Processing credit card payment for $%.2f%n", amount);
                    System.out.println("Validating card...");
                    System.out.println("Contacting payment gateway...");
                    System.out.println("Payment successful");
                    break;

                case PAYPAL:
                    System.out.printf("Processing PayPal payment for $%.2f%n", amount);
                    System.out.println("Redirecting to PayPal...");
                    System.out.println("Payment successful");
                    break;

                case BANK_TRANSFER:
                    System.out.printf("Processing bank transfer for $%.2f%n", amount);
                    System.out.println("Generating transfer reference...");
                    System.out.println("Payment pending confirmation");
                    break;

                case CASH:
                    System.out.printf("Processing cash payment for $%.2f%n", amount);
                    System.out.println("Payment received");
                    break;

                default:
                    throw new IllegalArgumentException("Unknown payment method: " + method);
            }
        }

        // Another switch on the same type
        public double calculateFee(PaymentMethod method, double amount) {
            switch (method) {
                case CREDIT_CARD:
                    return amount * 0.029; // 2.9% for credit cards
                case PAYPAL:
                    return amount * 0.034; // 3.4% for PayPal
                case BANK_TRANSFER:
                    return 2.5; // flat fee for bank transfers
                case CASH:
                    return 0; // no fee for cash
                default:
                    return 0;
            }
        }

        // Yet another switch
        public String getPaymentMethodName(PaymentMethod method) {
            switch (method) {
                case CREDIT_CARD:
                    return "Tarjeta de Crédito";
                case PAYPAL:
                    return "PayPal";
                case BANK_TRANSFER:
                    return "Transferencia Bancaria";
                case CASH:
                    return "Efectivo";
                default:
                    return "Desconocido";
            }
        }
    }

    /*
     * Exercise: Refactor this code to use polymorphism instead of switch statements
     * 1. Create a PaymentStrategy interface with methods: process, calculateFee, getName
     * 2. Implement the interface for each payment method
     * 3. Replace the switch statements with polymorphic calls
     * 4. Notice how adding a new payment method becomes easier
     */
}
