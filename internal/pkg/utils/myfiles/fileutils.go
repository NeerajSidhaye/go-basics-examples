package myfiles

import (
	"bufio"
	f "fmt"
	"os"
	"strings"
)

/*

Assignment

Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names. Each line of the text file
has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and
lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively
read each line of the text file and create a struct which contains the first and last names found
in the file. Each struct created will be added to a slice, and after all lines have been read from
the file, your program will have a slice containing one struct for each line in the file.

After reading all lines from the file, your program should iterate through your slice of structs and
print the first and last names found in each struct.

*/

type name struct{
    
	fname string
	lname string

}

//FileReaderAndPrintFileContentAsStructs :  reading names from file, and represet it in a slice of structs. Print the names from each struct.
func FileReaderAndPrintFileContentAsStructs() {
 
	var fileName string
	var filePath  = "../../internal/pkg/utils/myfiles/"

	f.Println("Please enter a file name")

	f.Scan(&fileName)

	if strings.EqualFold(fileName, "testfile.text") {
	  
		inputFile, err:= os.Open(filePath + fileName)  // opening the file for reading
		check(err)
		defer inputFile.Close()  // the file descriptor is closed at the end of the main function.

		fileScanner := bufio.NewScanner(inputFile) // creating a new scanner for reading/scanning file content
		
		nameSliceStruct := []name{}

		for fileScanner.Scan() {

			fullName := fileScanner.Text()
	
			nameArray := strings.Split(fullName, " ")
			
			nameStruct := name{nameArray[0], nameArray[1]}  // creating a new name struct with fname and lname
			

			nameSliceStruct = append(nameSliceStruct, nameStruct) // adding struct to the slice
		}
		
		for index, mystruct := range nameSliceStruct {   // printing struct from iterating over slice
			
			f.Println("Struct element from file at line->", index+1)
			f.Println("fName->", mystruct.fname)
			f.Println("lName->", mystruct.lname)

		}
		 
		f.Println("sliceStruct len and cap", len(nameSliceStruct), cap(nameSliceStruct))

	} else {
		f.Println("Invalid file name")
		os.Exit(0)
	}
	


}

func check(e error) {
	if e!=nil {
		f.Println(e)
	}
}
