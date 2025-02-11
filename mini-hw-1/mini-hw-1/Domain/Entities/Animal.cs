using mini_hw_1.Domain.Interfaces;

namespace mini_hw_1.Domain.Entities;

public abstract class Animal : IInventory, IAlive
{
    public int Number { get; set; }
    
    public string Name { get; set; }
    
    public int Food { get; set; }
    
    public bool IsHealthy { get; set; }

    protected Animal(string name, int food, int number)
    {
        Name = name;
        Food = food;
        Number = number;
    }
}