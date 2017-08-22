package main

import (
	"errors"
	"reflect"
	"testing"
)

type TestCase struct {
	list, mode []int
	min, max   int
	mean       float64
	err        error
}

//Test all type of valid inputs
func TestTempTracker(t *testing.T) {
	var cases = []TestCase{
		{[]int{3, 6, 1, 3, 1, 3}, []int{3}, 1, 6, float64(17) / float64(6), nil},
		{[]int{1, 1, 3, 6, 3, 1, 3}, []int{1, 3}, 1, 6, float64(18) / float64(7), nil},
	}

LoopCases:
	for _, c := range cases {
		var tracker = New()
		for _, temperature := range c.list {
			err := tracker.Insert(temperature)
			if err != nil {
				t.Errorf("Error inserting temp: %v from list: %v", temperature, c.list)
				break LoopCases
			}
		}

		if min, err := tracker.GetMin(); err != nil || min != c.min {
			t.Errorf(
				"Error getting minimum temperature. Expected: %v, received: %v, error: %v",
				c.min, min, err)
			break
		}

		if max, err := tracker.GetMax(); err != nil || max != c.max {
			t.Errorf(
				"Error getting maximum temperature. Expected: %v, received: %v, error: %v",
				c.max, max, err)
			break
		}

		if mean, err := tracker.GetMean(); err != nil || mean != c.mean {
			t.Errorf(
				"Error getting average temperature. Expected: %v, received: %v, error: %v",
				c.mean, mean, err)
			break
		}

		mode, err := tracker.GetMode()
		if len(mode) != len(c.mode) {
			t.Errorf(
				"Error getting mode. Expected: %v, received: %v, error: %v",
				c.mode, mode, err)
			break LoopCases
		}

		if err == nil {
			for i := 0; i < len(mode); i++ {
				if mode[i] != c.mode[i] {
					t.Errorf(
						"Error getting mode. Expected: %v, received: %v, error: %v",
						c.mode, mode, err)
					break LoopCases
				}
			}
		}

	}
}

//Test getting min, max, mean from empty list
func TestTempTrackerEmptyList(t *testing.T) {
	tracker := New()
	_, err := tracker.GetMax()
	if reflect.DeepEqual(err, errors.New("List is empty")) != true {
		t.Errorf("Unexpected error from GetMax: %v", err)
	}

	_, err = tracker.GetMin()
	if reflect.DeepEqual(err, errors.New("List is empty")) != true {
		t.Errorf("Unexpected error from GetMin: %v", err)
	}

	_, err = tracker.GetMean()
	if reflect.DeepEqual(err, errors.New("List is empty")) != true {
		t.Errorf("Unexpected error from GetMean: %v", err)
	}

	_, err = tracker.GetMode()
	if reflect.DeepEqual(err, errors.New("List is empty")) != true {
		t.Errorf("Unexpected error from GetMode: %v", err)
	}

}

//Test inputs might cause exception like empty list and negative temperatures
func TestTempTrackerWrongTemperature(t *testing.T) {
	var tracker = New()
	temps := []int{-1, 111}
	for _, temp := range temps {
		err := tracker.Insert(temp)
		if reflect.DeepEqual(err, errors.New("Temperature is not in valid range")) != true {
			t.Errorf("Unexpected result when inserting invalid temperature: %v, result: %v", temp, err)
			break
		}
	}

}
