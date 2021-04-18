// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
	"math"
)

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
		return 0, errors.New("Division by 0 is not allowed")
	}

	result = a / b

	for _, n := range nums {
		if n == 0 {
			return 0, errors.New("Division by 0 is not allowed")
		}
		result /= n

	}

	return result, nil
}

// Sqrt takes a number and returns its square root
func Sqrt(a float64) (result float64, err error) {
	if a < 0 {
		return 0, errors.New("You need a positive number or this won't work")
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
		return Add(a, b), err
	case "-":
		return Subtract(a, b), err
	case "*":
		return Multiply(a, b), err
	case "/":
		return Divide(a, b)
	default:
		return 0, fmt.Errorf("bad operator %q (must be +, -, * or /)", operator)
	}
}
