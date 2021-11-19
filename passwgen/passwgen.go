package passwgen

import "fmt"

type PasswGenCharType int

const (
	None        PasswGenCharType = 0
	Numeral     PasswGenCharType = 1
	UpperLetter PasswGenCharType = 2
	LowerLetter PasswGenCharType = 4
	Special     PasswGenCharType = 8
)

func NumeralChars() []byte {
	return []byte{
		'0', '1', '2', '3', '4',
		'5', '6', '7', '8', '9'}
}

func LowerLetterChars() []byte {
	return []byte{
		'a', 'b', 'c', 'd', 'e',
		'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o',
		'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y',
		'z'}
}

func UpperLetterChars() []byte {

	return []byte{
		'A', 'B', 'C', 'D', 'E',
		'F', 'G', 'H', 'I', 'J',
		'K', 'L', 'M', 'N', 'O',
		'P', 'Q', 'R', 'S', 'T',
		'U', 'V', 'W', 'X', 'Y',
		'Z'}
}

func SpecialLetterChars() []byte {
	return []byte{
		'!', '"', '#', '$', '%',
		'&', '\'', '(', ')', '*',
		'+', ',', '-', '.', '/',
		':', ';', '<', '=', '>',
		'?', '@', '[', '\\', ']',
		'^', '_', '`', '{', '|',
		'}', '~', ' '}
}

type PasswGen struct {
	CharactersType PasswGenCharType
	MinSize        int
	MaxSize        int
}

func main() {
	fmt.Printf("%v\n", SpecialLetterChars())
}
