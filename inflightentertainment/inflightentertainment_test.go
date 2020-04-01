package inflightentertainment

import (
	"testing"
)

var testCases = []struct {
	list   []int
	length int
	result bool
}{
	{[]int{8, 7, 5, 11, 10, 9, 9, 6}, 20, true},
	{[]int{3, 6}, 12, false},
	{[]int{8, 5, 11, 10, 9, 9, 6}, 15, true},
	{[]int{8, 7, 5, 11, 10}, 14, false},
	{[]int{8, 7, 5, 11, 10, 9}, 11, false},
	{[]int{1}, 20, false},
	{[]int{2, 4}, 1, false},
	{[]int{2, 4}, 6, true},
	{[]int{3, 8}, 6, false},
	{[]int{3, 8, 3}, 6, true},
	{[]int{1, 2, 3, 4, 5, 6}, 7, true},
	{[]int{4, 3, 2}, 5, true},
	{[]int{6}, 6, false},
	{[]int{}, 2, false},
}

func TestCanTwoMoviesFillFlight(t *testing.T) {
	for _, test := range testCases {
		result := CanTwoMoviesFillFlight(test.list, test.length)
		if result != test.result {
			t.Errorf("Wrong result %v for list: %v", result, test.list)
		}
	}
}
