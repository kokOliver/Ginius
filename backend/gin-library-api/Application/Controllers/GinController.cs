using Application.Interfaces;
using Microsoft.AspNetCore.Mvc;

namespace Application.Controllers;

[ApiController]
[Route("api/[controller]")]
public class GinController(IGinService ginService, ILogger<GinController> logger) : ControllerBase
{
    [HttpGet]
    public IActionResult Get()
    {
        var gin = ginService.GetGin();
        logger.LogInformation("Received request for {Gin}", gin);
        return Ok(gin);
    }
}