using System.Collections.Generic;
using FluentAssertions;
using Xunit;

namespace RefactoringAvanzado.Refactoring.ParallelChange.SproutChange;

public class SproutChangeTests
{
    private static decimal ExecuteSubject(List<CartItem> cart, Region region)
    {
        return TaxCalculator.CalculateTotal(cart, region);
    }

    private static readonly List<CartItem> Cart = new()
    {
        new CartItem { Id = "p1", Price = 10, Qty = 2, Category = "general" },
        new CartItem { Id = "b1", Price = 20, Qty = 1, Category = "books" },
        new CartItem { Id = "f1", Price = 15, Qty = 4, Category = "food" }
    };

    [Fact]
    public void ShouldCalculateTotalForEU()
    {
        var result = ExecuteSubject(Cart, Region.EU);

        result.Should().Be(104);
    }

    [Fact]
    public void ShouldCalculateTotalForUS()
    {
        var result = ExecuteSubject(Cart, Region.US);

        result.Should().Be(107);
    }
}
