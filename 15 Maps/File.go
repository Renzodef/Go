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
	// Maps are the equivalent of dictionaries in Python
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 7,
	}
	// Print the value of the key apple
	fmt.Println(mp["apple"]) // 5
	// Change the value of the key apple
	mp["apple"] = 10
	// Print again the value of the key apple
	fmt.Println(mp["apple"]) // 10
	// Assigning a boolean to the key apple
	val, ok := mp["apple"]
	// Returning the value and true if the key-value couple exists
	fmt.Println(val, ok) // 10 true
	// Delete the key-value couple where the key is apple
	delete(mp, "apple")
	// Assigning a new boolean to the key apple
	val2, ok2 := mp["apple"]
	// Returning 0 and false if the key-value couple not exists
	fmt.Println(val2, ok2) // 0 false
	// Printing the number of key-value couples of the map
	fmt.Println(len(mp)) // 2
	// Print all the key-value couples
	fmt.Println(mp)
}

/*
Output:
5
10
10 true
0 false
2
map[orange:7 pear:6]
*/
