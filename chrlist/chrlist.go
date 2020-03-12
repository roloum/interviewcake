package chrlist

//Reverse reverses a list of characters
func Reverse(list []string) []string {

	l := len(list)
	for i := 0; i < l/2; i++ {
		list[i], list[l-1-i] = list[l-1-i], list[i]
	}

	return list
}

//ReverseWords reverses the order of the words in a char list
func ReverseWords(list []string) {
	//Reverse whole list first and we'll get spaces in right place
	l := len(list)

	if l == 0 {
		return
	}

	reverseSubset(list, 0, l-1)

	//Then each word again
	firstWord := 0
	for i := 1; i < l; i++ {
		//end of list or next character is a space
		if i == l-1 || list[i+1] == " " {
			reverseSubset(list, firstWord, i)
			firstWord = i + 2
			i = firstWord
		}
	}
}

func reverseSubset(list []string, i, j int) {
	//median of subset
	m := ((j - i) / 2) + i

	for i <= m {
		list[i], list[j] = list[j], list[i]
		i++
		j--
	}

}
