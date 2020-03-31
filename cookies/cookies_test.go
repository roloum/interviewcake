package cookies

import (
	"reflect"
	"testing"
)

func TestMergeLists(t *testing.T) {

	empty := make([]int, 0, 0)

	cases := [][][]int{
		[][]int{empty, empty},
		[][]int{empty, []int{1, 2, 3}},
		[][]int{[]int{5, 6, 7}, empty},
		[][]int{[]int{2, 4, 6}, []int{1, 3, 7}},
		[][]int{[]int{2, 4, 6, 8}, []int{1, 7}},
	}
	results := [][]int{
		empty,
		[]int{1, 2, 3},
		[]int{5, 6, 7},
		[]int{1, 2, 3, 4, 6, 7},
		[]int{1, 2, 4, 6, 7, 8},
	}

	for i, c := range cases {
		t.Logf("myList: %v, Alicia's: %v", c[0], c[1])
		r := mergeLists(c[0], c[1])
		if !reflect.DeepEqual(results[i], r) {
			t.Errorf("Mine: %v, Alicia's: %v, expected: %v, result: %v", c[0], c[1],
				results[i], r)
		}
	}

}
