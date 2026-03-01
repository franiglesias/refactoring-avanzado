package com.refactoring.codesmells.bloaters;

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
}
