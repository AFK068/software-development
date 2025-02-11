using Spectre.Console;

namespace mini_hw_1.Infrastructure.UI;

public class UserInputHandler
{
    private static string GetStringInput(string prompt)
    {
        return AnsiConsole.Ask<string>(prompt);
    }

    private static int GetIntInputWithBounds(string prompt, int minValue = 1, int maxValue = 10)
    {
        while (true)
        {
            var value = AnsiConsole.Ask<int>(prompt);

            if (value >= minValue && value <= maxValue)
            {
                return value; 
            }
            
            AnsiConsole.MarkupLine($"[red]Ошибка: значение должно быть от {minValue} до {maxValue}. Пожалуйста, попробуйте снова.[/]");
        }
    }

    private static int GetIntInput(string prompt)
    {
        return AnsiConsole.Ask<int>(prompt);
    }
    
    public static (string name, int food, int number, int kindnessLevel) GetAnimalInput()
    {
        string name = GetStringInput("Введите имя животного: ");
        int food = GetIntInput("Введите количество еды (кг/день): ");
        int number = GetIntInput("Введите номер: ");
        int kindnessLevel = GetIntInputWithBounds("Введите уровень доброты (от 1 до 10): ");

        return (name, food, number, kindnessLevel);
    }
    
    public static (string name, int number) GetThingInput()
    {
        string name = GetStringInput("Введите название вещи: ");
        int number = GetIntInput("Введите номер: ");

        return (name, number);
    }
}