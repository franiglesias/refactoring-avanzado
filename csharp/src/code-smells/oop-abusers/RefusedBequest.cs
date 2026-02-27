namespace RefactoringAvanzado.CodeSmells.OopAbusers;

public class BaseController
{
    public virtual void Start()
    {
        Console.WriteLine("starting");
    }

    public virtual void Stop()
    {
        Console.WriteLine("stopping");
    }

    public virtual void Reset()
    {
        Console.WriteLine("resetting");
    }
}

public class ReadOnlyController : BaseController
{
    public override void Start()
    {
    }

    public override void Stop()
    {
    }
}

public static class RefusedBequest
{
    public static void DemoRefusedBequest(bool readOnly)
    {
        BaseController controller = readOnly ? new ReadOnlyController() : new BaseController();
        controller.Start();
        controller.Stop();
    }
}
