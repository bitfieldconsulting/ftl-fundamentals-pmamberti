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

type evaluateTestCase struct {
	expression  string
	want        float64
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

	testCases := []variadicTestCase{
		{nums: []float64{3, 2, 1}, want: 6, name: "3 * 2 * 1 = 6"},
		{nums: []float64{100, 8, 12}, want: 9600, name: "100 * 8 * 12 = 9600"},
		{nums: []float64{-2, 8, 10}, want: -160, name: "-1 * 8 * 10 = -160"},
		{nums: []float64{400, 13, 0, 4}, want: 0, name: "400 * 13 * 0 * 4 = 0"},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.nums...)

		if tc.want != got {
			t.Errorf("Error: unexpected value returned - want %.1f, got %.1f", tc.want, got)
		}
	}

}

func TestDivide(t *testing.T) {
	t.Parallel()

	testCases := []variadicTestCase{
		{nums: []float64{7, 0, 10, 44}, want: 999, name: "1 / 8 / -10 = -0.0125", errExpected: true},
		{nums: []float64{3, 2}, want: 1.5, name: "3 / 2 = 1.5", errExpected: false},
		{nums: []float64{100, 8, 12}, want: 1.0416666666666667, name: "100 / 8 / 12 = 1.041666667", errExpected: false},
		{nums: []float64{1, 8, -10}, want: -0.0125, name: "1 / 8 / -10 = -0.0125", errExpected: false},
		{nums: []float64{10}, want: 10, name: "10 - single value passed", errExpected: false},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.nums...)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			fmt.Printf("Expected test: %s\n", tc.name)
			t.Fatalf("Error: unexpected error status: %v)",
				errReceived)
		}

		if !tc.errExpected && tc.want != got {
			fmt.Printf("Expected test: %s\n", tc.name)
			t.Errorf("want %g(%T), got %g(%T)", tc.want, tc.want, got, got)
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

func TestEvaluate(t *testing.T) {
	t.Parallel()

	testCases := []evaluateTestCase{
		{expression: "11 + 7.3", want: 18.3, errExpected: false},
		{expression: "1.1 - 7.3", want: -6.199999999999999, errExpected: false},
		{expression: "11 * 2.5", want: 27.5, errExpected: false},
		{expression: "       11 / 7.3", want: 1.5068493150684932, errExpected: false},
		{expression: " 11      / 0", want: 999, errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Evaluate(tc.expression)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf("Unexpected Error - Expected %v, received %v", tc.errExpected, errReceived)
		}

		if !errReceived && tc.want != got {
			t.Errorf("Error: want %v, got %v", tc.want, got)
		}
	}
}
