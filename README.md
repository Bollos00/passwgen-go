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

This is a simple CLI random password generator written in go. It generates passwords given the type of characters and the intended password size.

## Types of characters

There are four types of characters shown in the table bellow.

| Type         | Available Characters |
|--------------|----------------------|
| Numeral      |   `'0', '1', '2', '3', '4','5', '6', '7', '8', '9'` |
| Lower Letter | `'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'`|
| Upper Letter | `'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'`|
| Special      | `` '!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[', '\\', ']', '^', '_', '`', '{', '|', '}', '~', ' ' `` |

