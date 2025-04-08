package main

import (
	"fmt"
	"go_basics/pkg/functions_and_control"
)

func main() {
	executeFunctionsAndControlsPkg()
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
