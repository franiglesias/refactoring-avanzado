namespace RefactoringAvanzado.CalisthenicExercises;

public class ReportService
{
    private Dictionary<string, string> cache = new Dictionary<string, string>();

    public string GenerateCsvReportFromJson(string jsonInput, string delimiter = ",")
    {
        dynamic data;
        try
        {
            data = System.Text.Json.JsonSerializer.Deserialize<object>(jsonInput);
        }
        catch
        {
            throw new Exception("Invalid JSON");
        }

        if (data is not System.Text.Json.JsonElement element || element.ValueKind != System.Text.Json.JsonValueKind.Array)
        {
            throw new Exception("Expected array");
        }

        var array = element.EnumerateArray().ToList();
        if (array.Count == 0)
        {
            return string.Empty;
        }

        var firstRow = array[0];
        var headers = firstRow.EnumerateObject().Select(p => p.Name).ToList();
        var lines = new List<string> { string.Join(delimiter, headers) };

        foreach (var row in array)
        {
            var values = headers.Select(h =>
            {
                if (row.TryGetProperty(h, out var prop))
                {
                    return prop.ToString();
                }
                return string.Empty;
            });
            lines.Add(string.Join(delimiter, values));
        }

        var result = string.Join("\n", lines);
        cache["last"] = result;
        return result;
    }
}
