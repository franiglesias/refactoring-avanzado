namespace RefactoringAvanzado.CodeSmells.Bloaters;

public class Invoice
{
    private readonly string customerName;
    private readonly string customerCity;
    private readonly string customerStreet;
    private readonly string customerZip;

    public Invoice(
        string customerName,
        string customerStreet,
        string customerCity,
        string customerZip)
    {
        this.customerZip = customerZip;
        this.customerStreet = customerStreet;
        this.customerCity = customerCity;
        this.customerName = customerName;
    }

    public string Print()
    {
        return
            $"Factura para: {customerName}\n" +
            $"Dirección: {customerStreet}, {customerCity}, {customerZip}";
    }
}

public class ShippingLabel
{
    private readonly string customerName;
    private readonly string customerStreet;
    private readonly string customerCity;
    private readonly string customerZip;

    public ShippingLabel(
        string customerName,
        string customerStreet,
        string customerCity,
        string customerZip)
    {
        this.customerZip = customerZip;
        this.customerCity = customerCity;
        this.customerStreet = customerStreet;
        this.customerName = customerName;
    }

    public string Print()
    {
        return
            $"Enviar a: {customerName}\n" +
            $"{customerStreet}, {customerCity}, {customerZip}";
    }
}
