package pointers

import (
	"fmt"
)

func Init() {
	println("Pointers package initialized")

	var p *int32 = new(int32)
	fmt.Printf("inital value of p is %v \n", *p)
	var i int32 = 10
	p = &i
	fmt.Printf("final value of p is %v \n", *p)
	*p = 20
	fmt.Printf("Value of i changed to %v \n", i)

	// pointers in slices
	var sliceInit = []int32{1, 2, 3}
	var sliceInitCopy = sliceInit
	sliceInitCopy[2] = 10
	// as 2nd slice points to 1st slice changing value in 2nd slice will reflect in first one
	fmt.Printf("value of 1st slice %v \n", sliceInit)
	fmt.Printf("value of 2nd slice %v \n", sliceInitCopy)

	// pointers while calling functions
	var thing1 = [3]float64{1, 2, 3}
	var thing2 [3]float64 = square(thing1)
	// value of thing1 will not changes as passed by copy
	fmt.Printf("value of thing1 is %v \n", thing1)
	fmt.Printf("value of thing2 is %v \n", thing2)

	// value of thing1 also get changes as passed by ref
	thing2ByRef := squareByRef(&thing1)
	fmt.Printf("value of thing1 is %v \n", thing1)
	fmt.Printf("value of thing2 is %v \n", *thing2ByRef)
}

// copying thing1 array to thing2 slice
func square(thing2 [3]float64) [3]float64 {
	for key := range thing2 {
		thing2[key] = thing2[key] * thing2[key]
	}

	return thing2
}

// passing by ref thing1 array to thing2 array
func squareByRef(thing2 *[3]float64) *[3]float64 {
	for key := range thing2 {
		thing2[key] = thing2[key] * thing2[key]
	}

	return thing2
}
