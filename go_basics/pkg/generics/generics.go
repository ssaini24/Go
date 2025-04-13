package generics

import "fmt"

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwh  float32
	mpkw float32
}

type car[T gasEngine | electricEngine] struct {
	name   string
	model  string
	engine T
}

func Init() {
	fmt.Println("generics init...")
	// generics at function level

	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice(intSlice))

	var floatSlice = []float32{10, 20, 30}
	fmt.Println(sumSlice(floatSlice))

	// generics at struct level
	carStruct := car[electricEngine]{
		"alto",
		"2024",
		electricEngine{
			20,
			220,
		},
	}

	fmt.Println(carStruct)
}

func sumSlice[T int | float32 | float64](sliceToAdd []T) T {
	var sum T
	for _, value := range sliceToAdd {
		sum += value
	}

	return sum
}
