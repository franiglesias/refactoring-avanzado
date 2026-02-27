namespace RefactoringAvanzado.Tests;

using Xunit;
using FluentAssertions;

public class HealthCheckTests
{
    [Fact]
    public void ShouldVerifyTestFrameworkIsWorking()
    {
        // Arrange
        var expected = 42;

        // Act
        var actual = 42;

        // Assert
        actual.Should().Be(expected);
    }

    [Fact]
    public void ShouldVerifyBasicArithmetic()
    {
        // Arrange
        var a = 2;
        var b = 2;

        // Act
        var result = a + b;

        // Assert
        result.Should().Be(4);
    }

    [Theory]
    [InlineData(1, 1, 2)]
    [InlineData(2, 3, 5)]
    [InlineData(-1, 1, 0)]
    [InlineData(0, 0, 0)]
    public void ShouldVerifyParameterizedTests(int a, int b, int expected)
    {
        // Act
        var result = a + b;

        // Assert
        result.Should().Be(expected);
    }
}
