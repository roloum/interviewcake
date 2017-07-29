package hical

//MergeRanges ...
func MergeRanges(meetings [][]int) [][]int {

	var merged [][]int
	meetings = Sort(meetings)

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

//Sort ...
func Sort(meetings [][]int) [][]int {
	length := len(meetings)
	if length < 2 {
		return meetings
	}
	var mid = length / 2
	result := merge(Sort(meetings[:mid]), Sort(meetings[mid:]))
	return result
}

func merge(left, right [][]int) [][]int {
	llen, rlen, l, r, i := len(left), len(right), 0, 0, 0

	var slice = make([][]int, len(left)+len(right))

	for l < llen || r < rlen {
		if (l < llen && r < rlen && left[l][0] <= right[r][0]) || (r == rlen && l < llen) {
			slice[i] = left[l]
			l++
		} else {
			slice[i] = right[r]
			r++
		}
		i++
	}

	return slice
}
