package passwgen

import (
	"fmt"
	"math/rand"
	"time"
)

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

func SpecialChars() []byte {
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

func (psg PasswGen) Generate() string {

	availableChars := []byte{}

	if psg.CharactersType&Numeral != 0 {
		availableChars = append(availableChars, NumeralChars()...)
	}
	if psg.CharactersType&UpperLetter != 0 {
		availableChars = append(availableChars, LowerLetterChars()...)
	}
	if psg.CharactersType&LowerLetter != 0 {
		availableChars = append(availableChars, UpperLetterChars()...)
	}
	if psg.CharactersType&Special != 0 {
		availableChars = append(availableChars, SpecialChars()...)
	}

	fmt.Printf("Available characters: %c\n\n", availableChars)

	rand.Seed(time.Now().UnixNano())

	password_size := psg.MinSize + rand.Intn(psg.MaxSize-psg.MinSize+1)
	fmt.Printf("Password size: %v\n\n", password_size)

	password := make([]byte, password_size)

	for i := range password {
		password[i] = availableChars[rand.Intn(int(len(availableChars)))]
	}

	return string(password)
}
