package arrayshuffle

import (
	"fmt"
	"testing"
)

func TestShuffle(t *testing.T) {
	numbers := [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{1},
		[]int{},
	}

	for _, n := range numbers {
		fmt.Println(n)
		Shuffle(n)
		fmt.Println(n)
	}
}
