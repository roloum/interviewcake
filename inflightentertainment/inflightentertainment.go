package inflightentertainment

import "errors"

var errMinElements = errors.New("List needs at least two elements")

//AreThereTwoMoviesWithLength ...
func AreThereTwoMoviesWithLength(list []int, length int) (bool, error) {
	count := len(list)

	if count < 2 {
		return false, errMinElements
	}

	//Unique map where we store the length of movies
	uniqueMap := make(map[int]int)

	//Iterate through half the list and check both sides of the list: i and count-i
	for i := 0; i <= int((count-1)/2); i++ {
		//Add right element to the map if it does not exist
		if uniqueMap[list[count-1-i]] == 0 {
			uniqueMap[list[count-1-i]] = 1
		}
		//Check list[i]. Is there an element in the map that adds up to length
		if uniqueMap[length-list[i]] == 1 {
			return true, nil
		}

		//Add left element to the map if it does not exist
		if uniqueMap[list[i]] == 0 {
			uniqueMap[list[i]] = 1
		}
		//Check list[count-1-i]. Is there an element in the map that adds up to length
		if uniqueMap[length-list[count-1-i]] == 1 {
			return true, nil
		}

	}

	return false, nil
}
