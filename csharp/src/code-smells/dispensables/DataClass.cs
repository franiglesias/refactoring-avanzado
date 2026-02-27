namespace RefactoringAvanzado.CodeSmells.Dispensables;

public class UserRecord
{
    public string Id { get; set; }
    public string Name { get; set; }
    public string Email { get; set; }
    public DateTime CreatedAt { get; set; }

    public UserRecord(string id, string name, string email, DateTime createdAt)
    {
        Id = id;
        Name = name;
        Email = email;
        CreatedAt = createdAt;
    }
}

public class UserService
{
    public UserRecord CreateUser(string name, string email)
    {
        if (!email.Contains('@'))
        {
            throw new ArgumentException("Invalid email");
        }

        return new UserRecord(Guid.NewGuid().ToString(), name, email, DateTime.Now);
    }

    public void UpdateUserEmail(UserRecord user, string newEmail)
    {
        if (!newEmail.Contains('@'))
        {
            throw new ArgumentException("Invalid email");
        }
        user.Email = newEmail;
    }
}

public class UserReportGenerator
{
    public string GenerateUserSummary(UserRecord user)
    {
        return $"User {user.Name} ({user.Email}) created on {user.CreatedAt.ToShortDateString()}";
    }
}

public static class DataClass
{
    public static string DemoDataClass()
    {
        var service = new UserService();
        var report = new UserReportGenerator();
        var user = service.CreateUser("Lina", "lina@example.com");
        service.UpdateUserEmail(user, "lina+news@example.com");
        return report.GenerateUserSummary(user);
    }
}
