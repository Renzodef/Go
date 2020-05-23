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
	// Explicit variable declaration
	var number uint16 = 260
	// Implicit variables declaration
	// Go will guess the type of the variable
	var number2 = 260
	var number3 = 2000.65
	// Implicit variables declaration with the walrus operator
	number4 := 6
	string1 := "String"
	// Explicit variables declaration without the value
	var number5 float64
	var bl bool

	// Printing the type of the variables
	fmt.Printf("%T", number) // uint16
	// Just to separate the outputs
	fmt.Println()
	fmt.Printf("%T", number2) // int
	fmt.Println()
	fmt.Printf("%T", number3) // float64
	fmt.Println()
	fmt.Printf("%T", number4) // int
	fmt.Println()
	fmt.Printf("%T", string1) // string
	fmt.Println()
	fmt.Printf("%T", string1) // string
	fmt.Println()

	// Just for a better output
	fmt.Println()
	fmt.Println("///////////////////////////")
	fmt.Println()

	// Printing the value of number5
	// The value wasn't declared
	// But to will assign it by default to 0
	fmt.Println(number5) // 0
	// Same for the boolean type
	// The default is false
	fmt.Println(bl) // false
}

/*
Output:
uint16
int
float64
int
string
string

///////////////////////////

0
false
*/