package chrlist

//Reverse reverses a list of characters
func Reverse(list []string) []string {

	l := len(list)
	for i := 0; i < l/2; i++ {
		list[i], list[l-1-i] = list[l-1-i], list[i]
	}

	return list
}
