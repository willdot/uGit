[![Go Report Card](https://goreportcard.com/badge/github.com/willdot/uGit)](https://goreportcard.com/report/github.com/willdot/uGit)


# uGit
A git command line tool written in Go. 

This is a command line app that makes running commonly used git commands easier and quicker.

## Installation

You will need to have installed GO on your machine and configured your paths correctly, for example:
```
export GOPATH="$HOME/Source/go"
export PATH=$PATH:$GOPATH/bin
```


Clicked [here](https://willdot.github.io/15SettingUpGo/) for a guide on how to do this.


Clone the repo and then run:

```
go install
```



## Usage


### Checkout

This will list all branches that you can checkout, and then you can select which one to checkout
```
ugit cko
```

### New branch
This will allow you to checkout a new branch
```
ugit cko-n
```

### Commit
This will display untracked files and allow you to add them to be tracked. It will then display files that have been changed, that you wish to commit and allow you to select them. It will then prompt you for a commit message
```
ugit com
```

### Commit and push
This will do the same as above, but will also push after
```
ugit com -p
```

### Delete
This will display all branches and then allow you to select which ones to delete. If git asked you if you want to delete it with -D (force) then it will prompt and ask you.
```
ugit d
```
