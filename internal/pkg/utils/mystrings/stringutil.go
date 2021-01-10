package mystrings

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
   Assignment:

   Write a program which prompts the user to enter a string.
   The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
   The program should print “Found!” if the entered string starts with the character ‘i’,
   ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise.
   The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.
*/

//FindSpecificCharInString : prints found! when a sourceString contains the characters ‘i’, ‘a’, and ‘n’, such as "iaaaan" is passed.
func FindSpecificCharInString() {

	var sourceString string
	fmt.Println("Please inupt string")
	//reading input from command promot. Input may contain spaces.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sourceString = scanner.Text()
		break
	}

	sourceString = strings.ToLower(sourceString)

	if strings.HasPrefix(sourceString, "i") &&
		strings.HasSuffix(sourceString, "n") &&
		strings.Contains(sourceString, "a") {

		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
