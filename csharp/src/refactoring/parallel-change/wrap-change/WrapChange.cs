namespace RefactoringAvanzado.Refactoring.ParallelChange.WrapChange;

// Dependencia externa que NO podemos modificar (simulada)
// En el mundo real, esto podría ser una librería de terceros o un servicio legacy
public class LegacyEmailService
{
    public virtual string SendEmail(string to, string subject, string body)
    {
        // Implementación rígida y limitada
        // Problema: no valida emails, no soporta retry, no tiene logging, etc.
        return $"Email sent to {to}, subject: {subject}, body: {body}";
    }
}

// Código de nuestra aplicación que usa directamente el servicio legacy
public static class EmailNotifications
{
    private static readonly LegacyEmailService LegacyService = new();

    public static string NotifyWelcome(string userEmail)
    {
        return LegacyService.SendEmail(userEmail, "Welcome!", "Thanks for joining our app.");
    }

    public static string NotifyPasswordReset(string userEmail)
    {
        return LegacyService.SendEmail(userEmail, "Reset your password", "Click the link to reset...");
    }

    public static string NotifyOrderConfirmation(string userEmail, string orderId)
    {
        return LegacyService.SendEmail(
            userEmail,
            "Order Confirmation",
            $"Your order {orderId} has been confirmed."
        );
    }
}
