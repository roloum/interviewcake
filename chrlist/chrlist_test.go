package chrlist

import (
	"reflect"
	"strings"
	"testing"
)

//TestReverseOne tests reverse function for one character list
func TestReverse(t *testing.T) {
	list := [][]string{
		[]string{"a"},
		[]string{"r", "o", "l", "a", "n", "d"},
		[]string{"r", "o", "l", "a", "n", "d", "o"},
	}
	expected := [][]string{
		[]string{"a"},
		[]string{"d", "n", "a", "l", "o", "r"},
		[]string{"o", "d", "n", "a", "l", "o", "r"},
	}

	for i, l := range list {

		Reverse(l)
		if !reflect.DeepEqual(l, expected[i]) {
			t.Errorf("List %v: %v, expected: %v", i+1, l, expected[i])
		}
	}
}

func TestReverseWords(t *testing.T) {
	list := [][]string{
		strings.Split("vault", ""),
		strings.Split("thief cake", ""),
		strings.Split("one another get", ""),
		strings.Split("rat the ate cat the", ""),
		strings.Split("yummy is cake bundt chocolate", ""),
		strings.Split("", "")}
	expected := [][]string{
		strings.Split("vault", ""),
		strings.Split("cake thief", ""),
		strings.Split("get another one", ""),
		strings.Split("the cat ate the rat", ""),
		strings.Split("chocolate bundt cake is yummy", ""),
		strings.Split("", "")}

	for i, l := range list {

		t.Logf("Case #%v", i+1)
		ReverseWords(l)
		if !reflect.DeepEqual(l, expected[i]) {
			t.Errorf("List %v: %v, expected: %v", i+1, l, expected[i])
		}
	}
}
