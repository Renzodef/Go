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
	age := 18
	if age >= 18 {
		fmt.Println("You can ride alone!")
		// Need to put else if in this way or it won't work
	}
	// Another if should be put in this way
	if age == 18 {
		fmt.Println("You are 18 now!")
	} else if age >= 14 {
		fmt.Print("You can ride with a parent!")
		// Need to put else in this way or it won't work
	} else {
		fmt.Println("You can't ride!")
	}
}

/*
Output:
You can ride alone!
You are 18 now!
*/
