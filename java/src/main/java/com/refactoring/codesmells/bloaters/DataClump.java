package com.refactoring.codesmells.bloaters;

/**
 * Example of Data Clump code smell.
 * The same group of data (street, city, postalCode, country) appears together repeatedly.
 */
public class DataClump {

    public static class ProductService {

        public void shipProduct(
                String productId,
                String street,
                String city,
                String postalCode,
                String country
        ) {
            System.out.printf("Shipping product %s to:%n", productId);
            System.out.printf("%s%n%s, %s%n%s%n", street, city, postalCode, country);
        }

        public boolean validateDeliveryAddress(
                String street,
                String city,
                String postalCode,
                String country
        ) {
            return street != null && !street.isEmpty() &&
                    city != null && !city.isEmpty() &&
                    postalCode != null && !postalCode.isEmpty() &&
                    country != null && !country.isEmpty();
        }

        public double calculateShippingCost(
                String street,
                String city,
                String postalCode,
                String country,
                double weight
        ) {
            double baseCost = 10.0;
            // International shipping costs more
            if (!"Spain".equals(country)) {
                baseCost += 15.0;
            }
            // City surcharge
            if ("Madrid".equals(city) || "Barcelona".equals(city)) {
                baseCost += 2.0;
            }
            return baseCost + (weight * 0.5);
        }
    }

    /*
     * Exercise: Refactor this code to eliminate the data clump
     * 1. Create an Address class to group these related fields
     * 2. Update all methods to use the Address class
     * 3. Consider adding address-related behavior to the Address class
     */
}
