namespace RefactoringAvanzado.CalisthenicExercises;

public class C
{
    public string U { get; set; }
    public string P { get; set; }
    public string S { get; set; }
    public string E { get; set; }

    public C(string u, string p, string s, string e)
    {
        U = u;
        P = p;
        S = s;
        E = e;
    }

    public string Cnx()
    {
        return $"{U}:{P}@{S}/{E}";
    }
}
