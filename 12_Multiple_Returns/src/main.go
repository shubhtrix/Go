package main

import  "fmt"

// Go has built-in support for multiple return values.
// This feature is used often in idiomatic Go, for example
// to return both result and error values from a function.

func vals (int) (int, int) {
	return 3, 7
}

func main() {

	a, b := vals(5)

	fmt.Println("Returned value :: a:  ", a)
	fmt.Println("Returned value :: b:  ", b)

	_, c := vals(9)
	fmt.Println("Returned value :: c:  ", c)
}
