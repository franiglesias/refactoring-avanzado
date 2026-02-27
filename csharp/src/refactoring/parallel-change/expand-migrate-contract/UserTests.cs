using System.Collections.Generic;
using FluentAssertions;
using Xunit;

namespace RefactoringAvanzado.Refactoring.ParallelChange.ExpandMigrateContract;

public class UserTests
{
    private readonly User _alice = new("u-1", "Alice Smith", "alice@example.com");
    private readonly User _bob = new("u-2", "Bob Jones", "bob@example.com");

    [Fact]
    public void FormatGreeting_ShouldGreetUserByName()
    {
        var result = UserOperations.FormatGreeting(_alice);

        result.Should().Be("Hello, Alice Smith!");
    }

    [Fact]
    public void FormatEmailHeader_ShouldFormatEmailHeaderWithNameAndEmail()
    {
        var result = UserOperations.FormatEmailHeader(_alice);

        result.Should().Be("From: Alice Smith <alice@example.com>");
    }

    [Fact]
    public void FormatDisplayName_ShouldFormatNameWithId()
    {
        var result = UserOperations.FormatDisplayName(_alice);

        result.Should().Be("Alice Smith (u-1)");
    }

    [Fact]
    public void BuildUserSummary_ShouldListAllUserNames()
    {
        var result = UserOperations.BuildUserSummary(new List<User> { _alice, _bob });

        result.Should().Be("- Alice Smith\n- Bob Jones");
    }

    [Fact]
    public void BuildUserSummary_ShouldReturnEmptyStringForEmptyList()
    {
        var result = UserOperations.BuildUserSummary(new List<User>());

        result.Should().Be("");
    }
}
