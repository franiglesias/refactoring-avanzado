// Code smell: Large Class [Clase grande].
// UserAccount acumula muchas responsabilidades no relacionadas
// como autenticación, perfil, notificaciones y gestión de administración,
// lo que dificulta el cambio.

// Ejercicio: Añade autenticación de dos factores (2FA) y preferencias de notificación.

// Tocarás autenticación, estado y notificaciones en una clase inflada,
// aumentando la probabilidad de romper comportamiento no relacionado.

namespace RefactoringAvanzado.CodeSmells.Bloaters;

public class UserAccount
{
    private string name;
    private string email;
    private string password;
    private DateTime lastLogin;
    private int loginAttempts = 0;
    private List<string> notifications = new();
    private bool isAdmin;

    public UserAccount(string name, string email, string password, bool isAdmin = false)
    {
        this.name = name;
        this.email = email;
        this.password = password;
        this.lastLogin = DateTime.Now;
        this.isAdmin = isAdmin;
    }

    // --- Autenticación ---
    public bool Login(string password)
    {
        if (this.password == password)
        {
            lastLogin = DateTime.Now;
            loginAttempts = 0;
            Console.WriteLine("Inicio de sesión exitoso");
            return true;
        }
        else
        {
            loginAttempts++;
            Console.WriteLine("Contraseña incorrecta");
            return false;
        }
    }

    public void ResetPassword(string newPassword)
    {
        password = newPassword;
        Console.WriteLine("Contraseña actualizada");
    }

    // --- Perfil ---
    public void UpdateEmail(string newEmail)
    {
        email = newEmail;
        Console.WriteLine("Correo actualizado");
    }

    public void UpdateName(string newName)
    {
        name = newName;
        Console.WriteLine("Nombre actualizado");
    }

    // --- Notificaciones ---
    public void AddNotification(string message)
    {
        notifications.Add(message);
    }

    public List<string> GetNotifications()
    {
        return notifications;
    }

    public void ClearNotifications()
    {
        notifications.Clear();
    }

    // --- Administración ---
    public void PromoteToAdmin()
    {
        isAdmin = true;
    }

    public void RevokeAdmin()
    {
        isAdmin = false;
    }
}
