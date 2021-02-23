package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
	name string
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 4, name: "2 + 2 should be 4"},
		{a: 1, b: 0.2, want: 1.2, name: "Adding decimals should be a breeze"},
		{a: 5, b: 0, want: 5, name: "Adding 0 returns the same number"},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%v - want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 0, name: "Subtracting a number to itself returns 0"},
		{a: 1, b: 0.2, want: 0.8, name: "Subtracting decimals should work, no?"},
		{a: 5, b: 7, want: -2, name: "Subtracting a bigger value will return a negative result"},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("%v - want %f, got %f", tc.name, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 1, want: 2, name: "Multiplication by 1 returns the same number"},
		{a: 10, b: 0.2, want: 2, name: "Multiplication between 0 and 1 will return a smaller number"},
		{a: 3, b: 0, want: 0, name: "Multiplication by 0 will return 0"},
		{a: 0, b: 0, want: 0, name: "0 times 0 is still 0"},
		{a: 8, b: 4, want: 32, name: "Multiplication will return the product of two numbers"},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)

		if tc.want != got {
			t.Errorf("%v - want %f, got %f", tc.name, tc.want, got)
		}
	}

}
