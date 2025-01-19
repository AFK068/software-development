namespace hw1;

public class FactoryAF
{
    public List<Customer> Customers { get; private set; }

    public List<Car> Cars { get; private set; }

    public FactoryAF(List<Customer> customers)
    {
        Customers = customers;
        Cars = [];
    }

    public void AddCar()
    {
        Cars.Add(new Car { Number = Cars.Count + 1 });
    }
    
    public void SaleCar()
    {
        foreach (var customer in Customers)
        {
            if (!Cars.Any())
            {
                break;
            }

            Car car = Cars.First();
            customer.Car = car;
            
            Cars.RemoveAt(0);
        }

        if (Cars.Any())
        {
            Cars.Clear();
        }
    }
}