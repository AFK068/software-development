namespace hw1;

public class Car
{
    public Engine Engine { get; set; }
    
    public required int Number { get; set; }
    
    private readonly Random _random = new Random();

    public Car()
    {
        Engine = new Engine {Size = _random.Next(1, 10)};
    }

    public override string ToString()
    {
        return $"Номер автомобиля: {Number}\nРазмер педалей: {Engine.Size}";
    }
}