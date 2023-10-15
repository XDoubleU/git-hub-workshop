namespace NotesTypes.DTOs.Notes
{
    public class UpdateNoteDto
    {
        public string? Title { get; set; }
        public string? Contents { get; set; }

        public UpdateNoteDto(string? title, string? contents)
        {
            Title = title;
            Contents = contents;
        }
    }
}
