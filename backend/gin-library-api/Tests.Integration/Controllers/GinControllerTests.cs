using Application;
using Microsoft.AspNetCore.Mvc.Testing;

namespace Tests.Integration.Controllers;

public class GinControllerTests(WebApplicationFactory<Program> factory) : IClassFixture<WebApplicationFactory<Program>>
{
    [Fact]
    public async Task Get_Scenario_ReturnsGin()
    {
        // Arrange
        var client = factory.CreateClient();
        
        // Act
        var response = await client.GetAsync("/api/gin");
        response.EnsureSuccessStatusCode();
        
        var responseString = await response.Content.ReadAsStringAsync();
        
        // Assert
        Assert.Equal("Gin", responseString);
     }
}