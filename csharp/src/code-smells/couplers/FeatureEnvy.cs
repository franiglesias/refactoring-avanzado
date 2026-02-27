namespace RefactoringAvanzado.CodeSmells.Couplers;

public class Customer
{
    public string Name { get; }
    public string Street { get; }
    public string City { get; }
    public string Zip { get; }

    public Customer(string name, string street, string city, string zip)
    {
        Name = name;
        Street = street;
        City = city;
        Zip = zip;
    }
}

public class ShippingCalculator
{
    public int Cost(Customer customer)
    {
        var baseAmount = customer.Zip.StartsWith("9") ? 10 : 20;
        var distant = customer.City.Length > 6 ? 5 : 0;
        return baseAmount + distant;
    }
}

public static class FeatureEnvyDemo
{
    public static int DemoFeatureEnvy(Customer c)
    {
        return new ShippingCalculator().Cost(c);
    }
}
