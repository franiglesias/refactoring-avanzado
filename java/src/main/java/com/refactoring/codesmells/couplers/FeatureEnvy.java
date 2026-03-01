package com.refactoring.codesmells.couplers;

public class FeatureEnvy {

    public static class Customer {
        private final String name;
        private final String email;
        private final String address;
        private final String phone;

        public Customer(String name, String email, String address, String phone) {
            this.name = name;
            this.email = email;
            this.address = address;
            this.phone = phone;
        }

        public String getName() {
            return name;
        }

        public String getEmail() {
            return email;
        }

        public String getAddress() {
            return address;
        }

        public String getPhone() {
            return phone;
        }
    }

    public static class Invoice {
        private final String invoiceNumber;
        private final double amount;
        private final Customer customer;

        public Invoice(String invoiceNumber, double amount, Customer customer) {
            this.invoiceNumber = invoiceNumber;
            this.amount = amount;
            this.customer = customer;
        }

        public String getInvoiceNumber() {
            return invoiceNumber;
        }

        public double getAmount() {
            return amount;
        }

        public Customer getCustomer() {
            return customer;
        }
    }

    public static class InvoiceService {
        // This method is "envious" of Customer's data
        public void sendInvoice(Invoice invoice) {
            // This method knows too much about Customer's internal structure
            // and uses Customer's data more than its own

            // Formatting customer info
            String customerInfo = String.format(
                    "Customer: %s%nEmail: %s%nAddress: %s%nPhone: %s",
                    invoice.getCustomer().getName(),
                    invoice.getCustomer().getEmail(),
                    invoice.getCustomer().getAddress(),
                    invoice.getCustomer().getPhone()
            );

            // Validating customer data
            if (invoice.getCustomer().getEmail() == null || invoice.getCustomer().getEmail().isEmpty()) {
                throw new IllegalArgumentException("customer email is required");
            }
            if (invoice.getCustomer().getAddress() == null || invoice.getCustomer().getAddress().isEmpty()) {
                throw new IllegalArgumentException("customer address is required");
            }

            System.out.printf("Sending invoice %s for %.2f to:%n%s%n",
                    invoice.getInvoiceNumber(), invoice.getAmount(), customerInfo);
        }
    }
}
