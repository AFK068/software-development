using Microsoft.Extensions.DependencyInjection;
using mini_hw_1.Application.Services;
using mini_hw_1.Domain.Interfaces;

namespace mini_hw_1.Infrastructure.DI;

public static class DiConfiguration
{
    public static ServiceProvider ConfigureServices()
    {
        return new ServiceCollection()
            .AddSingleton<IVetClinic, VetClinic>()
            .AddSingleton<Zoo>()
            .BuildServiceProvider();
    }
}