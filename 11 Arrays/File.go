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
	// array is the name of the array's variable
	// The type of the array has to be defined as below
	// By default the values are 0
	var array1 [5]int = [5]int{1, 2, 3, 4, 5}
	array1[0] = 100
	fmt.Println(array1)
	fmt.Println(len(array1))

	// Just for a better output
	fmt.Printf("\n")

	// Array with defined values
	array2 := [3]int{1, 2, 3}
	fmt.Println(array2)
	fmt.Println(len(array2))

	// For cycle with array
	sum := 0
	for i := 0; i < len(array2); i++ {
		sum += array2[i]
	}
	fmt.Printf("\n%d", sum)
}

/*
Output:
[100 2 3 4 5]
5

[1 2 3]
3

6
*/
