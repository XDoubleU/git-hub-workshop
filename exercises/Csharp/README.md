# C# exercise
## GitHub Action instructions
Have following jobs in your GitHub Action:
1. Build using `dotnet build NotesApi` command.
2. Lint using `dotnet format` action.
3. Test using `dotnet test NotesTest` command. Also make sure to start a postgres service.

Also make sure the `ASPNETCORE_ENVIRONMENT` environment variable is set to `CI`.

## Code change
Change [this line](https://github.com/XDoubleU/git-hub-workshop/blob/8e8768c1c2cbdaaf04be761a45d83ebc72869184/exercises/Csharp/NotesApi/NotesTests/Tests/NotesControllerTests/Create.cs#L35).
Afterwards make a PR to merge this change on your fork. The `test` check should fail. Afterwards fix this by pushing a new commit to the same branch.

## Merge open change
Try merging the `change/csharp` branch. This should cause a merge conflict and should be a broken change causing the `test` check to fail.

