using mini_hw_1.Domain.Entities;
using mini_hw_1.Domain.Interfaces;

namespace mini_hw_1.Application.Services;

public class Zoo
{
    private readonly IVetClinic _vetClinic;
    private readonly List<Animal> _animals = new();
    private readonly List<Thing> _things = new();
    
    public Zoo(IVetClinic vetClinic)
    {
        _vetClinic = vetClinic;
    }
    
    public bool AddAnimal(Animal animal)
    {
        if (_vetClinic.CheckHealth(animal))
        {
            _animals.Add(animal);
            return true;
        }

        return false;
    }
    
    public double CalculateTotalFood() => 
        _animals.Sum(a => a.Food);
    
    public int GetCountAnimals() =>
        _animals.Count;
    
    public IEnumerable<Animal> GetContactZooAnimals() => 
        _animals.OfType<Herbo>().Where(h => h.KindnessLevel > 5);
    
    public void AddThing(Thing thing) => 
        _things.Add(thing);
    
    public void PrintInventory()
    {
        Console.WriteLine("Животные:");
        foreach (var animal in _animals)
        {
            Console.WriteLine($"{animal.Name} (№{animal.Number}), Еда: {animal.Food} кг/день, Здоров: {animal.IsHealthy}");
        }

        Console.WriteLine("Вещи:");
        foreach (var thing in _things)
        {
            Console.WriteLine($"{thing.Name} (№{thing.Number})");
        }
    }
}