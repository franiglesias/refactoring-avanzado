namespace RefactoringAvanzado.CodeSmells.Bloaters;

public class ReportGenerator
{
    public void GenerateReport(
        string title,
        DateTime startDate,
        DateTime endDate,
        bool includeCharts,
        bool includeSummary,
        string authorName,
        string authorEmail)
    {
        Console.WriteLine($"Generando reporte: {title}");
        Console.WriteLine($"Desde {startDate.ToShortDateString()} hasta {endDate.ToShortDateString()}");
        Console.WriteLine($"Autor: {authorName} ({authorEmail})");
        if (includeCharts) Console.WriteLine("Incluyendo gráficos...");
        if (includeSummary) Console.WriteLine("Incluyendo resumen...");
        Console.WriteLine("Reporte generado exitosamente.");
    }
}

public static class LongParameterListDemo
{
    public static void DemoLongParameterList()
    {
        var gen = new ReportGenerator();
        gen.GenerateReport(
            "Ventas Q1",
            new DateTime(2025, 1, 1),
            new DateTime(2025, 3, 31),
            true,
            false,
            "Pat Smith",
            "pat@example.com"
        );
    }
}
