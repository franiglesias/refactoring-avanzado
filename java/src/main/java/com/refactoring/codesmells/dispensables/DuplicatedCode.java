package com.refactoring.codesmells.dispensables;

/**
 * Example of Duplicated Code smell.
 */
public class DuplicatedCode {

    public static class EmailService {

        public void sendWelcomeEmail(String email, String name) {
            // Duplicated validation
            if (email == null || email.isEmpty()) {
                throw new IllegalArgumentException("email is required");
            }
            if (!email.contains("@")) {
                throw new IllegalArgumentException("invalid email format");
            }

            // Duplicated formatting
            String subject = "Welcome!";
            String body = String.format("Hello %s, welcome to our service!", name);

            // Duplicated sending logic
            System.out.printf("Sending email to: %s%n", email);
            System.out.printf("Subject: %s%n", subject);
            System.out.printf("Body: %s%n", body);
            System.out.println("Email sent successfully");
        }

        public void sendPasswordResetEmail(String email, String resetToken) {
            // Duplicated validation
            if (email == null || email.isEmpty()) {
                throw new IllegalArgumentException("email is required");
            }
            if (!email.contains("@")) {
                throw new IllegalArgumentException("invalid email format");
            }

            // Duplicated formatting
            String subject = "Password Reset";
            String body = String.format("Use this token to reset your password: %s", resetToken);

            // Duplicated sending logic
            System.out.printf("Sending email to: %s%n", email);
            System.out.printf("Subject: %s%n", subject);
            System.out.printf("Body: %s%n", body);
            System.out.println("Email sent successfully");
        }

        public void sendOrderConfirmationEmail(String email, String orderId, double amount) {
            // Duplicated validation
            if (email == null || email.isEmpty()) {
                throw new IllegalArgumentException("email is required");
            }
            if (!email.contains("@")) {
                throw new IllegalArgumentException("invalid email format");
            }

            // Duplicated formatting
            String subject = "Order Confirmation";
            String body = String.format("Your order %s for $%.2f has been confirmed", orderId, amount);

            // Duplicated sending logic
            System.out.printf("Sending email to: %s%n", email);
            System.out.printf("Subject: %s%n", subject);
            System.out.printf("Body: %s%n", body);
            System.out.println("Email sent successfully");
        }
    }

    /*
     * Exercise: Refactor this code to eliminate duplication
     * 1. Extract the common validation logic
     * 2. Extract the common email sending logic
     * 3. Create a generic sendEmail method that the others can use
     */
}
