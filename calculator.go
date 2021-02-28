// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(nums ...float64) float64 {
	var result float64 = nums[0]

	for _, n := range nums[1:] {
		result = result + n
	}
	return result
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(nums ...float64) float64 {
	var result float64 = nums[0]

	for _, n := range nums[1:] {
		result = result - n
	}
	return result
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

// Sqrt takes a number and returns its square root
func Sqrt(a float64) (result float64, err error) {
	if a >= 0 {
		return math.Sqrt(a), nil
	}

	err = errors.New("You need a positive number or this won't work")
	return 999, err
}
