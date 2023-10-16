using System.Net;
using Microsoft.VisualBasic;
using NotesTests.Config;
using NotesTests.Tests;
using NotesTests;
using Xunit;
using Conversion = NotesApi.Helpers.Conversion;
using NotesTypes.Models;

namespace NotesTests.Tests.NotesControllerTests;

public class Delete : BaseTest
{
    public Delete(CustomWebApplicationFactory<Program> factory) : base(factory)
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
        var note = Fixtures.Notes.First();

        var response = await Fixtures.HttpClient.DeleteAsync($"{Endpoints.Notes}/{note.Id}");
        var data = await Conversion.GetData<Note>(response);

        Assert.Equal(HttpStatusCode.OK, response.StatusCode);
        Assert.Equal(note.Id, data?.Id);
    }

    [Fact]
    public async Task NoteNotFound()
    {
        var response = await Fixtures.HttpClient.DeleteAsync($"{Endpoints.Notes}/random");
        var data = await Conversion.GetData<string>(response);

        Assert.Equal(HttpStatusCode.NotFound, response.StatusCode);
        Assert.Equal("Note doesn't exist", data);
    }
}