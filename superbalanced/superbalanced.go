package main

import "math"

//TreeNode ...
type TreeNode struct {
	Data  interface{}
	Left  *TreeNode
	Right *TreeNode
}

//New ...
func New(data interface{}) *TreeNode {
	var node = new(TreeNode)
	node.Data = data
	node.Left, node.Right = nil, nil
	return node
}

//IsTreeSuperBalanced ...
func IsTreeSuperBalanced(node *TreeNode) bool {

	var queue []TreeNode

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

func main() {

}
