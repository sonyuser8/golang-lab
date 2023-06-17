package studyTypes

import (
	"log"
)

func studyMakeAndNew() {
	//built-in functions to allocate memory for a given target type T
	// make(T, args) & new(T)
	// - make(chan bool)
	// - make([]int, 5)

	// make() only works with slices, maps, and channels & new() works with any
	s := make([]int, 0)
	m := make(map[int]int)
	c := make(chan bool)

	log.Println(s, m, c)
}
