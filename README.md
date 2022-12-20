# CLI

## Using Cobra
1. stringer

```
stringer version 0.0.1

# inspect a string for digits
go run main.go inspect A1B2C3 --digits
> 'A2B2C3' has 3 digits

go run main.go insp A1B2C3 -d
> 'A2B2C3' has 3 digits

# check command help
go run main.go inspect --help


```


2. todo

```
cobra-cli init --viper
cobra-cli add add ## command to add a new todo to todo list
cobra-cli add remove ## command to remove a todo from todo list
cobra-cli add update ## command to update description of a todo
cobra-cli add list ## command to list all todos. optionally, with deadline


touch $HOME/.todos.yaml ## This is for 1st run only

todo add my1stTask --description "Tide my room" --deadline "30-10-2022"
todo add my2ndTask --description "Read xyz book" --deadline "31-10-2022"



todo list


todo remove my2ndTask


todo update my1stTask --description "This is an updated description2"

```

### commands

```
go mod init github.com/Jeejo-Joy/stringer

which go ## /usr/local/go/bin/go

go build
go install

use sudo chmod -R 777 <PATH_TO>/cmd

```

### Env

```

export GOROOT=/usr/local/
export GOPATH=$GOROOT/src //your-go-workspace
export GOBIN=$GOROOT/bin //where go-generate-executable-binaries

PATH=$PATH:$GOPATH:$GOBIN


source ~/.zshrc

source ~/.bash_profile

```
```
GOENV="/Users/<USER>/Library/Application Support/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOINSECURE=""
GONOPROXY=""
GONOSUMDB=""
GOOS="darwin"
GOPATH="/Users/<USER>/src/unbias"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/go"
```


# Error notes

keep the go export path .bash_profile,
moving it into .zshrc takes the default go installation.


# Reference
1. https://www.thorsten-hans.com/lets-build-a-cli-in-go-with-cobra/#:~:text=%23%20create%20the%20project%20folder%20mkdir,as%20a%20dependency%20go%20get

2. https://dev.to/divrhino/adding-flags-to-a-command-line-tool-built-with-go-and-cobra-34f1

3. https://blog.devgenius.io/how-to-build-cli-tool-with-go-and-cobra-31d5761937ed

4. https://github.com/spf13/cobra/blob/main/user_guide.md

5. https://blog.devgenius.io/how-to-build-cli-tool-with-go-and-cobra-31d5761937ed

6.
