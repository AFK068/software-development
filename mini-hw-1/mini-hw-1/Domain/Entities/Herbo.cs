namespace mini_hw_1.Domain.Entities;

public class Herbo : Animal
{
    public int KindnessLevel { get; set; }

    protected Herbo(string name, int food, int number, int kindnessLevel) : base(name, food, number)
    {
        KindnessLevel = kindnessLevel;
    }
}