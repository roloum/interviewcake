package permutations

//Get Generates all permutations of the input string
func Get(input string) []string {

	permutations := []string{}

	length := len(input)

	if length <= 1 {
		return []string{input}
	}

	chars := input[:length-1]
	last := input[length-1:]

	charsPermutations := Get(chars)

	for _, perm := range charsPermutations {
		for i := 0; i <= len(chars); i++ {
			output := perm[0:i] + last + perm[i:]
			permutations = append(permutations, output)
		}
	}

	return permutations
}
