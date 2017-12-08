package queue

import (
	"errors"
	"math"
)

var errFullStructure = errors.New("Structure is full")
var errEmptyStructure = errors.New("Structure is empty")

//Stack object
type Stack struct {

	//Array containing the elements
	stack []interface{}

	//Variable used to keep track of the array count
	count int64
}

//NewStack returns an instance of a stack
func NewStack() *Stack {
	s := new(Stack)
	s.init()

	return s
}

//init Initializes the array in the stack
func (s *Stack) init() {
	s.stack = []interface{}{}
	s.count = 0
}

//Push pushes a new element to the stack
func (s *Stack) Push(v interface{}) error {

	//Check if array is full
	if s.count == math.MaxInt64 {
		return errFullStructure
	}

	s.stack = append(s.stack, v)
	s.count++

	return nil
}

//Pop pops an element from the array
func (s *Stack) Pop() (interface{}, error) {

	if s.count == 0 {
		return nil, errEmptyStructure
	}

	//Get last element from the array
	v := s.stack[s.count-1]

	//Remove element
	s.stack = s.stack[:s.count-1]
	s.count--

	return v, nil

}
