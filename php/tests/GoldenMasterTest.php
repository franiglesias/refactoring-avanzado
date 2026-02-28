<?php

namespace Tests;

use PHPUnit\Framework\TestCase;
use Refactoring\GoldenMaster\ReceiptPrinter;
use function Refactoring\GoldenMaster\generateOrder;

class ReceiptPrinterWithoutDiscountForTest extends ReceiptPrinter
{
    protected function getCurrentDate(): \DateTime
    {
        return new \DateTime('2022-02-01');
    }

    protected function discount(): float
    {
        return 0.0;
    }
}

class ReceiptPrinterWithDiscountForTest extends ReceiptPrinter
{
    protected function getCurrentDate(): \DateTime
    {
        return new \DateTime('2022-02-01');
    }

    protected function discount(): float
    {
        return 0.05;
    }
}

class GoldenMasterTest extends TestCase
{
    private static int $counter = 0;

    public function testShouldPrintAReceipt(): void
    {
        // Given a customer
        $customer = 'Ana';
        // Given a number of items
        $item = 1;
        // Given quantity
        $quantity = 1;

        self::$counter++;
        $order = generateOrder('ORD-' . self::$counter, $customer, $item, $quantity);
        $receipt = (new ReceiptPrinterWithoutDiscountForTest())->print($order);

        $this->assertMatchesSnapshot($receipt);
    }
}
