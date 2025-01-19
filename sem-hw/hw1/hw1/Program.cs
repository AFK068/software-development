namespace hw1;

internal class Program
{
    static void Main(string[] args)
    {
        List<Customer> customers = new List<Customer>
        {
            new Customer {Name = "Customer1"},
            new Customer {Name = "Customer2"},
            new Customer {Name = "Customer3"}
        };
        
        var factory = new FactoryAF(customers);

        for (int i = 0; i < 5; i++)
        {
            factory.AddCar();
        }
            
        Console.WriteLine("Before sale");
        Console.WriteLine(string.Join(Environment.NewLine, factory.Cars));
        Console.WriteLine(string.Join(Environment.NewLine, factory.Customers));
        
        factory.SaleCar();

        Console.WriteLine("After sale");
        Console.WriteLine(string.Join(Environment.NewLine, factory.Cars));
        Console.WriteLine(string.Join(Environment.NewLine, factory.Customers));
    }
}