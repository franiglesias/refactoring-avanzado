namespace RefactoringAvanzado.CodeSmells.Couplers;

public class Catalog
{
    private readonly Dictionary<string, string> _items = new();

    public void Add(string id, string name)
    {
        _items[id] = name;
    }

    public string? Find(string id)
    {
        return _items.TryGetValue(id, out var name) ? name : null;
    }

    public string[] List()
    {
        return _items.Values.ToArray();
    }
}

public class Shop
{
    private readonly Catalog _catalog;

    public Shop(Catalog catalog)
    {
        _catalog = catalog;
    }

    public void Add(string id, string name)
    {
        _catalog.Add(id, name);
    }

    public string? Find(string id)
    {
        return _catalog.Find(id);
    }

    public string[] List()
    {
        return _catalog.List();
    }
}

public static class MiddlemanDemo
{
    public static string[] DemoMiddleman()
    {
        var c = new Catalog();
        var s = new Shop(c);
        s.Add("1", "Book");
        s.Add("2", "Pen");
        return s.List();
    }
}
