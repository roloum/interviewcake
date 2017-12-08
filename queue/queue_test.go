package queue

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {

	q := NewQueue()

	if _, err := q.Dequeue(); !reflect.DeepEqual(err, errEmptyStructure) {
		t.Errorf("Unexpected error %v. Expected: %v", err, errEmptyStructure)
	}

	for _, i := range []int{1, 2, 3, 4} {
		if err := q.Enqueue(i); err != nil {
			t.Fatalf("Error when adding element %v to the queue", i)
		}
	}

	for _, i := range []int{1, 2, 3} {
		if v, err := q.Dequeue(); err != nil {
			t.Fatalf("Error when adding element %v to the queue", i)
		} else if v != i {
			t.Errorf("Unexpected result: %v, expected: %v", v, i)
		}
	}

	for _, i := range []int{5, 6, 7, 8, 9} {
		if err := q.Enqueue(i); err != nil {
			t.Fatalf("Error when adding element %v to the queue", i)
		}
	}

	for _, i := range []int{4, 5, 6, 7, 8, 9} {
		if v, err := q.Dequeue(); err != nil {
			t.Fatalf("Error when adding element %v to the queue", i)
		} else if v != i {
			t.Errorf("Unexpected result: %v, expected: %v", v, i)
		}
	}

}

func TestStack(t *testing.T) {
	s := NewStack()
	if _, err := s.Pop(); !reflect.DeepEqual(err, errEmptyStructure) {
		t.Fatalf("Unexpected error when retrieving element from empty stack: %v",
			err)
	}

	ints := []int{1, 2, 3}
	for _, i := range ints {
		if err := s.Push(i); err != nil {
			t.Fatalf("Error pushing element %v to list: %v", i, err)
		}
	}

	for i := len(ints) - 1; i <= 0; i-- {
		if v, err := s.Pop(); err != nil {
			t.Fatalf("Unexpected error when retrieving element from empty stack: %v",
				err)
		} else if v != ints[i] {
			t.Errorf("Unexpected value when retrieving element from stac: %v, expected %v",
				v, ints[i])
		}
	}
}
