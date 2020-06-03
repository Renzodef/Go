// From terminal go in the directory of this file and
// run the file with:
// go run File.go
// Or build the file with:
// go build File.go
// and then launch it depending on your system OS

package main

// Needed libraries
import "fmt"

func main() {
	// Defining a variable
	x := 5
	// & is the pointer
	// This line will print the location of the variable x
	fmt.Println(&x)

	// Just for a better output
	fmt.Printf("\n")

	// Assigning the pointer of x to y
	y := &x
	// Printing x and its location
	fmt.Println(x, y)
	// * is the deference operator
	// In this way x will be edited from 7 to 8
	*y = 8
	// Printing x with the new value and its location
	fmt.Println(x, y)
}

/*
Output:
0xc0000120c8

5 0xc0000120c8
8 0xc0000120c8
*/
