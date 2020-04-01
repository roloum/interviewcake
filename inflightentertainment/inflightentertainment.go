package inflightentertainment

//CanTwoMoviesFillFlight returns boolean indicating whether there are two numbers
//in movieLenghts whose sum equals flightLength
func CanTwoMoviesFillFlight(movieLengths []int, flightLength int) bool {

	s := make(map[int]int)

	for _, length := range movieLengths {
		if _, ok := s[flightLength-length]; ok {
			return true
		}
		s[length] = 1
	}

	return false
}
