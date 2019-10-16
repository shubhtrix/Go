package main

import  "fmt"

// Go supports anonymous functions, which can form closures. 
// Anonymous functions are useful when you want to define a 
// function inline without having to name it.


// This function intSeq returns another function, which we 
// define anonymously in the body of intSeq. The returned 
// function closes over the variable i to form a closure.
func inSeq() func() int {

	i := 0
	return func () int {
		i++
		return i
	}
}

func main() {

	nextInt := inSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := inSeq()
	fmt.Println(newInt())
}
