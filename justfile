push: commit
    git push

commit: update
    git add .

update:
    git submodule foreach git pull