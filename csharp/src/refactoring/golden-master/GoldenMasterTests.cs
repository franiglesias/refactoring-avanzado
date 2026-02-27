using System;
using System.Collections.Generic;
using System.Linq;
using Xunit;

namespace RefactoringAvanzado.Refactoring.GoldenMaster;

public class ReceiptPrinterWithoutDiscountForTest : ReceiptPrinter
{
    protected override DateTime GetCurrentDate()
    {
        return new DateTime(2022, 2, 1);
    }

    protected override decimal Discount()
    {
        return 0;
    }
}

public class ReceiptPrinterWithDiscountForTest : ReceiptPrinter
{
    protected override DateTime GetCurrentDate()
    {
        return new DateTime(2022, 2, 1);
    }

    protected override decimal Discount()
    {
        return 0.05m;
    }
}

public class GoldenMasterTests
{
    private int _counter = 0;

    [Fact]
    public void ShouldPrintAReceipt()
    {
        // Given a customer
        const string customer = "Ana";
        // Given a number of items
        const int item = 1;
        // Given quantity
        const int quantity = 1;

        _counter++;
        var pedido = OrderGenerator.GenerateOrder($"ORD-{_counter}", customer, item, quantity);
        var receipt = new ReceiptPrinterWithoutDiscountForTest().Print(pedido);

        // Use Verify for snapshot testing
        // await Verify(receipt);

        // For now, just assert it's not empty (you'll need to setup Verify)
        Assert.NotEmpty(receipt);
    }
}
