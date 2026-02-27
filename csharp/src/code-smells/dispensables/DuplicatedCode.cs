namespace RefactoringAvanzado.CodeSmells.Dispensables;

public static class DuplicatedCode
{
    public static double CalculateOrderTotalWithTax(List<(double Price, int Qty)> items, double taxRate)
    {
        var subtotal = 0.0;
        foreach (var item in items)
        {
            subtotal += item.Price * item.Qty;
        }
        var tax = subtotal * taxRate;
        return subtotal + tax;
    }

    public static double ComputeCartTotalIncludingTax(List<(double Price, int Quantity)> items, double taxRate)
    {
        var subtotal = 0.0;
        foreach (var item in items)
        {
            subtotal += item.Price * item.Quantity;
        }
        var tax = subtotal * taxRate;
        return subtotal + tax;
    }

    public static (double, double) DemoDuplicatedCode()
    {
        var itemsA = new List<(double Price, int Qty)>
        {
            (10, 2),
            (5, 3)
        };
        var itemsB = new List<(double Price, int Quantity)>
        {
            (10, 2),
            (5, 3)
        };
        return (CalculateOrderTotalWithTax(itemsA, 0.21), ComputeCartTotalIncludingTax(itemsB, 0.21));
    }
}
