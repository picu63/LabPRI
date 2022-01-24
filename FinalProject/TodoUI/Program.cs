using Microsoft.AspNetCore.Components.Web;
using Microsoft.AspNetCore.Components.WebAssembly.Hosting;
using MudBlazor.Services;
using Serilog;
using TodoUI;
using TodoUI.Services;

var builder = WebAssemblyHostBuilder.CreateDefault(args);
builder.Logging.AddSerilog(new LoggerConfiguration()
    .WriteTo.Console()
    .WriteTo.Seq("http://localhost:5341/")
    .CreateLogger());
builder.RootComponents.Add<App>("#app");
builder.RootComponents.Add<HeadOutlet>("head::after");
builder.Services.AddHttpClient("TodoApi",
    client => { 
        client.BaseAddress = new Uri(builder.Configuration.GetValue<string>("TodoApi:BaseUrl"));
    });
builder.Services.AddScoped<TodoService>();
builder.Services.AddMudServices();
await builder.Build().RunAsync();
