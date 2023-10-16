using System.Net;
using Microsoft.VisualBasic;
using NotesTests.Config;
using NotesTests.Tests;
using NotesTests;
using Xunit;
using Conversion = NotesApi.Helpers.Conversion;
using NotesTypes.Models;

namespace NotesTests.Tests.NotesControllerTests;

public class GetAll : BaseTest
{
    public GetAll(CustomWebApplicationFactory<Program> factory) : base(factory)
    {
    }

    public override async Task InitializeAsync()
    {
        await base.InitializeAsync();
        await Fixtures.NotesSeed();
    }

    [Fact]
    public async Task Ok()
    {
        var response = await Fixtures.HttpClient.GetAsync(Endpoints.Notes);
        var data = await Conversion.GetData<List<Note>>(response);

        Assert.Equal(HttpStatusCode.OK, response.StatusCode);
        Assert.Equal(Fixtures.Notes.Count, data?.Count);
    }
}