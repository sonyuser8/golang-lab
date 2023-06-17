package studyTypes

import (
	"log"
)

func studySlice() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	// playground here
	var slice1 []int = []int{1, 2, 3, 4, 5} // []int means a int slice
	log.Println("Value of slice1 is ", slice1)
	log.Println("Reslice slice1: ", slice1[1:3]) // start from element 1 and ended in element3 but not include element3
	log.Println("Extend slice1: ", append(slice1, 10))

	slice2 := make([]int, 5)
	// slice2 := [2]int{1, 2}
	log.Printf("Type of slice2 is %T", slice2)
	log.Println("Type of slice2 is ", slice2)

	var slice3 []int
	log.Printf("Type of slice3 is %T", slice3)
	log.Println("Type of slice3 is ", slice3)
}
