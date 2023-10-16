# TypeScript exercise
## GitHub Action instructions
Have following jobs in your GitHub Action:
1. Build using `yarn build` command.
2. Lint using `yarn lint` command.
3. Test using `yarn db:test` and `yarn test:cov` commands. Also make sure to start a postgres service.

## Code change
TODO\
Change [this line](https://github.com/XDoubleU/git-hub-workshop/blob/5ffb01eaa2b94ab1524b2a189632106295de9a03/exercises/TypeScript/apps/api/tests/notes.test.ts#L100).
Afterwards make a PR to merge this change on your fork. The `test` check should fail. Afterwards fix this by pushing a new commit to the same branch.

## Merge open change
Try merging the `change/typescript` branch. This should cause a merge conflict and should be a broken change causing the `test` check to fail.
