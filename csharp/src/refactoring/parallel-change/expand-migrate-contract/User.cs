using System.Collections.Generic;
using System.Linq;

namespace RefactoringAvanzado.Refactoring.ParallelChange.ExpandMigrateContract;

public class User
{
    public string Id { get; }
    public string Name { get; }
    public string Email { get; }
    public string? FirstName { get; }
    public string? LastName { get; }

    public User(string id, string name, string email, string? firstName = null, string? lastName = null)
    {
        Id = id;
        Name = name;
        Email = email;
        FirstName = firstName;
        LastName = lastName;
    }
}

// --- Consumers of User.Name ---

public static class UserOperations
{
    public static string FormatGreeting(User user)
    {
        return $"Hello, {user.Name}!";
    }

    public static string FormatEmailHeader(User user)
    {
        return $"From: {user.Name} <{user.Email}>";
    }

    public static string FormatDisplayName(User user)
    {
        return $"{user.Name} ({user.Id})";
    }

    public static string BuildUserSummary(List<User> users)
    {
        return string.Join("\n", users.Select(u => $"- {u.Name}"));
    }
}
