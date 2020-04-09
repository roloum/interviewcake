package duplicate

import (
	"testing"
)

func TestFindRepeat(t *testing.T) {

	cases := []struct {
		numbers   []int
		duplicate int
	}{
		{[]int{1, 4, 5, 3, 8, 10, 6, 9, 2, 7, 9}, 9},
		{[]int{1, 1}, 1},
		{[]int{1, 2, 3, 2}, 2},
		{[]int{1, 2, 5, 5, 5, 5}, 5},
		{[]int{4, 1, 4, 8, 3, 2, 7, 6, 5}, 4},
	}

	for i, c := range cases {
		if result := FindRepeat(c.numbers); result != c.duplicate {
			t.Errorf("Row: %v, received: %v, expected: %v", i+1, result, c.duplicate)
		}
	}

}
