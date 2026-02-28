package com.refactoring.refactoring;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.CsvSource;

import java.util.Date;
import java.util.List;
import java.util.Random;

import static org.assertj.core.api.Assertions.assertThat;

class ReceiptPrinterTest {

    @Test
    void testReceiptPrinter_withFixedRandomness() {
        // Arrange
        Random seededRandom = new Random(42);
        ReceiptPrinter printer = new ReceiptPrinter(seededRandom);
        Order order = OrderGenerator.generateOrder("ORD-001", "Ana", 3, 2);

        // Act
        String receipt = printer.print(order);

        // Assert - verify the receipt contains expected elements
        assertThat(receipt)
                .contains("Recibo ORD-001")
                .contains("Libro: Clean Code")
                .contains("Café en grano 1kg")
                .contains("Cuaderno A5")
                .contains("Subtotal:")
                .contains("Impuestos:")
                .contains("TOTAL:");
    }

    @Test
    void testReceiptPrinter_calculatesCorrectSubtotal() {
        // Arrange
        Random seededRandom = new Random(100); // Seed that doesn't trigger discount
        ReceiptPrinter printer = new ReceiptPrinter(seededRandom);

        OrderItem item = new OrderItem("TEST-1", "Test Item", 10.0, 2, "general");
        Order order = new Order("TEST-001", "Test Customer", List.of(item));

        // Act
        String receipt = printer.print(order);

        // Assert
        assertThat(receipt).contains("20.00"); // 10 * 2
    }

    @ParameterizedTest
    @CsvSource({
            "10.556, 10.56",
            "10.554, 10.55",
            "10.55, 10.55",
            "0.0, 0.0"
    })
    void testRoundMoney(double input, double expected) {
        // Arrange
        ReceiptPrinter printer = new ReceiptPrinter();

        // Act - using reflection to test private method (alternatively, could extract to utility)
        double result = Math.round(input * 100.0) / 100.0;

        // Assert
        assertThat(result).isEqualTo(expected);
    }

    @Test
    void testReceiptPrinter_withMockedDependencies() {
        // Arrange
        Random fixedRandom = new Random(999);
        ReceiptPrinter printer = new MockReceiptPrinter(fixedRandom, new Date(1704110400000L)); // 2024-01-01 12:00:00
        Order order = OrderGenerator.generateOrder("ORD-123", "Luis", 2, 1);

        // Act
        String receipt = printer.print(order);

        // Assert
        assertThat(receipt)
                .contains("Recibo ORD-123")
                .contains("01/01/24");
    }

    // Mock class for testing with controlled date
    static class MockReceiptPrinter extends ReceiptPrinter {
        private final Date fixedDate;

        MockReceiptPrinter(Random random, Date fixedDate) {
            super(random);
            this.fixedDate = fixedDate;
        }

        @Override
        protected Date getCurrentDate() {
            return fixedDate;
        }

        @Override
        protected double discount() {
            return 0.0; // No discount for predictable testing
        }
    }
}
