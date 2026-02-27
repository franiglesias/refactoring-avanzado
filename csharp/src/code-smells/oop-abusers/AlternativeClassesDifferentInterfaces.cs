namespace RefactoringAvanzado.CodeSmells.OopAbusers;

public class TextLogger
{
    public void Log(string message)
    {
        Console.WriteLine($"[text] {message}");
    }
}

public class MessageWriter
{
    public void Write(string entry)
    {
        Console.WriteLine($"[text] {entry}");
    }
}

public static class AlternativeClassesDifferentInterfaces
{
    public static void UseAltClasses(string choice, string msg)
    {
        if (choice == "logger")
        {
            new TextLogger().Log(msg);
        }
        else
        {
            new MessageWriter().Write(msg);
        }
    }
}
