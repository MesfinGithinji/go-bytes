package main

import (
	"fmt"
)

// calculateResult computes the sum of all numbers between a and b (inclusive of a and b),
// then divides that sum by the sum of a and b.
func calculateResult(a, b int) float64 {
	// Ensure a is less than or equal to b
	if a > b {
		a, b = b, a
	}

	// Sum all numbers between a and b (inclusive)
	sumRange := 0
	for i := a; i <= b; i++ {
		sumRange += i
	}

	// Sum of the two entered numbers
	sumAB := a + b

	// Avoid division by zero
	if sumAB == 0 {
		return 0
	}

	// Return result as float
	return float64(sumRange) / float64(sumAB)
}

func main() {
	a := 4
	b := 10

	result := calculateResult(a, b)
	fmt.Printf("The result is: %.2f\n", result) // Output: 3.50
}
