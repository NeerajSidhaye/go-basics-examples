package mymath

import ( 
	
	"fmt"
    "math"
	 
)

//public function starts with Uppercase and should have comment on top of it.
// comment should start with function name itself, followed by colon and then description.

//ScanFloatValueAndTruncateToInteger : scanning user input in float and truncting float to int.
func ScanFloatValueAndTruncateToInteger() {  
	
	var myFloat float64
		
	for i:=0; i<2; i++ {
	
		fmt.Println("Enter float number")
	
		fmt.Scan(&myFloat)  // will promt to enter input on the command line.
		
		fmt.Printf("user input float value-> %f\n", myFloat)
		
		myFloat = math.Trunc(myFloat)
		fmt.Printf("Int value after match.Trunc-> %.f\n ", myFloat)  
		// in fmt.Printf, %.f -  the dot(.) represent that we want no value after decimal.
	
	}
	
	}