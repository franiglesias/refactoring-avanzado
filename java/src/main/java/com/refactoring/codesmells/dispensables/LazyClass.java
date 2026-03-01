package com.refactoring.codesmells.dispensables;

public class LazyClass {

    public static void main(String[] args) {
        printShippingLabel();
    }

    public static class Address {
        String name;
        String line1;
        String city;

        public Address(String name, String line1, String city) {
            this.name = name;
            this.line1 = line1;
            this.city = city;
        }
    }

    public static class ShippingLabelBuilder {
        public String build(Address a) {
            String cityPart = a.city != null ? ", " + a.city : "";
            return String.format("%s — %s%s", a.name, a.line1, cityPart);
        }
    }

    public static void printShippingLabel() {
        Address address = new Address("John Doe", "123 Main St", "New York");
        ShippingLabelBuilder labelBuilder = new ShippingLabelBuilder();
        String label = labelBuilder.build(address);
        System.out.println(label);
    }
}
