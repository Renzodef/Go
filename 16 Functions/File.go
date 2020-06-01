// From terminal go in the directory of this file and
// run the file with:
// go run File.go
// Or build the file with:
// go build File.go
// and then launch it depending on your system OS

package main

// Needed libraries
import "fmt"

// Void function
// without input and without output (so return isn't needed)
func void() {
	// Defer statements delay the execution of the function or method
	// If there is a return, defer it will be executed just before it (but after other outputs)
	// If the function hasn't a return, the defer will be executed as last output
	defer fmt.Println("This will be executed after Hello World!")
	fmt.Println("Hello World!")
}

// Put int after x is useless,
// because if you put the type only at the last variable,
// all the variables will take that type
// The first arguments are the input of the function
// The second block of arguments are the output of the function
// that should be returned from the function (with a return)
func function1(x, y int) (int, int) {
	return x + y, x - y
}

// Same function as before but with a different sintax
func function2(x, y int) (z1, z2 int) {
	z1 = x + y
	z2 = x - y
	return
}

func main() {
	// Calling the function void
	void()
	// Calling the function with the proper inputs
	ans1, ans2 := function1(14, 7)
	ans3, ans4 := function2(14, 7)
	fmt.Println(ans1, ans2) // 21 7
	fmt.Println(ans3, ans4) // 21 7
}

/*
Output:
Hello World!
This will be executed after Hello World!
21 7
21 7
*/
