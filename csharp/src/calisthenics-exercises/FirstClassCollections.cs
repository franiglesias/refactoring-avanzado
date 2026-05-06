namespace RefactoringAvanzado.CalisthenicExercises;

public class Product
{
    public string Id { get; set; }
    public double Price { get; set; }

    public Product(string id, double price)
    {
        Id = id;
        Price = price;
    }
}

public class ProductOperations
{
    private readonly List<Product> _products = new();

    public void AddProduct(Product product)
    {
        var exists = _products.Any(p => p.Id == product.Id);
        if (!exists) _products.Add(product);
    }

    public double TotalPrice()
    {
        return _products.Sum(p => p.Price);
    }

    public void RemoveProduct(string productId)
    {
        _products.RemoveAll(p => p.Id == productId);
    }
}
