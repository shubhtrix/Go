package main

import  "fmt"

// Unlike arrays, slices are typed only by the elements they contain
// (not the number of elements). To create an empty slice with non-zero
// length, use the builtin make. Here we make a slice of strings of 
// length 3 (initially zero-valued).

func main() {

	s:= make([]string, 3)
	fmt.Println("Empty: ", s)

	// set slice s
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("s:", s)
	fmt.Println("s[2]:", s[2])

	fmt.Println("length: ",len(s))

	s = append(s, "d")
	s = append(s, "d", "e")
	fmt.Println("New s:", s)

	//Slices can also be copy’d. Here we create an empty
	// slice c of the same length as s and copy into c from s
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("New slice c:", c)

	// Slice s[2], s[3], s[4]
	l := s[2:5]
	fmt.Println("New slice c:", l)

	// excluding 5th element
	l = s[:5]
	fmt.Println("New slice c:", l)

	// including 2nd element
	l = s[2:]
	fmt.Println("New slice c:", l)

	t := []string{"o","p","q"}
	fmt.Println("Single line declaration of new slice t :", t)

	twoD := make([][]int,3)

	for i := 0; i< 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D:", twoD)
}
