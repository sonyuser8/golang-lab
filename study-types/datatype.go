package studyTypes

import (
	"log"
)

func studyArray() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// playground here
	log.Println("arr0:", [2]int{99, 98})

	var arr1 = [2]int{1, 2}
	log.Println("arr1:", arr1)

	arr2 := [2]int{3, 4}
	log.Println("arr2:", arr2)

	var arr3 [2]string
	log.Println(arr3)
}
