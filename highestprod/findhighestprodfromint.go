package highestprod

//FindHighestProductFrom3Ints returns the highest product of 3 integers from an array
//Time: o(n), n=size of numbers
//Space: o(1)
func FindHighestProductFrom3Ints(numbers []int) int {
	if len(numbers) < 3 {
		panic("Array of ints needs at least 3 numbers")
	}

	var highest, lowest int = numbers[0], numbers[0]

	firstProd2 := numbers[0] * numbers[1]
	var lowest2, highest2, highest3 int = firstProd2, firstProd2, firstProd2 * numbers[2]

	for i := 1; i < len(numbers); i++ {
		current := numbers[i]

		//check product of 3
		if i > 1 {
			curLowProd3 := lowest2 * current
			curHighProd3 := highest2 * current

			if curHighProd3 > highest3 {
				highest3 = curHighProd3
			}
			if curLowProd3 > highest3 {
				highest3 = curLowProd3
			}
		}

		curLowProd2 := lowest * current
		curHighProd2 := highest * current

		if curHighProd2 > curLowProd2 {
			if curHighProd2 > highest2 {
				highest2 = curHighProd2
			}
			if curLowProd2 < lowest2 {
				lowest2 = curLowProd2
			}
		} else {
			if curLowProd2 > highest2 {
				highest2 = curLowProd2
			}
			if curHighProd2 < lowest2 {
				lowest2 = curHighProd2
			}
		}

		if current < lowest {
			lowest = current
		}
		if current > highest {
			highest = current
		}

	}
	return highest3
}

//FindHighestProductFromKInts returns the highest product of k integers from an array
//Time: o(n*k), n=size of numbers, k=number of ints to multiply
//space: o(k)
func FindHighestProductFromKInts(numbers []int, k int) int {
	if k < 3 {
		panic("Array of ints needs at least 3 numbers")
	}
	length := len(numbers)
	if length < k {
		panic("Length of array is less than k")
	}

	var highestK int
	var lowHigh = make([][]int, k-1)

	//index constants for low and high values in the 2d array
	const low = 0
	const high = 1

	//initialize low high array
	for i := 0; i < k-1; i++ {
		lowHigh[i] = make([]int, 2)
		if i > 0 {
			lowHigh[i][low] = lowHigh[i-1][low] * numbers[i]
			lowHigh[i][high] = lowHigh[i-1][high] * numbers[i]
		} else {
			lowHigh[i][low], lowHigh[i][high] = numbers[i], numbers[i]
		}
	}
	//initialize highestK
	highestK = lowHigh[k-2][high] * numbers[k-1]

	for i := 1; i < length; i++ {
		current := numbers[i]

		if i >= k-1 {
			curLowProdK := lowHigh[k-2][low] * current
			curHighProdK := lowHigh[k-2][high] * current

			if curLowProdK > highestK {
				highestK = curLowProdK
			}
			if curHighProdK > highestK {
				highestK = curHighProdK
			}
		}

		for tempK := k - 2; tempK >= 0; tempK-- {
			//low2-high2, low3-high2, ... low(k-1)-high(k-1)
			if tempK > 0 {
				curLowProd := lowHigh[tempK-1][low] * current
				curHighProd := lowHigh[tempK-1][high] * current

				if curHighProd > curLowProd {
					if curHighProd > lowHigh[tempK][high] {
						lowHigh[tempK][high] = curHighProd
					}
					if curLowProd < lowHigh[tempK][low] {
						lowHigh[tempK][low] = curLowProd
					}
				} else {
					if curLowProd > lowHigh[tempK][high] {
						lowHigh[tempK][high] = curLowProd
					}
					if curHighProd < lowHigh[tempK][low] {
						lowHigh[tempK][low] = curHighProd
					}
				}
				//lowest and highest value
			} else {
				if current < lowHigh[tempK][low] {
					lowHigh[tempK][low] = current
				}
				if current > lowHigh[tempK][high] {
					lowHigh[tempK][high] = current
				}
			}
		}
	}

	return highestK
}
