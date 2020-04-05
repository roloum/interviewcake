package permutations

import (
	"reflect"
	"sort"
	"testing"
)

func TestGet(t *testing.T) {

	cases := []struct {
		input        string
		permutations []string
	}{
		{"", []string{""}},
		{"a", []string{"a"}},
		{"ab", []string{"ab", "ba"}},
		{"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
	}

	for i, c := range cases {
		result := Get(c.input)
		sort.Strings(result)
		sort.Strings(c.permutations)
		if !reflect.DeepEqual(result, c.permutations) {
			t.Errorf("Row: %v, Input: %v, Expected: %v, Received: %v", i+1, c.input,
				c.permutations, result)
		}
	}
}
