package main

import  "fmt"

// Go supports methods defined on struct types.

type rect struct {
	height, width int
}

// This area method has a receiver type of *rect.
func (r *rect) area () int {

    return r.height * r.width
}

func (r rect) perim () int {

    return 2*r.height + 2*r.width
}

func main() {

    r := rect{height:10, width:4}

	fmt.Println(r.area())

	fmt.Println(r.perim())

    rp := &r

    fmt.Println(rp.area())

	fmt.Println(rp.perim())

}
