using System.Net.Http.Json;
using System.Text.Json.Serialization;

namespace TodoUI.Services;

public class TodoService
{
    private readonly ILogger<TodoService> logger;
    private readonly HttpClient httpClient;

    public TodoService(IHttpClientFactory httpClientFactory, ILogger<TodoService> logger)
    {
        this.logger = logger;
        this.httpClient = httpClientFactory.CreateClient("TodoApi");
    }

    public async Task<Todo[]> GetAllTodos()
    {
        var requestUri = "api/todos";
        logger.LogInformation($"Getting list of todos from {new Uri(httpClient.BaseAddress, requestUri)}");
        return await httpClient.GetFromJsonAsync<Todo[]>(requestUri);
    }
}

public class Todo
{

    [JsonPropertyName("_id")]
    public Guid Id { get; }
    public string Task { get; init; }
    public bool Completed { get; init; }

    public void Deconstruct(out Guid Id, out string Task, out bool Completed)
    {
        Id = this.Id;
        Task = this.Task;
        Completed = this.Completed;
    }
}