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

package main

import (
	"github.com/Bollos00/passwgen-go/src/passwgen"
	flag "github.com/ogier/pflag"
)

func initFlags() passwgen.PasswGen {

	// Create the flags of the program
	var characters *string = flag.StringP("characters", "c", "1aA%",
		"Allowed characters type on the password")
	var min *int = flag.IntP("min", "m", 8, "Minimum size of the password")
	var max *int = flag.IntP("max", "M", 16, "Maximum size of the password")
	var length *int = flag.IntP("length", "l", -1,
		"Fixed length of the password (overwrites 'min' and 'max' flags)")
	var ignoredChars *string = flag.StringP("ignored", "i", "",
		"Characters ignored in the creation of the password")
	var copyClipboard *bool = flag.BoolP("clipboard", "C", false,
		"Copy the generated password to the clipboard instead of printing it")

	flag.Parse()

	pssgenType := passwgen.PasswGenCharType(passwgen.None)

	// Detect the allowed characters type of the password
	for _, v := range []byte(*characters) {

		if passwgen.FindItemInSlice(passwgen.NumeralChars(), v) != -1 {
			pssgenType = pssgenType | passwgen.Numeral
		} else if passwgen.FindItemInSlice(passwgen.LowerLetterChars(), v) != -1 {
			pssgenType = pssgenType | passwgen.LowerLetter
		} else if passwgen.FindItemInSlice(passwgen.UpperLetterChars(), v) != -1 {
			pssgenType = pssgenType | passwgen.UpperLetter
		} else if passwgen.FindItemInSlice(passwgen.SpecialChars(), v) != -1 {
			pssgenType = pssgenType | passwgen.Special
		}
	}

	if *min > *max {
		*max = *min
	}

	if *length >= 0 {
		*max = *length
		*min = *length
	}

	passgen_instance := passwgen.PasswGen{
		CharactersType: pssgenType,
		MinSize:        *min,
		MaxSize:        *max,
		IgnoredChars:   []byte(*ignoredChars),
		CopyClipboard:  *copyClipboard,
	}

	return passgen_instance
}

func main() {

	passgen_instace := initFlags()

	passgen_instace.Generate()
	// password := passgen_instace.Generate()
	// fmt.Printf("Generated password:\n%v\n", password)
}
