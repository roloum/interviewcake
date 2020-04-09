package duplicate

//FindRepeat Finds a number that appears more than once
//Can't sort in place
func FindRepeat(numbers []int) int {

	length := len(numbers)

	floor, ceiling := 1, length-1

	for floor < ceiling {

		mid := floor + (ceiling-floor)/2
		leftFloor := floor
		leftCeiling := mid
		rightFloor := mid + 1
		rightCeiling := ceiling

		uniqueIntsLeft := leftCeiling - leftFloor + 1

		numbersInLeft := 0
		for _, n := range numbers {
			if n >= leftFloor && n <= leftCeiling {
				numbersInLeft++
			}
		}

		if numbersInLeft > uniqueIntsLeft {
			floor, ceiling = leftFloor, leftCeiling
		} else {
			floor, ceiling = rightFloor, rightCeiling
		}
	}

	return floor
}
