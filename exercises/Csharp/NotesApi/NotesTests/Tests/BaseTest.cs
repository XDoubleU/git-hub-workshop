using NotesTests.Config;
using Xunit;

namespace NotesTests.Tests
{
    public abstract class BaseTest : IClassFixture<CustomWebApplicationFactory<Program>>, IAsyncLifetime
    {
        protected readonly CustomWebApplicationFactory<Program> Factory;


        protected BaseTest(CustomWebApplicationFactory<Program> factory)
        {
            Factory = factory;
        }

        protected Fixtures Fixtures { get; private set; } = null!;

        public virtual async Task InitializeAsync()
        {
            await Task.Run(() => Fixtures = new Fixtures(Factory));
        }

        public Task DisposeAsync()
        {
            return Task.CompletedTask;
        }
    }
}
