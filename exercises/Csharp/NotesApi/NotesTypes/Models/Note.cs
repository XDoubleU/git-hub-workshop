using System.ComponentModel.DataAnnotations.Schema;
using System.ComponentModel.DataAnnotations;

namespace NotesTypes.Models
{
    public class Note
    {
        [Key]
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        public string Id { get; set; } = null!;
        public string Title { get; set; }
        public string Contents { get; set; }

        public Note(string title, string contents)
        {
            Title = title;
            Contents = contents;
        }
    }
}
