
- Install Golang --> go.dev 

- Install VS Code --> https://code.visualstudio.com/download

- Install extensions in VS Code --> Go extension

- Create a directory --> mkdir golang-training
- File/Add Folder or Workspace from VS Code 

- To check go installation

```bash 
go version
```
go version go1.25.4 windows/amd64


https://github.com/JitenPalaparthi/golang-batch3

- To start a go project

```bash
mkdir 1-hello && cd 1-hello
go mod init demo
```

- A file called go.mod is created and then the module name demo is given in that file

- create a file main.go

```bash
go run main.go 
```
- To build -- compiles, assembles , links and generates the binary

```bash
go build main.go
```

```bash
go build -o demo main.go 
```

- Cross compilation

```bash
GOARCH=amd64 GOOS=windows go build -o demo.exe main.go
```