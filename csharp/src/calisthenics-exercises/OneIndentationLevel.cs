namespace RefactoringAvanzado.CalisthenicExercises;

public class OrderItem
{
    public double Price { get; set; }

    public OrderItem(double price)
    {
        Price = price;
    }
}

public class OrderCustomer
{
    public bool IsVip { get; set; }

    public OrderCustomer(bool isVip)
    {
        IsVip = isVip;
    }
}

public class Order
{
    public string Id { get; set; }
    public List<OrderItem> Items { get; set; }
    public OrderCustomer Customer { get; set; }

    public Order(string id, List<OrderItem> items, OrderCustomer customer)
    {
        Id = id;
        Items = items;
        Customer = customer;
    }
}

public static class OrderProcessor
{
    public static double ProcessOrdersWithDiscounts(List<Order> orders)
    {
        double total = 0;
        foreach (var order in orders)
        {
            if (order.Items != null && order.Items.Count > 0)
            {
                foreach (var item in order.Items)
                {
                    if (order.Customer != null && order.Customer.IsVip)
                    {
                        if (item.Price > 100)
                        {
                            total += item.Price * 0.8;
                        }
                        else
                        {
                            total += item.Price * 0.9;
                        }
                    }
                    else
                    {
                        if (item.Price > 100)
                        {
                            total += item.Price * 0.95;
                        }
                        else
                        {
                            total += item.Price;
                        }
                    }
                }
            }
        }
        return total;
    }
}
