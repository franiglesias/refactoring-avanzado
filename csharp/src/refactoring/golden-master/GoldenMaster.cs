using System;
using System.Collections.Generic;
using System.Linq;

namespace RefactoringAvanzado.Refactoring.GoldenMaster;

public class OrderItem
{
    public string Sku { get; init; } = string.Empty;
    public string Description { get; init; } = string.Empty;
    public decimal UnitPrice { get; init; }
    public int Quantity { get; init; }
    public string? Category { get; init; }
}

public class Order
{
    public string Id { get; init; } = string.Empty;
    public string CustomerName { get; init; } = string.Empty;
    public List<OrderItem> Items { get; init; } = new();
}

public class ReceiptPrinter
{
    // Do not change this function at the beginning of the exercise; first create the Golden Master.
    public string Print(Order order)
    {
        var now = GetCurrentDate();

        var header = $"Recibo {order.Id} - {now:dd/MM/yyyy} {now:HH:mm:ss}";

        decimal subtotal = 0;
        var lines = order.Items.Select((it, idx) =>
        {
            var lineTotal = Round(it.UnitPrice * it.Quantity);
            subtotal = Round(subtotal + lineTotal);
            return $"{idx + 1}. {it.Description} ({it.Sku}) x{it.Quantity} = ${lineTotal:F2}";
        }).ToList();

        var luckyDiscountPct = Discount();
        var luckyDiscount = Round(subtotal * luckyDiscountPct);

        var taxableGeneral = order.Items
            .Where(i => i.Category != "books")
            .Sum(i => i.Category == "food" ? 0 : i.UnitPrice * i.Quantity);

        var foodTax = order.Items
            .Where(i => i.Category == "food")
            .Sum(i => i.UnitPrice * i.Quantity * 0.03m);

        var generalTax = taxableGeneral * 0.07m;
        var taxes = Round(generalTax + foodTax);

        var total = Round(subtotal - luckyDiscount + taxes);

        var summary = new List<string>
        {
            $"Subtotal: ${subtotal:F2}",
            luckyDiscount > 0
                ? $"Descuento de la suerte: -${luckyDiscount:F2} ({luckyDiscountPct * 100:F2}%)"
                : "Descuento de la suerte: $0.00 (0.00%)",
            $"Impuestos: ${taxes:F2}",
            $"TOTAL: ${total:F2}"
        };

        var allLines = new List<string> { header };
        allLines.AddRange(lines);
        allLines.Add("---");
        allLines.AddRange(summary);

        return string.Join("\n", allLines);
    }

    protected virtual decimal Discount()
    {
        var random = new Random();
        decimal luckyDiscountPct = 0;
        if (random.NextDouble() < 0.1)
        {
            luckyDiscountPct = (decimal)random.NextDouble() * 0.05m;
        }
        return luckyDiscountPct;
    }

    protected virtual DateTime GetCurrentDate()
    {
        return DateTime.Now;
    }

    private static decimal Round(decimal n)
    {
        return Math.Round(n, 2);
    }
}

public static class OrderGenerator
{
    private static readonly List<OrderItem> Products = new()
    {
        new OrderItem { Sku = "BK-001", Description = "Libro: Clean Code", UnitPrice = 30, Category = "books" },
        new OrderItem { Sku = "FD-010", Description = "Café en grano 1kg", UnitPrice = 12.5m, Category = "food" },
        new OrderItem { Sku = "GN-777", Description = "Cuaderno A5", UnitPrice = 5.2m, Category = "general" },
        new OrderItem { Sku = "GN-123", Description = "Bolígrafos (pack 10)", UnitPrice = 3.9m, Category = "general" },
        new OrderItem { Sku = "FD-222", Description = "Té verde 200g", UnitPrice = 6.75m, Category = "food" }
    };

    public static readonly List<string> Customers = new() { "Ana", "Luis", "Mar", "Iván", "Sofía" };

    // Utility to generate Orders
    public static Order GenerateOrder(string id, string customerName, int numItems, int quantity)
    {
        var items = new List<OrderItem>();
        for (int i = 0; i < numItems; i++)
        {
            var p = Products[i];
            items.Add(new OrderItem
            {
                Sku = p.Sku,
                Description = p.Description,
                UnitPrice = p.UnitPrice,
                Category = p.Category,
                Quantity = quantity
            });
        }

        return new Order { Id = id, CustomerName = customerName, Items = items };
    }
}
