namespace RefactoringAvanzado.CalisthenicExercises;

public class BankAccount
{
    private double _balance;

    public BankAccount(double initialBalance = 0)
    {
        _balance = initialBalance;
    }

    public double Balance
    {
        get => _balance;
        set
        {
            if (value < 0) throw new Exception("Negative");
            _balance = value;
        }
    }

    public void Transfer(double amount, BankAccount to)
    {
        Balance -= amount;
        to.Balance += amount;
    }
}

public static class BankAccountOperations
{
    public static void Pay(BankAccount account, double amount)
    {
        account.Balance = account.Balance - amount;
    }
}
