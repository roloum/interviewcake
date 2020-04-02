package topscore

import (
	"reflect"
	"testing"
)

func TestSortScores(t *testing.T) {

	cases := []struct {
		scores   []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{55}, []int{55}},
		{[]int{30, 60}, []int{60, 30}},
		{[]int{37, 89, 41, 65, 91, 53}, []int{91, 89, 65, 53, 41, 37}},
		{[]int{20, 10, 30, 30, 10, 20}, []int{30, 30, 20, 20, 10, 10}},
	}

	for i, c := range cases {
		result := SortScores(c.scores, 100)
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("Row: %v, scores: %v, expected: %v", i+1, c.scores, result)
		}
	}
}
