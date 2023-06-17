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

	arr2 := [3]int{3, 4, 5}
	log.Println("arr2:", arr2)

	log.Println("loop method 1")
	// [+] loop method1
	for i := 0; i <= len(arr2)-1; i++ {
		log.Println("arr2 idx i:", i, "=", arr2[i])
	}
	// [-]

	log.Println("loop method 2")
	// [+] loop method2
	for idx, el := range arr2 {
		log.Println("arr2 idx:", idx, "=", el)
	}
	// [-]

	log.Println("loop method 2")
	// [+] loop method3: I don't need a index
	for _, el := range arr2 {
		log.Println("arr2 el", "=", el)
	}
	// [-]

}
