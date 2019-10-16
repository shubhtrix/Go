package main

import  "fmt"

type Person struct {
	name string
	age int
}

func NewPerson (n string) *Person {

	p := Person{name: n}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(Person{"shubh", 29})

	fmt.Println(Person{name: "shubh", age: 29})

	fmt.Println(Person{name: "Fred"})

	fmt.Println(&Person{name: "next", age: 39})

	fmt.Println(NewPerson("newname"))

	s := Person{name: "struct", age: 12}

	fmt.Println(s.name)

	//You can also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &s
}
