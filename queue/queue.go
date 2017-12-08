package queue

//Queue object made out of two stacks, FIFO behavior
type Queue struct {

	//All new elements are pushed to the in stack
	in *Stack

	//When dequeue is called, the elements are popped from the out stack, unless
	//it is empty, in which case the elements of the in stack will be popped and
	//pushed in the out stack
	out *Stack
}

//NewQueue Creates an instance of Queue, initializes both stacks and returns
//the new queue
func NewQueue() *Queue {
	q := new(Queue)

	q.in = NewStack()
	q.out = NewStack()

	return q
}

//Enqueue adds an element to the in queue
func (q *Queue) Enqueue(v interface{}) error {

	return q.in.Push(v)
}

//Dequeue returns the first inserted element
func (q *Queue) Dequeue() (interface{}, error) {

	//If out stack is empty, insert the elements from the in stack into the out stack
	if q.out.count == 0 {

		for q.in.count != 0 {

			v, err := q.in.Pop()
			if err != nil {
				return nil, err
			}

			if err := q.out.Push(v); err != nil {
				return nil, err
			}
		}

	}

	return q.out.Pop()

}
