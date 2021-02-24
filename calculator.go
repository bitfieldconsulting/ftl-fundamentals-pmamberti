// Package calculator provides a library for simple calculations in Go.
package calculator

import "errors"

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result of multiplying them
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of dividing the first by the second
func Divide(a, b float64) (result float64, err error) {
	if b != 0 {
		return a / b, err
	}

	err = errors.New("you cannot divide by 0")

	return 555, err

}
