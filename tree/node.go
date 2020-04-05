package tree

//Node ...
type Node struct {
	value interface{}
	//left and right have to be references
	left, right *Node
}

//NewNode ...
func NewNode(value interface{}) *Node {
	n := Node{value, nil, nil}
	return &n
}

//InsertLeft ...
func (n *Node) InsertLeft(value interface{}) *Node {
	n.left = NewNode(value)
	return n.left
}

//InsertRight ...
func (n *Node) InsertRight(value interface{}) *Node {
	n.right = NewNode(value)
	return n.right
}
