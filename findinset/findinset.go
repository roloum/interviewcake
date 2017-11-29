package findinset

import "errors"

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

	leftIndex, rightIndex, middle := 0, count-1, int(count/2)

	for {
		if list[middle] == element { //Found element
			return element, nil
		} else if leftIndex == middle { //No more elements to check, errNotFound
			return 0, errNotFound
		} else if list[middle] > element { //Find in right portion or array
			middle, leftIndex = int((rightIndex-leftIndex)/2), middle
		} else { //Find in left portion of array
			middle, rightIndex = int((rightIndex-leftIndex)/2), middle
		}
	}

}

//FindElementRecursive Finds element in asc sorted list of ints recursively
//Using a helper function that receives a receives the reference of the array
//So it does not make additional copies of the data structure
func FindElementRecursive(list []int, element int) (int, error) {

	count := len(list)
	if count == 0 {
		return 0, errors.New("List is empty")
	}

	//If element is < first element or > last element return immediately
	if element < list[0] || element > list[count-1] {
		return 0, errNotFound
	}

	return findElementHelper(&list, element, 0, int(count/2), count-1)
}

//findElementHelper recursively finds an element in a asc sorted int list
func findElementHelper(list *[]int, element, leftIndex, middle,
	rightIndex int) (int, error) {

	if (*list)[middle] == element { //Found?
		return element, nil
	} else if leftIndex == middle { //No more elements to check
		return 0, errNotFound
	} else if (*list)[middle] > element { //Search right side
		return findElementHelper(list, element, middle,
			int((rightIndex-leftIndex)/2), rightIndex)
	}
	//Search left side
	return findElementHelper(list, element, leftIndex,
		int((rightIndex-leftIndex)/2), middle)

}
