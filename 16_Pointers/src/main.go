package main

import  "fmt"

func zeroval( a int) {
	a = 0
}

func zeroptr( a *int) {
	*a = 0
}

func main() {

	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	// To print pointerval of i
	fmt.Println("Address of i: ", &i)
}
