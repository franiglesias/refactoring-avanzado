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

}

public static class TaxCalculator
{
    // Regla existente: un único impuesto plano por región; los libros y la comida están exentos en la UE
    public static decimal CalculateTotal(List<CartItem> cart, Region region)
    {
        var subtotal = cart.Sum(it => it.Price * it.Qty);

        var tax = CalculateTax(subtotal);

        return RoundCurrency(subtotal + tax);
    }

    public static decimal RoundCurrency(decimal amount)
    {
        return Math.Round(amount, 2);
    }

    public static decimal CalculateTax(decimal subtotal)
    {
        if (region == Region.US || region == Region.EU) {
             return NewCalculateTax(subtotal);
        }

         return OldCalculateTax(subtotal);
    }

    public static decimal NewCalculateTax(decimal subtotal)
    {
        decimal tax = 0;
        if (region == Region.US)
        {
            tax = subtotal * 0.17m;
        }
    }

    public static decimal OldCalculateTax(decimal subtotal)
    {
        decimal tax = 0;
        if (region == Region.US)
        {
            tax = subtotal * 0.07m; // 7% plano
        }
        else if (region == Region.EU)
        {
            // exenciones ingenuas en línea
            var taxable = cart
                .Where(it => it.Category != "books" && it.Category != "food")
                .Sum(it => it.Price * it.Qty);
            tax = taxable * 0.2m; // 20% plano solo sobre los ítems gravables
        }
        return tax;
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
