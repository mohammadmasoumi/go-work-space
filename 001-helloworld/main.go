package main

import "fmt"

/**
create go module
 - go mod init PATH_TO_THE_MAIN.GO
 - go mod tidy
*/

/**
control flow:
  (1) sequence - code runs sequentially! line by line
  (2) loop: iterative
  (3) conditional
*/

// the entrypoint to your project
func main() {
	fmt.Println("Hello world")

	foo()

	// it takes a number of values
	fmt.Println("something more.", 23, true)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()

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
	n, _ := fmt.Println("Hello", 42, true)
	fmt.Println(n) // number of bytes, it's written

}

func foo() {
	fmt.Println("I am in foo.")
}

func bar() {
	fmt.Println("and then we excited!")
}
