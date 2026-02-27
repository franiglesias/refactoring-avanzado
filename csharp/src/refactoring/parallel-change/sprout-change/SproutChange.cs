using System;
using System.Collections.Generic;
using System.Linq;

namespace RefactoringAvanzado.Refactoring.ParallelChange.SproutChange;

public class CartItem
{
    public string Id { get; init; } = string.Empty;
    public decimal Price { get; init; }
    public int Qty { get; init; }
    public string? Category { get; init; }
}

public enum Region
{
    US,
    EU
}

public interface ITaxPolicy
{
    // Interface placeholder for future implementation
}

public static class TaxCalculator
{
    // Regla existente: un único impuesto plano por región; los libros y la comida están exentos en la UE
    private static decimal CalculateUSTax(List<CartItem> cart)
    {
        var usSubtotal = cart.Sum(it => it.Price * it.Qty);
        return usSubtotal * 0.07m; // 7% plano
    }

    private static decimal CalculateEUTax(List<CartItem> cart)
    {
        var taxable = cart
            .Where(it => it.Category != "books" && it.Category != "food")
            .Sum(it => it.Price * it.Qty);
        return taxable * 0.2m; // 20% plano solo sobre los ítems gravables
    }

    private static decimal CalculateDefault(List<CartItem> cart)
    {
        return 0;
    }

    // (reglas embebidas en línea)
    public static decimal CalculateTotal(List<CartItem> cart, Region region)
    {
        var subtotal = cart.Sum(it => it.Price * it.Qty);

        decimal tax = region switch
        {
            Region.US => CalculateUSTax(cart),
            Region.EU => CalculateEUTax(cart),
            _ => CalculateDefault(cart)
        };

        return RoundCurrency(subtotal + tax);
    }

    public static decimal RoundCurrency(decimal amount)
    {
        return Math.Round(amount, 2);
    }
}

// Uso de ejemplo, mantenido simple para estudiantes
public static class SproutChangeDemo
{
    public static decimal DemoSprout()
    {
        var cart = new List<CartItem>
        {
            new() { Id = "p1", Price = 10, Qty = 2, Category = "general" },
            new() { Id = "b1", Price = 20, Qty = 1, Category = "books" }
        };
        return TaxCalculator.CalculateTotal(cart, Region.EU);
    }
}
