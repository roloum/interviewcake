package palindromepermutation

//HasPalindromePermutation checks if any permutation of the input is a palindrome
func HasPalindromePermutation(word string) bool {

	s := make(map[rune]int)

	//Iterate word
	for _, char := range word {
		if _, ok := s[char]; !ok {
			//Character does not exist? add to map
			s[char] = 1
		} else {
			//Character exists? remove from map
			delete(s, char)
		}
	}

	//It is palindrome only if there is 1 char left in the map at most
	return len(s) < 2
}
