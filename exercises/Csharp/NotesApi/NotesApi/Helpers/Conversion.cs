using Newtonsoft.Json;

namespace NotesApi.Helpers
{
    public static class Conversion
    {
        public static async Task<T?> GetData<T>(HttpResponseMessage response)
        {
            var responseContent = await response.Content.ReadAsStringAsync();
            return JsonConvert.DeserializeObject<T>(responseContent);
        }
    }
}
