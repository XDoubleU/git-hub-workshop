using System.Net;
using Microsoft.VisualBasic;
using NotesTests.Config;
using NotesTests.Tests;
using NotesTests;
using Xunit;
using Conversion = NotesApi.Helpers.Conversion;
using NotesTypes.DTOs.Notes;
using System.Net.Http.Json;
using NotesTypes.Models;

namespace NotesTests.Tests.NotesControllerTests;

public class Update : BaseTest
{
    public Update(CustomWebApplicationFactory<Program> factory) : base(factory)
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

        var updateNoteDto = new UpdateNoteDto("New title", null);

        var response =
            await Fixtures.HttpClient.PatchAsJsonAsync($"{Endpoints.Notes}/{note.Id}", updateNoteDto);
        var data = await Conversion.GetData<Note>(response);

        Assert.Equal(HttpStatusCode.OK, response.StatusCode);
        Assert.Equal(updateNoteDto.Title, data?.Title);
        Assert.Equal(note.Contents, data?.Contents);
    }

    [Fact]
    public async Task NoteNotFound()
    {
        var updateNoteDto = new UpdateNoteDto("Test2", "Trainee2");

        var response = await Fixtures.HttpClient.PatchAsJsonAsync($"{Endpoints.Notes}/random", updateNoteDto);
        var data = await Conversion.GetData<string>(response);

        Assert.Equal(HttpStatusCode.NotFound, response.StatusCode);
        Assert.Equal("Note doesn't exist", data);
    }
}