package findinset

import (
	"errors"
	"math"
)

var errNotFound = errors.New("Element does not exist")

//FindElement Finds an element in a list of n integers sorted in ascending order
//Using an interative approach
func FindElement(list []int, element int) (int, error) {

	count := len(list)
	if count == 0 {
		return 0, errors.New("List is empty")
	}

	//If element is < first element or > last element return immediately
	if element < list[0] || element > list[count-1] {
		return 0, errNotFound
	}

	left, right := 0, count-1

	for {

		mid := left + int(math.Ceil(float64(right-left)/2))

		if list[mid] == element { //Found the element
			return element, nil
		} else if left == right { //No more elements to check, not found
			return 0, errNotFound
		} else if element > list[mid] { //Look to the right
			left = mid + 1
		} else if element < list[mid] { //Look to the left
			right = mid - 1
		}

	}
}

//FindElementRecursive Finds element in asc sorted list of ints recursively
func FindElementRecursive(list []int, element int) (int, error) {

	count := len(list)
	if count == 0 {
		return 0, errors.New("List is empty")
	}

	//If element is < first element or > last element return immediately
	if element < list[0] || element > list[count-1] {
		return 0, errNotFound
	}

	return findElementHelper(list, element)
}

//findElementHelper recursively finds an element in a asc sorted int list
func findElementHelper(list []int, element int) (int, error) {

	length := len(list)
	if length > 0 {

		//Found element
		if middle := int(length / 2); list[middle] == element {
			return element, nil
		} else if element > list[middle] { //Look to the right
			return findElementHelper(list[middle+1:], element)
		} else { //Look to the left
			return findElementHelper(list[:middle], element)
		}
	}
	return 0, errNotFound

}
