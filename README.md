
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
- Go environment variables

```bash
go env
```

- GOROOT, GOPATH, GOBIN, GOOS, GOARCH 

    - if GOBIN is empty, it considers GOPATH/bin directory is the GOBIN directory


- Go install 
    -  if it is from internet, or from local it pulls the code , compilers it, build it and move it to GOBIN path

- Go supported target machines

```bash
go tool dist list 
```

GOOS/GOARCH
aix/ppc64
android/386
android/amd64
android/arm
android/arm64
darwin/amd64
darwin/arm64
dragonfly/amd64
freebsd/386
freebsd/amd64
freebsd/arm
freebsd/arm64
freebsd/riscv64
illumos/amd64
ios/amd64
ios/arm64
js/wasm
linux/386
linux/amd64
linux/arm
linux/arm64
linux/loong64
linux/mips
linux/mips64
linux/mips64le
linux/mipsle
linux/ppc64
linux/ppc64le
linux/riscv64
linux/s390x
netbsd/386
netbsd/amd64
netbsd/arm
netbsd/arm64
openbsd/386
openbsd/amd64
openbsd/arm
openbsd/arm64
openbsd/ppc64
openbsd/riscv64
plan9/386
plan9/amd64
plan9/arm
solaris/amd64
wasip1/wasm
windows/386
windows/amd64
windows/arm64