package main

import (
	"fmt"
	m "math" // So cute! we can use a alias for import
)

/*
 Every variable or constants declared must be used or the compiler complains
*/
func main() {
	const PI float64 = 3.1415

	fmt.Printf("PI is %.2f \n", PI)

	var radius = 3.2 // type float is inserted by compiler

	// reduced form to declare variable := Declare and initialize
	area := PI * m.Pow(radius, 2)
	fmt.Println("The circumference area is:", area)

	// Blocks: It's ok if you desire declare constants or variables like this ...
	const (
		x = 1
		y = 2
	)

	fmt.Println(x + y)

	var (
		a = 5
		b = 2
	)

	fmt.Println(a * b)

	// And if you prefer declare at the same time
	var t, f bool = true, false
	fmt.Println(t, f)

	// Maybe you're lazy and everything in the same line looks great
	first, second, third := 1, false, "third"

	fmt.Println(first, second, third)

}
