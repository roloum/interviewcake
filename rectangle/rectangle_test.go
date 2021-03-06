package rectangle

import (
	"errors"
	"reflect"
	"testing"
)

func TestFindRectancularIntersection(t *testing.T) {
	cases := []struct {
		rectangle1, rectangle2, expected map[string]int
		err                              error
	}{
		{
			map[string]int{"leftX": -1, "bottomY": -3, "width": 3, "height": 2},
			map[string]int{"leftX": -2, "bottomY": -2, "width": 3, "height": 4},
			map[string]int{"leftX": -1, "bottomY": -2, "width": 2, "height": 1},
			nil,
		},
		{
			map[string]int{"leftX": -2, "bottomY": 1, "width": 1, "height": 1},
			map[string]int{"leftX": 1, "bottomY": -2, "width": 2, "height": 4},
			map[string]int{},
			errors.New("No overlap"),
		},
		{
			map[string]int{"leftX": -2, "bottomY": 1, "width": 2, "height": 1},
			map[string]int{"leftX": 0, "bottomY": 1, "width": 2, "height": 1},
			map[string]int{},
			errors.New("No overlap"),
		},
		{
			map[string]int{"leftX": -2, "bottomY": -1, "width": 2, "height": 1},
			map[string]int{"leftX": -1, "bottomY": 0, "width": 2, "height": 1},
			map[string]int{},
			errors.New("No overlap"),
		},
	}

	for _, c := range cases {
		result, err := FindRectancularIntersection(c.rectangle1, c.rectangle2)
		t.Logf("%v => %T, %v => %T", c.err, c.err, err, err)
		if reflect.DeepEqual(c.err, err) != true {
			t.Errorf("Error returned. %v", err)
		} else if reflect.DeepEqual(c.expected, result) != true {

			t.Errorf("Error on input rec1: %v, rec2: %v. Expected: %v, received: %v",
				c.rectangle1, c.rectangle2, c.expected, result)
		}
	}
}
