push: commit
    git push

commit: update
    git add .

update:
    git submodule foreach git pull

set positional-arguments

@add problemNames:
    echo git submodule add -b main https://github.com/SaffatHasan/$1.git $1