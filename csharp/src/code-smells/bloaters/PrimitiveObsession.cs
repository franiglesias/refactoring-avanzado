namespace RefactoringAvanzado.CodeSmells.Bloaters;

public class Order
{
    private readonly string customerName;
    private readonly string customerEmail;
    private readonly string address;
    private readonly double totalAmount;
    private readonly string currency;

    public Order(
        string customerName,
        string customerEmail,
        string address,
        double totalAmount,
        string currency)
    {
        this.customerName = customerName;
        this.customerEmail = customerEmail;
        this.address = address;
        this.totalAmount = totalAmount;
        this.currency = currency;
    }

    public void SendInvoice()
    {
        if (!customerEmail.Contains("@"))
        {
            throw new Exception("Email inválido");
        }
        if (string.IsNullOrEmpty(address))
        {
            throw new Exception("No se ha indicado dirección");
        }
        if (totalAmount <= 0)
        {
            throw new Exception("El monto debe ser mayor que cero");
        }
        Console.WriteLine($"Factura enviada a {customerName} in {address} por {totalAmount} {currency}");
    }
}
