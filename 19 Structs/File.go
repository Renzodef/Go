// From terminal go in the directory of this file and
// run the file with:
// go run File.go
// Or build the file with:
// go build File.go
// and then launch it depending on your system OS

package main

import "fmt"

// Needed libraries

// Point is a struct with two properties
type Point struct {
	x int
	y int
}

// Function that takes as input
// a Point and a value and change the x value
// of the point with the new value taken from input
func changeX(p Point, y int) Point {
	p.x = y
	return p
}

func main() {
	// Creating a new Point
	// With x = 1 and y = 3
	var p1 Point = Point{1, 3}
	// Printing the property x of the Point p1
	fmt.Println(p1.x)
	// Creating a new Point
	// With x = 4
	// By defauly y will be 0
	var p2 Point = Point{x: 4}
	// Printing the new Point
	fmt.Println(p2) // {4 0}
	// Assigning to p3 the new Point with the changed value
	p3 := changeX(p2, 9)
	// Printing the new Point
	fmt.Println(p3) // {9 0}
}

/*
Output:
1
{4 0}
{9 0}
*/
