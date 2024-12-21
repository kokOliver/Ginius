

using Application.Services;

namespace Tests.Unit.Services;

public class GinServiceTests
{
    private readonly GinService _uut = new();

    [Fact]
    public void GetGin_ReturnsGin()
    {
        var result = _uut.GetGin();
        
        Assert.Equal("Gin", result);
    }
}