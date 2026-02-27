namespace RefactoringAvanzado.CodeSmells.Couplers;

public class Team
{
    private string _name;
    private Budget _budget;
    private Manager? _manager;

    public Team(string name, int budget)
    {
        _budget = new Budget(budget);
        _name = name;
    }

    public void AssignManager(Manager m)
    {
        _manager = m;
        m.AssignTeam(this);
    }

    public void RaiseBudget(int amount)
    {
        _budget = new Budget(_budget.Amount + amount);
    }

    public void Rename(string newName)
    {
        _name = newName;
    }
}

public class Manager
{
    private readonly string _name;
    private Team? _team;

    public Manager(string name)
    {
        _name = name;
    }

    public void AssignTeam(Team t)
    {
        if (_team != null)
        {
            throw new InvalidOperationException("Team already assigned");
        }
        _team = t;
    }

    public void RaiseTeamBudget(int amount)
    {
        _team?.RaiseBudget(amount);
    }

    public void RenameTeam(string newName)
    {
        _team?.Rename(newName);
    }
}

public class Budget
{
    public int Amount { get; }

    public Budget(int amount)
    {
        if (amount < 900)
        {
            throw new InvalidOperationException("Budget too low");
        }
        if (amount > 3000)
        {
            throw new InvalidOperationException("Budget too high");
        }
        Amount = amount;
    }

    public void Raise(int amount)
    {
        if (Amount + amount > 3000)
        {
            throw new InvalidOperationException("Budget exceeded");
        }
    }
}

public static class InappropriateIntimacyDemo
{
    public static Team DemoInappropriateIntimacy()
    {
        var t = new Team("Core", 1000);
        var m = new Manager("Alice");
        m.AssignTeam(new Team("Frontend", 2000));
        t.AssignManager(m);
        m.RaiseTeamBudget(200);
        m.RenameTeam("Platform");
        return t;
    }
}
