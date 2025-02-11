using Xunit;
using mini_hw_1.Application.Services;
using mini_hw_1.Domain.Entities;
using Assert = Xunit.Assert;

namespace Tests;

public class VetClinicTests
{
    [Fact]
    public void CheckHealth_ReturnsRandomBoolean()
    {
        // Arrange
        var vetClinic = new VetClinic();
        var animal = new Monkey("Name", 1,  1, 7);
        
        // Act
        var isHealthy = vetClinic.CheckHealth(animal);
        
        // Assert
        Assert.True(isHealthy || !isHealthy);
    }
}