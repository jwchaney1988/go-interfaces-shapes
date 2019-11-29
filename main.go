package main

// Import the needed packages
import "fmt"

// Create an interface named shape with a Method Name getArea
type shape interface {
	getArea() float64
}

// Create a struct named triangle with Field Names base and height
type triangle struct {
	base   float64
	height float64
}

// Create a struct named square with a Field Name sideLength
type square struct {
	sideLength float64
}

func main() {
	// Declare variables
	t := triangle{base: 10, height: 10}
	s := square{sideLength: 10}

	// Since both triangle and square satisfy the shape interface they can now call the func printArea
	// passing in their respective types
	printArea(t)
	printArea(s)
}

// Same method name and return type as what is in the shape interface, so it satisfies the the shape interface
// and can be used anywhere the shape interface is called for
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

// Same method name and return type as what is in the shape interface, so it satisfies the the shape interface
// and can be used anywhere the shape interface is called for
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

// Method to accept a type shape (or types that satisfy the shape interface) and print their value
func printArea(s shape) {
	// Since both triangle and square satisfy the shape interface Go will replace the s shape type with
	// the respective type that was passed in and call the correct GetArea func
	fmt.Println(s.getArea())
}
