package com.refactoring.calisthenics;

public class OneDotPerLine {

    public static void main(String[] args) {
        Order order = new Order(
            new Customer("John Doe",
                new Address("Elm Street", "Madrid")
            )
        );

        // Violación: múltiples puntos en una línea
        String destination = order.getCustomer().getAddress().getCity();
        System.out.println("Destination: " + destination);
    }

    public static class Address {
        private final String street;
        private final String city;

        public Address(String street, String city) {
            this.street = street;
            this.city = city;
        }

        public String getCity() {
            return city;
        }
    }

    public static class Customer {
        private final String name;
        private final Address address;

        public Customer(String name, Address address) {
            this.name = name;
            this.address = address;
        }

        public Address getAddress() {
            return address;
        }
    }

    public static class Order {
        private final Customer customer;

        public Order(Customer customer) {
            this.customer = customer;
        }

        public Customer getCustomer() {
            return customer;
        }
    }
}
