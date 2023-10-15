# Go exercise
## GitHub Action instructions
Have following jobs in your GitHub Action:
1. Build using `make build` command.
2. Lint using `golangci/golangci-lint-action` action.
3. Test using `make init`, `make db/migrations/up`, `make test/cov/report` commands. Also make sure to start a postgres service.

## Code change
Change [this line](https://github.com/XDoubleU/git-hub-workshop/blob/33f7129ef49edcb1bcc8d30d08d33d088f69fd26/exercises/Go/cmd/api/notes_test.go#L105).
Afterwards make a PR to merge this change on your fork. The `test` check should fail. Afterwards fix this by pushing a new commit to the same branch.

## Merge open change
Try merging the `change/go` branch. This should cause a merge conflict and should be a broken change causing the `test` check to fail.
