# [passwgen-go](https://github.com/Bollos00/passwgen-go)

```
$ ./passwgen-go --help
Usage of ./passwgen-go:
  -c, --characters="1aA%": Allowed characters type on the password
  -i, --ignored="": Characters ignored in the creation of the password
  -l, --length=-1: Fixed length of the password (overwrites 'min' and 'max' flags)
  -M, --max=16: Maximum size of the password
  -m, --min=8: Minimum size of the password
```

This is a simple CLI random password generator written in golang. It generates the password given the type of characters and the intended password size.

# Installation

## Download executable

Binary releases are available in [here](https://github.com/Bollos00/passwgen-go/releases) for different platforms, so you can just download and run it.

## Building from source code

It is also possible to build it from the source code if you have [go](https://go.dev/dl/) and git installed on your system. To do so, follow those steps:

```
$ git clone https://github.com/Bollos00/passwgen-go.git
$ cd passwgen-go
```

To install:
```
$ go install .
```

To just build it:
```
$ go build .
```

To run it:
```
$ go run ./main.go
```

# Usage

## Types of characters

There are four types of characters shown in the table bellow.

| Type         | Available Characters |
|--------------|----------------------|
| Numeral      |   `'0', '1', '2', '3', '4','5', '6', '7', '8', '9'` |
| Lower Letter | `'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'`|
| Upper Letter | `'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'`|
| Special      | <code>'!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[', '\\', ']', '^', '_', '\`', '{', '\|', '}', '~', ' '</code> |

By default it will use all of them, but you can change it passing the `-c` a and `-i` parameters to the program. Examples:

* Generates a password with only numerical characters
```
$ passwgen-go -c "0"
```

* Generates a password with only lower and upper letter characters
```
$ passwgen-go -c "aA"
```

* Generates a password with numeral, lower and upper letter characters but ignoring 01oOil
```
$ passwgen-go -c "gH5" -i "01oOil"
```

* Generates a password only with special characters
```
$ passwgen-go -c "@"
```

## Password size
It is also possible to set the desired length of the password, being it a fixed value or a number between a range with the `-m`, `-M` and `-l` flags. By default it generates a password with a length between 8 nad 16. Examples:

* Generates a password with length between 20 and 30
```
$ passwgen-go -m 20 -M 30
```

* Generates a password with fixed length of 50
```
$ passwgen-go -l 50
```

* Generates a password with fixed length of 6 with only numbers
```
$ passwgen-go -l 6 -c "1"
```