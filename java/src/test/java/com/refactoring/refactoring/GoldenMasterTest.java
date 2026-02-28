package com.refactoring.refactoring;

import au.com.origin.snapshots.Expect;
import au.com.origin.snapshots.junit5.SnapshotExtension;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;

import java.util.Date;

@ExtendWith(SnapshotExtension.class)
class GoldenMasterTest {

    private static int counter = 0;
    private Expect expect;

    @Test
    void shouldPrintAReceipt() {
        // Given a customer
        String customer = "Ana";
        // Given a number of items
        int item = 1;
        // Given quantity
        int quantity = 1;

        counter++;
        Order order = OrderGenerator.generateOrder("ORD-" + counter, customer, item, quantity);
        String receipt = new ReceiptPrinterWithoutDiscountForTest().print(order);

        expect.toMatchSnapshot(receipt);
    }

    static class ReceiptPrinterWithoutDiscountForTest extends ReceiptPrinter {
        @Override
        protected Date getCurrentDate() {
            return new Date(122, 1, 1); // Year 2022, Month February (0-indexed), Day 1
        }

        @Override
        protected double discount() {
            return 0.0;
        }
    }

    static class ReceiptPrinterWithDiscountForTest extends ReceiptPrinter {
        @Override
        protected Date getCurrentDate() {
            return new Date(122, 1, 1); // Year 2022, Month February (0-indexed), Day 1
        }

        @Override
        protected double discount() {
            return 0.05;
        }
    }
}
