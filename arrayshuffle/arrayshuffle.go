package arrayshuffle

import (
	"math/rand"
	"time"
)

//Shuffle Shuffles an array in place
func Shuffle(numbers []int) {
	length := len(numbers)

	if length < 2 {
		return
	}

	rand.Seed(time.Now().UnixNano())

	max := length - 1
	for i := 0; i < length-1; i++ {
		min := i
		limit := max - min + 1
		random := rand.Intn(limit) + i

		numbers[i], numbers[random] = numbers[random], numbers[i]
	}

}
