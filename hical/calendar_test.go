package hical

import "testing"

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
		{
			[][]int{{1, 3}, {2, 4}},
			[][]int{{1, 4}},
		},
		{
			[][]int{{5, 6}, {6, 8}},
			[][]int{{5, 8}},
		},
		{
			[][]int{{1, 8}, {2, 5}},
			[][]int{{1, 8}},
		},
		{
			[][]int{{1, 3}, {4, 8}},
			[][]int{{1, 3}, {4, 8}},
		},
		{
			[][]int{{1, 4}, {2, 5}, {5, 8}},
			[][]int{{1, 8}},
		},
		{
			[][]int{{5, 8}, {1, 4}, {6, 8}},
			[][]int{{1, 4}, {5, 8}},
		},
		{
			[][]int{{1, 10}, {2, 5}, {6, 8}, {9, 10}, {10, 12}},
			[][]int{{1, 12}},
		},
		{
			[][]int{{0, 1}, {3, 5}, {4, 8}, {10, 12}, {9, 10}},
			[][]int{{0, 1}, {3, 8}, {9, 12}},
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
