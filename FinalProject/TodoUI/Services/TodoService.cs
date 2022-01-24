﻿using System.Net.Http.Json;
using System.Text.Json;
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

    public async Task ChangeStatus(string todoId, bool isCompleted)
    {
        var requestUri = $"api/todos/{todoId}";
        var response = await httpClient.PutAsJsonAsync(requestUri, new {Completed = isCompleted });
        if (!response.IsSuccessStatusCode)
        {
            throw new Exception(response.StatusCode.ToString());
        }
    }

    public async Task Delete(string todoId)
    {
        var requestUri = $"api/todos/{todoId}";
        var response = await httpClient.DeleteAsync(requestUri);
        if (!response.IsSuccessStatusCode)
        {
            throw new Exception(response.StatusCode.ToString());
        }
    }

    public async Task<Todo> CreateNewByName(string taskName)
    {
        var requestUri = $"api/todos";
        var newTodo = new { Completed = false, Task = taskName};
        var response = await httpClient.PostAsJsonAsync(requestUri, newTodo);
        if (!response.IsSuccessStatusCode)
        {
            throw new Exception(response.StatusCode.ToString());
        }
        var id = (await response.Content.ReadFromJsonAsync<CreateResponse>()).Id;

        return new Todo(){Completed = newTodo.Completed, Task = newTodo.Task, Id = id};
    }

    record CreateResponse(string Id);
}

public class Todo
{

    [JsonPropertyName("_id")]
    public string Id { get; set; }
    public string Task { get; init; }
    public bool Completed { get; set; }

    public void Deconstruct(out string Id, out string Task, out bool Completed)
    {
        Id = this.Id;
        Task = this.Task;
        Completed = this.Completed;
    }
}