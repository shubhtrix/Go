package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is devisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println("Num is -ve")
	} else if num < 10 {
		fmt.Println("Num is having single digit")
	} else {
		fmt.Println("Num is having multiple digit")
	}
}
