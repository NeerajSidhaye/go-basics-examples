package mystrings

import (
	"fmt"
	"strings"
)

//FindSpecificCharInString : prints found! when a sourceString contains the characters ‘i’, ‘a’, and ‘n’, such as "iaaaan" is passed.
func FindSpecificCharInString(sourceString string) {

	fmt.Println("looking for char i,a,n in the source string->", sourceString)

	sourceString = strings.ToLower(sourceString)

	if strings.HasPrefix(sourceString, "i") && strings.HasSuffix(sourceString, "n") &&
		strings.Contains(sourceString, "a") {
			
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
