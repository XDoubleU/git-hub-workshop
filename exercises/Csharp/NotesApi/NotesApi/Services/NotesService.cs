using Microsoft.EntityFrameworkCore;
using NotesApi.Context;
using NotesTypes.DTOs.Notes;
using NotesTypes.Models;

namespace NotesApi.Services;

// ReSharper disable once UnusedMember.Global
public class NotesService : INotesService
{
    private readonly NotesContext _context;

    public NotesService(NotesContext context)
    {
        _context = context;
    }

    public async Task<List<Note>> GetAll()
    {
        return await _context.Notes.ToListAsync();
    }

    public async Task<Note?> GetById(string id)
    {
        return await _context.Notes.SingleOrDefaultAsync(u => u.Id == id);
    }

    public async Task<Note?> Create(CreateNoteDto createNoteDto)
    {
        var note = new Note(createNoteDto.Title, createNoteDto.Contents);
        _context.Notes.Add(note);
        await _context.SaveChangesAsync();

        return note;
    }

    public async Task<Note> Update(Note note, UpdateNoteDto updateNoteDto)
    {
        note.Title = updateNoteDto.Title ?? note.Title;
        note.Contents = updateNoteDto.Contents ?? note.Contents;

        await _context.SaveChangesAsync();

        return note;
    }

    public async Task<Note?> Delete(Note note)
    {
        _context.Notes.Remove(note);
        await _context.SaveChangesAsync();

        return note;
    }
}

public interface INotesService
{
    Task<List<Note>> GetAll();
    Task<Note?> GetById(string id);
    Task<Note?> Create(CreateNoteDto createNoteDto);
    Task<Note> Update(Note note, UpdateNoteDto updateNoteDto);
    Task<Note?> Delete(Note note);
}