package main

import "fmt"

// it's like global variable to this file
// the scope of this variable is the current file
// the package level scope
var outsideParams = "Hello"

/**
DECLARE there is a VARIABLE with the IDENTIFIER "z"
and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
ASSIGNS the ZERO VALUE of TYPE int to "z"
*/
var z int

/**
Zero values
Variables declared without an explicit initial value are given their zero value.

The zero value is:

0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.
*/

// DECLARE the variable "initialization"
// ASSIGN the value 25
// declare & assign = initialization
var initialization = 25

func main() {

	// short decoration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	// := is an operator because is operating on these operands
	x := 42
	fmt.Println("x is:", x)

	// variable
	var y = 43
	fmt.Println("y is:", y)

	// assign value to the variable z
	z = 25

	// function call
	foo()
	zeroValues()

	/**
	What is the difference between var and gopher operator?
		- non-declaration statement outside function body.
		- we can assign vars later, vars will hold their default values
	*/
}

func foo() {
	fmt.Println("outsideParams:", outsideParams)
	fmt.Println("initialization:", initialization)
	fmt.Println("z:", z)
}

func zeroValues() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("i: %v,  f: %v, b: %v, s: %q\n", i, f, b, s)
}
