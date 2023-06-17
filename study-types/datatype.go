package studyTypes

import (
	"log"
)

func studyArray() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// playground here
	log.Println("m0:", [2]int{99, 98})

	var m1 = [2]int{1, 2}
	log.Println("m1:", m1)

	m2 := [2]int{3, 4}
	log.Println("m2:", m2)
}

func studyArray2() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// playground here
	log.Println("m0:")

}
