package main

import "fmt"

// it's like global variable to this file
// the scope of this variable is the current file
var outsideParams = "Hello"

func main() {

	// short decoration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	// := is an operator because is operating on these operands
	x := 42
	fmt.Println("x is: ", x)

	var y = 43
	fmt.Println("y is: ", y)

	/**
	What is the difference between var and gopher operator?
		- non-declaration statement outside function body.

	*/
}

func foo() {
	fmt.Println(outsideParams)
}
