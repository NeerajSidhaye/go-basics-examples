package main

import (
	
	//"internal/pkg/utils/mymath"
	//"internal/pkg/utils/myslice"
	//"internal/pkg/utils/mystrings"

	//"github.com/BeTheCodeWithYou/go-basics-examples/helloworld/internal/pkg/utils/mymath"
    //"github.com/BeTheCodeWithYou/go-basics-examples/helloworld/internal/pkg/utils/mymap"
    "github.com/BeTheCodeWithYou/go-basics-examples/helloworld/internal/pkg/utils/myfiles"

	"fmt"
)

func main() {

	fmt.Println("Hello, world!")

	//mymath.ScanFloatValueAndTruncateToInteger()

	//mystrings.FindSpecificCharInString()

	//myslice.SortSliceOfIntAndPrint()

	//mymap.ConvertMapToJSON()

	myfiles.FileReaderAndPrintFileContentAsStructs()
}
