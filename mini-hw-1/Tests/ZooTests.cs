using mini_hw_1.Application.Services;
using mini_hw_1.Domain.Entities;
using mini_hw_1.Domain.Interfaces;
using Xunit;
using Moq;
using Assert = Xunit.Assert;

namespace Tests;

public class ZooTests
{
    private readonly Mock<IVetClinic> _vetClinicMock;
    private readonly Zoo _zooTests;
    
    public ZooTests()
    {
        _vetClinicMock = new Mock<IVetClinic>();
        _zooTests = new Zoo(_vetClinicMock.Object);
    }
    
    [Fact]
    public void AddAnimal_HealthyAnimal_ReturnsTrue()
    {
        // Arrange
        var animal = new Monkey("Name", 1, 1, 7);
        _vetClinicMock.Setup(v => v.CheckHealth(animal)).Returns(true);
        
        // Act
        var result = _zooTests.AddAnimal(animal);
        
        // Assert
        Assert.True(result);
    }
    
    [Fact]
    public void AddAnimal_UnhealthyAnimal_ReturnsFalse()
    {
        // Arrange
        var animal = new Monkey("Name", 1, 1, 7);
        _vetClinicMock.Setup(v => v.CheckHealth(animal)).Returns(false);
        
        // Act
        var result = _zooTests.AddAnimal(animal);
        
        // Assert
        Assert.False(result);
    }
    
    [Fact]
    public void CalculateTotalFood_ReturnsCorrectSum()
    {
        // Arrange
        var monkey = new Monkey("Name", 2, 1, 7);
        var tiger = new Tiger("Name", 8, 2);
        
        _vetClinicMock.Setup(v => v.CheckHealth(It.IsAny<Animal>())).Returns(true);
        
        _zooTests.AddAnimal(monkey);
        _zooTests.AddAnimal(tiger);

        // Act
        var totalFood = _zooTests.CalculateTotalFood();

        // Assert
        Assert.Equal(10.0, totalFood);
    }
    
    [Fact]
    public void GetContactZooAnimals_ReturnsOnlyHerbivores()
    {
        // Arrange
        var monkey = new Monkey("Name", 1, 1, 7);
        var rabbit = new Rabbit("Name", 1, 2, 6);
        var tiger = new Tiger("Name",1, 3);
        
        _vetClinicMock.Setup(v => v.CheckHealth(It.IsAny<Animal>())).Returns(true);
        
        _zooTests.AddAnimal(monkey);
        _zooTests.AddAnimal(rabbit);
        _zooTests.AddAnimal(tiger);

        // Act
        var contactAnimals = _zooTests.GetContactZooAnimals();

        // Assert
        Assert.Equal(2, contactAnimals.Count());
    }
}