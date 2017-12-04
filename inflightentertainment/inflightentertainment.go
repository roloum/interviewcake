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

	for _, firstMovieLength := range list {
		secondMovieLength := length - firstMovieLength

		if uniqueMap[secondMovieLength] == 1 {
			return true, nil
		}

		uniqueMap[secondMovieLength] = 1

	}

	return false, nil
}
