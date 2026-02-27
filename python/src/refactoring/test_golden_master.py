import pytest
from datetime import datetime
from .golden_master import ReceiptPrinter, generate_order


class ReceiptPrinterWithoutDiscountForTest(ReceiptPrinter):
    def get_current_date(self) -> datetime:
        return datetime(2022, 2, 1)

    def discount(self) -> float:
        return 0.0


class ReceiptPrinterWithDiscountForTest(ReceiptPrinter):
    def get_current_date(self) -> datetime:
        return datetime(2022, 2, 1)

    def discount(self) -> float:
        return 0.05


class TestReceiptPrinter:
    counter = 0

    def test_should_print_a_receipt(self, snapshot):
        """Given a customer, number of items and quantity, should print a receipt"""
        TestReceiptPrinter.counter += 1
        customer = 'Ana'
        item = 1
        quantity = 1

        order = generate_order(f'ORD-{TestReceiptPrinter.counter}', customer, item, quantity)
        receipt = ReceiptPrinterWithoutDiscountForTest().print(order)

        assert receipt == snapshot
