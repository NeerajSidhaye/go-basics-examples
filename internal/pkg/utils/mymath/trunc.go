package mymath

import (
	"fmt"
	"math"
)

/*
	Assignment:
	Write a program which prompts the user to enter a floating point number and
	prints the integer which is a truncated version of the floating point number that was entered.

*/

//public function starts with Uppercase and should have comment on top of it.
// comment should start with function name itself, followed by colon and then description.

//ScanFloatValueAndTruncateToInteger : scanning user input in float and truncting float to int.
func ScanFloatValueAndTruncateToInteger() {

	var myFloat float64

	fmt.Println("Enter float number")

	fmt.Scan(&myFloat) // will prompt to enter input on the command line.

	myInt := math.Trunc(myFloat)
	fmt.Println("Int value after match.Trunc-> ", myInt)
}
