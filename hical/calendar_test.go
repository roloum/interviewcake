package hical

import "testing"

func TestSort(t *testing.T) {
	cases := []struct {
		in, out [][]int
	}{
		{
			[][]int{{2, 3}, {8, 10}, {1, 2}, {0, 1}},
			[][]int{{0, 1}, {1, 2}, {2, 3}, {8, 10}},
		},
		{
			[][]int{{4, 6}, {10, 12}, {1, 3}},
			[][]int{{1, 3}, {4, 6}, {10, 12}},
		},
	}

	for _, c := range cases {
		result := Sort(c.in)
		length := len(result)
		if length == 0 {
			t.Errorf("Error on input %v, result is empty", c.in)
		} else if length != len(c.out) {
			t.Errorf("Different size. Input: %v, result: %v, expected: %v", c.in, result, c.out)
		}
		for i, meeting := range result {
			if meeting[0] != c.out[i][0] || meeting[1] != c.out[i][1] {
				t.Errorf("Error sorting array. Input: %v, result: %v, expected: %v", c.in, result, c.out)
				break
			}
		}
	}
}

func TestMergeRanges(t *testing.T) {
	cases := []struct {
		meetings, ranges [][]int
	}{
		{
			[][]int{{0, 1}, {3, 5}, {4, 8}, {9, 10}, {10, 12}},
			[][]int{{0, 1}, {3, 8}, {9, 12}},
		},
		{
			[][]int{{0, 2}, {10, 12}, {7, 8}, {1, 3}, {4, 6}, {1, 2}, {5, 7}, {6, 9}},
			[][]int{{0, 3}, {4, 9}, {10, 12}},
		},
		{
			[][]int{{4, 6}, {10, 12}, {1, 3}},
			[][]int{{1, 3}, {4, 6}, {10, 12}},
		},
	}

	for _, c := range cases {
		result := MergeRanges(c.meetings)
		if len(result) == 0 {
			t.Errorf("Error on input %v, result is empty", c.meetings)
		}
		for i, meeting := range result {
			if meeting[0] != c.ranges[i][0] || meeting[1] != c.ranges[i][1] {
				t.Errorf("Error on input %v, received: %v, expected %v", c.meetings, result, c.ranges)
				break
			}
		}
	}
}
