package rotationpoint

import (
	"errors"
	"math"
)

var errMinElements = errors.New("List must contain at least two elements")

//FindRotationPoint ...
func FindRotationPoint(list []string) (int, error) {

	count := len(list)
	if count < 2 {
		return 0, errMinElements
	}

	left, right := 0, count-1
	min := right

	for {

		middle := left + int(math.Ceil(float64((right-left)/2)))

		if list[middle] < list[min] {
			min = middle
		}

		if left == middle || right == middle { //No more elements to check
			return min, nil
		} else if list[left] > list[middle] {
			right = middle
		} else {
			left = middle
		}

	}
}
