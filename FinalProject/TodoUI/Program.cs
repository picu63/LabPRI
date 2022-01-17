using Microsoft.AspNetCore.Components.Web;
using Microsoft.AspNetCore.Components.WebAssembly.Hosting;
using TodoUI;
using TodoUI.Services;

var builder = WebAssemblyHostBuilder.CreateDefault(args);
builder.RootComponents.Add<App>("#app");
builder.RootComponents.Add<HeadOutlet>("head::after");
builder.Services.AddHttpClient("TodoApi",client =>
    client.BaseAddress = new Uri(builder.Configuration.GetValue<string>("TodoApi:BaseUrl")));
builder.Services.AddScoped<TodoService>();
await builder.Build().RunAsync();
