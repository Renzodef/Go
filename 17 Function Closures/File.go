// From terminal go in the directory of this file and
// run the file with:
// go run File.go
// Or build the file with:
// go build File.go
// and then launch it depending on your system OS

package main

// Needed libraries
import "fmt"

// This function return a closure
// A closure is a function value that references variables from outside its body
// Closures are similar to anonymous function
func adder() func(int) int {
	sum := 0
	// Returning a closure
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// We assign the function adder to a variable
	// This variable will became a function
	// that returns the input number
	number := adder()
	// Calling the function assigned to the variable
	fmt.Println(number(5))
}

/*
Output:
5
*/
