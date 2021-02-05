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
	fmt.Printf("Hello world")

	foo()

	fmt.Println("something more.")

	for i := 0; i < 100; i++ {

	}

}

func foo() {
	fmt.Println("I am in foo.")
}
