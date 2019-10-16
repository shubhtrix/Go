package main

import  "fmt"

func plus (a int, b int) int {

	return a + b
}

func plus_m (a, b, c int) int {
	return a + b + c
}

func main() {

	var a, b int = 2, 3

	res := plus (a, b)
	fmt.Println("Plus output : ", res)

	res = plus_m (a, b, 3)
	fmt.Println("Plus_m output : ", res)
}
