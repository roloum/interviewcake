package main

import (
	"container/list"
	"errors"
)

//TempTracker keeps track of registered temperatures
type TempTracker struct {
	list                         *list.List
	min, max, sum, maxOccurrence int
	mode                         []int
	occurrences                  map[int]int
}

//New Returns the pointer to a new instance of TempTracker
func New() *TempTracker {
	var t = new(TempTracker)

	t.occurrences = make(map[int]int)

	t.list = list.New()

	return t
}

//Insert Inserts a new temperature in the list
func (t *TempTracker) Insert(temperature int) error {
	if temperature < 0 || temperature > 110 {
		return errors.New("Temperature is not in valid range")
	}

	t.list.PushBack(temperature)

	length := t.list.Len()

	if length != 1 {
		if temperature < t.min {
			t.min = temperature
		}
		if temperature > t.max {
			t.max = temperature
		}
	} else {
		t.min, t.max = temperature, temperature
	}

	t.sum += temperature

	t.occurrences[temperature]++
	if t.occurrences[temperature] > t.maxOccurrence {
		t.maxOccurrence = t.occurrences[temperature]
		t.mode = []int{temperature}
	} else if t.occurrences[temperature] == t.maxOccurrence {
		t.mode = append(t.mode, temperature)
	}

	return nil
}

//GetMin Returns min temperature
func (t *TempTracker) GetMin() (int, error) {
	if t.list.Len() == 0 {
		return 0, errors.New("List is empty")
	}
	return t.min, nil
}

//GetMax Returns maximum temperature
func (t *TempTracker) GetMax() (int, error) {
	if t.list.Len() == 0 {
		return 0, errors.New("List is empty")
	}
	return t.max, nil
}

//GetMean Returns temperature average
func (t *TempTracker) GetMean() (float64, error) {
	length := t.list.Len()
	if length == 0 {
		return 0, errors.New("List is empty")
	}
	avg := float64(t.sum) / float64(length)
	return avg, nil
}

//GetMode Returns the numbers which appear the most times
func (t *TempTracker) GetMode() ([]int, error) {
	if t.list.Len() == 0 {
		return []int{}, errors.New("List is empty")
	}
	return t.mode, nil
}

func main() {
}
