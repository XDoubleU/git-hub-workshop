namespace NotesTypes.DTOs.Notes
{
    public class CreateNoteDto
    {
        public string Title { get; set; }
        public string Contents { get; set; }

        public CreateNoteDto(string title, string contents)
        {
            Title = title;
            Contents = contents;
        }
    }
}
