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
		"orange": 9, // must to have comma
	}
	mp["banana"] = 10       // add el into map
	delete(mp, "apple")     // delete el of map
	val, ok := mp["orange"] // check el existence
	log.Println(mp)
	log.Println("val:", val, ", ok:", ok)
	// val: 9 , ok: true

	var mp2 = make(map[string]int) // makes a empty map
	log.Println(mp2)
}
