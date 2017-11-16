package balancedbtree

import "math"

//Node ...
type Node struct {
	//Val ...
	Val int
	//Left ...
	Left *Node
	//Right ...
	Right *Node
}

//IsTreeBalanced Finds out if a tree is balanced by adding the depths of the
//leafs to an array and comparing if there are more than 2 different depths
//or if the difference between depths is > 1
func IsTreeBalanced(root *Node) bool {

	//Lets create a slice of nodes that will ser
	type node struct {
		Node  *Node
		Depth int
	}
	queue := []node{{root, 0}}

	//map that will tell us if the array contains a depth already
	depthsMap := make(map[int]int)

	depths := []int{}

	for len(queue) > 0 {

		//extract node,depth and remove node from the slice
		n := queue[0].Node
		depth := queue[0].Depth
		queue = queue[1:]

		//If node is a leaf lets evaluate the depths
		if n.Left == nil && n.Right == nil {

			//Not in the map yet
			if depthsMap[depth] == 0 {
				//add to the map
				depthsMap[depth] = depth
				//add to the array
				depths = append(depths, depth)

			}

			if len(depths) > 2 || (len(depths) == 2 &&
				math.Abs(float64(depths[0]-depths[1])) > 1) {
				return false
			}
		} else {
			//lets process the node
			if n.Left != nil {
				queue = append(queue, node{n.Left, depth + 1})
			}
			if n.Right != nil {
				queue = append(queue, node{n.Right, depth + 1})
			}
		}
	}

	return true
}

//IsTreeSuperBalancedBF Checks if a tree is balanced with BF approach
func IsTreeSuperBalancedBF(node *Node) bool {

	//Create a queue and
	var queue []Node
	queue = append(queue, *node)

	for len(queue) > 0 {
		n := queue[0]

		var left, right int

		if n.Left != nil {
			queue = append(queue, *n.Left)
			left++
			if n.Left.Left != nil || n.Left.Right != nil {
				left++
			}
		}
		if n.Right != nil {
			queue = append(queue, *n.Right)
			right++
			if n.Right.Left != nil || n.Right.Right != nil {
				right++
			}
		}

		if math.Abs(float64(left-right)) > 1 {
			return false
		}

		queue = queue[1:]

	}
	return true
}
