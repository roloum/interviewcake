package cafeorders

import (
	"testing"
)

func TestIsFirstComeFirstServed(t *testing.T) {

	cases := [][][]int{
		[][]int{[]int{1, 4, 5}, []int{2, 3, 6}, []int{1, 2, 3, 4, 5, 6}},
		[][]int{[]int{1, 5}, []int{2, 3, 6}, []int{1, 2, 6, 3, 5}},
		[][]int{[]int{}, []int{2, 3, 6}, []int{2, 3, 6}},
		[][]int{[]int{1, 5}, []int{2, 3, 6}, []int{1, 6, 3, 5}},
		[][]int{[]int{1, 5}, []int{2, 3, 6}, []int{1, 2, 3, 5, 6, 8}},
		[][]int{[]int{1, 9}, []int{7, 8}, []int{1, 7, 8}},
		[][]int{[]int{55, 9}, []int{7, 8}, []int{1, 7, 8, 9}},
	}

	results := []bool{true, false, true, false, false, false, false}

	for i, c := range cases {
		r := IsFirstComeFirstServed(c[0], c[1], c[2])
		if r != results[i] {
			t.Fatalf("row: %v, takeOut: %v, dineIn: %v, served: %v, expected: %v, received: %v",
				i+1, c[0], c[1], c[2], results[i], r)
		}
	}

}
