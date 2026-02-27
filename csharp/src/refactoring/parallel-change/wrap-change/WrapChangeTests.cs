using FluentAssertions;
using Xunit;

namespace RefactoringAvanzado.Refactoring.ParallelChange.WrapChange;

public class WrapChangeTests
{
    [Fact]
    public void ShouldSendWelcomeEmail()
    {
        var message = EmailNotifications.NotifyWelcome("john@example.com");

        message.Should().Be("Email sent to john@example.com, subject: Welcome!, body: Thanks for joining our app.");
    }

    [Fact]
    public void ShouldSendPasswordResetEmail()
    {
        var message = EmailNotifications.NotifyPasswordReset("jane@example.com");

        message.Should().Be("Email sent to jane@example.com, subject: Reset your password, body: Click the link to reset...");
    }

    [Fact]
    public void ShouldSendOrderConfirmationEmail()
    {
        var message = EmailNotifications.NotifyOrderConfirmation("alice@example.com", "ORD-12345");

        message.Should().Be("Email sent to alice@example.com, subject: Order Confirmation, body: Your order ORD-12345 has been confirmed.");
    }
}
