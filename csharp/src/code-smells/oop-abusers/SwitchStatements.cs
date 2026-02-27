namespace RefactoringAvanzado.CodeSmells.OopAbusers;

public enum EmployeeKind
{
    Engineer,
    Manager,
    Sales
}

public record EmployeeRecord(EmployeeKind Kind, double Base, double? Bonus = null, double? Commission = null);

public static class SwitchStatements
{
    public static double CalculatePay(EmployeeRecord rec)
    {
        return rec.Kind switch
        {
            EmployeeKind.Engineer => rec.Base,
            EmployeeKind.Manager => rec.Base + (rec.Bonus ?? 0),
            EmployeeKind.Sales => rec.Base + (rec.Commission ?? 0),
            _ => throw new ArgumentException($"Unknown employee kind: {rec.Kind}")
        };
    }

    public static double[] DemoSwitchStatements()
    {
        return
        [
            CalculatePay(new EmployeeRecord(EmployeeKind.Engineer, 1000)),
            CalculatePay(new EmployeeRecord(EmployeeKind.Manager, 1000, Bonus: 200)),
            CalculatePay(new EmployeeRecord(EmployeeKind.Sales, 800, Commission: 500))
        ];
    }
}
