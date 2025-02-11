using mini_hw_1.Domain.Entities;
using mini_hw_1.Domain.Interfaces;

namespace mini_hw_1.Application.Services;

public class VetClinic : IVetClinic
{
    private readonly Random _random = new();
    
    public bool CheckHealth(Animal animal)
    {
        animal.IsHealthy = _random.Next(5) != 0;
        return animal.IsHealthy;
    }
}