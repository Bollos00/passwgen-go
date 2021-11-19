package main

import (
	"Bollos00/passwgen-go/passwgen"
	"fmt"

	flag "github.com/ogier/pflag"
)

func Unused(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

func FindItemInSlice(slice []byte, val byte) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func initFlags() passwgen.PasswGen {

	// Create the flags of the program
	var characters *string = flag.StringP("characters", "c", "1aA%", "Type of characters on the password")
	var min *int = flag.IntP("min", "m", 8, "Minimum size of the password")
	var max *int = flag.IntP("max", "M", 16, "Maximum size of the password")

	flag.Parse()

	pssgenType := passwgen.PasswGenCharType(0)

	for _, v := range []byte(*characters) {

		if FindItemInSlice(passwgen.NumeralChars(), v) {
			pssgenType = pssgenType | passwgen.Numeral
		} else if FindItemInSlice(passwgen.LowerLetterChars(), v) {
			pssgenType = pssgenType | passwgen.LowerLetter
		} else if FindItemInSlice(passwgen.UpperLetterChars(), v) {
			pssgenType = pssgenType | passwgen.UpperLetter
		} else if FindItemInSlice(passwgen.SpecialChars(), v) {
			pssgenType = pssgenType | passwgen.Special
		}
	}

	if *min > *max {
		*max = *min
	}

	passgen_instance := passwgen.PasswGen{
		CharactersType: pssgenType,
		MinSize:        *min,
		MaxSize:        *max,
	}

	return passgen_instance
}

func main() {

	passgen_instace := initFlags()

	password := passgen_instace.Generate()
	fmt.Printf("Generated password:\n%v\n", password)
}
