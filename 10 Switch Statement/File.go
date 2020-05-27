// From terminal go in the directory of this file and
// run the file with:
// go run File.go
// Or build the file with:
// go build File.go
// and then launch it depending on your system OS

package main

// Needed libraries
import (
	"fmt"
)

func main() {
	x := -1
	switch x {
	case 1, -1:
		fmt.Println("Case 1")
	case 2, -2:
		fmt.Println("Case 2")
	default:
		fmt.Println("Not a case")
	}

	// Another way for Switch
	switch {
	case x > 0:
		fmt.Println("Positive Number")
	case x < 0:
		fmt.Println("Negative Number")
	default:
		fmt.Println("The number is 0")
	}
}

/*
Output:
Case 1
Negative Number
*/
