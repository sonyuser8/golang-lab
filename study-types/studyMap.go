package studyTypes

import (
	"log"
)

func studyMap() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	// playground here
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}
	log.Println(mp)

	var mp2 = make(map[string]int) // makes a empty map
	log.Println(mp2)
}
