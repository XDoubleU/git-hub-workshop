using Microsoft.AspNetCore.Mvc;

namespace NotesApi.Controllers;

[ApiController]
[Consumes("application/json")]
[Produces("application/json")]
[Route("[controller]")]
public abstract class BaseController : Controller
{
}