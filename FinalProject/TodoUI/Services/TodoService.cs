using System.Net.Http.Json;
using System.Text.Json.Serialization;

namespace TodoUI.Services;

public class TodoService
{
    private readonly HttpClient httpClient;

    public TodoService(IHttpClientFactory httpClientFactory)
    {
        this.httpClient = httpClientFactory.CreateClient("TodoApi");
    }

    public async Task<Todo[]> GetAllTodos()
    {
        return await httpClient.GetFromJsonAsync<Todo[]>(string.Empty);
    }
}

public record Todo(Guid Id, string Task, bool Completed)
{
    [JsonPropertyName("_id")]
    public Guid Id { get; } = Id;
}