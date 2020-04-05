package tree

import (
	"math"

	"github.com/golang-collections/collections/queue"
)

//IsBinarySearch Determines if the tree is a valid binary search tree
func IsBinarySearch(root *Node) bool {

	//struct for the queue
	type node struct {
		node     *Node
		min, max int
	}

	//create queue and enqueue root
	q := queue.New()
	q.Enqueue(node{root, math.MinInt64, math.MaxInt64})

	//Until there are no elements in the queue
	for q.Len() > 0 {

		//Dequeue node
		n := q.Dequeue().(node)

		//left node must be less than node but more than node's parent
		if n.node.left != nil && (n.node.left.value.(int) >= n.node.value.(int) ||
			n.node.left.value.(int) <= n.min) {
			return false
			//right node must be more than node but less than node's parent
		} else if n.node.right != nil &&
			(n.node.right.value.(int) <= n.node.value.(int) ||
				n.node.right.value.(int) >= n.max) {
			return false
		}

		//Add children to queue
		if n.node.left != nil {
			q.Enqueue(node{n.node.left, n.min, n.node.value.(int)})
		}
		if n.node.right != nil {
			q.Enqueue(node{n.node.right, n.node.value.(int), n.max})
		}

	}

	return true
}
