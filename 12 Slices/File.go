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
	// Defining the array
	var array [5]int = [5]int{1, 2, 3, 4, 5}
	// Creating a slice from the array
	// A slice is an array withouth a defined number of values
	var slice []int = array[0:3]
	fmt.Println(slice)
	// Append can only be executed with a slice
	// and not with an array
	a := append(slice, 4)
	fmt.Println(a)

	// Just for a better output
	fmt.Printf("\n")

	// Create a slice not from an array
	var slice1 []int = []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)
	b := append(slice1, 6)
	fmt.Println(b)

	// Just for a better output
	fmt.Printf("\n")

	// Another convenient way to create a slice
	// A slice with 5 values by default as 0
	slice2 := make([]int, 5)
	// Assigning values to the slice
	// so we change the default 0 values
	slice2 = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice2)
	c := append(slice2, 8)
	fmt.Println(c)
}

/*
Output:
[1 2 3]
[1 2 3 4]

[1 2 3 4 5]
[1 2 3 4 5 6]

[1 2 3 4 5 6 7]
[1 2 3 4 5 6 7 8]
*/
