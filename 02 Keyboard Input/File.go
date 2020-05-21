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
    fmt.Println(text)
    // These lines are required in order to not close the terminal
    // on Windows when we make an executable file
    // On linux we won't have this problem
    fmt.Print("Press Enter to exit: ")
    reader.ReadString('\n')
}