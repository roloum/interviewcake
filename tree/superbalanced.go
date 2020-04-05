package tree

import (
	"math"

	"github.com/golang-collections/collections/stack"
)

//IsBalanced Determines if the tree is superbalanced
//A tree is "superbalanced" if the difference between the depths of any two
//leaf nodes ↴ is no greater than one.
func IsBalanced(root *Node) bool {

	//Struct for stack
	type node struct {
		node  *Node
		depth int
	}

	depths := []int{}

	nodes := stack.New()
	//Push root to stack
	nodes.Push(node{root, 0})

	//Until there are no more leafs in the stack
	for nodes.Len() > 0 {

		n := nodes.Pop().(node)

		//If it's a leaf
		if n.node.left == nil && n.node.right == nil {

			//Add depth to array if it does not exist yet
			if !inArray(depths, n.depth) {
				depths = append(depths, n.depth)
			}

			//More than 2 depths or difference between depths>1, it's not balanced
			if len(depths) > 2 || (len(depths) == 2 &&
				math.Abs(float64(depths[0]-depths[1])) > 1) {
				return false
			}
		} else {
			//Add left node to stack
			if n.node.left != nil {
				nodes.Push(node{n.node.left, n.depth + 1})
			}
			//Add right node to stack
			if n.node.right != nil {
				nodes.Push(node{n.node.right, n.depth + 1})
			}
		}

	}

	return true
}

func inArray(array []int, number int) bool {
	for _, n := range array {
		if number == n {
			return true
		}
	}
	return false
}
