using Microsoft.Extensions.DependencyInjection;
using mini_hw_1.Application.Services;
using mini_hw_1.Domain.Entities;
using mini_hw_1.Infrastructure.DI;
using mini_hw_1.Infrastructure.UI;
using Spectre.Console;

class Program
{
    static void Main()
    {
        var services = DiConfiguration.ConfigureServices();
        var zooServices = services.GetRequiredService<Zoo>();

        while (true)
        {
            var choice = MenuWork.ShowMenu("Главное меню:", "Добавить животное", "Добавить вещь", "Показать отчет", "Показать всю инвентаризацию", "Выход");

            switch (choice)
            {
                case "Добавить животное":
                    var animalData = UserInputHandler.GetAnimalInput();
                    var animalType = MenuWork.ShowAnimalTypeMenu();

                    Animal newAnimal = animalType switch
                    {
                        "Обезьяна" => new Monkey(animalData.name, animalData.food, animalData.number,
                            animalData.kindnessLevel),
                        "Кролик" => new Rabbit(animalData.name, animalData.food, animalData.number,
                            animalData.kindnessLevel),
                        "Тигр" => new Tiger(animalData.name, animalData.food, animalData.number),
                        "Волк" => new Wolf(animalData.name, animalData.food, animalData.number),
                        _ => throw new ArgumentOutOfRangeException()
                    };
                    
                    if (zooServices.AddAnimal(newAnimal))
                    {
                        AnsiConsole.MarkupLine("[green]Животное успешно добавлено![/]");
                    }
                    else
                    {
                        AnsiConsole.MarkupLine("[red]Животное не прошло проверку здоровья и не добавлено.[/]");
                    }
                    break;
                
                case "Добавить вещь":
                    var thingData = UserInputHandler.GetThingInput();
                    var newThing = new Thing(thingData.number, thingData.name);
                    zooServices.AddThing(newThing);
                    AnsiConsole.MarkupLine("[green]Вещь успешно добавлена![/]");
                    break;
                
                case "Показать отчет":
                    AnsiConsole.MarkupLine($"[bold]Всего еды: {zooServices.CalculateTotalFood()} кг/день[/]");
                    
                    AnsiConsole.MarkupLine("[bold]Животные для контактного зоопарка:[/]");
                    foreach (var animal in zooServices.GetContactZooAnimals())
                    {
                        AnsiConsole.MarkupLine($"[cyan]{animal.Name} (№{animal.Number})[/]");
                    }
                    break;
                
                case "Показать всю инвентаризацию":
                    zooServices.PrintInventory();
                    break;
                
                case "Выход":
                    Environment.Exit(0);
                    break;
            }
            
            AnsiConsole.WriteLine("Нажмите любую клавишу для продолжения...");
            Console.ReadKey();
            Console.Clear();
        }
    }
}