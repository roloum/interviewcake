package findinset

import (
	"errors"
	"reflect"
	"testing"
)

var testCases = []struct {
	list    []int
	element int
	err     error
	result  int
}{
	{[]int{1, 3, 5, 7, 9}, 5, nil, 5},
	{[]int{1}, 1, nil, 1},
	{[]int{1}, 2, errNotFound, 0},
	{[]int{1, 2}, 1, nil, 1},
	{[]int{1, 2}, 3, errNotFound, 0},
	{[]int{1, 3, 5, 7, 9}, 6, errNotFound, 0},
	{[]int{1, 3, 5, 7, 9}, -1, errNotFound, 0},
	{[]int{}, 0, errors.New("List is empty"), 0},
}

func TestFindElement(t *testing.T) {
	for _, test := range testCases {
		if result, err := FindElement(test.list, test.element); !reflect.DeepEqual(err, test.err) {
			t.Errorf("Unexpected error %v, list: %v, element: %v, expected: %v",
				err, test.list, test.element, test.err)
		} else if result != test.result {
			t.Errorf("Unexpected result %v, list: %v, element: %v",
				result, test.list, test.element)
		}
	}
}

func TestFindElementRecursive(t *testing.T) {
	for _, test := range testCases {
		if result, err := FindElementRecursive(test.list, test.element); !reflect.DeepEqual(err, test.err) {
			t.Errorf("Unexpected error %v, list: %v, element: %v, expected: %v",
				err, test.list, test.element, test.err)
		} else if result != test.result {
			t.Errorf("Unexpected result %v, list: %v, element: %v",
				result, test.list, test.element)
		}
	}
}
