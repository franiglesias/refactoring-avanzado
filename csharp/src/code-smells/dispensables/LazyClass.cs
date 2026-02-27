namespace RefactoringAvanzado.CodeSmells.Dispensables;

public record Address(string Name, string Line1, string? City);

public class ShippingLabelBuilder
{
    public string Build(Address a)
    {
        return $"{a.Name} — {a.Line1}{(a.City != null ? $", {a.City}" : "")}";
    }
}

public static class LazyClass
{
    public static void PrintShippingLabel()
    {
        var address = new Address("John Doe", "123 Main St", "New York");

        var labelBuilder = new ShippingLabelBuilder();
        var label = labelBuilder.Build(address);
        Console.WriteLine(label);
    }
}
