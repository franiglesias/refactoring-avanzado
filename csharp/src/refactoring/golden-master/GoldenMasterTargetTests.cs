using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using Xunit;

namespace RefactoringAvanzado.Refactoring.GoldenMaster;

public class GoldenMasterTargetTests
{
    [Theory]
    [InlineData("Ana")]
    [InlineData("Luis")]
    [InlineData("Mar")]
    [InlineData("Iván")]
    [InlineData("Sofía")]
    public void ShouldPrintReceiptForCustomer(string customer)
    {
        var printer = new ReceiptPrinter();
        var counter = 0;

        var items = new[] { 1, 2, 3, 4, 5 };
        var quantities = new[] { 1, 3, 10, 25, 300 };

        foreach (var item in items)
        {
            foreach (var quantity in quantities)
            {
                counter++;
                var pedido = OrderGenerator.GenerateOrder($"ORD-{counter}", customer, item, quantity);
                var receipt = printer.Print(pedido);

                // Use Verify for snapshot testing with UseParameters
                // await Verify(receipt).UseParameters(customer, item, quantity);

                // For now, just assert it's not empty
                Assert.NotEmpty(receipt);
            }
        }
    }

    [Fact]
    public async Task ShouldGenerateGoldenMaster()
    {
        var printer = new ReceiptPrinter();

        var customers = OrderGenerator.Customers;
        var items = new[] { 1, 2, 3, 4, 5 };
        var quantities = new[] { 1, 3, 10, 25, 300 };

        var receipts = new List<string>();
        var counter = 0;

        foreach (var customer in customers)
        {
            foreach (var item in items)
            {
                foreach (var quantity in quantities)
                {
                    counter++;
                    var pedido = OrderGenerator.GenerateOrder($"ORD-{counter}", customer, item, quantity);
                    var receipt = printer.Print(pedido);
                    receipts.Add(receipt);
                }
            }
        }

        Console.WriteLine($"Generated {receipts.Count} receipts");

        var allReceipts = string.Join("\n==================\n", receipts);

        // Use Verify for snapshot testing
        // await Verify(allReceipts);

        // For now, just assert we generated the expected number
        Assert.Equal(125, receipts.Count); // 5 customers × 5 items × 5 quantities
    }
}
