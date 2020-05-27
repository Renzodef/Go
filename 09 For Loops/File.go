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
	x := 0
	// For cycle
	for x <= 5 {
		if x == 2 {
			// If we put continue here the loop will fail
			break
		}
		fmt.Println(x)
		// Increment at every iteration
		x++
	}
	// Just for a better output
	fmt.Printf("\n")

	// For cycle with definition and increment
	for x := 0; x <= 5; x++ {
		if x == 2 {
			// Continue should be used in a loop defined in this way
			continue
		}
		fmt.Println(x)
	}
}

/*
Output:
0
1

0
1
3
4
5
*/
