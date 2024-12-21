using Application.Interfaces;

namespace Application.Services;

public class GinService : IGinService
{
    public string GetGin()
    {
        return "Gin";
    }
}