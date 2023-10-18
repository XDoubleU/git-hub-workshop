# stash local changes
git add . &> /dev/null
git stash &> /dev/null

branches=("csharp" "go" "python" "typescript")
for branch in "${branches[@]}"
do
  printf "checking if local branch change/$branch exists\n"
  if [[ $(git show-ref refs/heads/change/$branch) ]]; then
    printf "deleting local branch change/$branch\n"
    git branch -D change/$branch &> /dev/null
  fi

  printf "checking if remote branch change/$branch exists\n"
  if [[ $(git ls-remote --heads origin refs/heads/change/$branch) ]]; then
    printf "deleting remote branch change/$branch\n"
    git push origin -d change/$branch &> /dev/null
  fi

  printf "create change/$branch branch\n"
  git checkout -b change/$branch &> /dev/null

  case $branch in
    csharp)
      sed -i "s/Assert.Equal(\"Title\"/Assert.Equal(\"XDoubleU was here\"/" exercises/Csharp/NotesApi/NotesTests/Tests/NotesControllerTests/Create.cs
      ;;
    
    go)
      sed -i "s/\"NewNote\")/\"XDoubleU was here\")/" exercises/Go/cmd/api/notes_test.go
      ;;

    python)
      sed -i "s/== \"Title0\"/== \"NewTitle0\"/" exercises/Python/note/tests.py
      ;;
    
    typescript)
      sed -i "s/toBe(\"NewNote\")/toBe(\"XDoubleU was here\")/" exercises/TypeScript/apps/api/tests/notes.test.ts
      ;;
  esac
  
  git add . &> /dev/null
  git commit -m "Open change" &> /dev/null
  git push -u origin change/$branch &> /dev/null
  printf "\n"
done

git checkout main &> /dev/null

# apply stashed local changes back
git stash pop &> /dev/null
git reset &> /dev/null