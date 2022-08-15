# fibonacci-go

## Installation

Either download pre-compiled binary for your operating system from
[GitHub releases](https://github.com/vaeu/fibonacci-go/releases) or
install the program using the `go` tool:

```
$ go install github.com/vaeu/fibonacci-go/cmd/...@latest
```

## Usage

```
$ fibonacci [-min int] [-max int]
```

## Examples

Print Nth number from 0 through 10.

```
$ fibonacci
```

Print Nth number from 15 through 18.

```
$ fibonacci -min 15 -max 18
```

Print Nth number of 26.

```
$ fibonacci -min 26 -max 26
```
