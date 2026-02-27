namespace RefactoringAvanzado.CalisthenicExercises;

public static class ShippingCostCalculator
{
    public static double ShippingCost(double weightKg, string destination)
    {
        if (destination == "DOMESTIC")
        {
            if (weightKg <= 1)
            {
                return 5;
            }
            else if (weightKg <= 5)
            {
                return 10;
            }
            else
            {
                return 20;
            }
        }
        else
        {
            if (weightKg <= 1)
            {
                return 15;
            }
            else if (weightKg <= 5)
            {
                return 25;
            }
            else
            {
                return 40;
            }
        }
    }
}
