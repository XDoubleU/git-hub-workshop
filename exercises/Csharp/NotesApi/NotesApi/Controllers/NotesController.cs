using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using NotesApi.Services;
using NotesTypes.DTOs.Notes;

namespace NotesApi.Controllers;

public class NotesController : BaseController
{
    private readonly INotesService _notesService;

    public NotesController(INotesService notesService)
    {
        _notesService = notesService;
    }

    [HttpGet("")]
    public async Task<IActionResult> GetAll()
    {
        var notes = await _notesService.GetAll();
        return Ok(notes);
    }

    [HttpPost("")]
    public async Task<IActionResult> Create(CreateNoteDto createNoteDto)
    {
        var note = await _notesService.Create(createNoteDto);
        return Ok(note);
    }

    [HttpPatch("{id}")]
    public async Task<IActionResult> Update(string id, UpdateNoteDto updateNoteDto)
    {
        var existingNote = await _notesService.GetById(id);
        if (existingNote is null) return NotFound("Note doesn't exist");

        var note = await _notesService.Update(
            existingNote,
            updateNoteDto
        );

        return Ok(note);
    }

    [HttpDelete("{id}")]
    public async Task<IActionResult> Delete(string id)
    {
        var note = await _notesService.GetById(id);
        if (note is null) return NotFound("Note doesn't exist");
        return Ok(await _notesService.Delete(note));
    }
}