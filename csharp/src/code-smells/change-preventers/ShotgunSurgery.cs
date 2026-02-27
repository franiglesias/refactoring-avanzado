namespace RefactoringAvanzado.CodeSmells.ChangePreventers;

public record LineItem(string Name, double Price, int Qty);

public class PriceCalculator
{
    public double TotalWithTax(LineItem[] items)
    {
        var subtotal = items.Sum(i => i.Price * i.Qty);
        var tax = subtotal * 0.21;
        return subtotal + tax;
    }
}

public class InvoiceService
{
    public double CreateTotal(LineItem[] items)
    {
        var baseAmount = items.Sum(i => i.Price * i.Qty);
        var vat = baseAmount * 0.21;
        return baseAmount + vat;
    }
}

public class SalesReport
{
    public string Summarize(LineItem[] items)
    {
        var sum = items.Sum(i => i.Price * i.Qty);
        var tax = sum * 0.21;
        var total = sum + tax;
        return $"total={total:F2}";
    }
}

public class LoyaltyPoints
{
    public int Points(LineItem[] items)
    {
        var baseAmount = items.Sum(i => i.Price * i.Qty);
        var withTax = baseAmount + baseAmount * 0.21;
        return (int)Math.Floor(withTax / 10);
    }
}

public static class ShotgunDemo
{
    public static (double, double) DemoShotgun(LineItem[] items)
    {
        var calc = new PriceCalculator().TotalWithTax(items);
        var inv = new InvoiceService().CreateTotal(items);
        return (calc, inv);
    }
}
