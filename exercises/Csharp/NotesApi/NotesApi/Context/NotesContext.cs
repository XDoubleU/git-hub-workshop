using Microsoft.EntityFrameworkCore;
using NotesTypes.Models;

#pragma warning disable CS8618

namespace NotesApi.Context;

public class NotesContext : DbContext
{
    public NotesContext(DbContextOptions<NotesContext> options)
        : base(options)
    {
    }

    public virtual DbSet<Note> Notes { get; set; }
}