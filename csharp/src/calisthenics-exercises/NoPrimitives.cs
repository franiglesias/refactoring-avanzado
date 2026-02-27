namespace RefactoringAvanzado.CalisthenicExercises;

public static class TransferOperations
{
    public static string Transfer(
        double amount,
        string fromIban,
        string toIban,
        string currency)
    {
        if (string.IsNullOrEmpty(fromIban) || string.IsNullOrEmpty(toIban) || string.IsNullOrEmpty(currency))
        {
            throw new Exception("Missing data");
        }
        if (amount <= 0)
        {
            throw new Exception("Invalid amount");
        }
        if (fromIban.Length < 24)
        {
            throw new Exception("Invalid IBAN From");
        }
        else if (toIban.Length < 24)
        {
            throw new Exception("Invalid IBAN To");
        }
        if (fromIban == toIban)
        {
            throw new Exception("Same account");
        }
        return $"{amount} {currency} from {fromIban} to {toIban}";
    }
}
