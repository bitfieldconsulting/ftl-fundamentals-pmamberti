// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(nums ...float64) float64 {
	var result float64 = nums[0]

	for _, n := range nums[1:] {
		result += n
	}
	return result
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(nums ...float64) float64 {
	var result float64 = nums[0]

	for _, n := range nums[1:] {
		result -= n
	}
	return result
}

// Multiply takes two numbers and returns the result of multiplying them
func Multiply(nums ...float64) float64 {
	var result float64 = 1

	for _, n := range nums {
		result *= n
	}

	return result
}

// Divide takes two numbers and returns the result of dividing the first by the second
func Divide(nums ...float64) (float64, error) {
	var result float64 = nums[0]
	var err error

	for _, n := range nums[1:] {
		if n == 0 {
			err = errors.New("Error: division by 0 is not allowed")
			result = 555
		}
		result /= n

	}

	return result, err
}

// Sqrt takes a number and returns its square root
func Sqrt(a float64) (result float64, err error) {
	if a >= 0 {
		return math.Sqrt(a), nil
	}

	err = errors.New("You need a positive number or this won't work")
	return 999, err
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
