namespace RefactoringAvanzado.CalisthenicExercises;

public class CartItem
{
    public string Id { get; set; }
    public double Price { get; set; }
    public int Qty { get; set; }

    public CartItem(string id, double price, int qty)
    {
        Id = id;
        Price = price;
        Qty = qty;
    }
}

public class CheckoutSession
{
    private List<CartItem> cartItems = new List<CartItem>();
    private string? customerId = null;
    private string? shippingAddress = null;
    private string? billingAddress = null;
    private string? couponCode = null;
    private string? paymentMethod = null;
    private string currency = "USD";
    private double taxRate = 0.21;

    public void AddItem(string id, double price, int qty)
    {
        cartItems.Add(new CartItem(id, price, qty));
    }

    public double Total()
    {
        var subtotal = cartItems.Sum(i => i.Price * i.Qty);
        var discount = couponCode != null ? 10 : 0;
        var taxed = (subtotal - discount) * (1 + taxRate);
        return currency == "USD" ? taxed : taxed * 0.9;
    }
}
