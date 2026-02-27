namespace RefactoringAvanzado.CodeSmells.ChangePreventers;

public record User(string Id, string Name, string Email);

public class ProfileManager
{
    private readonly Dictionary<string, User> _store = new();

    public void Register(User user)
    {
        if (string.IsNullOrWhiteSpace(user.Name))
            throw new InvalidOperationException("invalid name");
        if (!user.Email.Contains('@'))
            throw new InvalidOperationException("invalid email");
        _store[user.Id] = user;
    }

    public void UpdateEmail(string id, string newEmail)
    {
        if (!newEmail.Contains('@'))
            throw new InvalidOperationException("invalid email");
        if (!_store.TryGetValue(id, out var u))
            throw new InvalidOperationException("not found");
        _store[id] = u with { Email = newEmail };
    }

    public string ExportAsJson()
    {
        return System.Text.Json.JsonSerializer.Serialize(_store.Values);
    }

    public string ExportAsCsv()
    {
        var rows = new List<string> { "id,name,email" };
        rows.AddRange(_store.Values.Select(u => $"{u.Id},{u.Name},{u.Email}"));
        return string.Join('\n', rows);
    }

    public string SendWelcomeEmail(User user)
    {
        return $"Welcome {user.Name}! Sent to {user.Email}";
    }
}

public static class DivergentChangeDemo
{
    public static string DemoDivergentChange(ProfileManager pm, User u)
    {
        pm.Register(u);
        pm.UpdateEmail(u.Id, u.Email);
        return pm.ExportAsJson();
    }
}
