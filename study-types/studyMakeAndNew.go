package studyTypes

import (
	"log"
)

func studyMakeAndNew() {
	//built-in functions to allocate memory for a given target type T
	// make(T, args) & new(T)
	// - make(chan bool)
	// - make([]int, 5)
	slice := make([]int, 5)
	log.Printf("Type of slice2 is %T", slice)
	log.Println("Type of slice2 is ", slice)
}
