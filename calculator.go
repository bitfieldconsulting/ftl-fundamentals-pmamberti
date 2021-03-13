// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
	"math"
)

// CloseEnough takes the result of a division between two float-point
// numbers together with the requested tolerance, and returns true if
// the rounding is within the tolerance.
func CloseEnough(roundResult, divisionResult, tolerance float64) bool {
	return math.Abs(roundResult-divisionResult) <= tolerance
}

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64, nums ...float64) float64 {
	result := a + b

	for _, n := range nums {
		result += n
	}
	return result
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64, nums ...float64) float64 {
	result := a - b

	for _, n := range nums {
		result -= n
	}
	return result
}

// Multiply takes two numbers and returns the result of multiplying them
func Multiply(a, b float64, nums ...float64) float64 {
	result := a * b

	for _, n := range nums {
		result *= n
	}

	return result
}

// Divide takes two numbers and returns the result of dividing the first by the second
func Divide(a, b float64, nums ...float64) (result float64, err error) {

	if b == 0 {
		err = errors.New("Division by 0 is not allowed")
		result = 0
		return result, err
	}

	result = a / b

	for _, n := range nums {
		if n == 0 {
			err = errors.New("Division by 0 is not allowed")
			result = 0
			return result, err
		}
		result /= n

	}

	return result, err
}

// Sqrt takes a number and returns its square root
func Sqrt(a float64) (result float64, err error) {
	if a < 0 {
		err = errors.New("You need a positive number or this won't work")
		return 0, err
	}
	return math.Sqrt(a), nil
}

// Evaluate takes a string, and returns the result
// of evaluating the expression in the string
func Evaluate(s string) (result float64, err error) {
	var a, b float64
	var operator string

	fmt.Sscanf(s, "%f %s %f", &a, &operator, &b)

	switch operator {
	case "+":
		result = Add(a, b)
	case "-":
		result = Subtract(a, b)
	case "*":
		result = Multiply(a, b)
	case "/":
		result, err = Divide(a, b)
	}
	return result, err

}
