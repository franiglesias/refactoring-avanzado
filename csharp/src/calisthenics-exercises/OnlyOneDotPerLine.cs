namespace RefactoringAvanzado.CalisthenicExercises;

public class Address
{
    private readonly string street;
    private readonly string city;

    public Address(string street, string city)
    {
        this.street = street;
        this.city = city;
    }

    public string GetCity()
    {
        return city;
    }
}

public class Customer
{
    private readonly string name;
    private readonly Address address;

    public Customer(string name, Address address)
    {
        this.name = name;
        this.address = address;
    }

    public Address GetAddress()
    {
        return address;
    }
}

public class Order
{
    private readonly Customer customer;

    public Order(Customer customer)
    {
        this.customer = customer;
    }

    public Customer GetCustomer()
    {
        return customer;
    }
}

public static class OrderExample
{
    public static string GetOrderDestination()
    {
        var order = new Order(new Customer("John Doe", new Address("Elm Street", "Madrid")));
        var destination = order.GetCustomer().GetAddress().GetCity();
        return destination;
    }
}
