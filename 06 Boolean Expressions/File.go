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
	x := 5
	y := 5.0
	// val and val2 will be a boolean
	val := x < 4
	// Converting x to float64
	// or the operation won't be possible
	// since we would compare an integer with a float
	val2 := float64(x) == y
	// %t returns a boolean
	fmt.Printf("%t\n", val) // false
	fmt.Printf("%t", val2)  // true
}

/*
Output:
false
true
*/
