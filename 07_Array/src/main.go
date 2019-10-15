package main

import  "fmt"

func main() {

	var arr [5]int
	fmt.Println("Empty: ", arr)


	arr[4] = 100
	fmt.Println("Set: ", arr)
	fmt.Println("Get: ", arr[4])

	fmt.Println("Len: ", len(arr))

	arr1 := [5]int{1,2,3,4,5}
	fmt.Println("dcl:", arr1)

	var arr2 [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(">")
			arr2[i][j] = i + j
		}
		fmt.Println("*")
	}

	fmt.Println("2d:", arr2)
}
