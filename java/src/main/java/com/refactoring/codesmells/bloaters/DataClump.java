package com.refactoring.codesmells.bloaters;

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
}
