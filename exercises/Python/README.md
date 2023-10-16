# Python exercise
## GitHub Action instructions
Have following jobs in your GitHub Action:
1. Lint using `isort` and `black`.
2. Test using `pytest` command. Also make sure to start a postgres service.

## Code change
Change [this line](https://github.com/XDoubleU/git-hub-workshop/blob/5579769b332191ca9fc1ad66d3eaad175fa5eae0/exercises/Python/note/tests.py#L40).
Afterwards make a PR to merge this change on your fork. The `test` check should fail. Afterwards fix this by pushing a new commit to the same branch.

## Merge open change
Try merging the `change/python` branch. This should cause a merge conflict and should be a broken change causing the `test` check to fail.
