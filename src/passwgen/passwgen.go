//
// MIT License
//
// Copyright (c) 2020-2021  Bruno Bollos Correa
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

package passwgen

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/atotto/clipboard"
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

func RemoveInSlice(slice []byte, index int) []byte {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func Unused(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

func FindItemInSlice(slice []byte, val byte) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}

	return -1
}

type PasswGen struct {
	CharactersType PasswGenCharType
	MinSize        int
	MaxSize        int
	IgnoredChars   []byte
	CopyClipboard  bool
}

func (psg PasswGen) Validate() PasswGen {
	if psg.MinSize <= 0 || psg.MaxSize <= 0 {
		fmt.Printf("The size of the password should be a positive integer.\n")
		fmt.Printf("Changing the password length to a random number between 8 and 16.\n\n")
		psg.MinSize = 8
		psg.MaxSize = 16
	}
	if psg.CharactersType == None {
		fmt.Printf("No valid characters to the password found. Changing to default.\n\n")
		psg.CharactersType = Numeral | UpperLetter | LowerLetter | Special
	}

	return psg
}

func (psg PasswGen) Generate() string {

	availableChars := []byte{}

	psg = psg.Validate()

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

	for _, v := range psg.IgnoredChars {
		index := FindItemInSlice(availableChars, v)
		if index != -1 {
			availableChars = RemoveInSlice(availableChars, index)
		}
	}

	fmt.Printf("Available characters: %c\n\n", availableChars)

	rand.Seed(time.Now().UnixNano())

	password_size := psg.MinSize + rand.Intn(psg.MaxSize-psg.MinSize+1)
	fmt.Printf("Password size: %v\n\n", password_size)

	password := make([]byte, password_size)

	for i := range password {
		password[i] = availableChars[rand.Intn(int(len(availableChars)))]
	}

	if psg.CopyClipboard {
		clipboard.WriteAll(string(password))
	} else {
		fmt.Printf("Generated password:\n%v\n", string(password))
	}

	return string(password)
}
