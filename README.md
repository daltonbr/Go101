# Go 101

FreeCodeCamp.org
https://www.youtube.com/watch?v=YS4e4q9oBaU

Go (golang) - [golang.org](https://golang.org)

## Characteristics of Go

* Strong and statically typed
* Excellent community
* Key features
  * Simplicity
  * Fast compile times
  * Garbage collected (still you can manage your memory)
  * Built-in concurrency
  * Compile to standalone binaries

## Variables

* Variable declaration

```go
var foo int
var foo int = 42
foo := 42
```

* Can't redeclare variables,  but can shadow them
* All variables must be used
* Visibility
  * lower case first letter for package scope
  * upper case first letter to export
  * no private scope

* Naming conventions
  * HighCamelCase or lowerCamelCase
  * Capitalize acronyms (HTTP, URL)

* Type conversion
  * destinationType (variable)
  * use ```strconv``` package for strings
  * no implicit type conversion in Go

## Go Global Variables

Exporting to .bashrc, .zshrc, etc.

```shell
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin

# Create this folder
export GOPATH=$HOME/golib
export PATH=$PATH:$GOPATH/bin
```

### Go auto-completion

```
go get github.com/nsf/gocode
```
