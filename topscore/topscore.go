package topscore

//SortScores Sort the scores in O(n) time
func SortScores(scores []int, highestPossible int) []int {

	length := len(scores)
	sorted := make([]int, length, length)

	m := make(map[int]int)

	for _, score := range scores {
		m[score]++
	}

	i := 0
	for score := highestPossible; score >= 0; score-- {
		count := m[score]

		for j := 0; j < count; j++ {
			sorted[i] = score
			i++
		}

	}

	return sorted
}
