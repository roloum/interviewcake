package worldcloud

import (
	"strings"
	"unicode"
)

//Data ...
type Data struct {
	input string
	m     map[string]int
}

//WordsToCounts ...
func (d *Data) WordsToCounts() map[string]int {

	d.m = make(map[string]int)

	//First and last index used for determining where word begins/ends
	wordStart, length := 0, 0
	inputLen := len(d.input)

	for i := 0; i <= inputLen; i++ {
		if i == inputLen {
			if length > 0 {
				updateMap(d, wordStart, length)
			}
			break
		}

		if unicode.IsLetter(rune(d.input[i])) {
			if length == 0 {
				wordStart = i
			}
			length++
		} else {

			if length == 0 {
				continue
			}

			chr := string(d.input[i])
			if chr == "-" || chr == "'" {

				if i > 0 && unicode.IsLetter(rune(d.input[i-1])) && i < inputLen-1 &&
					unicode.IsLetter(rune(d.input[i+1])) {
					length++
					continue
				}

			}

			updateMap(d, wordStart, length)
			length = 0

		}

	}

	return d.m
}

func updateMap(d *Data, wordStart, length int) {
	key := string(d.input[wordStart : wordStart+length])

	if _, exists := d.m[key]; exists {
		d.m[key]++
	} else if _, exists := d.m[strings.ToLower(key)]; exists {
		d.m[strings.ToLower(key)]++
	} else if count, exists := d.m[strings.Title(key)]; exists {
		d.m[key] = count + 1
		delete(d.m, strings.Title(key))
	} else {
		d.m[key]++
	}
}
