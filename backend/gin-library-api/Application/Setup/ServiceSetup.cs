using Application.Interfaces;
using Application.Services;

namespace Application.Setup;

public static class ServiceSetup
{
    public static void ConfigureServices(this IServiceCollection services)
    {
        services.AddScoped<IGinService, GinService>();
    }
}