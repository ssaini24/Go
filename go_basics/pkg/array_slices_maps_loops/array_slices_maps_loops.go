package array_slices_maps_loops

import (
	"fmt"
	"strconv"
)

func Init() {
	// array init
	intArray := [3]int{1, 2, 3}
	fmt.Println(intArray[0])
	fmt.Println(intArray)

	// slice init
	var initSlice = []int{10, 20, 30}
	fmt.Printf("The length of the slice is %v with %v capacity \n", len(initSlice), cap(initSlice))
	initSlice = append(initSlice, 100)
	fmt.Printf("The length of the slice is %v with %v capacity ", len(initSlice), cap(initSlice))

	// appending slices
	var initSlice2 []int = []int{8, 10}
	initSlice = append(initSlice, initSlice2...)
	fmt.Println(initSlice)

	// slices defined with size and capacity
	// has great performance impacts if we know the capacity
	var intSlice3 []int32 = make([]int32, 10, 20)
	fmt.Println(intSlice3[0])

	// Maps
	var myMap map[string]uint8 = make(map[string]uint8)
	println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23}
	println(myMap2["Adam"])

	// if a key exist map will always return a value!
	println(myMap2["eve"])
	// or more accurate way
	var value, ok = myMap2["eve"]
	if !ok {
		println("key does not exists in map")
	} else {
		println("key has value: " + strconv.Itoa(int(value)))
	}

	// loops
	for key, value := range myMap2 {
		fmt.Printf("map has key: %v with value %v \n", key, value)
	}
}
