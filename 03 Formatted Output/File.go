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
	// Variables declaration
	number := 3
	number2 := 3.2432

	// Printf enables the output formatter
	// /n will make a new line
	// %v will print the value of a variable
	fmt.Printf("Number \n%v\nis cool", number)
	
	fmt.Printf("\n\n////////////////////\n\n")

	// Sprintf enables to assign a formatted string to a variable
	// %.f will round the number selected
	var string1 = fmt.Sprintf("Number \n%.f\nis cool", number2)
	fmt.Println(string1)
}

/*
Output:
Number 
3
is cool

////////////////////

Number 
3
is cool
*/