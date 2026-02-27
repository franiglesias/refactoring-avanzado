namespace RefactoringAvanzado.CodeSmells.ChangePreventers;

public abstract class Component
{
    public abstract string Draw(Renderer renderer);
}

public class Button : Component
{
    public string Label { get; }

    public Button(string label)
    {
        Label = label;
    }

    public override string Draw(Renderer renderer)
    {
        return renderer.RenderButton(this);
    }
}

public class TextBox : Component
{
    public string Text { get; }

    public TextBox(string text)
    {
        Text = text;
    }

    public override string Draw(Renderer renderer)
    {
        return renderer.RenderTextBox(this);
    }
}

public abstract class Renderer
{
    public abstract string RenderButton(Button b);
    public abstract string RenderTextBox(TextBox t);
}

public class HtmlRenderer : Renderer
{
    public override string RenderButton(Button b)
    {
        return $"<button>{b.Label}</button>";
    }

    public override string RenderTextBox(TextBox t)
    {
        return $"<input value=\"{t.Text}\"/>";
    }
}

public class MarkdownRenderer : Renderer
{
    public override string RenderButton(Button b)
    {
        return $"[{b.Label}]";
    }

    public override string RenderTextBox(TextBox t)
    {
        return $"_{t.Text}_";
    }
}

public static class ParallelHierarchyDemo
{
    public static string[] DemoParallelHierarchy()
    {
        Component[] comps = { new Button("Save"), new TextBox("name") };
        var renderer = new HtmlRenderer();
        return comps.Select(c => c.Draw(renderer)).ToArray();
    }
}
