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
	// Install it from terminal with:
	// got get gonum.org/v1/gonum/mat
	"gonum.org/v1/gonum/mat"
)

func main() {
	// Defining the rows
	row1 := []float64{1, 2, 3}
	row2 := []float64{4, 5, 6}
	row3 := []float64{7, 8, 9}
	row4 := []float64{10, 11, 12}
	// Set the dimensions of the matrix
	A := mat.NewDense(4, 3, nil)
	// Assigning the rows to the matrix
	A.SetRow(0, row1)
	A.SetRow(1, row2)
	A.SetRow(2, row3)
	A.SetRow(3, row4)
	// Print the formatted matrix
	fmt.Printf("\n%v\n\n", mat.Formatted(A, mat.Prefix(""), mat.Excerpt(0)))
	// Get the dimensions of the matrix
	rows, cols := A.Dims()
	fmt.Println("rows: ", rows)
	fmt.Println("cols: ", cols)
}

/*
Output:

⎡ 1   2   3⎤
⎢ 4   5   6⎥
⎢ 7   8   9⎥
⎣10  11  12⎦

rows:  4
cols:  3
*/
