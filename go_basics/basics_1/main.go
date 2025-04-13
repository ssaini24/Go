package main

import (
	"fmt"
	"go_basics/pkg/array_slices_maps_loops"
	"go_basics/pkg/channels"
	"go_basics/pkg/functions_and_control"
	"go_basics/pkg/generics"
	"go_basics/pkg/goroutines"
	"go_basics/pkg/pointers"
	"go_basics/pkg/strings_runes_bytes"
	"go_basics/pkg/structs_interfaces"
)

func main() {
	executeFunctionsAndControlsPkg()
	executeArraysSlicesMapsLoopsPkg()
	strings_runes_bytes.Init()
	executeStructsInterfaces()
	pointers.Init()
	goroutines.Init()
	channels.Init()
	generics.Init()
}

func executeStructsInterfaces() {
	var gasEngine structs_interfaces.GasEngine
	gasEngine.Mpg = 10
	gasEngine.Gallons = 1
	structs_interfaces.CanMakeIt(gasEngine, 100)
}

func executeArraysSlicesMapsLoopsPkg() {
	array_slices_maps_loops.Init()
}

func executeFunctionsAndControlsPkg() {
	// Using AddNumbers
	sum := functions_and_control.AddNumbers(5, 10)
	fmt.Println("Sum:", sum)

	// Using PrintMessage
	functions_and_control.PrintMessage("Hello from the new package!")

	quotient, remainder, err := functions_and_control.IntDivision(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder)
	}

	quotient, remainder, err = functions_and_control.IntDivision(10, 3)
	switch {
	case err != nil:
		fmt.Println("Error:", err)
	default:
		fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder)
	}

	switch remainder {
	case 0:
		fmt.Println("No remainder")
	case 1:
		fmt.Println("Remainder is 1")
	case 2:
		fmt.Println("Remainder is 2")
	default:
		fmt.Println("Remainder is something else")
	}
}
