package main

import  "fmt"

// Maps are Go’s built-in associative data type
// (sometimes called hashes or dicts in other languages).

func main() {

	// To create an empty map, use the builtin
	// make: make(map[key-type]val-type).
	m := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax.
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("New map m : ", m)

	v1 := m["k1"]
	fmt.Println("New map v1 : ", v1)

	fmt.Println("Len of m : ", len(m))

	delete(m, "k2")
	fmt.Println("After delte map m : ", m)

	// The optional second return value when getting a value from
	// a map indicates if the key was present in the map. This can
	// be used to disambiguate between missing keys and keys with 
	// zero values like 0 or "". Here we didn’t need the value itself,
	// so we ignored it with the blank identifier _.
	_, prs := m["k2"]
	fmt.Println("New map prs : ", prs)

	// You can also declare and initialize a new map in the same
	// line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("Single line declaration of map n : ", n)
}
