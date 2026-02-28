package com.refactoring.codesmells.bloaters;

/**
 * Example of Long Parameter List code smell.
 */
public class LongParameterList {

    public static void createUser(
            String username,
            String email,
            String firstName,
            String lastName,
            int age,
            String country,
            String city,
            String address,
            String postalCode,
            String phoneNumber,
            boolean isActive,
            String role
    ) {
        // Validation
        if (username == null || username.isEmpty() || email == null || email.isEmpty()) {
            throw new IllegalArgumentException("username and email are required");
        }

        // Business logic
        System.out.printf("Creating user: %s (%s)%n", username, email);
        System.out.printf("Name: %s %s%n", firstName, lastName);
        System.out.printf("Location: %s, %s, %s, %s%n", address, city, postalCode, country);
        System.out.printf("Contact: %s%n", phoneNumber);
        System.out.printf("Age: %d, Role: %s, Active: %b%n", age, role, isActive);
    }

    /*
     * Exercise: Refactor this code to reduce the parameter list
     * 1. Group related parameters into classes (e.g., Address, PersonalInfo)
     * 2. Consider using the Builder pattern
     * 3. Think about which parameters are truly required vs optional
     */
}
