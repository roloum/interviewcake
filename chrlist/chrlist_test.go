package chrlist

import (
	"reflect"
	"testing"
)

//TestReverseOne tests reverse function for one character
func TestReverseOne(t *testing.T) {
	list, expected := []string{"a"}, []string{"a"}

	result := Reverse(list)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("List: %v, expected: %v, received: %s", list, expected, result)
	}
}

//TestReverseEven tests reverse function for list with even # of characters
func TestReverseEven(t *testing.T) {
	list, expected := []string{"r", "o", "l", "a", "n", "d"},
		[]string{"d", "n", "a", "l", "o", "r"}

	result := Reverse(list)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("List: %v, expected: %v, received: %s", list, expected, result)
	}
}

//TestReverseOdd tests reverse function for list with odd # of characters
func TestReverseOdd(t *testing.T) {
	list, expected := []string{"r", "o", "l", "a", "n", "d", "o"},
		[]string{"o", "d", "n", "a", "l", "o", "r"}

	result := Reverse(list)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("List: %v, expected: %v, received: %s", list, expected, result)
	}
}
