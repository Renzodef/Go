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
	// Defining two floats
	// If we define them withouth the float32 flag
	// The division wouldn't be possible
	// because the result would be a float
	var num1 float32 = 9
	var num2 float32 = 4
	answer1 := num1 / num2
	// %g returns a float
	fmt.Printf("%g\n", answer1) // 2.25
	// Defining two other numbers
	// By default they will be considered as integer
	var num3 = 9
	var num4 = 4
	// Modulo operation that returns the reminder
	answer2 := num3 % num4
	// %d returns a decimal
	// If we puy %g, a warning will pop up
	// and the result will be converted to integer in output
	fmt.Printf("%d", answer2) // 1
}

/*
Output:
2.25
1
*/
