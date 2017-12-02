package inflightentertainment

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	list   []int
	length int
	result bool
	err    error
}{
	{[]int{8, 7, 5, 11, 10, 9, 9, 6}, 20, true, nil},
	{[]int{8, 5, 11, 10, 9, 9, 6}, 15, true, nil},
	{[]int{8, 7, 5, 11, 10}, 14, false, nil},
	{[]int{8, 7, 5, 11, 10, 9}, 11, false, nil},
	{[]int{1}, 20, false, errMinElements},
}

func TestAreThereTwoMoviesWithLength(t *testing.T) {
	for _, test := range testCases {
		if result, err := AreThereTwoMoviesWithLength(test.list,
			test.length); !reflect.DeepEqual(err, test.err) {
			t.Errorf("Unexpected error: %v. Expected: %v. List: %v. Length: %v",
				err, test.err, test.list, test.length)
		} else if result != test.result {
			t.Errorf("Wrong result %v for list: %v", result, test.list)
		}
	}
}
