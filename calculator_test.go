package calculator_test

import (
	"calculator"
	"fmt"
	"math/rand"
	"testing"
)

type testCase struct {
	a, b        float64
	want        float64
	name        string
	errExpected bool
}

type variadicTestCase struct {
	nums        []float64
	want        float64
	name        string
	errExpected bool
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []variadicTestCase{
		{nums: []float64{1, 2, 3}, want: 6, name: "1+2+3 = 6"},
		{nums: []float64{1, 2, 3, 4}, want: 10, name: "1+2+3+4 = 10"},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.nums...)

		if tc.want != got {
			t.Errorf("Error: unexpected value returned - want %.1f, got %.1f", tc.want, got)
		}
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()
	for i := 0; i < 10000000; i++ {
		a := rand.Float64()
		b := rand.Float64()
		want := a + b
		got := calculator.Add(a, b)

		if want != got {
			t.Errorf("Want: %.2f, got %.2f", want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []variadicTestCase{
		{nums: []float64{3, 2, 1}, want: 0, name: "3 - 2 - 1 = 0"},
		{nums: []float64{100, 8, 12}, want: 80, name: "100 - 8 - 12 = 10"},
		{nums: []float64{1, 8, 10}, want: -17, name: "1 - 8 - 10 = -17"},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.nums...)

		if tc.want != got {
			t.Errorf("Error: unexpected value returned - want %.1f, got %.1f", tc.want, got)
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

func TestDivide(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 4, b: 0, want: 0, errExpected: true, name: "(fail) Error Expected: Division by 0 is not allowed"},
		{a: 5, b: 2, want: 2.5, errExpected: false, name: "(fail) Not Expecting an Error, got one instead"},
		{a: 3, b: 1, want: 3, errExpected: false, name: "(fail) Error Expected: wrong result returned"},
		{a: 100, b: 10, want: 10, errExpected: false, name: "(pass) Divide a number by another returns something"},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			fmt.Printf("Expected test: %s\n", tc.name)
			t.Fatalf("Divide(%.0f, %.0f: unexpected error status: %v)",
				tc.a, tc.b, errReceived)
		}

		if !tc.errExpected && tc.want != got {
			fmt.Printf("Expected test: %s\n", tc.name)
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}

}

func TestDivideRandom(t *testing.T) {
	t.Parallel()

	for i := 0; i < 1000; i++ {
		var a float64 = rand.Float64()
		var b float64 = rand.Float64()
		errExpected := false

		if b == 0 {
			errExpected = true
		}

		got, err := calculator.Divide(a, b)
		errReceived := err != nil

		if errExpected != errReceived {
			t.Fatalf("Unexpected Error Status: %v", errReceived)
		}

		want := a / b

		if !errReceived && want != got {
			t.Errorf("want %.2f, got %.2f", want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 100, want: 10, errExpected: false, name: "Square root of 100 is 10"},
		{a: 0, want: 0, errExpected: false, name: "Square root of 0 is 0"},
		{a: -1, want: 999, errExpected: true, name: "Square root can only be calculated for positive numbers"},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("Unexpected Error - Expected %v, received %v", tc.errExpected, errReceived)
		}

		if tc.want != got {
			t.Errorf("Error: want %v, got %v", tc.want, got)
		}
	}

}
