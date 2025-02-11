using mini_hw_1.Domain.Interfaces;

namespace mini_hw_1.Domain.Entities;

public class Thing : IInventory
{
    public int Number { get; set; }
    
    public string Name { get; set; }

    public Thing(int number, string name)
    {
        Number = number;
        Name = name;
    }
}