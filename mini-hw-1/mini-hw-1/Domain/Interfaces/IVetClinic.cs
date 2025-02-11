using mini_hw_1.Domain.Entities;

namespace mini_hw_1.Domain.Interfaces;

public interface IVetClinic
{
    bool CheckHealth(Animal animal); 
}