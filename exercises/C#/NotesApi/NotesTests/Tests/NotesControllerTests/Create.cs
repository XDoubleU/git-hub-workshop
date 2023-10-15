using System.Net;
using System.Net.Http.Json;
using Microsoft.VisualBasic;
using NotesTests.Config;
using NotesTests.Tests;
using NotesTests;
using Xunit;
using NotesTypes.DTOs.Notes;
using NotesTypes.Models;
using Conversion = NotesApi.Helpers.Conversion;

namespace NotesTests.Tests.NotesControllerTests;

public class Create : BaseTest
{
    public Create(CustomWebApplicationFactory<Program> factory) : base(factory)
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
        var createNoteDto = new CreateNoteDto("Title", "Some text");

        var response = await Fixtures.HttpClient.PostAsJsonAsync(Endpoints.Notes, createNoteDto);
        var data = await Conversion.GetData<Note>(response);

        Assert.Equal(HttpStatusCode.OK, response.StatusCode);
        Assert.Equal(createNoteDto.Title, data?.Title);
        Assert.Equal(createNoteDto.Contents, data?.Contents);
    }
}