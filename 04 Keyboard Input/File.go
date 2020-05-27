// From terminal go in the directory of this file and
// run the file with:
// go run File.go
// Or build the file with:
// go build File.go
// and then launch it depending on your system OS

package main

// Needed libraries
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Defining the input function
	reader := bufio.NewReader(os.Stdin)
	// Output message
	fmt.Print("Enter text: ")
	// Keyboard's input that we assign to a variable
	text, _ := reader.ReadString('\n')
	// Println will add an empty line after the print
	// Printing the variable that contains the keyboard's input
	fmt.Print("You typed: ")
	fmt.Println(text)

	// Just for a better output
	fmt.Printf("////////////////////\n\n")

	// Another similar way to achieve the same aim
	// But the input will be put between ""
	scanner := bufio.NewScanner(os.Stdin)
	// Output message
	fmt.Print("Enter text: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("You typed: %q", input)

	// Just for a better output
	fmt.Printf("\n\n////////////////////\n\n")

	// If we want to input a number
	scanner2 := bufio.NewScanner(os.Stdin)
	// Output message
	fmt.Print("Enter the year you were born: ")
	scanner2.Scan()
	input2, _ := strconv.ParseInt(scanner2.Text(), 10, 64)
	// If input2 is not a number, it will be considered as 0
	fmt.Printf("You will be %d years old at the end of 2020!", 2020-input2)

	// Just for a better output
	fmt.Printf("\n\n////////////////////\n\n")

	// These lines are required in order to not close the terminal
	// on Windows when we maker an executable file
	// On linux we won't have this problem
	fmt.Print("Press Enter to exit: ")
	reader.ReadString('\n')
}
