using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc.Testing;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Storage;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.DependencyInjection.Extensions;
using NetCore.AutoRegisterDi;
using NotesApi.Context;

namespace NotesTests.Config;

[RegisterAsTransient]
public sealed class CustomWebApplicationFactory<TProgram>
    : WebApplicationFactory<TProgram>, IDisposable where TProgram : class
{
    private readonly IDbContextTransaction _transaction;
    public readonly NotesContext Context;

    public CustomWebApplicationFactory()
    {
        var scope = Services.GetRequiredService<IServiceScopeFactory>().CreateScope();
        Context = scope.ServiceProvider.GetRequiredService<NotesContext>();

        Context.Database.EnsureCreated();

        _transaction = Context.Database.BeginTransaction();
    }

    public new void Dispose()
    {
        GC.SuppressFinalize(this);

        _transaction.Rollback();
        _transaction.Dispose();
    }

    protected override void ConfigureWebHost(IWebHostBuilder builder)
    {
        builder.ConfigureServices(services =>
        {
            services.RemoveAll<NotesContext>();
            services.RemoveAll<DbContextOptions>();

            foreach (var option in services.Where(s => s.ServiceType.BaseType == typeof(DbContextOptions)).ToList())
                services.Remove(option);

            var test0 = GetConfiguration();
            var test = GetConfiguration().GetConnectionString("Database");
            services.AddDbContext<NotesContext>(
                options => options.UseNpgsql(GetConfiguration().GetConnectionString("Database")),
                ServiceLifetime.Singleton
            );
        });

        base.ConfigureWebHost(builder);
    }

    public static IConfiguration GetConfiguration()
    {
        var environment = Environment.GetEnvironmentVariable("ASPNETCORE_ENVIRONMENT");

        var jsonFile = "appsettings.json";
        if (environment == "CI") jsonFile = "appsettings.CI.json";

        return new ConfigurationBuilder()
            .AddEnvironmentVariables()
            .AddJsonFile(jsonFile)
            .Build();
    }

    ~CustomWebApplicationFactory()
    {
        Dispose();
    }
}