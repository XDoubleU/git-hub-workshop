using System.Data;
using System.Net.Http.Headers;
using NotesTests.Config;
using NotesTypes.Models;

namespace NotesTests;

public class Fixtures
{
    public Fixtures(CustomWebApplicationFactory<Program> factory)
    {
        Factory = factory;
        HttpClient = factory.CreateClient();
    }

    protected CustomWebApplicationFactory<Program> Factory { get; }

    public HttpClient HttpClient { get; private set; }

    public List<Note> Notes { get; private set; } = new List<Note>();

    public async Task NotesSeed()
    {
        for(var i = 0; i < 10; i++)
        {
            var note = new Note($"title{i}", "some text");
            Factory.Context.Notes.Add(note);
            Notes.Add(note);
            await Factory.Context.SaveChangesAsync();
        }
    }
}