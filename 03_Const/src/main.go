package main

import (
	"fmt"
	"math"
)

const str string = "Shubham"

func main() {

	fmt.Println(str)

	const n = 5000000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
