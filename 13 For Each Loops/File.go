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
	// Defining a slice
	var slice []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// Simple for cycle that prints all the elements of the slice
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// Just for a better output
	fmt.Printf("\n")

	// For each loop (only to be used with slices, arrays or other iterables)
	// _ is used to define a variable that we don't want to use later
	// In a for each loop the first element is the index
	// The second one represents the element of the array in that index
	for _, element := range slice {
		fmt.Println(element)
	}

	// Just for a better output
	fmt.Printf("\n")

	// For each cycle 2
	// i in this case will be assigned to the index of the cycle
	for i, element := range slice {
		fmt.Printf("%d: %d\n", i, element)
	}
}

/*
Output:
1
2
3
4
5
6
7
8
9

1
2
3
4
5
6
7
8
9

0: 1
1: 2
2: 3
3: 4
4: 5
5: 6
6: 7
7: 8
8: 9
*/
