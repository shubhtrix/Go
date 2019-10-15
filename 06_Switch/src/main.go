package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2

	fmt.Print ("Write ", i, " as ")
	fmt.Println("i is ", i)
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's a week day")
	}

	// Switch without an expression is an alternate way to express if/else logic
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// A type switch compares types instead of values. You can use this to
	// discover the type of an interface value. In this example, the 
	// variable t will have the type corresponding to its clause.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println("Don't know type %T\n",t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("Hey!!!")
}
