package cakethief

import (
	"reflect"
	"testing"
)

func TestMaxDuffelBagValue(t *testing.T) {
	for _, test := range []struct {
		list            [][]int
		capacity, value int
		err             error
	}{
		{[][]int{
			[]int{7, 160},
			[]int{4, 110},
			[]int{3, 90},
			[]int{2, 15},
		}, 20, 555, nil},
		{[][]int{
			[]int{7, 160},
			[]int{4, 110},
			[]int{3, 90},
		}, 20, 550, nil},
		{[][]int{
			[]int{7, 160},
			[]int{3, 90},
			[]int{2, 15},
		}, 20, 555, nil},
		{[][]int{}, 20, 0, errEmptyList},
		{[][]int{[]int{0, 100}}, 20, 0, errCeroValue},
		{[][]int{[]int{-1, -100}}, 20, 0, errNegativeValue},
	} {
		if value, err := MaxDuffelBagValue(test.list, test.capacity); !reflect.DeepEqual(
			err, test.err) {
			t.Errorf("Unexpected error: %v. List: %v", err, test.list)
		} else if value != test.value {
			t.Errorf("Unexpected value: %v. List: %v, expected: %v", value, test.list,
				test.value)
		}
	}
}
