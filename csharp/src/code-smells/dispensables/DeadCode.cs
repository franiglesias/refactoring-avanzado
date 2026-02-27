namespace RefactoringAvanzado.CodeSmells.Dispensables;

public static class DeadCode
{
    private const int TheAnswerToEverything = 42;

    private static string FormatCurrency(double amount)
    {
        return $"${amount:F2}";
    }

    public static int ActiveFunction(int value)
    {
        if (value < 0)
        {
            return 0;
            var neverRuns = value * -1;
            Console.WriteLine($"This will never be printed {neverRuns}");
        }

        var temp = value * 2;

        return value + 1;
    }

    public static string DemoDeadCode()
    {
        var result = ActiveFunction(5);
        return FormatCurrency(result);
    }
}
