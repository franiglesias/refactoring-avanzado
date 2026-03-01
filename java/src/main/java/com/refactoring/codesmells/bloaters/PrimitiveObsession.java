package com.refactoring.codesmells.bloaters;

public class PrimitiveObsession {

    public static class Order {
        private final String customerName;
        private final String customerEmail;
        private final String address;
        private final double totalAmount;
        private final String currency;

        public Order(String customerName, String customerEmail, String address,
                     double totalAmount, String currency) {
            this.customerName = customerName;
            this.customerEmail = customerEmail;
            this.address = address;
            this.totalAmount = totalAmount;
            this.currency = currency;
        }

        public void sendInvoice() {
            // Validation logic scattered everywhere instead of being in domain objects
            if (!customerEmail.contains("@")) {
                throw new IllegalArgumentException("Email inválido");
            }
            if (address == null || address.isEmpty()) {
                throw new IllegalArgumentException("No se ha indicado dirección");
            }
            if (totalAmount <= 0) {
                throw new IllegalArgumentException("El monto debe ser mayor que cero");
            }
            System.out.printf("Factura enviada a %s in %s por %.2f %s%n",
                    customerName, address, totalAmount, currency);
        }
    }
}
