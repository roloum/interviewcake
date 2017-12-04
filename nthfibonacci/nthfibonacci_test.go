package nthfibonacci

import (
	"reflect"
	"testing"
)

func TestNthFibonacci(t *testing.T) {

	for _, test := range []struct {
		n, result int
		err       error
	}{
		{3, 2, nil},
		{0, 0, nil},
		{1, 1, nil},
		{-1, 0, errNegativeInt},
		{10, 55, nil},
		{7, 13, nil}} {
		if result, err := nthFibonacci(test.n); !reflect.DeepEqual(err, test.err) {
			t.Errorf("Unexpected error: %v, n=%v", err, test.n)
		} else if result != test.result {
			t.Errorf("Unexpected result: %v, n=%v, expected: %v", result, test.n,
				test.result)
		}
	}
}
