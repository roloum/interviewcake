package cookies

func mergeLists(mylist, aliciasList []int) []int {

	//mylen := myList length, alen := aliciasList length
	mylen, alen := len(mylist), len(aliciasList)
	l := mylen + alen
	merged := make([]int, l, l)
	var i, j, idx int

	for i < mylen || j < alen {
		if i < mylen && (j == alen || mylist[i] < aliciasList[j]) {
			merged[idx] = mylist[i]
			i++
		} else {
			merged[idx] = aliciasList[j]
			j++
		}
		idx++
	}

	return merged
}
