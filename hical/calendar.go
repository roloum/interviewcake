package hical

import (
	"sort"
)

//MergeRanges merges meeting ranges
func MergeRanges(meetings [][]int) [][]int {

	//Store merged meetings in new array
	var merged [][]int

	//Sort meetings
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	var min, max int = meetings[0][0], meetings[0][1]

	for i := 1; i < len(meetings); i++ {
		if meetings[i][0] <= max {
			if meetings[i][1] > max {
				max = meetings[i][1]
			}
		} else {
			merged = append(merged, []int{min, max})
			min, max = meetings[i][0], meetings[i][1]
		}
	}

	merged = append(merged, []int{min, max})

	return merged

}
