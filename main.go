package main

import (
	"fmt"

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
	}

	return passgen_instance
}

func main() {

	passgen_instace := initFlags()

	password := passgen_instace.Generate()
	fmt.Printf("Generated password:\n%v\n", password)
}
