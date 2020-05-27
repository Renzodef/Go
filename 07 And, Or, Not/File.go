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
	// && means And,
	// so the boolean expression will be true
	// if both the conditions are true
	val := 5 < 7 && 6 < 7
	fmt.Printf("%t\n", val) // true
	val2 := 6 < 7 || 8 < 7
	// || means Or,
	// so the boolean expression will be true
	// if one of the conditions is true
	fmt.Printf("%t\n", val2) // true
	// ! means Not,
	// so the boolean expression will return
	// the opposite of its original value
	val3 := !(6 < 7 || 8 < 7)
	fmt.Printf("%t", val3) // false
}

/*
Output:
true
true
false
*/
