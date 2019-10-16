package main

import  "fmt"

// Range iterates over elements in a variety of data structures.
func main() {

	nums := []int{1,2,3}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("Sum of nums is : ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("Index: ", i)
		}
	}

	// range on map iterates over key/value pairs.
	kvs := map[string]string{"key1": "val1", "key2": "val2"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k , v)
	}

	for k := range kvs {
		fmt.Println("key: ", k)
	}

	// range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the rune
	// and the second the rune itself.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
