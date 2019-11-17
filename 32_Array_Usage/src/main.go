package main

import  "fmt"

const (
	maxSessIDs int = 10
)

var (
	curFreeSessID int = 0
	sessCount int = 0
	SessIdMap [maxSessIDs]int
)

func initSessIdMap () {

	for i:=0; i<maxSessIDs; i++ {
		SessIdMap[i] = i + 1
	}

	fmt.Println("Initialization Done!!!")
}

func dumpSessIdMap () {

	for i:=0; i<maxSessIDs; i++ {
		fmt.Print(SessIdMap[i])
	}

	fmt.Print("\n")
}

func freeSessID (sessID int) {

	SessIdMap[sessID - 1] = curFreeSessID
	curFreeSessID = sessID - 1
	sessCount--
}

func allocSessID () int {

	//var allocID uint16

	allocID := curFreeSessID + 1
	curFreeSessID = SessIdMap[curFreeSessID]
	sessCount++

	return allocID
}

func main() {

	initSessIdMap()

	dumpSessIdMap()

	Id := allocSessID()
	fmt.Println("Allocated ID: ", Id)
	Id = allocSessID()
	fmt.Println("Allocated ID: ", Id)
	Id = allocSessID()
	fmt.Println("Allocated ID: ", Id)
	Id = allocSessID()
	fmt.Println("Allocated ID: ", Id)

	fmt.Println("Allocated ID count: ", sessCount)

	freeSessID(3)
	dumpSessIdMap()

	Id = allocSessID()
	fmt.Println("Allocated ID: ", Id)
	Id = allocSessID()
	fmt.Println("Allocated ID: ", Id)
	dumpSessIdMap()

	freeSessID(2)
	dumpSessIdMap()
	Id = allocSessID()
	fmt.Println("Allocated ID: ", Id)
	Id = allocSessID()
	fmt.Println("Allocated ID: ", Id)
	dumpSessIdMap()
/*	arr[4] = 100
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
*/
}
