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

	// Returned Types
	// make() returns the variable target type & new() returns the address
	var v map[int]int = make(map[int]int) // Initializes a map
	var p *map[int]int = new(map[int]int) // Zeros a map pointer to nil
	log.Println(v, p)
}
