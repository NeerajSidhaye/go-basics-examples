package mymap

import (
	"encoding/json"
	"fmt"
)

/*
   Assignment:

   Write a program which prompts the user to first enter a name, and then enter an address. 
   Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. 
   Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

*/

//ConvertMapToJSON : converting map to JSON
func ConvertMapToJSON() {

	var firstName string
	var address string

	fmt.Println("Enter Name:-> ")

	fmt.Scan(&firstName)  // reading from commandline. Entered value will be assigned to firstName string

	fmt.Println("Enter address:-> ")

	fmt.Scan(&address)  // reading address from command line

	var userAttribute = map[string]string{"fname": firstName, "address" : address}  // dclaring and initializing a map
	
	userAttributeJSON , err := json.Marshal(userAttribute) // marshalling map - converting to byte array or byte slice
	
	if err!=nil{
		fmt.Println("error whie marshalling userAttribute map")
	}

	fmt.Println("UserAttribute JSON-> ", string(userAttributeJSON)) // converting byte array to string and printing the JSON

}