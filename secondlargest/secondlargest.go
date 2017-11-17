package secondlargest

//Node Tree node
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

//FindSecondLargest Finds the second largest element in a binary search tree
//Function assumes the tree is a Binary search tree
func FindSecondLargest(root *Node) int {

	//Validate the tree has at least two nodes
	if root == nil || (root.Left == nil && root.Right == nil) {
		panic("Not enough nodes")
	}

	secondLargest, largest := root, root

	//Find the parent of the largest element
	for largest.Right != nil {
		secondLargest, largest = largest, largest.Right
	}

	//If the largest element has a left node, set secondLargest as the left node
	//And find the new largest
	if largest.Left != nil {
		secondLargest = largest.Left

		for secondLargest.Right != nil {
			secondLargest = secondLargest.Right
		}
	}

	return secondLargest.Value
}
