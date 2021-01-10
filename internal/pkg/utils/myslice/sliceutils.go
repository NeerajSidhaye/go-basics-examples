package myslice

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
   Assignment:

   Write a program which prompts the user to enter integers and stores the integers in a sorted slice.

   The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3.
   During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
   The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
   The slice must grow in size to accommodate any number of integers which the user decides to enter.
   The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

*/

//SortSliceOfIntAndPrint : adding integers and sorting them into the slice until use enter character X
func SortSliceOfIntAndPrint() {

	intSlice := make([]int, 3) // initializing integer slice with length and capcity of 3

	inputScanner := bufio.NewScanner(os.Stdin)

	var userInputStr string

	fmt.Println("Enter a number OR type X to exit the program.")
	var index int
	for inputScanner.Scan() {

		userInputStr = inputScanner.Text()
		userInputInt, err := strconv.Atoi(userInputStr) // converting string to int

		if strings.EqualFold(userInputStr, "X") { // compares with case sensitivity
			fmt.Println("Input is X - existing the program.")
			break
		} else if err != nil { // err is taking care of any other string than X and not a valid number.
			fmt.Println("Input is not a valid number: ", err.Error())
			break
		} else {

			intSlice[index] = userInputInt // assigning user input to slice
			index++
			fmt.Println("index and len of slice->", index, len(intSlice))
			if len(intSlice) == index {
				intSlice = append(intSlice, index) // increasing slice lenght & capacity when user input reach to existing cap of slice.
				fmt.Println("slice appedned ->", intSlice, len(intSlice), cap(intSlice))
			}
			sort.Ints(intSlice) // sorting the slice
			fmt.Println("sorted slice is->", intSlice)
		}

		fmt.Println("enter a number OR type X to exit the program.")

	}

}
