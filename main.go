package main

import (

	//"github.com/BeTheCodeWithYou/go-basics-examples/helloworld/mymath"

	"github.com/BeTheCodeWithYou/go-basics-examples/helloworld/mystrings"

	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hello, world!")

	//mymath.ScanFloatValueAndTruncateToInteger()

	fmt.Println("Please inupt string")
	//reading input from command promot. Input may contain spaces.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		mystrings.FindSpecificCharInString(scanner.Text())
		os.Exit(0)
	}
}
