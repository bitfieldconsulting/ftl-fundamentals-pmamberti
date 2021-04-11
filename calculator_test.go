package calculator_test

import (
	"calculator"
	"math"
	"math/rand"
	"testing"
)

func TestCloseEnough(t *testing.T) {
	testCases := []struct {
		a, b, roundedResult, tolerance float64
		want                           bool
	}{
		{a: 1.66666666, b: 1.6666, tolerance: 0.0001},
		{a: 0.33333333, b: 0.333, tolerance: 0.001},
		{a: 0.66666666, b: 0.66, tolerance: 0.01},
	}

	for _, tc := range testCases {
		if !closeEnough(tc.a, tc.b, tc.tolerance) {
			t.Errorf(
				"Result(%g) is not close enough(%.5f)",
				tc.b,
				tc.tolerance,
			)
		}
	}
}

func TestAddSubtractMultiply(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		operator   func(float64, float64, ...float64) float64
		a, b, want float64
		nums       []float64
		name       string
	}{
		{
			name:     "5 + 6 = 11. Empty nums.",
			operator: calculator.Add,
			a:        5,
			b:        6,
			nums:     []float64{},
			want:     11,
		},
		{
			name:     "0 + 0 = 0. Empty nums.",
			operator: calculator.Add,
			a:        0,
			b:        0,
			nums:     []float64{},
			want:     0,
		},
		{
			name:     "1 + 2 + 3 = 6",
			operator: calculator.Add,
			a:        1,
			b:        2,
			nums:     []float64{3},
			want:     6,
		},
		{
			name:     "1 + 2 + 3 + 4 = 10",
			operator: calculator.Add,
			a:        1,
			b:        2,
			nums:     []float64{3, 4},
			want:     10,
		},
		{
			name:     "5 - 5 = 0. Empty nums.",
			operator: calculator.Subtract,
			a:        5,
			b:        5,
			nums:     []float64{},
			want:     0,
		},
		{
			name:     "3 - 2 - 1 = 0",
			operator: calculator.Subtract,
			a:        3,
			b:        2,
			nums:     []float64{1},
			want:     0,
		},
		{
			name:     "100 - 8 - 12 = 79",
			operator: calculator.Subtract,
			a:        100,
			b:        8,
			nums:     []float64{1, 12},
			want:     79,
		},
		{
			name:     "1 - 8 - 10 = -17",
			operator: calculator.Subtract,
			a:        1,
			b:        8,
			nums:     []float64{10},
			want:     -17,
		},
		{
			name:     "4 * 11 = 44. Empty nums.",
			operator: calculator.Multiply,
			a:        4,
			b:        11,
			nums:     []float64{},
			want:     44,
		},
		{
			name:     "100 * 8 * 12 = 9600",
			operator: calculator.Multiply,
			a:        100,
			b:        8,
			nums:     []float64{12},
			want:     9600,
		},
		{
			name:     "3 * 2 * -10 = -60",
			operator: calculator.Multiply,
			a:        3,
			b:        2,
			nums:     []float64{1, -10},
			want:     -60,
		},
		{
			name:     "0 * 10 * 55 * 10 * 2 = 0. 0 passed as first parameter.",
			operator: calculator.Multiply,
			a:        0,
			b:        10,
			nums:     []float64{55, 10, 2},
			want:     0,
		},
		{
			name:     "400 * 13 * 0 * 4 = 0",
			operator: calculator.Multiply,
			a:        400,
			b:        13,
			nums:     []float64{0, 4},
			want:     0,
		},
	}

	for _, tc := range testCases {
		got := tc.operator(tc.a, tc.b, tc.nums...)

		if tc.want != got {
			t.Errorf(
				"%v - want %.1f, got %.1f",
				tc.name,
				tc.want,
				got,
			)
		}
	}
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		a, b, want float64
		nums       []float64
		name       string
	}{
		{
			name: "5 + 6 = 11. Empty nums.",
			a:    5,
			b:    6,
			nums: []float64{},
			want: 11,
		},
		{
			name: "1 + 2 + 3 = 6",
			a:    1,
			b:    2,
			nums: []float64{3},
			want: 6,
		},
		{
			name: "1 + 2 + 3 + 4 = 10",
			a:    1,
			b:    2,
			nums: []float64{3, 4},
			want: 10,
		},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b, tc.nums...)

		if tc.want != got {
			t.Errorf(
				"%v - want %.1f, got %.1f",
				tc.name,
				tc.want,
				got,
			)
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
			t.Errorf(
				"%g + %g - want: %.2f, got %.2f",
				a,
				b,
				want,
				got,
			)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		a, b, want float64
		nums       []float64
		name       string
	}{
		{
			name: "5 - 5 = 0. Empty nums.",
			a:    5,
			b:    5,
			nums: []float64{},
			want: 0,
		},
		{
			name: "3 - 2 - 1 = 0",
			a:    3,
			b:    2,
			nums: []float64{1},
			want: 0,
		},
		{
			name: "100 - 8 - 12 = 79",
			a:    100,
			b:    8,
			nums: []float64{1, 12},
			want: 79,
		},
		{
			name: "1 - 8 - 10 = -17",
			a:    1,
			b:    8,
			nums: []float64{10},
			want: -17,
		},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b, tc.nums...)

		if tc.want != got {
			t.Errorf(
				"%v - want %.1f, got %.1f",
				tc.name,
				tc.want,
				got,
			)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		a, b, want float64
		nums       []float64
		name       string
	}{
		{
			name: "4 * 11 = 44. Empty nums.",
			a:    4,
			b:    11,
			nums: []float64{},
			want: 44,
		},
		{
			name: "100 * 8 * 12 = 9600",
			a:    100,
			b:    8,
			nums: []float64{12},
			want: 9600,
		},
		{
			name: "3 * 2 * -10 = -60",
			a:    3,
			b:    2,
			nums: []float64{1, -10},
			want: -60,
		},
		{
			name: "0 * 10 * 55 * 10 * 2 = 0. 0 passed as first parameter.",
			a:    0,
			b:    10,
			nums: []float64{55, 10, 2},
			want: 0,
		},
		{
			name: "400 * 13 * 0 * 4 = 0",
			a:    400,
			b:    13,
			nums: []float64{0, 4},
			want: 0,
		},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b, tc.nums...)

		if tc.want != got {
			t.Errorf(
				"%v - want %.1f, got %.1f",
				tc.name,
				tc.want,
				got,
			)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		a, b, want  float64
		nums        []float64
		name        string
		errExpected bool
	}{
		{
			name:        "1 / 8 / -10 = -0.0125",
			a:           7,
			b:           0,
			nums:        []float64{10, 44},
			want:        0,
			errExpected: true,
		},
		{
			name:        "3 / 2 = 1.5",
			a:           3,
			b:           2,
			nums:        []float64{},
			want:        1.5,
			errExpected: false,
		},
		{
			name:        "Rational value that doesn't have an exact floating-point representation",
			a:           2,
			b:           3,
			nums:        []float64{},
			want:        0.6666666666666666,
			errExpected: false,
		},
		{
			name:        "100 / 8 / 12 = 1.041666667",
			a:           100,
			b:           8,
			nums:        []float64{12},
			want:        1.0416666666666667,
			errExpected: false,
		},
		{
			name:        "1 / 8 / -10 = -0.0125",
			a:           1,
			b:           8,
			nums:        []float64{-10, 2},
			want:        -0.00625,
			errExpected: false,
		},
		{
			name:        "10 / 1 = 10",
			a:           10,
			b:           1,
			nums:        []float64{1},
			want:        10,
			errExpected: false,
		},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b, tc.nums...)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf(
				"%v - Unexpected error status: %v)",
				tc.name,
				errReceived,
			)
		}

		if !tc.errExpected && tc.want != got {
			t.Errorf("%v - want %g, got %g", tc.name, tc.want, got)
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
			t.Fatalf(
				"%g / %g - Unexpected Error Status: %v",
				a,
				b,
				errReceived,
			)
		}

		want := a / b

		if !errReceived && want != got {
			t.Errorf("%g, %g, want %.2f, got %.2f", a, b, want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		a, want, tolerance float64
		errExpected        bool
		name               string
	}{
		{
			name:        "Square root of 100 is 10",
			a:           100,
			want:        10,
			errExpected: false,
		},
		{
			name:        "Square root of 0 is 0",
			a:           0,
			want:        0,
			errExpected: false,
		},
		{
			name:        "Square root can only be calculated for positive numbers",
			a:           -1,
			want:        0,
			errExpected: true,
		},
		{
			name:        "Floating point rounding for GoodEnough",
			a:           3,
			want:        1.7325,
			errExpected: false,
			tolerance:   0.001,
		},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf(
				"Unexpected Error - Expected %v, received %v",
				tc.errExpected,
				errReceived,
			)
		}

		if !errReceived && !closeEnough(tc.want, got, tc.tolerance) {
			t.Errorf(
				"Sqrt(%g) -  Result(%g) is not close enough(%.9f)",
				tc.a,
				tc.want,
				tc.tolerance,
			)
		}
	}
}

func TestEvaluate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		expression  string
		want        float64
		errExpected bool
	}{
		{expression: "11 + 7.3", want: 18.3, errExpected: false},
		{
			expression:  "1.1 - 7.3",
			want:        -6.199999999999999,
			errExpected: false,
		},
		{expression: "11 * 2.5", want: 27.5, errExpected: false},
		{
			expression:  "       11 / 7.3",
			want:        1.5068493150684932,
			errExpected: false,
		},
		{expression: " 11      / 0", want: 999, errExpected: true},
		{expression: " 11      a 0", want: 999, errExpected: true},
	}

	for _, tc := range testCases {
		got, err := calculator.Evaluate(tc.expression)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf(
				"Unexpected Error - Expected %v, received %v",
				tc.errExpected,
				errReceived,
			)
		}

		if !errReceived && tc.want != got {
			t.Errorf(
				"%v - want %v, got %v",
				tc.expression,
				tc.want,
				got,
			)
		}
	}
}

func closeEnough(
	roundResult, divisionResult, tolerance float64,
) bool {
	return math.Abs(roundResult-divisionResult) <= tolerance
}
