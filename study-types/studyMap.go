package studyTypes

import (
	"log"
)

func studyMap() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	// playground here
	var slice1 []int = []int{1, 2, 3, 4, 5} // []int means a int slice
	log.Println("Value of slice1 is ", slice1)
}
