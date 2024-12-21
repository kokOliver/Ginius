using Application.Interfaces;

namespace Tests.Unit.Services;

public class GinServiceTests(IGinService ginService)
{
    [Fact]
    public void GetGin_ReturnsGin()
    {
        var result = ginService.GetGin();
        
        Assert.Equal("Gin", result);
    }
}