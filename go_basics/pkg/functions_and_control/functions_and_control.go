package functions_and_control

import (
	"errors"
	"fmt"
)

// Function with parameters and return value
func AddNumbers(a int, b int) int {
	return a + b
}

// Function with no return value
func PrintMessage(message string) {
	fmt.Println(message)
}

// Function with multiple return values
func IntDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("division by zero")
		return 0, 0, err
	}

	result := (int)(numerator / denominator)
	remainder := numerator % denominator
	return result, remainder, err
}
