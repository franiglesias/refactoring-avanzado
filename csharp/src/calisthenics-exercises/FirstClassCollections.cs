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

public static class ProductOperations
{
    public static void AddProduct(List<Product> products, Product product)
    {
        var exists = products.Any(p => p.Id == product.Id);
        if (!exists) products.Add(product);
    }

    public static double TotalPrice(List<Product> products)
    {
        return products.Sum(p => p.Price);
    }

    public static List<Product> RemoveProduct(List<Product> products, string productId)
    {
        return products.Where(p => p.Id != productId).ToList();
    }
}
