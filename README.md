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

## Primitive Types

* Numeric types
  * Integers
    * Signed Integers
    * ```int``` type has varying size, but min 32 bits
    * ```int8``` through ```int64```
  * Unsigned integers
    * 8 bit (```byte``` or ```uint8```) through ```uint32```
  * Arithmetic operations
    * Addition, subtraction, multiplication, division, remainder
  * Bitwise operations
    * and (```&```) or (```|```), xor(```^```), and not (```&^```)
  * Zero value is 0
  * Can't mix types in same family (e.g. ```uint16``` + ```uint32``` = error)
  * Floating points numbers
    * Follows IEEE-754 standard
    * Zero value is 0
    * 32 and 64 bit versions
    * Literal styles
      * Decimal (3.14)
      * Exponential (13e18 or 2E10)
      * Mixed (13.7e12)
    * Arithmetic operations
      * Addition, subtraction, multiplication, division
  * Complex numbers
    * Zero value is (0+0i)
    * 64 and 128 bit versions
    * Built-in functions
      * ```complex``` - make complex number from two floats
      * ```real``` - get real part as float
      * ```imag``` - get imaginary part as float
    * Arithmetic operations
      * Addition, subtraction, multiplication, division
  * Text types
    * Strings
      * UTF-8
      * Immutable
      * Can be concatenated with plus (```+```) operator
      * Can be converted to []byte
    * Rune
      * UTF-32
      * Alias for ```int32```
      * Special methods normally required to process
        * e.g. ```strings.Reader#ReadRune```

