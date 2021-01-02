package mymath

import ( 
	
	"fmt"
    "math"
	 
)

//ScanFloatValueAndTruncateToInteger : scanning user input in float and truncting float to int
func ScanFloatValueAndTruncateToInteger() {  
	
	var myFloat float64
		
	for i:=0; i<2; i++ {
	
		fmt.Println("Enter float number")
	
		fmt.Scan(&myFloat)
		
		fmt.Printf("user input float value-> %f\n", myFloat)
		
		myFloat = math.Trunc(myFloat)
		fmt.Printf("Int value after match.Trunc-> %.f\n ", myFloat)  // %.f - here the dot(.) represent that we want no value after decimal.
	
	}
	
	}