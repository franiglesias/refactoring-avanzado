namespace RefactoringAvanzado.CodeSmells.Couplers;

public class Level2
{
    private readonly int _value;

    public Level2(int value)
    {
        _value = value;
    }

    public int GetValue()
    {
        return _value;
    }
}

public class Level1
{
    private readonly Level2 _next;

    public Level1(Level2 next)
    {
        _next = next;
    }

    public Level2 GetNext()
    {
        return _next;
    }
}

public class Root
{
    private readonly Level1 _next;

    public Root(Level1 next)
    {
        _next = next;
    }

    public Level1 GetNext()
    {
        return _next;
    }
}

public static class MessageChainsDemo
{
    public static int ReadDeep(Root root)
    {
        return root.GetNext().GetNext().GetValue();
    }
}
