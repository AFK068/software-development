using Spectre.Console;

namespace mini_hw_1.Infrastructure.UI;

public class MenuWork
{
    public static string ShowMenu(string title, params string[] items)
    {
        var selection = AnsiConsole.Prompt(
            new SelectionPrompt<string>()
                .Title(title)
                .PageSize(10)
                .AddChoices(items)
        );

        return selection;
    }
    
    public static string ShowAnimalTypeMenu()
    {
        return ShowMenu("Выберите тип животного:", "Обезьяна", "Кролик", "Тигр", "Волк");
    }
}