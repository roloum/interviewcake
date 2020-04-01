package palindromepermutation

import (
	"testing"
)

func TestHasPalindromePermutation(t *testing.T) {

	cases := map[string]bool{
		"aabcbcd":  true,
		"aabccbdd": true,
		"aabcd":    false,
		"aabbcd":   false,
		"":         true,
		"a":        true,
	}

	for word, result := range cases {
		if r := HasPalindromePermutation(word); r != result {
			t.Errorf("Word: %v, expected: %v", word, result)
		}
	}

}
