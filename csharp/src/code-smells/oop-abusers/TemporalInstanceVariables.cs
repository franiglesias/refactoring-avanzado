namespace RefactoringAvanzado.CodeSmells.OopAbusers;

public class PizzaOrder
{
    private string? _size;
    private List<string> _toppings = [];
    private string? _address;

    public void Start(string size)
    {
        _size = size;
        _toppings = [];
        _address = null;
    }

    public void AddTopping(string topping)
    {
        if (_size == null)
        {
            return;
        }
        _toppings.Add(topping);
    }

    public void SetDeliveryAddress(string address)
    {
        _address = address;
    }

    public string Place()
    {
        var summary = $"Pizza {_size ?? "?"} to {_address ?? "UNKNOWN"} with [{string.Join(", ", _toppings)}]";
        _size = null;
        _address = null;
        _toppings = [];
        return summary;
    }
}

public static class TemporalInstanceVariables
{
    public static string DemoPizzaOrder()
    {
        var o = new PizzaOrder();
        o.Start("L");
        o.AddTopping("pepperoni");
        o.AddTopping("mushroom");
        o.SetDeliveryAddress("123 Main St");
        return o.Place();
    }
}
