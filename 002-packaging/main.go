package main

import "fmt"

func main() {

	/**
	godoc.org - packages
	func Println(a ...interface{}) (n int, err error)

	[variadic parameter]

	"... <some types>" is how we signify a variadic parameter
	"interface{}" is the empty interface
	so the parameter "... interface{}" means "give me as many arguments of any type as you'd like"

	Println formats using the default formats for its operands and writes to standard output. Spaces are always added between
	operands and a newline is appended. It returns the number of bytes written and any write error encountered.
	*/
	// use the "_" underscore character to throw away returns
	n, _ := fmt.Println("Hello", 42, true)
	fmt.Println(n) // number of bytes, it's written
}
